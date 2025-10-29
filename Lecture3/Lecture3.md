## Objectives
- Working with Functions
- Working with the Packages and Modules
## Functions
### Introduction
A function is a reusable block of code designed to perform a specific task. Functions are fundamental to writing efficient and well-structured programs. By following the DRY (Don't Repeat Yourself) principle, functions help eliminate code duplication and promote cleaner code. They also make it easier to break down complex problems into smaller, manageable parts, resulting in code that is more organized, easier to understand, and simpler to maintain.
### Creating Functions
In Go, we define a function using the `func` keyword, followed by the function name, parameters, and return type(s). If a function does not take parameters and does not return any value, we can omit both the parameters and the return type.    
All Go programs begin execution in a package named `main`, and the entry point of the program is the `main` function. We declare it using the `func` keyword, followed by the function name `main`, and empty parentheses `()` since it does not take any arguments. The code that the function executes is placed inside curly braces `{ }`.
```go
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```
Let's create a new function that greets the user. We will declare this function outside the `main` function. It will take no arguments and return no values it will simply print `"Hello, user!"`.
```go
package main

import "fmt"

func sayHello() {
	fmt.Println("Hello, user!")
}

func main() {
	sayHello()
}
```
To execute the code inside a function, we need to call it. We do this by writing the function's name followed by parentheses. In our example, the parentheses are empty because the function does not take any arguments.
### Function with Parameters
To make a function more powerful, we can allow it to accept values that change its behavior depending on what the user passes in. We do this by adding parameters inside the parentheses after the function name. In Go, we **must** specify the type for each parameter.
```go
package main

import "fmt"

func main() {
	sayHello("Ali")
}

func sayHello(name string) {
	fmt.Println("Hello,", name)
}
```
In this example, we call the `sayHello` function and pass the value `"Ali"` as an argument. The function receives this value in its parameter `name`, which has the type `string`, then it  use this value to display a personalized greeting.
### Returning Value
Functions can do even more they can also return a value using the `return` keyword. For example, a function can perform a calculation and return the result. To enable a function to return a value in Go, we must specify the return type after the parameter list.
```go
package main

import "fmt"

func main() {
	fmt.Println(addTwoNumbers(4, 5))
}

// Parameters and return type are explicitly typed
func addTwoNumbers(num1 int, num2 int) int {
	var result int = num1 + num2
	return result
}
```
If consecutive parameters share the same type, we can specify the type once at the end. 
```go
package main

import "fmt"

func main() {
	fmt.Println(addTwoNumbers(4, 5))
}

// Parameters and return type are explicitly typed
func addTwoNumbers(num1, num2 int) int {
	var result int = num1 + num2
	return result
}
```
We can also make our functions return  multiple values, we do that by list all return types inside parentheses and separate them with commas.
```go
package main

import (
	"fmt"
	"errors"
)

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

func divide(num1, num2 int) (int, error) {
	if num2 == 0 {
		// errors.New creates a new error message
		return 0, errors.New("cannot divide by zero")
	}
	return num1 / num2, nil // nil means "no error"
}
```
In this example, the `divide` function returns two values: an `int` result and an `error`. If the second number (`num2`) is zero, the function cannot perform the division, so it returns `0` along with a new error message created using `errors.New`.    
If `num2` is not zero, the function performs the division and returns the result along with `nil`, which means “no error.”   
In the `main` function, we handle both return values. We check whether `err` is `nil`.  
- If it's not `nil`, that means an error occurred, so we print the error.
- Otherwise, we print the result.
### Variable Scope
Scope defines the region within a program where a particular variable or function can be accessed. Go's scoping rules are primarily based on blocks.  
#### Package Scope (Exported vs. Unexported)
In Go, there isn't a "global" scope in the traditional sense. Instead, Go has **package scope**. Visibility outside of a package is controlled by **capitalization**.
- If an identifier (variable, function, type) starts with a **capital letter**, it is **exported** and can be accessed by other packages.
- If it starts with a **lowercase letter**, it is **unexported** (private) and only accessible within its own package.
```python
package mypackage

// This variable is exported because it's capitalized
var ExportedMessage = "This is exported"

// This variable is unexported (private)
var internalMessage = "This is internal"
```
#### Local (Block) Scope
Variables declared inside a function are local to that function. Variables declared inside a block (like a `for` loop or `if` statement) are local to that block.
```go
package main

import "fmt"

func printSum(a, b int) {
	var result = a + b // 'result' is local to printSum
	fmt.Println(result)

	if result > 5 {
		var inBlock = "Only visible here" // 'inBlock' is local to this if-block
		fmt.Println(inBlock)
	}
	// fmt.Println(inBlock) // Error: 'inBlock' is not defined in this scope
}

func main() {
	printSum(3, 5)
	// fmt.Println(result) // Error: 'result' is not defined in this scope
}
```
In this example, variables are only accessible within the scope where they are defined. The variable `result` is created inside the `printSum` function, so it can only be used within that function. Inside the `if` block, we create another variable called `inBlock`, which is only accessible inside that block.  
This means:
- We can use `result` inside the `if` block because the block is inside the function’s scope.
- We **cannot** access `inBlock` outside the `if` block because it belongs only to that inner scope.
- Similarly, the `main` function cannot access the `result` variable, since it was declared inside `printSum`.
### Arbitrary Number of Arguments
Go gives us more flexibility when working with functions. Sometimes, we don't know in advance how many arguments a user will pass, but we still want our function to handle all of them. To solve this, Go allows us to create **variadic functions**, which can accept an arbitrary number of arguments. We do this by adding `...` (ellipsis) before the type of the last parameter.
```go
package main

import "fmt"

func main() {
	// We can pass any number of arguments
	sayHelloToAll("Mohamed", "Ali", "Ahmed")
	sayHelloToAll("Just_One")
}
// 'names' will be a slice of strings ([]string)
func sayHelloToAll(names ...string) {
	for _, name := range names {
		fmt.Println("Hello", name)
	}
}
```
We can pass  slice  to a variadic function by appending `...`.
```go
package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2, 3)) // Output: 6

	// Pass a slice
	mySlice := []int{4, 5, 6}
	fmt.Println(sum(mySlice...)) // Output: 15
}
```
### Recursive Functions
Recursive function are special function that have ability to call theirself untill a condition (that we call base state) is valid.  
lets suppose we want to create a function that calculate factorial of numbers  
we know that:

- 0! is equal to 1
- 1! is equal to 1
- 4! is equal to 4\*3\*2\*1
- 5! is equal to 5\*4\*3\*2\*1 = 5\*4!

with that in mind, we can set the base condition as if n == 0 we return 1, else we return n multiplied by the factorial of n-1 and so on
```go
package main

import "fmt"

func main() {
	fmt.Println(factorial(5)) // Output: 120
}

func factorial(n int) int {
	if n == 0 { // Base case
		return 1
	}
	// Recursive step
	return n * factorial(n-1)
}
```
### Anonymous Functions
Anonymous functions (or function literals) are functions without a name. They're useful for short, one-time-use functionality, especially for callbacks or goroutines.  
In Go, they can be assigned to variables, passed as arguments, or executed immediately.
```go
package main

import "fmt"

func main() {
	// Assign an anonymous function to a variable
	var square = func(n int) int {
		return n * n
	}

	fmt.Println(square(4)) // Output: 16
}
```
### Passing Functions as Arguments
In Go, functions are **first-class citizens**, meaning they can be passed as arguments just like variables. 
To make a function accept another function as a parameter, we specify the parameter type using the `func` signature.
```go
package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello,", name)
}

// 'f' is a parameter of type 'func(string)'
// It accepts any function that takes one string and returns nothing
func runCallback(f func(string), value string) {
	f(value)
}

func main() {
	// Pass a named function
	runCallback(greet, "Go") // Output: Hello, Go

	// Pass an anonymous function directly
	runCallback(func(name string) {
		fmt.Println("Hi,", name)
	}, "Ali") // Output: Hi, Ali
}
```
Here we created two functions. The first one takes a string as an argument and prints `"Hello"` followed by the name.   
The second function takes two parameters:
- The first parameter is a function with the signature `func(string)`, which means it must be a function that accepts a string argument and returns nothing.
- The second parameter is a normal string value.
    

Inside this second function, we simply **call** the function passed as the first parameter, and pass the string value to it. This allows us to execute different functions depending on what we pass in.
### Returning Functions
What’s amazing about Go is that we can return a function from another function. This is useful for creating **“function factories”** functions that generate other functions, often with customized behavior.  
To return a function, we simply set the return type to a **`func` signature**, specifying the parameters and return type of the function we want to return.
```go
package main

import "fmt"

func add(a, b int) int      { return a + b }
func multiply(a, b int) int { return a * b }

// This function *returns* a function of type 'func(int, int) int'
func getOperation(opType string) func(int, int) int {
	if opType == "add" {
		return add
	} else if opType == "multiply" {
		return multiply
	}
	// Return nil or a default function if the type is unknown
	return nil
}

func main() {
	operation := getOperation("multiply")
	if operation != nil {
		fmt.Println(operation(3, 4)) // Output: 12
	}
}
```
Here we created three functions.  
The first function, `add`, takes two integers and returns their sum.  
The second function, `multiply`, takes two integers and returns their product.   
The third function, `getOperation`, takes a string called `opType` as a parameter and **returns a function** with the signature `func(int, int) int`.
- This means `getOperation` will give us back a function that takes two integers and returns an integer.
- Inside `getOperation`, we check the value of `opType`:
    - If it is `"add"`, we return the `add` function.
    - If it is `"multiply"`, we return the `multiply` function.
    - If the type is unknown, we return `nil` (or we could return a default function).

In `main`, we call `getOperation("multiply")` and store the returned function in the variable `operation`.  
Then we **call** `operation(3, 4)` just like a normal function, which executes the `multiply` function and prints `12`.  
This allows us to **choose which function to execute at runtime** based on a string or some condition, making our code flexible and reusable.
### Functional Programming Concepts
While Go is not a purely functional language, it embraces several key functional concepts:
- **Pure functions:** You can (and should) write functions that always produce the same output for the same arguments and have no side effects. This makes code easier to test and reason about.
- **Recursive functions:** Recursion is fully supported.
- **First-Class and Higher-Order functions:** As shown, functions are first-class citizens. Functions that take other functions as arguments or return them (higher-order functions) are very common in Go. A classic example is the `http.HandleFunc` function from the standard library, which takes a function to handle an HTTP request.

Go does **not** enforce immutability. Variables are mutable by default. While `const` exists, it's for compile-time constants, not for creating immutable variables at runtime.
## Packages and Modules
### Introduction to Packages
A **package** is Go's fundamental unit for organizing code. We can think of it as a directory containing Go files that all work together to perform a related set of tasks. Every Go file must declare which package it belongs to using the `package` keyword at the very top of the file (e.g., `package main`).   
Packages serve two primary purposes:
1. **Organization**: They keep related code (functions, types, variables) grouped together, making large projects manageable.
2. **Encapsulation**: They control which parts of our code are "public" and which are "private." This is directly related to the **Variable Scope** we discussed earlier.
    - **Exported** (Public): If an identifier (like a function, type, or variable) starts with a **capital letter**, it is exported. This means any other package that imports it can access it.
    - **Unexported** (Private): If it starts with a **lowercase letter**, it is unexported and can only be accessed by other code within the same package.
### Using Standard Library Packages
Go comes with a powerful **Standard Library**, which is a collection of built-in packages that provide core functionality. We don't need to install them; we just need to tell Go we want to use them with the `import` keyword.
For example, we've been using `fmt` to print to the console. To use it, we import it. If we want to generate random numbers, we can import the `math/rand` package.
```go
package main

import "fmt" // Import a single package
import "math/rand"

func main() {
	fmt.Println("Hello, Go!")
	fmt.Println("My favorite number is", rand.Intn(10))
}
```
When importing multiple packages, it's standard Go style to group them in a block using parentheses. This is the preferred way.
```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// We need to "seed" the random generator to get different numbers
	// or it will produce the same number every time.
	rand.Seed(time.Now().UnixNano())
	
	fmt.Println("My favorite number is", rand.Intn(10))
}
```
When we use a function from an imported package, we **must** prefix it with the package name, like `fmt.Println()` or `rand.Intn()`.
### Introduction to Modules
While packages organize code, modules organize packages. A Go module is a collection of related Go packages that are versioned together as a single unit. Modules are how Go manages dependencies the other packages and libraries our project needs to run.   
A module is defined by a `go.mod` file, which lives in the root directory of our project. This file tracks our project's name (its "module path") and all its dependencies and their specific versions.
### Creating Our Own Module
To start any new project in Go, Our first step should be to create a module. We do this from our terminal using the `go mod init` command. This command creates the `go.mod` file for us.  
We just need to give our module a name, called a **module path**. This path is often a URL where our code will be hosted (like `github.com/your-username/myproject`), but for local projects, a simple name is fine.
```shell
# First, create a new directory for our project
mkdir myawesomeproject
cd myawesomeproject

# Now, initialize the module
go mod init myawesomeproject
```
This command will create a `go.mod` file that looks very simple:
```go
module myawesomeproject

go 1.21 // Or whatever your current version of Go is
```
Our project is now officially a Go module. The name `myawesomeproject` is the root path for all packages we create inside it.
### Creating and Using Our Own Packages
Now that we have a module, we can easily create and use our own packages. Let's create a `calculator` package inside our `myawesomeproject` module.   
We start by creating a new directory inside our project named `calculator`. 
Then innside that `calculator` directory, we create a new file `add.go` That will store our code.  
Our project structure should now look like this:
```shell
myawesomeproject/
├── go.mod
├── calculator/
│   └── add.go
└── main.go       (We'll create this one next)
```
After this we add code to ``calculator/add.go``. in the first name the package declaration should be same as the directory name and the capitalized ``Add`` function name, which makes it exported.
```go
// File: calculator/add.go
package calculator


func Add(a, b int) int {
	return a + b
}


func subtract(a, b int) int {
	return a - b
}
```
In Go, if we want a function to be exported and accessible from other packages, we start its name with a capital letter. If we want the function to be internal and unexported, we start its name with a lowercase letter.  

Now, we can use our new package from `main.go`. We import it using its full module path: `myawesomeproject/calculator`.
```go
// File: main.go
package main

import (
	"fmt"
	"myawesomeproject/calculator" // Import our local package
)

func main() {
	sum := calculator.Add(10, 5)
	fmt.Println("The sum is:", sum) // Output: The sum is: 15

	// This line would cause a compile error because 'subtract' is unexported:
	// diff := calculator.subtract(10, 5)
}
```
When we run `go run main.go`, Go sees the import, looks for the `calculator` directory inside the `myawesomeproject` module, and uses the exported `Add` function.
### Package Aliasing
In Go, we can use an **alias** when importing a package to avoid naming conflicts or to use a shorter name. we do that by specifying the alias directly before the package path.
```go
package main

import (
	"fmt"
	mathutils "math" // Alias 'math' to 'mathutils'
	str "strings"    // Alias 'strings' to 'str'
)

func main() {
	fmt.Println(mathutils.Sqrt(25))  // Use the alias
	fmt.Println(str.ToUpper("hello")) // Use the alias
}
```
### Managing Third-Party Modules
One of the biggest advantages of Go modules is how easily they let us use external libraries (third-party modules) written by other developers.  
For example, let's say we want to use a popular package to generate UUIDs:  `github.com/google/uuid`  
To use this module we don’t need to manually install it, we just import it directly in our `main.go` file:
```go
package main

import (
	"fmt"
	"myawesomeproject/calculator"
	"github.com/google/uuid" // Third-party package import
)

func main() {
	sum := calculator.Add(10, 5)
	fmt.Println("The sum is:", sum)

	newID := uuid.New()
	fmt.Println("Generated UUID:", newID.String())
}
```
Our code now references a package, but our module doesn’t have it yet. to fix this we use go special command
```shell
go mod tidy
```
When we run `go mod tidy`, Go automatically scans our project to detect all the packages we are importing in our code. If any required dependencies are missing, such as `github.com/google/uuid`, the command downloads them and ensures they are available for our module. In addition to installing missing modules, `go mod tidy` also updates your `go.mod` file by adding the exact versions of each dependency, keeping our module configuration accurate and reproducible. Finally, it removes any unused dependencies to keep our project clean and avoid unnecessary bloat
## Tasks
### Task 1
Create a Go function that tests whether a given number is a prime number or not.
### Task 2
Write a Go function that converts a decimal number to its binary representation **using recursion**, it take ``int`` as argument and return ``int``.
### Task 3
Create a **Go module** and inside it, create a **package** that contains the following functions:
- A function that returns the largest element in a slice of integers.
- A function that returns the smallest element in a slice of integers.
- A function that calculates and returns the average of all elements in the slice (as a `float64`).

Then, create a `main` program file that prompts the user to enter elements, stores them in a slice, and uses the previously defined functions to find and display the largest element, smallest element, and the average.
