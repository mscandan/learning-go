- Structs in Go
  Structs is a configurable data type. You can declare a struct just like the code block below.

  ```go
    type structName struct {
      key typeofKey
      ...
      ...
    }
  ```

  Using struct is like using slices you can create a new variable typed struct that you wanted.

  - syntax 1
    ```go
      varName := structName{"value of key1", "value of key2"...}
    ```
    or
  - syntax 2

  ```go
     varName := structName{key1: valueOfKey1, key2: valueOfKey2}
  ```

  To access that value in a struct you can simply use the code below.

  ```go
    fmt.Println(varName.keyName)
  ```

- Updating struct values
  We can update the struct values by reassinging the struct variable that we want to change

  ```go
    package main

    type person struct {
      firstname string
      lastname string
    }
    func main() {
      var selcuk person
      fmt.Printf("%+v\n", selcuk)
      selcuk.firstname = "Selcuk"
      selcuk.lastname = "Candan"
      fmt.Printf("%+v\n", selcuk)
    }
  ```

- Using struct in struct

  - decleration
    ```go
      type contactInfo struct {
        email string
        zipCode int
      }
      type person struct {
        firstname string
        lastname string
        contanct contactInfo
      }
    ```
  - assinging values
    ```go
      selcuk := person{
        firstname: "Selcuk",
        lastname: "Candan",
        contact: contactInfo{
          email: "mehmetselcukcandan@icloud.com",
          zipCode: 11111
        }
      }
    ```

- Receiver functions with structs

  ```go
    type person struct {
      firstname string
      lastname string
    }
    func (p person) print() {
      fmt.Printf("%+v\n", p)
    }
  ```

- Pointers
  If you are beginner to the programming world. You may not know but heard about pointers. Pointers points to the address of a variable on ram. If you want to know more about them check out this [link](https://tour.golang.org/moretypes/1)

  ```go
    selcuk := 3 //declare a variable called selcuk and assign 3
    addressOfSelcuk := &selcuk // take the address of selcuk variable in the ram
    fmt.Println(*addressOfSelcuk) // print the value of addressOfSelcuk points
  ```

  - Why do we need pointers:
    <br />
    For example in go when you pass some variables as a parameter in a function and you change the value of the parameter in the function you will see the actually value didn't change. When you pass a variable as a parameter, go creates a copy of a variable that you passed as a parameter. So actually you are reassinging to a copy of a parameter, not the parameter itself. But if you use the pointer, that points to a parameter and if you use that parameter to assing a new value to a parameter, that function will work like how you want it to work.

- Value Types and Reference Types in Go
  - Value Types -> use pointers to change these
    - int
    - float
    - string
    - bool
    - struct
  - Reference Types
    - slices
    - maps
    - channels
    - pointers
    - functions
