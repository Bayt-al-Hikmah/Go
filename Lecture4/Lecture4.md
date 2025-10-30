## Objectives
- Work with pointers
- Use structs and methods
- Implement interfaces and polymorphism
- Creating Enums With ``iota``
- Handle Errors
## Pointers
### Introduction
In Go, every variable is stored at a specific location in the computer's memory. Each location has a unique address that allows us to access it this is called a memory address. A **pointer** is a special type of variable that doesn’t store a normal value (like `5` or `"hello"`). Instead, it stores the **memory address** of another variable.

We use two special operators with pointers:
- The **`&`** (address-of) operator: Placed before a variable name, it gives us the memory address of that variable.
- The **`*`** (dereference) operator: Placed before a pointer variable, it gives us the value stored at the memory address the pointer holds. It's like "following" the pointer to its target.
### Creating Pointer
We create a pointer in Go using the `var` keyword, followed by the pointer name. Then we specify the type of value the pointer will point to, and place the `*` symbol before the type:
```go
var ptr *int  // A pointer to an integer
```
After that, we can store the memory address of a variable inside the pointer by assigning it the variable’s address. To do that, we place the `&` operator before the variable name:
```go
num := 42
ptr = &num
```
Now, if we want to access the value that the pointer is referring to, we use the `*` operator before the pointer name. This is called **dereferencing** the pointer.
```go
fmt.Println(*ptr) // this will display 42
```
**Example:**
```go
package main

import "fmt"

func main() {
    num := 42
    ptr := &num  // ptr holds the address of num

    fmt.Println("Value of num:", num)     // 42
    fmt.Println("Address of num:", &num)  // e.g., 0xc00001a0a8
    fmt.Println("Pointer ptr holds:", ptr) // same address
    fmt.Println("Value pointed by ptr:", *ptr) // 42
}
```
- &num → **address of** num
- \*ptr → **value at** the address stored in ptr (dereferencing)

If we don’t assign a value to a pointer, it will hold the value `nil`. A `nil` pointer does not point to any memory location or variable.
### Pointers in Functions
The most important use for pointers is to share memory with functions. By default, Go passes by value, which means when we pass a variable to a function, Go makes a copy of it. The function operates on the copy, and the original variable is unchanged.    
If we want a function to modify the original variable, we must pass a **pointer** to that variable instead.
```go
package main

import "fmt"

func doubleValue(val int) {
	val = val * 2
}

func doubleValueByPointer(valPtr *int) {
	*valPtr = *valPtr * 2
}

func main() {
	num := 5
	
	// Pass by value (passing 'num')
	doubleValue(num)
	fmt.Println("After doubleValue:", num) // Still 5
	
	// Pass by reference (passing the address of 'num')
	doubleValueByPointer(&num)
	fmt.Println("After doubleValueByPointer:", num) // Now 10
}
```
The first function `doubleValue` receives an `int` parameter. Since Go passes values by copy, the function only modifies a **local copy** of the variable. The original variable in `main` remains unchanged.  
The second function `doubleValueByPointer` receives a pointer to an `int` (`*int`). Using the dereference operator `*`, it modifies the **actual value in memory** that the pointer refers to. Because of this, the change persists after the function call.
### Allocating Memory Using `new()`
We can do more then that with pointers, Go provides us with the built-in **`new()`** function to allocate memory for a variable and return its **pointer**. When we use `new()`, Go allocates memory and gives us the **address** of that space even if we didn’t assign a value yet.
```go
ptr := new(int)  // allocates memory for an int, returns a pointer to it
```
The `ptr` holds a memory address, that contains the default value for an int (which is `0`), we can update the value using dereferencing:
```go
*ptr = 100
fmt.Println(*ptr) // 100
```
## Structs and Methods
### Introduction to Structs
A struct (short for structure) is a composite data type in Go that groups together one or more related variables, called fields, under a single name. Structs are used to represent real-world objects or concepts with multiple attributes for example, a person with a name and age, or a car with brand and model.   
Structs are Go’s way of creating custom data types, similar to classes in other languages, but without built-in inheritance. When we define a struct, we are actually creating a new type that we can use just like any built-in type (e.g., `int`, `string`).
### Defining and Using a Struct
We define a struct in Go using the `type` and `struct` keywords. First, we start with the `type` keyword, followed by the **name of the struct**, which becomes the name of our **new custom type**. Then we use the `struct` keyword and add a pair of curly braces `{}`. Inside the braces, we list the fields that the struct will contain, along with their types.
```go
type Person struct {
    Name string
    Age  int
}
```
Now that we have defined a struct, Go provides two common ways to create and use it:  
Declare a variable of the struct type and assign values using dot notation, Here we creates an empty `Person` struct, and we fill in each field one by one using the `.` operator.
```go
    var p2 Person
    p2.Name = "Bob"
    p2.Age = 30
```
Use a struct literal to initialize the fields immediately, this approach lets us set all field values right when we create the struct cleaner and more concise when we already know the values.
```go
p1 := Person{Name: "Alice", Age: 25}
```
### Adding Methods To Our Structs
A **method** in Go is just like a function, but it is **associated with a specific type**. We define a method by adding a **receiver parameter** between the `func` keyword and the method name.
The receiver acts like the “owner” of the method it allows the method to access the fields and behavior of the value it belongs to.
```go
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
```
In this example, we define a custom type called `Rectangle` using a struct. The `Rectangle` struct has two fields: `Width` and `Height`, both of type `float64`.    
Below the struct, we define a method named `Area` for the `Rectangle` type. The method uses a receiver `(r Rectangle)`, which means the method is associated with the `Rectangle` type and can access its fields. Inside the method, we calculate the area by multiplying `r.Width` by `r.Height`, and then we return that result.   
Because this method is tied to the `Rectangle` type, we can call it on any rectangle instance to easily compute its area.
```go
rect := Rectangle{Width: 10, Height: 5}
	
// Call the method using the dot '.' operator
fmt.Println("Area:", rect.Area())
```
### Pointer with Structs
In Go, we can also create **pointers for structs**, this allows us to directly modify the fields of a struct inside functions or methods.  
We create a pointer to a struct using the `&` operator, which gives us the address of a struct variable.
```go
type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{Name: "Charlie", Age: 28}

    // Create a pointer to the struct
    personPointer := &p

    // Access and modify fields using the pointer
    personPointer.Age = 29

    fmt.Println("Updated Age:", p.Age) // Output: 29
}
```
Even though `personPointer` holds the address of `p`, Go allows us to directly use the dot `.` operator to access and modify fields no need for manual dereferencing.
### Pointer Receivers in Methods
When defining methods, we can choose between value receivers (copy of the struct) and pointer receivers (reference to the original struct).  
A pointer receiver allows the method to modify the original struct’s data.
```go
type Counter struct {
	Value int
}

// Method with pointer receiver
func (c *Counter) Increment() {
	c.Value++
}
```
Here, the `Increment` method uses `*Counter` as the receiver, meaning it has access to the original struct and can update its fields. If we used a value receiver (without `*`), it would modify only a copy and the original value would stay unchanged.
```go
func main() {
	c := Counter{Value: 10}
	
	// Call method that modifies the original struct
	c.Increment()
	
	fmt.Println("Counter Value:", c.Value) // Output: 11
}
```
Also, even if we call the method using a value like c.Increment(), Go automatically converts it to a pointer call under the hood when the method requires a pointer receiver so we don’t have to manually write (&c).Increment().

## Interfaces and Polymorphism

### Introduction
An **interface** in Go is a type that defines a contract or a set of method signatures. It specifies what methods a type must have, but not how they are implemented.  
For example, an "Animal" interface might require a `Speak()` method. A "Shape" interface might require an `Area()` method.
### Implementing an Interface
To create an interface in Go, we start by defining a new type using the `type` keyword, followed by the **name of the interface** and the `interface` keyword. Inside the curly braces `{}`, we list one or more method signatures that any type must implement in order to satisfy the interface.
```go
type Speaker interface {
    Speak() string
}
```
Now we can create types like `Dog` or `Cat` and attach a `Speak` method to them.
```go
type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return d.Name + " says Woof!"
}

type Cat struct {
    Name string
}

func (c Cat) Speak() string {
    return c.Name + " says Meow!"
}
```
Both the **Cat** and **Dog** classes have the `Speak` method. If we want to create a function that takes one of these animals and makes it speak, it would be repetitive to write a separate function for each animal (one for cat, one for dog).  
Now imagine we have even more animal types repeating the same function logic would become inefficient and harder to maintain.  
To solve this, we use the **interface** we created. Instead of making the function accept a specific animal type, we make it accept the **interface** as the argument. This way, any class that implements the interface can be passed to the function, making the code reusable and scalable.
```go
func MakeSound(s Speaker) {
    fmt.Println(s.Speak())
}
```
Even when we want to create a slice, we can use the interface to store different types of animals in the same slice. This allows us to keep objects of different types (like Cat and Dog) together as long as they all implement the same interface.
```go
func main() {
    pets := []Speaker{
        Dog{Name: "Buddy"},
        Cat{Name: "Whiskers"},
    }

    for _, pet := range pets {
        MakeSound(pet)
    }
}
```
### The Empty Interface
An interface with no methods, `interface{}`, is called the empty interface. Since every type has zero or more methods, every type implicitly implements the empty interface. This means a variable of type `interface{}` can hold a value of any type (e.g., `int`, `string`, `Person`, etc.).
```go
package main

import "fmt"

// 'i' can hold any value
func describe(i interface{}) {
	fmt.Printf("Type: %T, Value: %v\n", i, i)
}

func main() {
	describe(10)
	describe("hello")
	describe(true)
}

```
In Go 1.18 and later, the keyword `any` was introduced as an alias for `interface{}`, making code cleaner and easier to understand.
```go
package main

import "fmt"

// 'i' can hold any value
func describe(i any) {
	fmt.Printf("Type: %T, Value: %v\n", i, i)
}

func main() {
	describe(10)
	describe("hello")
	describe(true)
}
```
## Enums
### Introduction
An enum (short for _enumeration_) is a special type used to represent a fixed set of related values for example, days of the week, directions, user roles, or program states. Enums help make code more readable, organized, and type-safe by replacing arbitrary numbers or strings with meaningful names.  

However, Go does **not** have a built-in `enum` keyword like some other languages (e.g., C, Java, or Rust).  
Instead, Go programmers **simulate enums** using **constants** (`const`) together with a special identifier called `iota`. It starts at `0` for the first constant in the block and automatically increments by `1` for each subsequent constant. This makes it easy to create sequences of related constants that behave like enums.
### Creating an Enum
To make our "enum" type-safe, we first define our own custom type (e.g., `Direction` based on `int`). Then, we create a `const` block using `iota` to assign the values.
```go
type Direction int

// 2. Create the enum values using const and iota
const (
    North Direction = iota // 0
    East                   // 1
    South                  // 2
    West                   // 3
)

```
In this example, we first define a new custom type called `Direction` based on the `int` type. Then, inside a `const` block, we create our enum values using `iota`. The first value, `North`, is set to `iota`, which starts at `0`, and each subsequent line automatically increments by 1, giving us `East = 1`, `South = 2`, and `West = 3`. This pattern allows us to represent a fixed set of related values our directions just like an enum in other languages.
### Working with Enum
Once we have defined our enum values, we can use them just like any other type in Go. A common way to work with enums is through a `switch` statement, which lets us handle each possible value in a clear and readable way.
```go
func PrintDirection(d Direction) {
    switch d {
    case North:
        fmt.Println("Moving North")
    case East:
        fmt.Println("Moving East")
    case South:
        fmt.Println("Moving South")
    case West:
        fmt.Println("Moving West")
    default:
        fmt.Println("Unknown direction")
    }
}

func main() {
    var myDirection Direction = North
    PrintDirection(myDirection) // Output: North
}
```
In this example, the `PrintDirection` function takes a value of type `Direction` and uses a `switch` statement to check which enum value it matches. Depending on the direction, it prints the appropriate message. In `main`, we create a variable `myDirection` and assign it the value `North`, then pass it to `PrintDirection`, which prints the corresponding message. 

## Error Handling

### Introduction
As programmers, we're human, which means our code will inevitably have to deal with errors. Some errors are simple syntax mistakes that stop our code from compiling. We fix those, and move on.

But a more complex class of errors happens at runtime, and these are often out of our direct control. What if a user enters an invalid value? What if a file we're trying to read doesn't exist, or a network request to an external resource fails?  
Unlike other languages that use `try-catch` blocks to "throw" exceptions, Go treats errors as values. This is "the Go way" of handling errors: functions that can fail simply return the error as their last return value. This makes error handling explicit and a first-class citizen in the code.
### The `error` Type
The `error` type is a built-in interface in Go. It's incredibly simple and defines just one method:
```go
type error interface {
    Error() string
}
```
Any type that implements this `Error() string` method satisfies the `error` interface. When a function returns a non-`nil` error, we can get the error message by calling this method (or just by printing the error, which calls the `Error()` method implicitly).
### Returning and Creating Errors
Now the Go pattern to handel Error is making our functions return  normal result value plus an error. If the operation is successful, the error is ``nil``. If it fails, the function returns a zero-value for the result (e.g., 0, "", or nil) and a non-nil error.  
To create a simple error, we use the built-in `errors` package or the `fmt` package.
- **`errors.New(message string)`**: Creates a basic error with a static message.
- **`fmt.Errorf(format string, ...)`**: Creates a formatted error, which is more flexible.

The idiomatic way to handle this is to check the error immediately with an `if err != nil` block.
```go
package main

import (
    "errors"
    "fmt"
)

// divide returns a result AND an error
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10.0, 2.0)
    
    if err != nil {
        fmt.Println("Error occurred:", err)
    } else {
        fmt.Println("Result 1:", result) 
    }

    result, err = divide(10.0, 0)
    if err != nil {
        fmt.Println("Error occurred:", err) 
    } else {
        fmt.Println("Result 2:", result)
    }
}
```
Here we returning two values from the `divide` function the result of the division and an error. If the second number (`b`) is zero, we return an error instead of dividing.   
In `main`, we call `divide` twice: the first call succeeds and prints the result, while the second call tries to divide by zero, triggering the error path. 
### Custom Errors and Error Wrapping
We can do more then that we  can create our own error types by defining a `struct` and implementing the `error` interface. This is useful for passing more context (like an error code) than just a string message.
```go
// MyError is a custom error type
type MyError struct {
    StatusCode int
    Message    string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("status %d: %s", e.StatusCode, e.Message)
}
```
Here we defined a custom error type called `MyError` with two fields: `StatusCode` and `Message`. By implementing the `Error()` method, this struct satisfies Go’s built-in `error` interface. That means values of `MyError` can be returned and handled just like standard errors, but with extra context such as an HTTP status code making error messages more informative.
### Error Chaining
Sometimes one error happens because of another error. In Go, we can **wrap an error** to keep that original context. This helps us understand where the problem started and makes debugging easier.  
We wrap errors using `fmt.Errorf` with the **`%w`** verb.  
Go then lets us check if the wrapped error chain contains a specific error type or value.
- **`errors.As(err, &target)`** checks if the error chain contains a specific **type** (e.g., `*FileError`) and stores it in `target` if found.
- **`errors.Is(err, target)`** checks if the error chain contains a specific **error value**.

**Example**
```go
package main

import (
    "errors"
    "fmt"
)

type DivideByZeroError struct {
    A, B float64
}

func (e *DivideByZeroError) Error() string {
    return fmt.Sprintf("cannot divide %.2f by %.2f", e.A, e.B)
}


func divideSafe(a, b float64) (float64, error) {
    if b == 0 {
        return 0, &DivideByZeroError{A: a, B: b}
    }
    return a / b, nil
}
func main() {
	a := 4.0
	b := 0.0

	result, err := divideSafe(a, b)

	if err != nil {
		// Check if error is our custom type
		var divErr *DivideByZeroError
		if errors.As(err, &divErr) {
			fmt.Println("Custom handling: division by zero is not allowed")
			fmt.Println("Details:", err)
		}
		return
	}

	fmt.Println("Result:", result)
}
```
We created a custom error type `DivideByZeroError` and returned it when someone tries to divide by zero. In `main`, we call `divideSafe`, check if a custom error occurred using `errors.As`, and handle it with a helpful message. If no error happens, we print the result.

## Tasks
### Task 1
Create a `BankAccount` struct with a `balance` field (type `float64`).

1. Implement a `Deposit` method with a pointer receiver that takes an `amount` and adds it to the balance.
    
2. Implement a `Withdraw` method with a pointer receiver that takes an `amount`. It should **return an error** (using `errors.New`) if the `amount` is greater than the `balance` ("insufficient funds"). Otherwise, it should subtract the amount and return `nil`.
    
3. In `main`, create an account, deposit some money, and then try to withdraw both a valid and an invalid amount, checking for errors each time.

### Task 2

Define an interface called `Writer` with one method: `Write(data string) (int, error)`.

1. Create a struct `ConsoleWriter`. Implement the `Writer` interface for it. The `Write` method should just print the `data` to the console using `fmt.Println` and return the length of the string and a `nil` error.
    
2. Create a struct `StringWriter` that has a field `content string`. Implement the `Writer` interface for it. The `Write` method should append the `data` to the `content` field and return the length of the _new_ `data` and a `nil` error.
    
3. In `main`, create an instance of both `ConsoleWriter` and `StringWriter` and use them (and a `fmt.Println` to show the `StringWriter`'s content).

### Task 3

Write a function `SafeSqrt(num float64) (float64, error)` that calculates the square root of a number.

1. If the number is negative, the function should return `0` and an `error` created with `fmt.Errorf("cannot calculate sqrt of negative number: %f", num)`.
    
2. If the number is non-negative, it should return the square root (using `math.Sqrt` from the `math` package) and a `nil` error.
    
3. In `main`, test this function with both a positive number and a negative number, checking the `err` value each time.