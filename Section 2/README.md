# After our first lines of codes. Now we will be going deep on GoLang.

## Declaring variables

- Go is a staticly typed language. When we create a new variable, we have to set it's type.
  ```go
    var name string = "Selcuk"
  ```
  With this line of code. We are creating a variable typed 'string' and named 'name' and assing 'Selcuk' to it. The 'var' keyword means this is a variable.
  But in today's world the code above for declaring variable is not preferred. The preferred methot to declare variable is:
  ```go
    name := "Selcuk
  ```
  With this line of code you will create a variable named 'name' and assing 'Selcuk' to it. You don't have to specify the type of the variable. Go will figure it out for you.
- Basic (fundamental) variable types in Go
  - bool -> true or false
  - string -> sequence of characters
  - int -> integer numbers
  - float64 -> 64bit floating point numbers
- Re-assinging to variable
  When you re-assing a value to a variable, you need to use '=' not ':='. ':=' is used for only when declaring a variable.
  ```go
    name = "John"
  ```
