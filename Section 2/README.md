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
- Functions

  - Creating a function in go is very simple. The func keyword means that we are creating a function.
    ```go
      func printHelloWorld() {
        fmt.Println("Hello world")
      }
    ```
  - The code above is creates a new function called 'printHelloWorld' and in the function we are printing "Hello world".
  - Functions can have return values. What does that mean you might ask. It means in some functions we can make some operations on variables or else.
    ```go
      func returnHelloWorld() string {
        return "Hello world"
      }
    ```
    With this code above. We are creating a function called 'returnHelloWorld' and in the end of this function we are returning a "Hello world" string. You can assign the return value to the new variable in wherever you want you use the function.
    ```go
      helloWorld := returnHelloWorld() // this line of code creates a new variable and assigns the return value of returnHelloWorld function to it.
      fmt.Println(helloWorld)
      fmt.Println(returnHelloWorld()) // this line of code prints the return value of returnHelloWorld function.
    ```
  - When creating a function if you don't specify the type of return value you don't have to return anything from the function.
    ```go
      func sayMyName() {
        fmt.Println("Selcuk")
      }
    ```

- Array and Slice

  - Array
    - Array means fixed length list of things. When we creating an array we need to specify it's length.
  - Slice
    - Slice means an array that can grow or shrink.
    - Every element in slice must be of same type.
    - Declaring a slice
      ```go
        cars := []string{"Audi", "BMW", "Ferrari"}
      ```
    - Appending value to a slice
      ```go
        cars = append(cars, "Mercedes Benz")
      ```
  - They both must be defined with data type.
  - Indexes starts from 0.

- For loops

  - We can use for loops to iterate in a slice.
    ```go
      for i, car := range cars {
        fmt.Println(i, car)
      }
    ```
    In the code above, the i equals the current index of the current element, it starts from 0 and goes off to slice's length - 1 and the car means the slice's i indexed value.

- Custom Type Declerations
  We can declare a specific types for different occasions. As we can see in the "deck.go" file we've created a new type called deck and it is a slice of strings

- Receiver Functions
  In this code block below, we've created a function named 'print' for spesificly for variables in deck type. This is called 'receiver function' and we are enable to use it with every variable typed 'deck'. The '(d deck)' part is specifies that function is a recevier.

  ```go
    func (d deck) print(){
      for i, card := range d {
        fmt.Println(i, card)
      }
    }
  ```

- Multiple Return Values
  We've learned that the functions can return the value. They can return more than one. In the code block below the '(deck, deck)' part means that we are going to return 2 values typed deck.

  ```go
    func deal(d deck, handSize int) (deck, deck) {
      // logic
    }
    returnVar1, returnVar2 := deal(newDeck, 5)
  ```

- Testing in go
  We need to create a file that specific to test functions. In testing we have to write a meaningfull test to test our functions to check if the functions works how we wanted them to work. The "go test" command below is runs the tests that we wrote to check our functions are working properly
  ```sh
    go test
  ```
  We can run the test with the command below
  ```sh
    go test deck_test.go deck.go
  ```
