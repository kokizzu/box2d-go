package main

import (
	"bufio"
	"flag"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"golang.org/x/tools/go/ast/astutil"
	ccgo "modernc.org/ccgo/v4/lib"
)

// PrefixField is a unicode F symbol that is not part of the box2d code.
// We use this to easily identify & capitalize field names
const PrefixField = "ｆ"

var PREDEF_ATOMIC = `
#include <stdbool.h>
#include <stdint.h>
bool __builtin___atomic_compare_exchange_n(int32_t *ptr, int32_t *expected, int32_t desired, bool weak, int32_t success_memorder, int32_t failure_memorder);
`

var PREDEF_ALLOC = `
#include <stddef.h>
#define aligned_alloc __builtin_aligned_alloc
extern void* __builtin_aligned_alloc(size_t align, size_t size);
`

func main() {
	// defer profile.Start(profile.CPUProfile).Stop()

	var noTranspile bool

	flag.BoolVar(&noTranspile, "no-transpile", false, "do not transpile the c sources")
	flag.Parse()

	goos := runtime.GOOS
	goarch := runtime.GOARCH

	files, err := filepath.Glob("../box2d-c/src/*.c")
	must(err)

	args := []string{
		"ccgo",
		"-verify-types",
		"--package-name", "box2d",
		"-D", "BOX2D_DISABLE_SIMD",
		"--predef", PREDEF_ATOMIC,
		"--predef", PREDEF_ALLOC,
		"--prefix-field", PrefixField,
		"-o", "../box2d_c.go",
		"-I", "../box2d-c/include/",
	}

	args = append(args, files...)

	if !noTranspile {
		slog.Info("Transpile c to go")
		err := ccgo.NewTask(goos, goarch, args, os.Stdout, os.Stderr, nil).Main()
		must(err)
	}

	slog.Info("Read public API from header files")
	api := readAPI()

	slog.Info("Read generated go file")
	bytesSrc, err := os.ReadFile("../box2d_c.go")
	must(err)

	source := string(bytesSrc)

	slog.Info("Perform code cleanup via regexp")
	source = regexp.MustCompile("//go:build .*").ReplaceAllLiteralString(source, "")

	source = regexp.MustCompile(PrefixField+"([a-z])([A-Za-z0-9_]*)").ReplaceAllStringFunc(source, func(s string) string {
		name := strings.TrimPrefix(s, PrefixField)
		name = strings.ToUpper(name[:1]) + name[1:]
		return name
	})

	slog.Info("Parse generated go source")
	fset := token.NewFileSet()
	var root ast.Node
	root, err = parser.ParseFile(fset, "box2d_c.go", source, parser.ParseComments|parser.SkipObjectResolution)
	must(err)

	slog.Info("Apply code migrations")
	root = astutil.Apply(root, nil, removeModernC())
	root = astutil.Apply(root, nil, hideSupportFuncs())
	root = astutil.Apply(root, nil, fixQSortInvocation())
	root = astutil.Apply(root, nil, useApiTypes(api.ExposedTypes))
	ast.Inspect(root, renameApiTypes(api.ExposedTypes))
	ast.Inspect(root, removeTypeAliases())

	writeSource("../box2d_c.go", root)

	var decls []ast.Decl

	decls = append(decls, &ast.GenDecl{
		Tok: token.IMPORT,
		Specs: []ast.Spec{
			&ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: `"unsafe"`,
				},
			},
			&ast.ImportSpec{
				Path: &ast.BasicLit{
					Kind:  token.STRING,
					Value: `"runtime"`,
				},
			},
		},
	})

	createClassTypes(&decls)

	slog.Info("Generate go source for box2d api")
	ast.Inspect(root, createWrappers(&decls, api))
	ast.Inspect(root, exportVars(&decls, api.Exposed))

	writeSource("../box2d_api.go", &ast.File{
		Name:  ast.NewIdent("box2d"),
		Decls: decls,
	})
}

func writeSource(target string, src ast.Node) {
	slog.Info("Write " + target)

	fp, err := os.Create(target)
	must(err)
	defer fp.Close()

	w := bufio.NewWriter(fp)
	defer func() { must(w.Flush()) }()

	must(format.Node(w, token.NewFileSet(), src))
}

func useApiTypes(api map[string]bool) astutil.ApplyFunc {
	return func(cursor *astutil.Cursor) bool {
		if field, ok := cursor.Node().(*ast.Field); ok {
			field.Type = renameApiIdent(field, api)
		}

		return true
	}
}

func removeModernC() astutil.ApplyFunc {
	return func(cursor *astutil.Cursor) bool {
		n, ok := cursor.Node().(*ast.ImportSpec)
		if !ok {
			return true
		}

		if strings.Contains(n.Path.Value, "modernc.org/") {
			cursor.Delete()
		}

		return true
	}
}

type Visitor = func(ast.Node) bool

func renameApiTypes(api map[string]bool) Visitor {
	return func(node ast.Node) bool {
		switch node := node.(type) {
		case *ast.Ident:
			if api[node.Name] {
				node.Name = node.Name[2:]
			}
		}

		return true
	}
}

func removeTypeAliases() Visitor {
	return func(node ast.Node) bool {
		switch node := node.(type) {
		case *ast.TypeSpec:
			_, ok := node.Type.(*ast.StructType)
			if ok {
				node.Assign = token.NoPos
			}
		}

		return true
	}
}

func exportVars(decls *[]ast.Decl, api map[string]bool) Visitor {
	return func(node ast.Node) bool {
		if co, ok := node.(*ast.GenDecl); ok && (co.Tok == token.CONST || co.Tok == token.VAR) {
			for _, spec := range co.Specs {
				if dec, ok := spec.(*ast.ValueSpec); ok {
					ident := dec.Names[0]

					if api[ident.Name] {
						if !strings.HasPrefix(ident.Name, "b2_") {
							panic("expected ident to have b2_ prefix, got " + ident.Name)
						}

						name := ident.Name[3:]
						name = strings.ToUpper(name[:1]) + name[1:]

						if strings.HasSuffix(name, "Count") {
							continue
						}

						var ty ast.Expr

						switch {
						case strings.HasPrefix(name, "Color"):
							ty = ast.NewIdent("HexColor")

						case strings.HasSuffix(name, "Joint"):
							ty = ast.NewIdent("JointType")
							name = "JointType" + strings.TrimSuffix(name, "Joint")

						case strings.HasSuffix(name, "Body"):
							ty = ast.NewIdent("BodyType")

						case strings.HasSuffix(name, "Shape"):
							ty = ast.NewIdent("ShapeType")

						case strings.HasPrefix(name, "ToiState"):
							ty = ast.NewIdent("TOIState")

						default:
							panic("unknown enum type for " + name)
						}

						*decls = append(*decls, &ast.GenDecl{
							Tok: co.Tok,
							Specs: []ast.Spec{
								&ast.ValueSpec{
									Names:  []*ast.Ident{ast.NewIdent(name)},
									Values: []ast.Expr{ident},
									Type:   ty,
								},
							},
						})
					}
				}
			}
		}

		return true
	}
}

func createClassTypes(decls *[]ast.Decl) {
	for _, ty := range classTypes {
		var fields ast.FieldList

		if ty.Extends == "" {
			fields.List = append(fields.List, &ast.Field{
				Names: []*ast.Ident{ast.NewIdent("Id")},
				Type:  ast.NewIdent(ty.IdType),
			})

			fields.List = append(fields.List, &ast.Field{
				Names: []*ast.Ident{ast.NewIdent("tls")},
				Type:  &ast.StarExpr{X: ast.NewIdent("TLS")},
			})
		} else {
			fields.List = append(fields.List, &ast.Field{
				Type: ast.NewIdent(ty.Extends),
			})
		}

		*decls = append(*decls, &ast.GenDecl{
			Tok: token.TYPE,
			Specs: []ast.Spec{
				&ast.TypeSpec{
					Name: ast.NewIdent(ty.Name),
					Type: &ast.StructType{
						Fields: &fields,
					},
				},
			},
		})
	}
}

func createWrappers(decls *[]ast.Decl, api API) Visitor {
	return func(node ast.Node) bool {
		decl, ok := node.(*ast.FuncDecl)
		if !ok {
			return true
		}

		if !api.Exposed[decl.Name.Name] {
			return true
		}

		var classType *ClassTypes

		for idx := range classTypes {
			ty := &classTypes[idx]
			//if strings.HasPrefix(decl.Name.Name, ty.FnPrefix) {
			if decl.Type.Params != nil && len(decl.Type.Params.List) >= 2 {
				arg0, ok := decl.Type.Params.List[1].Type.(*ast.Ident)
				if ok && arg0.Name == ty.IdType {
					classType = ty
					break
				}
			}
			//}
		}

		var args []ast.Expr

		args = append(args, &ast.SelectorExpr{
			X:   ast.NewIdent("b"),
			Sel: ast.NewIdent("tls"),
		})

		var body []ast.Stmt

		var paramTypes []*ast.Field
		for _, param := range decl.Type.Params.List[1:] {
			for _, name := range param.Names {
				var arg ast.Expr = name

				paramType := renameApiIdent(param, api.ExposedTypes)

				if isUintptr(param.Type) && api.PointerTypes[decl.Name.Name][name.Name].IsValid() {
					actualType := api.PointerTypes[decl.Name.Name][name.Name]

					if actualType.Const {
						paramType = ast.NewIdent(actualType.GoType(&api))

						stackIdent := ast.NewIdent(name.Name + "Stack")

						body = append(body,

							&ast.AssignStmt{
								Tok: token.DEFINE,
								Lhs: []ast.Expr{stackIdent},
								Rhs: []ast.Expr{
									&ast.CallExpr{
										Fun: ast.NewIdent("copyToStack"),
										Args: []ast.Expr{
											&ast.SelectorExpr{
												X:   ast.NewIdent("b"),
												Sel: ast.NewIdent("tls"),
											},
											name,
										},
									},
								},
							},
							&ast.DeferStmt{
								Call: &ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X:   stackIdent,
										Sel: ast.NewIdent("Free"),
									},
								},
							},
						)

						arg = &ast.SelectorExpr{
							X:   stackIdent,
							Sel: ast.NewIdent("Addr"),
						}
					} else {
						paramType = &ast.StarExpr{
							X: ast.NewIdent(actualType.GoType(&api)),
						}

						body = append(body, &ast.ExprStmt{
							X: &ast.CallExpr{
								Fun:  ast.NewIdent("escapes"),
								Args: []ast.Expr{name},
							},
						})

						// prepend defer
						body = append([]ast.Stmt{
							&ast.DeferStmt{
								Call: &ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X:   ast.NewIdent("runtime"),
										Sel: ast.NewIdent("KeepAlive"),
									},
									Args: []ast.Expr{name},
								},
							},
						}, body...)

						arg = &ast.CallExpr{
							Fun: ast.NewIdent("uintptr"),
							Args: []ast.Expr{
								&ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X:   ast.NewIdent("unsafe"),
										Sel: ast.NewIdent("Pointer"),
									},

									Args: []ast.Expr{name},
								},
							},
						}
					}
				}

				if classType := getClassTypeById(identOf(paramType)); classType != nil {
					name.Name = strings.TrimSuffix(name.Name, "Id")

					paramType = ast.NewIdent(classType.Name)

					arg = &ast.SelectorExpr{
						X:   arg,
						Sel: ast.NewIdent("Id"),
					}
				}

				paramTypes = append(paramTypes, &ast.Field{
					Names: []*ast.Ident{name},
					Type:  paramType,
				})

				args = append(args, arg)
			}
		}

		var returnTypes []*ast.Field
		if decl.Type.Results != nil {
			for _, param := range decl.Type.Results.List {
				ty := renameApiIdent(param, api.ExposedTypes)

				wrapTy := getClassTypeById(identOf(ty))
				if wrapTy != nil {
					ty = ast.NewIdent(wrapTy.Name)
				}

				if identOf(ty) == "Joint" && strings.HasPrefix(decl.Name.Name, "b2Create") {
					jointType := strings.TrimPrefix(decl.Name.Name, "b2Create")
					ty = ast.NewIdent(jointType)
				}

				for _, name := range param.Names {
					returnTypes = append(returnTypes, &ast.Field{
						Type:  ty,
						Names: []*ast.Ident{name},
					})
				}
			}
		}

		call := &ast.CallExpr{
			Fun:  decl.Name,
			Args: args,
		}

		if decl.Type.Results != nil && len(decl.Type.Results.List) > 0 {
			assign := &ast.AssignStmt{
				Tok: token.ASSIGN,
				Rhs: []ast.Expr{call},
			}

			for _, ret := range decl.Type.Results.List {
				for _, name := range ret.Names {
					isWrapper := getClassTypeById(identOf(ret.Type)) != nil
					if isWrapper {
						assign.Lhs = append(assign.Lhs, &ast.SelectorExpr{
							X:   name,
							Sel: ast.NewIdent("Id"),
						})

						// body.tls = b.tls
						body = append(body, &ast.AssignStmt{
							Tok: token.ASSIGN,
							Lhs: []ast.Expr{
								&ast.SelectorExpr{
									X:   name,
									Sel: ast.NewIdent("tls"),
								},
							},
							Rhs: []ast.Expr{
								&ast.SelectorExpr{
									X:   ast.NewIdent("b"),
									Sel: ast.NewIdent("tls"),
								},
							},
						})
					} else {
						assign.Lhs = append(assign.Lhs, name)
					}
				}
			}

			body = append(body, assign)

			body = append(body, &ast.ReturnStmt{})
		} else {
			body = append(body, &ast.ExprStmt{X: call})
		}

		var fdecl *ast.FuncDecl

		if classType != nil {
			// adjust the id parameter value to come from the wrapper type
			args[1] = &ast.SelectorExpr{
				X:   ast.NewIdent("b"),
				Sel: ast.NewIdent("Id"),
			}

			name := strings.TrimPrefix(decl.Name.Name, classType.FnPrefix)
			name = strings.TrimPrefix(name, "b2")

			fdecl = &ast.FuncDecl{
				Name: ast.NewIdent(name),
				Recv: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{ast.NewIdent("b")},
							Type:  ast.NewIdent(classType.Name),
						},
					},
				},
				Type: &ast.FuncType{
					TypeParams: decl.Type.TypeParams,
					Params:     &ast.FieldList{List: paramTypes[1:]},
					Results:    &ast.FieldList{List: returnTypes},
				},
				Body: &ast.BlockStmt{List: body},
			}
		} else {
			slog.Info("Wrap normal method", slog.String("method", decl.Name.Name))

			// normal method
			fdecl = &ast.FuncDecl{
				Name: ast.NewIdent(strings.ReplaceAll(decl.Name.Name[2:], "_", "")),
				Recv: &ast.FieldList{
					List: []*ast.Field{
						{
							Names: []*ast.Ident{ast.NewIdent("b")},
							Type:  ast.NewIdent("Box2D"),
						},
					},
				},
				Type: &ast.FuncType{
					TypeParams: decl.Type.TypeParams,
					Params:     &ast.FieldList{List: paramTypes},
					Results:    &ast.FieldList{List: returnTypes},
				},
				Body: &ast.BlockStmt{List: body},
			}
		}

		*decls = append(*decls, fdecl)

		return true
	}
}

func needClassTypeWrap(list []*ast.Field) bool {
	for _, f := range list {
		if getClassTypeById(identOf(f.Type)) != nil {
			return true
		}
	}

	return false
}

func identOf(expr ast.Expr) string {
	if ident, ok := expr.(*ast.Ident); ok {
		return ident.Name
	}

	return ""
}

func isUintptr(expr ast.Expr) bool {
	ident, ok := expr.(*ast.Ident)
	return ok && ident.Name == "uintptr"
}

func hideSupportFuncs() astutil.ApplyFunc {
	return func(cursor *astutil.Cursor) bool {
		if node, ok := cursor.Node().(*ast.SelectorExpr); ok {
			if ident, ok := node.X.(*ast.Ident); ok {
				if ident.Name == "libc" || ident.Name == "iqlibc" {
					if node.Sel.Name[0] == 'X' {
						node.Sel.Name = node.Sel.Name[1:]
					}

					if node.Sel.Name != "TLS" {
						node.Sel.Name = strings.ToLower(node.Sel.Name[:1]) + node.Sel.Name[1:]
					}

					cursor.Replace(node.Sel)
				}
			}
		}

		return true
	}
}

func fixQSortInvocation() astutil.ApplyFunc {
	return func(cursor *astutil.Cursor) bool {
		if call, ok := cursor.Node().(*ast.CallExpr); ok {
			if name, ok := call.Fun.(*ast.Ident); ok {
				if name.Name == "qsort" {
					ccgoCall, ok := call.Args[4].(*ast.CallExpr)
					if ok {

						if ccgoCall.Fun.(*ast.Ident).Name != "__ccgo_fp" {
							panic("expected call to __ccgo_fp")
						}

						call.Args[4] = ccgoCall.Args[0]
					}
				}
			}
		}

		return true
	}
}

func renameApiIdent(param *ast.Field, api map[string]bool) ast.Expr {
	ty, ok := param.Type.(*ast.Ident)
	if !ok {
		return param.Type
	}

	name := cleanApiName(api, ty.Name)
	if api[name] {
		tyCopy := *ty
		tyCopy.Name = name[2:]
		ty = &tyCopy
	}

	if name, ok := renameSimpleType(ty.Name); ok {
		tyCopy := *ty
		tyCopy.Name = name
		ty = &tyCopy
	}

	return ty
}

type ParamType struct {
	CType string
	Const bool
}

func (p ParamType) GoType(api *API) string {
	if api.ExposedTypes[p.CType] {
		return p.CType[2:]
	}

	return p.CType
}

func (p ParamType) IsValid() bool {
	return strings.HasPrefix(p.CType, "b2")
}

type API struct {
	Exposed      map[string]bool
	ExposedTypes map[string]bool

	// function to param name to type
	PointerTypes map[string]map[string]ParamType
}

func (a *API) addPointerType(fn, param string, ty ParamType) {
	fnt := a.PointerTypes[fn]

	if fnt == nil {
		fnt = map[string]ParamType{}
		a.PointerTypes[fn] = fnt
	}

	fnt[param] = ty
}

func readAPI() API {
	var api API

	api.Exposed = map[string]bool{}
	api.ExposedTypes = map[string]bool{}
	api.PointerTypes = map[string]map[string]ParamType{}

	headers, err := filepath.Glob("../box2d-c/include/box2d/*.h")
	if err != nil {
		panic(err)
	}

	reType := regexp.MustCompile(`typedef (?:struct|enum) (b2[A-Za-z0-9_]+)`)
	reFunc := regexp.MustCompile(`B2_API \S+ (b2[A-Za-z0-9_]+)`)
	reEnumValues := regexp.MustCompile(`(?m)^\s*(b2_[A-Za-z0-9_]+)`)
	rePointerTypes := regexp.MustCompile(`(const )?([A-Za-z0-9]+)\* ([A-Za-z0-9]+)`)

	for _, header := range headers {
		src, err := os.ReadFile(header)
		must(err)

		lines := strings.Split(string(src), "\n")

		for _, match := range reType.FindAllStringSubmatch(string(src), -1) {
			api.ExposedTypes[match[1]] = true
			api.Exposed[match[1]] = true
		}

		for _, match := range reFunc.FindAllStringSubmatch(string(src), -1) {
			api.Exposed[match[1]] = true
		}

		for _, match := range reEnumValues.FindAllStringSubmatch(string(src), -1) {
			api.Exposed[match[1]] = true
		}

		for _, line := range lines {
			match := reFunc.FindStringSubmatch(line)
			if len(match) == 0 {
				continue
			}

			fn := match[1]

			for _, match := range rePointerTypes.FindAllStringSubmatch(line, -1) {
				co, ty, param := match[1], match[2], match[3]
				api.addPointerType(fn, param, ParamType{
					CType: ty,
					Const: co != "",
				})
			}
		}
	}

	api.ExposedTypes["b2InternalAssertFcn"] = false
	api.Exposed["b2SetAssertFcn"] = false
	api.Exposed["b2SetAllocator"] = false

	api.Exposed["b2Body_SetName"] = false
	api.Exposed["b2Body_GetName"] = false

	api.Exposed["b2Body_GetJoints"] = false
	api.Exposed["b2Body_GetShapes"] = false
	api.Exposed["b2Chain_GetSegments"] = false
	api.Exposed["b2Shape_GetSensorOverlaps"] = false

	api.Exposed["b2World_CollideMover"] = false
	api.Exposed["b2World_SetRestitutionCallback"] = false
	api.Exposed["b2World_SetFrictionCallback"] = false
	api.Exposed["b2World_OverlapAABB"] = false
	api.Exposed["b2World_SetCustomFilterCallback"] = false
	api.Exposed["b2World_SetPreSolveCallback"] = false

	api.ExposedTypes["b2BodyEvents"] = false
	api.Exposed["b2World_GetBodyEvents"] = false

	api.ExposedTypes["b2SensorEvents"] = false
	api.Exposed["b2World_GetSensorEvents"] = false

	api.ExposedTypes["b2ContactEvents"] = false
	api.Exposed["b2World_GetContactEvents"] = false

	for key := range api.Exposed {
		if strings.Contains(key, "DynamicTree") {
			api.Exposed[key] = false
		}
	}

	return api
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

var reSimpleType = regexp.MustCompile("u?int(?:8|16|32|64)_t")

func renameSimpleType(name string) (string, bool) {
	if reSimpleType.MatchString(name) {
		return strings.TrimSuffix(name, "_t"), true
	}

	return name, false
}

func cleanApiName(api map[string]bool, name string) string {
	clean := strings.TrimRight(name, "0123456789")

	if !api[name] && api[clean] {
		return clean
	}

	return name
}

type ClassTypes struct {
	Name     string
	IdType   string
	FnPrefix string
	Extends  string
}

var classTypes = []ClassTypes{
	{
		Name:     "Body",
		IdType:   "BodyId",
		FnPrefix: "b2Body_",
	},
	{
		Name:     "Shape",
		IdType:   "ShapeId",
		FnPrefix: "b2Shape_",
	},
	{
		Name:     "World",
		IdType:   "WorldId",
		FnPrefix: "b2World_",
	},
	{
		Name:     "Joint",
		IdType:   "JointId",
		FnPrefix: "b2Joint_",
	},
	{
		Name:     "Chain",
		IdType:   "ChainId",
		FnPrefix: "b2Chain_",
	},
	{
		Name:     "PrismaticJoint",
		Extends:  "Joint",
		IdType:   "JointId",
		FnPrefix: "b2PrismaticJoint_",
	},
	{
		Name:     "MotorJoint",
		Extends:  "Joint",
		IdType:   "JointId",
		FnPrefix: "b2MotorJoint_",
	},
	{
		Name:     "RevoluteJoint",
		Extends:  "Joint",
		IdType:   "JointId",
		FnPrefix: "b2RevoluteJoint_",
	},
	{
		Name:     "DistanceJoint",
		Extends:  "Joint",
		IdType:   "JointId",
		FnPrefix: "b2DistanceJoint_",
	},
	{
		Name:     "WheelJoint",
		Extends:  "Joint",
		IdType:   "JointId",
		FnPrefix: "b2WheelJoint_",
	},
	{
		Name:     "WeldJoint",
		Extends:  "Joint",
		IdType:   "JointId",
		FnPrefix: "b2WeldJoint_",
	},
	{
		Name:     "MouseJoint",
		Extends:  "Joint",
		IdType:   "JointId",
		FnPrefix: "b2MouseJoint_",
	},
	{
		Name:     "FilterJoint",
		Extends:  "Joint",
		IdType:   "JointId",
		FnPrefix: "b2FilterJoint_",
	},
}

func getClassTypeById(idType string) *ClassTypes {
	for idx := range classTypes {
		if classTypes[idx].IdType == idType {
			return &classTypes[idx]
		}
	}

	return nil
}
