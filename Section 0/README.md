# Section 0 - Setup for Go Development

## First of all we need to install Go.

- [Install from here](https://golang.org/dl/)
- If you are a Mac user and want to use terminal for installation run
  ```bash
    brew install go
  ```

## After the installation process, to check if installation is succesfully completed simply run this command in terminal

```bash
  go
```

- If the installation process succesfully completed terminal output should be like this:

```bash
  Go is a tool for managing Go source code.

  Usage:

        go <command> [arguments]

  The commands are:

        bug         start a bug report
        build       compile packages and dependencies
        clean       remove object files and cached files
        doc         show documentation for package or symbol
        env         print Go environment information
        fix         update packages to use new APIs
        fmt         gofmt (reformat) package sources
        generate    generate Go files by processing source
        get         add dependencies to current module and install them
        install     compile and install packages and dependencies
        list        list packages or modules
        mod         module maintenance
        run         compile and run Go program
        test        test packages
        tool        run specified go tool
        version     print Go version
        vet         report likely mistakes in packages

  Use "go help <command>" for more information about a command.

  Additional help topics:

        buildconstraint build constraints
        buildmode       build modes
        c               calling between Go and C
        cache           build and test caching
        environment     environment variables
        filetype        file types
        go.mod          the go.mod file
        gopath          GOPATH environment variable
        gopath-get      legacy GOPATH go get
        goproxy         module proxy protocol
        importpath      import path syntax
        modules         modules, module versions, and more
        module-get      module-aware go get
        module-auth     module authentication using go.sum
        packages        package lists and patterns
        private         configuration for downloading non-public code
        testflag        testing flags
        testfunc        testing functions
        vcs             controlling version control with GOVCS

  Use "go help <topic>" for more information about that topic.
```

## Install the text editor of your choice.

- My personal favourite editor is VSCode so you can download VSCode from [here](https://code.visualstudio.com/)

## Install the Go Extension for VSCode

- This extension simply makes easier to write Go code in VSCode. You can check the original documentation from [here](https://code.visualstudio.com/docs/languages/go). To install extension go to extensions tab in VSCode and Search for Go. The first result should be the extension you are installing.

- After installation of extension open a Go file and click the extension should be pop-up some errors because we need to install some tools. Click install all and your terminal should propmt this. Remember I am using MacOS so your output might be different then mine.

```bash
Tools environment: GOPATH=/Users/<your_username>/go
Installing 5 tools at /Users/mscandan/go/bin in module mode.
  gopkgs
  go-outline
  dlv
  staticcheck
  gopls

Installing github.com/uudashr/gopkgs/v2/cmd/gopkgs (/Users/mscandan/go/bin/gopkgs) SUCCEEDED
Installing github.com/ramya-rao-a/go-outline (/Users/mscandan/go/bin/go-outline) SUCCEEDED
Installing github.com/fatih/gomodifytags (/Users/mscandan/go/bin/gomodifytags) SUCCEEDED
Installing github.com/josharian/impl (/Users/mscandan/go/bin/impl) SUCCEEDED
Installing github.com/haya14busa/goplay/cmd/goplay (/Users/mscandan/go/bin/goplay) SUCCEEDED
Installing github.com/go-delve/delve/cmd/dlv (/Users/mscandan/go/bin/dlv) SUCCEEDED
Installing github.com/go-delve/delve/cmd/dlv (/Users/mscandan/go/bin/dlv) SUCCEEDED
Installing honnef.co/go/tools/cmd/staticcheck (/Users/mscandan/go/bin/staticcheck) SUCCEEDED
Installing github.com/go-delve/delve/cmd/dlv@master (/Users/mscandan/go/bin/dlv-dap) SUCCEEDED
Installing honnef.co/go/tools/cmd/staticcheck (/Users/mscandan/go/bin/staticcheck) SUCCEEDED
Installing golang.org/x/tools/gopls (/Users/mscandan/go/bin/gopls) SUCCEEDED

All tools successfully installed. You are ready to Go :).
Installing golang.org/x/tools/gopls (/Users/mscandan/go/bin/gopls) SUCCEEDED

All tools successfully installed. You are ready to Go :).
```