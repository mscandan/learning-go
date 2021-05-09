# After our text editor and tools setup, we are ready to write our first lines of codes with Go.

## Open your folder and create 'main.go' file.

### Every tutorial starts with a "Hello World!"

```go
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```

### How to run the code that we wrote? This is simple just run:

```sh
go run main.go
```

### If you want to see other go commands just type 'go' in terminal and hit return.

```sh
go
```

### If you want to checkout other commands and their documentation check out this [link.](https://golang.org/cmd/go/)

### Let's compile our code and get executable output for this just run

```sh
go build main.go
```

### This command should create file named 'main'. If you want to execute this file simply run:

```sh
./main
```

### We wrote our first lines of code and we executed it but what did we do really?

- We created a package named main.
  - What does it mean? Package is simply our project/workspace.
  - Why we named it main? In go there are 2 types of packages. Executable and reusable packages.
  - Executable packages generates a file that we can run. The word main is used for creating executable package.
  - Main package is must have a function called main.
  - Reusable packages, used as 'helpers' and we will put our reusable codes there. They will be our libraries.
- We imported a library called fmt
  - What does it mean? We use import for accessing the code that imported package contains.
  - If we import a library we can use their functions. Just like when we wrote:
    ```go
    fmt.Println("hello world")
    ```
  - If you want to learn more about packages and discover them click [here](https://golang.org/pkg/)
- We wrote our main function.
  - func is short for function. Functions are the same with in other programming languages.
  - Keyword 'func' means that we are creating function.
  - main is our function's name.
  - () is where will our parameters for functions will go.
  - Our main functions is special. When we try to run our go file. main functions automatically runs. So if we want to action of any code we need to put it in main func.