# digital ocean -- how to use go modules
[link](https://www.digitalocean.com/community/tutorials/how-to-use-go-modules)

## getting started

0. create a projects folder to contain the module, then create the module folder `mymodule`
1. `go mod init <name>` creates a go module <name> with the go.mod file

## understanding the go.mod file

sample:

```
module <name>

go <version>
```

0. the first line, the `module` directive, tells go the name of your module so that when it is looking at import paths, it knows not to look elsewhere for <name>
1. the only other directive is the `go` directive, which tells go which version the module targets

## adding go code

0. `main.go` signifies the starting point of a program; the name isn't truly as important as the `main` function inside of it
1. you may name packages whatever you like, but main is a special case, as when go sees the main function inside of the main package, it knows the main function is the first function to run - otherwise known as the entry point

## adding a package to the module

a module may contain a number of packages and sub-packages; or nonoe

0. a package is created with a dir <name>, with <name>.go containing the package <name>
1. ALL EXPORTED FUNCTIONS MUST START WITH A CAPITAL LETTER, ie `PrintHello`
2. use the `import` command to import the module to the main file

```main.go
import "<module_name/package_name>"
```

## adding a remote module as a dependency

go modules are distributed from version controlled repos; adding new modules as a dependency is easy, just use the repo's path as the reference for the module, and import packages as normal

```
import "github.com/spf13/cobra"
//               ^module/^package
```

0. use `go get` before importing, so for the previous example, `go get github.com/spf13/cobra`
- using go get, go makes an HTTP/GET request for the repo link, determines which version of cobra is latest by looking at the branches and tags, and downloads that version
- it keeps track of which one to choose by adding the version in the go.mod file
- ADDITIONALLY, modules are stored by version in `$GO_PATH/pkg/mod/<source>/<organization>/<module>/<version>`
1. once a module has been gotten by way of `go get`, go stores the imports in the go.mod file
- once importing github.com/spf13/cobra, you will see the following in your go.mod

```
require (
	gitub.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/cobra v1.2.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)
```

the `require` directive tells go which modules you want, the version, and sometimes // indirect
- this says that, at the time the require directive was added, the module is not referenced directly in any of the module's source files

## using a specific version of a module



## conclusion









