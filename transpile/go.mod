module github.com/oliverbestmann/box2d-go/transpile

go 1.24.5

require (
	golang.org/x/tools v0.36.0
	modernc.org/ccgo/v4 v4.28.0
)

require (
	github.com/dave/dst v0.27.3 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	golang.org/x/mod v0.27.0 // indirect
	golang.org/x/sync v0.16.0 // indirect
	modernc.org/cc/v4 v4.26.4 // indirect
	modernc.org/gc/v2 v2.6.5 // indirect
	modernc.org/mathutil v1.7.1 // indirect
	modernc.org/opt v0.1.4 // indirect
	modernc.org/sortutil v1.2.1 // indirect
	modernc.org/strutil v1.2.1 // indirect
	modernc.org/token v1.1.0 // indirect
)

// contains bugfix for int -> bool conversion
replace modernc.org/ccgo/v4 v4.28.0 => gitlab.com/oliver.bestmann/ccgo/v4 v4.0.0-20250827192611-8265e45e6562
