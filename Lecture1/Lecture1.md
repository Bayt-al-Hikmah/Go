## Objectives :
- Learn About How Computers Work
- Introduction to Programming Languages
- Introduction to Go
- User Input and Output
## How Computers Work :
We all have computers, and we use them for a variety of purposes, including watching videos, playing games, performing mathematical calculations, communicating with friends, and many other applications. But the fundamental question is: how do these devices actually work?

The answer lies in their electrical nature. Computers are essentially electrical devices that perform all calculations using electrical signals. The central processing unit (CPU), often referred to as the "brain" of the computer, executes these calculations. To temporarily store data while the computer is in operation, it relies on memory, specifically Random Access Memory (RAM).

Crucially, all information within a computer is represented by electrical signals. This includes data stored in memory and the instructions that the CPU executes. These electrical signals can have only two distinct states: presence, typically represented by the digit "1," and absence, represented by "0."

With this binary representation in mind, we can understand that CPU instructions are essentially sequences of 1s and 0s. This sequence of binary digits is known as machine code, which is the most fundamental level of programming language that the CPU can directly understand.

## Introduction to Programming Languages

Programming languages are tools that were developed to facilitate communication with computers. Instead of writing instructions directly in binary code, which can become incredibly cumbersome for large programs, we can use programming languages with their more user-friendly syntax. This simplifies the coding process, making it easier to read, understand, and debug code.

However, computers cannot directly understand the syntax of these high-level programming languages. To bridge this gap, we use a program called a **compiler**. The compiler first checks the code for any errors and then translates it into machine code  a low-level language consisting of binary instructions that the computer can execute.
## Go programming:
### Introduction

Go (often called Golang) is a modern, **statically typed**, **compiled** programming language developed at Google. It is designed to be simple, efficient, reliable, and easy to learn. Go is widely used for building high-performance backend systems, cloud-native tools (like Docker and Kubernetes), command-line applications, and microservices.

Go's philosophy emphasizes simplicity and readability. It provides powerful features, especially for **concurrency** (running multiple tasks at once), directly within the language.

Go source code is written in files that end with the `.go` extension. All Go code is organized into **packages**. A package is a directory of Go files that work together. An executable program must have a `package main` and a special `main` function.

Go is a compiled language, which means the `go` tool compiles our human-readable source code directly into a single executable machine code file. This file can then be run on any computer with the same operating system and architecture, without needing any other dependencies or virtual machines.
### Go Toolchain
The Go distribution, which we can download and install, comes with collection of tools and libraries  that allow developers to write, format, test, and build Go applications. It includes:
- go command-line tool
- Standard library
- Garbage collector
- Built-in formatter (gofmt)
- Testing framework
- Package manager (go mod)

The toolchain enables us to write clean, consistent code, manage dependencies, run tests, and produce optimized binaries.
Part of the toolchain is the go command-line simply called `go`. This single tool handles all the essential tasks for Go development:
- `go run`: Compiles and runs our program in one step (great for development).
- `go build`: Compiles our program into an executable file.
- `go fmt`: Automatically formats our code to match Go's standard style.
- `go mod`: Manages our project's dependencies (other packages our code uses).
- `go test`: Runs tests for our code.  

To get started with Go, we need to install the Go toolchain on our system. The official guide for installation is available on the Go website at: [https://go.dev/doc/install](https://go.dev/doc/install)  
Once installed, we can verify the setup by running `go version` in the terminal or command prompt.    
### Text editor Vs Ide:
**A text editor** represents software where we can write and stylize our text. We can create our Go programs using any simple text editor. To run them, we use the Go tools (part of the Go installation) to compile and run the program (e.g., using the `go run` command) or build an executable (using `go build`).  
**An IDE (Integrated Development Environment)** is a text editor with many powerful tools integrated within it, such as a compiler, debugger, and code-completion features. This means that with an IDE, we simply write our programs and click "Run," and the IDE will compile and execute the program for us. Popular choices for Go include **VS Code** (with the Go extension) and **GoLand**.
#### Note :
Go program files have the extension **`.go`**. If you're using a text editor to write your Go code, remember to save the file with the `.go` extension. This tells the Go tools that it's a Go source code file.
### Running our First Program
Once we have the Go tools installed, we are ready to run our first Go program! Go programs typically begin execution in a package named `main` with a special function called `func main()`. This function serves as the **entry point** of every executable Go application.    
Go programs have a very clear and simple structure:   
- **Package Declaration:** Every Go file starts with a `package` declaration. For an executable program, this is always `package main`.
- **Import Section:** This section uses the `import` keyword to include other packages (libraries) that provide functions our program needs. For example, the `fmt` package provides functions for formatting and printing.
- **Main Function:** The `func main()` function is the entry point of every executable Go program. It contains the instructions that the computer will execute when the program is run.
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello Reader")
}
```
The first line, `package main`, declares that this file belongs to the **main package**. In Go, a package is a way to organize and reuse code, but when a program uses the `main` package, it tells the compiler that this code should be compiled into an **executable program**, rather than a reusable library.  
Next, the line `import "fmt"` brings in the **`fmt` package** from Go’s standard library. This package provides functions for formatting and printing text, among other things. By importing it, we gain access to useful functions such as `Println`, `Printf`, and `Sprintf`, which are commonly used for displaying output.  
Then we have `func main()`, which defines the **main function** the starting point of every Go program. When we run our Go application, execution always begins with this function.  
Finally, the line `fmt.Println("Hello Reader")` calls the **`Println` function** from the `fmt` package. This function prints the text `"Hello Reader"` to the console, followed by a newline character. Essentially, this is Go’s version of the classic “Hello, World!” program, and it demonstrates how to print output using a standard library function.  

To run this program, save it in a file with a `.go` extension, such as `hello.go`. Then, we open the terminal, navigate to the directory where the file is located, and run the following command:
```shell
go run hello.go
```
This command compiles and runs the file, which is ideal for development and testing, but if we want to compile our Go program into a native executable, we can use the `build` command:
```shell
go build hello.go
```
This creates an executable file (e.g., `hello` or `hello.exe`) in the same directory. we can then run this file directly:
```shell
./hello
```
### Variables
Variables are fundamental building blocks in Go programs; they act like labeled containers that store data. These containers can hold different kinds of information, such as numbers, text (strings), and true/false values (booleans).    
Go is a **statically typed** language. This means that once a variable is declared to hold a certain type of data (like a number), it can _only_ hold that type of data. This helps prevent many common bugs.
#### DataTypes
Before we start working with variables, it’s important to understand the different types of information that Go allows us to store inside them. Go provides a rich set of **data types** that help us represent numbers, text, truth values, collections, and more. Each data type serves a specific purpose and determines how the data is stored and manipulated in memory, We can divide datatypes into two categories
#### Basic Data Types
**Basic Types** (or "primitive types") are the simplest building blocks. They represent a **single value**, like one number (`int`, `float64`), one logical state (`bool`), or one sequence of text (`string`), they are divided into: 

**Numeric Types:**   
**Integer** values are whole numbers. A standard `int` could be `42`, `-1000`, or `0`. An unsigned integer (`uint`) cannot be negative, so its values might be `500` or `0`. Specific-sized integers like `int8` are limited (e.g., `127` or `-128`), while a `uint8` would be between `0` and `255`.
- **Signed:** `int`, `int8`, `int16`, `int32`, `int64`. The `int` type's size (32 or 64 bits) depends on the system architecture. Examples: `5`, `-10`, `0`.
- **Unsigned:** `uint`, `uint8`, `uint16`, `uint32`, `uint64`. These can only store non-negative values.

**Floating-point** numbers represent decimals. A `float64` value, which is the default for decimals in Go, could be `3.14159`, `-0.01`, `123.0`, or `6.022e23`. we also have `float32` it take less storage and have less precision.   

**Complex numbers** store both real and imaginary parts. A `complex128` value is written with an `i` for the imaginary component, such as `(3 + 5i)` or `(-2.2 - 7.1i)`, we also have ``complex64`` it take less storage in memory but hold smaller values.  

**Boolean Type:**   

The **`bool`** type is straightforward; it only has two possible values. A boolean variable can be set to either **`true`** or **`false`**. These are typically used to check conditions, such as `isLoggedIn = true` or `isValid = false`.    

**Text & Character Types:**    

A **`string`** value is a sequence of characters enclosed in double quotes, such as `"Hello, World!"` or `"user@example.com"`. We can also use backticks to create raw strings where newlines and escape characters are preserved literally, like `` `This is a \n raw string.` ``.   

A **`rune`** represents a single character or Unicode code point and is written in single quotes. Examples include `'a'`, `'?'`, or a Unicode character like `'€'` or `'中'`.   

A **`byte`** is an alias for `uint8` and represents a single byte value. It can be represented as a number (e.g., `65`) or as a character literal (e.g., `'A'`, which corresponds to ASCII value 65).   

#### Composite Data Types
**Composite Types** are more complex structures that **group multiple values** together. They are "composed" of other types (which can be basic or even other composite types). Examples include arrays, slices, maps, and structs, which all hold collections or combinations of data.  

An **array** is a fixed-size collection. When we define one, we must specify its length and type. An example value for an array of three integers (`[3]int`) would be `{10, 20, 30}`. An array of two strings (`[2]string`) might look like `{"apple", "banana"}`.   

A **slice** is a flexible, dynamic view of an array's elements. Its literal notation looks similar to an array but without the size specified. An integer slice value could be `[]int{2, 4, 6, 8}`, and a string slice could be `[]string{"one", "two", "three"}`.   

A **map** stores key-value pairs. An example value for a map storing string keys and integer values (`map[string]int`) could be `map[string]int{"age": 30, "id": 12345}`. Another example, mapping product IDs to prices (`map[int]float64`), might be `map[int]float64{101: 49.99, 102: 8.50}`.  

A **struct** is a custom data type that groups related fields. If we defined a `Person` struct (e.g., `type Person struct { Name string; Age int }`), an example value (or "instance") of that struct would be `Person{Name: "Alice", Age: 28}`. We can also omit the field names if we provide them in order, like `Person{"Bob", 42}`.
### Creating Variables
Go provides us with several ways to create variables. Choosing the right method often depends on the context, such as whether we are inside a function or at the package level, and whether we want to initialize the variable immediately.
#### Using The `var` Keyword
This is the most formal way to declare a variable. It can be used both inside and outside of functions. We start with the `var` keyword, followed by the variable name, and then the type. We can also provide an initial value using the equals sign (`=`).  
If we don't provide an initial value, Go assigns a **"zero value"** (like `0` for numbers, `false` for booleans, or `""` for strings).
```go
var n int = 4 
var age int
```
If we provide an initial value, Go can often infer the type, allowing us to omit it. This is called **type inference**.
```go
var m = 5
var name = "Alice"
```
#### Using the Short Declaration Operator (`:=`)
This is the most common and concise way to **declare and initialize** a variable at the same time. This operator can **only** be used inside functions.  
The `:=` operator automatically infers the variable's type from the value we provide.
```go
// Go infers 'count' is an int and 'message' is a string
count := 10
message := "Hello, Go!"
```
#### Declaring Multiple Variables
Go provides a few ways to declare multiple variables in a single statement, which helps reduce clutter. We can declare multiple variables of the **same type** in one line.
```go
// Declares three integers, all set to their zero value (0)
var x, y, z int
```
We can also declare and initialize multiple variables at once. Go will infer the types for each. This works with both `var` and `:=`.
```go
// Using 'var' (can be inside or outside functions)
var name, age, active = "Bob", 42, true

// Using ':=' (must be inside a function)
host, port, err := "localhost", 8080, false
```
Finally, we can use a `var` **block** to group related variable declarations. This is very common for package-level variables.
```go
var (
    firstName = "Jane"
    lastName  = "Doe"
    userID    = 101
    )
```
### Constants:
A constant is a fixed value that cannot be changed while your program is running. An example of a constant is **Pi**, which has a fixed known value. We declare constants using the `const` keyword.  
When creating constants, it is a good practice in Go to use `PascalCase` (e.g., `Pi`) for constants that should be visible outside the current package and `camelCase` (e.g., `gravity`) for constants used only within the current package.
```go
package main

import "fmt"

const PI = 3.14159

func main() {
    const Gravity = 9.81
    fmt.Println(PI, Gravity)
}
```
### Arithmetic Operators
Go provides basic arithmetic operators for manipulating numbers. These include **Addition (`+`)**, **Subtraction (`-`)**, and **Multiplication (`*`)**, which work as we would expect.  
Two operators require special attention: **Division (`/`)** and **Modulo (`%`)**.
- When using **Division (`/`)**, the behavior depends on the type. For `float64` types, it performs a standard decimal division (e.g., `5.0 / 2.0` is `2.5`). For `int` types, it performs **integer division**, which discards any fractional part (e.g., `5 / 2` is `2`).
- The **Modulo (`%`)** operator returns the remainder of an integer division (e.g., `5 % 2` is `1`).

Finally, Go has **Increment (`++`)** and **Decrement (`--`)** operators, which add or subtract one from a variable.
```go
package main

import "fmt"

func main() {
    a := 10
    b := 3
    // Basic arithmetic
    c := a + b // c = 13
    d := a - b // d = 7
    e := a * b // e = 30
    // Integer division (10 / 3 = 3.33..., truncates to 3)
    f := a / b 
    // Modulo (10 divided by 3 is 3 with a remainder of 1)
    g := a % b 
    
    fmt.Println("c (add) =", c)
    fmt.Println("d (sub) =", d)
    fmt.Println("e (mul) =", e)
    fmt.Println("f (div) =", f)
    fmt.Println("g (mod) =", g)
    // Increment statement
    a++ // a is now 11
    fmt.Println("new a =", a)
}
```

### Escape Sequences
**Escape sequences** are special character combinations that let us represent hard-to-type or invisible characters within a string or rune literal. They always start with a backslash (`\`), which signals to Go that the next character has a special meaning.    
Common examples include:
- `\n` (Newline): Moves the cursor to the next line.
- `\t` (Tab): Inserts a tab character.
- `\"` (Double quote): Inserts a literal double quote inside a double-quoted string.
- `\'` (Single quote): Inserts a literal single quote inside a rune.    
- `\\` (Backslash): Inserts a literal backslash.

```go
package main

import "fmt"

func main() {
    // \n creates a new line
    fmt.Println("Hello\nWorld")
    // \t adds a tab
    fmt.Println("Before tab\tAfter tab")
    // \" allows you to include quotes in your string
    fmt.Println("He said, \"This is a quote.\"")
}
```

### Compound Assignment Operators
Go provides **compound assignment operators** (also called shorthand operators) that combine an arithmetic operation with an assignment (`=`). This provides a more concise way to modify a variable in place.
For example, `a += 4` is the shorthand equivalent of `a = a + 4`.   
This pattern works for all the main arithmetic operators:
- `+=` (Add and assign)
- `-=` (Subtract and assign)
- `*=` (Multiply and assign)
- `/=` (Divide and assign)
- `%=` (Modulo and assign)

```go
package main

import "fmt"

func main() {
    a := 10
    fmt.Println("Start:", a) // 10
    a += 5 // a = a + 5
    fmt.Println("+= 5:", a) // 15
    a -= 3 // a = a - 3
    fmt.Println("-= 3:", a) // 12
    a *= 2 // a = a * 2
    fmt.Println("*= 2:", a) // 24
    a /= 6 // a = a / 6
    fmt.Println("/= 6:", a) // 4
    a %= 3 // a = a % 3
    fmt.Println("%= 3:", a) // 1
}
```
### Comments:
Comments are lines of code that the compiler will ignore. They are used to add explanatory notes and improve the readability of our code. There are two ways to create comments in Go:    
**Single-line comments:** Begin with `//` and continue until the end of the current line. All text following `//` on the same line is ignored by the compiler.
```go
// this single line comment
```
**Multi-line comments:** Begin with `/*` and end with `*/`. Any text between these two symbols, including multiple lines, will be ignored by the compiler.
```go
/*
	This
	multi-line
	comment
*/
```
### Type Conversion
Go is a **statically-typed language**, which means it is very strict about types. It will not automatically convert between different types (a process called implicit conversion). For example, we cannot assign an `int` to a `float64` directly.  
We must perform an **explicit conversion** using the syntax `T(v)`, where `T` is the target type and `v` is the value you want to convert.
#### Numeric Conversion
Converting between numeric types (like `int`, `float64`, `int32`, etc.) is straightforward. We simply use the target type as if it were a function.  
Be aware that converting from a type with more precision to one with less (like `float64` to `int`) can result in a **loss of data**, as the decimal part will be truncated (not rounded).
```go
package main

import "fmt"

func main() {
	// --- Numeric Conversion ---
	i := 42
	f := float64(i) // Convert int to float64
	
	fmt.Println("Integer:", i) // 42
	fmt.Println("Float:", f)   // 42.0

	pi := 3.14159
	radius := int(pi) // Convert float64 to int (loses decimal info)
	
	fmt.Println("Float PI:", pi)       // 3.14159
	fmt.Println("Integer Radius:", radius) // 3 (decimal is truncated)
}
```
#### String Conversion
We cannot directly convert a `string` to a number (or vice-versa) using the `T(v)` syntax. Instead, we must use the built-in **`strconv`** (string conversion) package.  
**Number to String:**
To convert a number to a string, we "format" it. The `strconv` package provides simple functions for this:
- `strconv.Itoa(i int) string`: The most common. `Itoa` stands for **I**nteger **to** **A**SCII (a way to say "to text").
- `strconv.FormatFloat(...)`: A more general function for formatting `float64` values.

Here's a code example demonstrating how to convert numbers to strings.
```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 1. Using strconv.Itoa() for an integer
	age := 30
	ageString := strconv.Itoa(age)
	
	// 'ageString' is now the string "30"
	fmt.Println("The string value is:", ageString)

	fmt.Println("---")
	// 2. Using strconv.FormatFloat() for a float
	price := 49.95
	
	// FormatFloat(value, format, precision, bitSize)
	// 'f' = standard decimal format
	// 2   = two decimal places
	// 64  = for a float64
	priceString := strconv.FormatFloat(price, 'f', 2, 64)

	// 'priceString' is now the string "49.95"
	fmt.Println("The string value is:", priceString)
}
```
In this example, we import the `strconv` package, which is Go's standard library for string conversions. We use the `strconv.Itoa` function (Integer to ASCII) to convert the `age` variable into a string. For the `price` variable, we use `strconv.FormatFloat`, which gives us more control. We tell it to format the number as a standard decimal (`'f'`), keep two decimal places (`2`), and that the original number is a `float64` (`64`).
#### String to Number:
Converting a string to a number is called "parsing." This operation can easily fail (for example, if the string is `"hello"` instead of `"123"`). Because of this, parsing functions in Go always return **two** values: the resulting number and an `error`. You must always check this error variable to see if the conversion was successful.
- `strconv.Atoi(s string) (int, error)`: The most common. `Atoi` stands for **A**SCII **to** **I**nteger.
- `strconv.ParseFloat(s string, bitSize int) (float64, error)`: A general function for parsing floats (e.g., `bitSize = 64` for a `float64`).

Here's a code example demonstrating how to convert strings to numbers.
```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 1. Using strconv.Atoi() for an integer
	ageStr := "42"
	age, err := strconv.Atoi(ageStr)
	// The "Go way" is to always check the error!
	if err != nil {
		fmt.Println("Error converting 'ageStr':", err)
	} else {
		fmt.Println("The integer value is:", age)
	}
	fmt.Println("---")
	// 2. Using strconv.ParseFloat() for a float
	priceStr := "199.99"
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		fmt.Println("Error converting 'priceStr':", err)
	} else {
		fmt.Println("The float value is:", price)
	}
	fmt.Println("---")
	// 3. Example of a failed conversion
	badStr := "not-a-number"
	num, err := strconv.Atoi(badStr)
	
	if err != nil {
		fmt.Println("Error converting 'badStr':", err)
	} else {
		fmt.Println("The number is:", num)
	}
}
```
In this example, we use the `(value, err)` pattern, which is standard in Go for operations that can fail. For both `strconv.Atoi` and `strconv.ParseFloat`, we get two return values. We immediately check `if err != nil`. If `err` is not `nil`, it means the conversion failed (as seen in the `badStr` example), and we should print the error instead of using the returned value. If `err` is `nil`, we know the conversion was successful and we can safely use the `age` or `price` variables.

## User Input and Output
Our programs aren't complete without interacting with users. So far, we've seen how to create variables, but all our data has been "hard-coded" (written directly into the program). Go offers us ways to get information and data from the user and also ways to display information back _to_ the user's screen.
### Displaying Output to the User
The most common way to display output in Go is by using functions from the **`fmt`** (format) package.
- **`fmt.Println()`**: This is the most basic function. It prints its arguments followed by a **new line**.
- **`fmt.Print()`**: This function prints its arguments without adding a new line at the end.

A more powerful function is **`fmt.Printf()`**, which stands for "print formatted." This function lets us embed values inside a string using special **format specifiers** (or "verbs"). These specifiers start with a `%` and tell Go how to format the variable.  
Here are the most common format specifiers:
- **`%v`**: The default format (a good general-purpose choice).
- **`%s`**: For a string.
- **`%d`**: For a decimal integer (`int`).
- **`%f`**: For a floating-point number (`float64`).
- **`%t`**: For a boolean (`bool`).
- **`%T`**: To print the **T**ype of the variable.

```go
package main

import "fmt"

func main() {
    name := "Alice"
    age := 30
    price := 19.99

    // Using Println
    fmt.Println("Hello and welcome!")
    // Using Printf to embed values
    fmt.Printf("User: %s\n", name)
    fmt.Printf("Age: %d, Price: %f\n", age, price)
    // Using %v for default formatting
    fmt.Printf("Default format: %v (age) and %v (name)\n", age, name)
    // Using %T to see the types
    fmt.Printf("The type of 'name' is %T and 'price' is %T\n", name, price)
}
```

### Getting Input from the User
To get data from the user, we can use functions from the `fmt` package, such as `fmt.Scan()` and `fmt.Scanln()`.

When we use either of these functions, we must pass it a **pointer** to the variable where we want to store the user's input. A pointer is simply the memory address of the variable. We get a variable's pointer by using the ampersand (`&`) symbol in front of it (e.g., `&myVar`). This tells the function "store the value you read at this memory location."

The main difference between the two functions is how they handle input:

- **`fmt.Scan()`**: This function reads input, separating values by **spaces**. It stops reading after it has filled all the variables you passed to it. It's great for reading multiple space-separated values in one go.
    
- **`fmt.Scanln()`**: This function is similar, but it stops scanning at a **newline**. It expects the user to press Enter after providing the input.
    

Both functions stop reading a string value at the first space. This makes them ideal for single words or numbers, but not for reading a full sentence with spaces.
```go 
package main

import "fmt"

func main() {
    // 1. Declare a variable to store the user's input
    var name string

    // 2. Prompt the user for their name
    fmt.Print("Please enter your name: ")

    // 3. Read the user's input and store it
    // We pass the memory address (pointer) of 'name'
    fmt.Scanln(&name)

    // 4. Use the input
    fmt.Printf("Hello, %s! Welcome to the program.\n", name)

    // --- Example with a number ---
    var age int
    fmt.Print("Please enter your age: ")
    fmt.Scanln(&age)
    fmt.Printf("Next year, you will be %d.\n", age + 1)
}
```
#### Reading a Full Line 
To read a full line of text (including spaces) until the user presses Enter, we need to use the **`bufio`** package. `fmt.Scanln()` and `fmt.Scan()` cannot do this because they stop at the first space.  
We firs create a new "reader" that is buffered (it reads data in efficient chunks) and attach it to the standard input (our keyboard). Then, we tell it to read everything until it finds a newline character (`\n`), which is what the Enter key sends.    
```go
package main

import (
	"bufio"   // Import the buffered I/O package
	"fmt"
	"os"      // Import the OS package to access Stdin
	"strings" // Import the strings package to clean up the input
)

func main() {
	fmt.Print("Please enter a full sentence (e.g., hello world): ")

	reader := bufio.NewReader(os.Stdin)
	

	text, err := reader.ReadString('\n')
	
	// Check for errors!
	if err != nil {
		fmt.Println("Error reading input:", err)
		return 
	}
	
	text = strings.TrimSpace(text)
	fmt.Printf("You entered: \"%s\"\n", text)
}
```

We import `bufio` and `os` to get access to the tools we need.
1. **`bufio.NewReader(os.Stdin)`** creates a new reader that takes its input from `os.Stdin` (the "standard input," which is your keyboard).
2. **`reader.ReadString('\n')`** tells the reader to read every character it sees, including spaces, until it finally sees the `\n` (newline) character.
3. This function returns two values: the `text` it read (including the `\n`) and a possible `err`. We check the error just in case.
4. Because `text` still has the newline character at the end (e.g., `"hello world\n"`), we use **`strings.TrimSpace(text)`** to remove it and any other surrounding whitespace, leaving us with just `"hello world"`.
## Tasks:

### Task 1:
Write a Go program that reads the radius of a circle from the user (as a `float64`) and then displays its surface area. (Area = $\pi \times \text{radius}^2$). You can use the constant `math.Pi` from the `math` package.
### Task 2:
Write a Go program that reads the name of a student (as a string). Then, read three separate grades (as float64 numbers) for that student using a separate variable for each grade. Finally, display the student's name, their three individual grades, and the calculated average of those grades.