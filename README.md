# Hello World
Sample code, commands and Dockerfile to build and run a containerized app in Go.

# Go basics
Overview of the language and how to work with it.

## Installation
See [the tutorial](https://go.dev/doc/install) on how to install Go.

## Modules
Go code is contained in modules which can be shared across development projects.

A Go module contains its own code, the required packages and the Go version to run it.
Those dependencies are managed with the [go.mod](go.mod) file.
See [this overview](https://go.dev/ref/mod#modules-overview) for more details.

go.mod file example:
```
// Arbitrary name for the module of your code.
// Has to be a downloadable git or Go module proxy URL
// if the module should be shared across projects.
module github.com/OlGe404/kochorganisator9000

// Description of required Go modules for your own module.
// Explanation of available directives:

// require: Necessary modules in their minimum version.
// exclude: Ignore modules from being loaded.
// replace: Overwrite a version of a module with other content.
// retract: Marks a module version as faulty to prevent future work depending on it.

require (
    example.com/new/thing/v2 v2.3.4
    example.com/old/thing v1.2.3
)

exclude example.com/old/thing v1.2.3
replace example.com/bad/thing v1.4.5 => example.com/good/thing v1.4.5
retract [v1.9.0, v1.9.5]

// go version to use.
go 1.19
```

The go.mod file can be generated with `go mod init` when initializing a new project.

### Packages
Go packages are a way to group functions within the same directory.
Packages can be added with the "import" statement. There are default packages that are
included in the Go installation.

Those packages can be found at https://pkg.go.dev/std and don't have to be included as
modules in the go.mod file.
