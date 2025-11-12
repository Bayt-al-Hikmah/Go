## Objectives

- Command-line arguments and Flags
- Write data to and read data from files
- Concurrency In GO

## Command-Line Arguments and Flags

### Introduction

So far, our programs run the same way every time we execute them. But what if we want to change their behavior from the outside, without recompiling the code? For example, we might want to tell our program _which_ file to open, what URL to connect to, or how many times to run a loop.  
This is where command-line arguments come in. They are values we pass to our program right when we launch it from the terminal, allowing us to configure its behavior at runtime.
### Basic Arguments with `os.Args`
The simplest way to get arguments is by using the `os` package. When your program runs, Go automatically populates a slice of strings called `os.Args` with all the arguments provided.
- `os.Args[0]` is always the path to the program itself.
- `os.Args[1]` is the first argument, `os.Args[2]` is the second, and so on.
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args is a slice of strings
	fmt.Println("Program name:", os.Args[0])

	// Check if any other arguments were provided
	if len(os.Args) > 1 {
		fmt.Println("First argument:", os.Args[1])
		fmt.Println("All arguments:", os.Args[1:])
	} else {
		fmt.Println("No arguments provided.")
	}
}
```
Here we using ``os.Args`` to access to command line arguments, the first value is always the file name then after it we find the arguments that we included in our terminal.  
If we save this as `main.go` and run it:
```bash
# Run with no arguments
$ go run main.go
Program name: /var/folders/..../main
No arguments provided.

# Run with arguments
$ go run main.go hello world 123
Program name: /var/folders/..../main
First argument: hello
All arguments: [hello world 123]
```

### Better Arguments with the `flag` Package
While `os.Args` is simple, it's not very robust. What if we want named options (like `-name="John"`)? What if we want default values or help messages?   
For this, Go provides the built-in **`flag`** package. It's the standard way to handle command-line options.  
The basic pattern is:
1. **Define** your flags (e.g., `flag.String()`, `flag.Int()`, `flag.Bool()`). This tells the package what to look for.
2. **Parse** the arguments by calling `flag.Parse()`. This loads the values from the command line into your flag variables.
```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	// 1. Define the flags
	// flag.String(name, default_value, help_message)
	wordPtr := flag.String("word", "default", "The word to say")
	countPtr := flag.Int("n", 1, "Number of times to say it")
	
	// 2. Parse the command line
	flag.Parse()

	// The flags return pointers, so we must dereference them with *
	fmt.Println("Word:", *wordPtr)
	fmt.Println("Count:", *countPtr)

	for i := 0; i < *countPtr; i++ {
		fmt.Println(*wordPtr)
	}
}
```
We started by defining flags:    
- `flag.String("word", "default", "The word to print")`  
	This creates a string flag `-word`. If the user does not provide it, the value defaults to `"default"`. The last argument is the help message shown when running `go run main.go -help`.
- `flag.Int("n", 1, "Number of times to print the word")`  
	This creates an integer flag `-n`. If the user does not provide it, the default is `1`.
	
After that we Parse the flags:
- `flag.Parse()` reads the command-line arguments and stores them in the pointers.
        
Now we have the flags we can their values on our code: 
- Since flags return pointers, you access their values with `*wordPtr` and `*countPtr`.

Now, we can run this program in much more flexible ways:
```shell
# Run with default values
$ go run main.go
Word: default
Count: 1
default

# Run with custom values
$ go run main.go -word=hello -n=3
Word: hello
Count: 3
hello
hello
hello
```
We can get information about how the program works by running it with the `-help` flag. This command displays all the flags that the program accepts, along with the description messages that we provided when defining each flag.
```Shell
# The flag package also provides a -help message for free
$ go run main.go -help
Usage of /var/folders/..../main:
  -n int
    	Number of times to say it (default 1)
  -word string
    	The word to say (default "default")
```

## File Handling
### Introduction
We worked with many programs on our course but all this time we where storing all the data and the result as variables and this store them in memory. When the program finishes or restarts, we lose all the data we worked with. This is fine for simple calculations, but what if we want our program to save its results? Or what if we need to read a large amount of data from an external source?   
To make our data persistent (meaning it lasts after the program stops), we need to work with files. Go provides powerful and simple tools for file I/O (Input/Output) in its standard library, primarily in the `os` and `io` packages.
### Writing to Files
To write to a file in Go, we use `os.Create("hello.txt")`. This command creates a new file if it doesn't exist. If the file already exists, Go clears its contents before opening it. So `os.Create` always gives us a fresh file to write to.

When we call `os.Create`, Go returns two things: a `*os.File` value and an `error`. The `*os.File` is like a handle to the file, allowing us to read from or write to it. The `error` tells us if anything went wrong. If the error is `nil`, everything worked fine. If not, the file wasn’t created successfully.

After opening a file, we must always close it using `file.Close()`. Closing the file makes sure everything we wrote is saved and frees system resources. Forgetting to close files can lead to bugs and data loss, especially in larger programs.

Go provides a helpful tool called `defer` to ensure cleanup happens. When we write `defer file.Close()`, we tell Go to run `file.Close()` at the end of the function. This means we don’t need to remember to close the file in multiple places Go does it for us automatically.

`defer` works even if the function returns early or if something unexpected happens, like a panic. This makes our programs safer and easier to maintain. If we use multiple `defer` statements, Go runs them in reverse order, like a stack.
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return 
	}
	defer file.Close()

	data := "Hello, persistent data!"
	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File 'hello.txt' created and written successfully.")
}
```
To write text into the file, we use `file.WriteString(data)`. This function takes a string and writes it directly into the open file. If the file was created or opened earlier with `os.Create` or `os.OpenFile`, calling `file.WriteString("Hello")` will store the text `"Hello"` inside it. The function also returns how many bytes were written and an `error` value. If the error is `nil`, the write was successful; otherwise, something went wrong during the write operation.
### Appending to File
Writing to file helped us to store data. But there’s one small problem: every time we open the file with `os.Create`, we overwrite its contents. That means if we run our program again, the old data disappears and only the new data remains. Sometimes this is what we want, but many times we want to add new data to the end of the file instead of replacing everything. This is called appending.

To append to a file in Go, we don’t use `os.Create`. Instead, we use `os.OpenFile` with special flags. The key one here is `os.O_APPEND`, which tells Go to add data to the end of the file. We also include `os.O_WRONLY` to open the file for writing, and `os.O_CREATE` to make sure the file is created if it doesn't already exist.

Here’s an example of opening a file for appending and writing to it:
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("hello.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	data := "New line of data!\n"
	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Data appended successfully.")
}
```
In this code:
- `os.O_APPEND` makes sure new data is added to the end of the file.
- `os.O_WRONLY` opens the file for writing.
- `os.O_CREATE` creates the file if it doesn't already exist.
- `0644` sets the file's permissions (standard read/write settings).

Now, each time we run the program, it adds a new line instead of replacing everything. This is perfect for logging, saving history, or building files step by step.

### Reading from Files
To read a file, we use `os.Open()`. This also returns a `*os.File` and an `error`. Just like with writing, we should `defer file.Close()` immediately after opening.

While you _can_ read a file byte by byte, this is often complicated. A much more common and easier way to read a text file is one line at a time. For this, Go provides a powerful tool called the `bufio.Scanner`. We create a new scanner and point it at our open file.
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fmt.Println("Contents of 'hello.txt':")

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
```
In this code:
- `scanner := bufio.NewScanner(file)` creates a new scanner that "wraps" our open file.
- `for scanner.Scan()` is a loop. `scanner.Scan()` moves to the next line and returns `true` if it found one. When it reaches the end of the file, it returns `false`, and the loop stops.
- `scanner.Text()` returns the line that was just read as a string.
- `scanner.Err()` is an important final check to see if any errors (other than the end of the file) happened while reading.
###  `os.ReadFile` and `os.WriteFile`
Sometimes we don't need to read a file line by line we just want the entire content at once. Or we want to quickly save a whole block of data. For these simple cases, Go gives us two very convenient functions:
- `os.ReadFile()` reads the entire file in one call
- `os.WriteFile()` writes all data to a file in one call

These functions automatically handle opening and closing the file internally, which makes your code shorter and easier to read.
```go
package main

import (
	"fmt"
	"os"
)

func main() {

	data := []byte("This is a simple way to write.")
	err := os.WriteFile("simple.txt", data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("Wrote to simple.txt")

	readData, err := os.ReadFile("simple.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Read from simple.txt:", string(readData))
}
```
In this code:
- `data := []byte(...)` creates a "slice of bytes." These functions work with `[]byte` data, not plain strings.
- `os.WriteFile("simple.txt", data, 0644)` takes the filename, the byte slice, and the file permissions. It creates or overwrites the file in one simple step.
- `readData, err := os.ReadFile("simple.txt")` takes just the filename and returns the _entire_ file's contents as a new byte slice (`readData`).
- `string(readData)` is the standard Go way to convert the `[]byte` slice back into a readable string for printing.
## Concurrency
### Introduction
Modern computers have multiple CPUs (or cores), which let them perform many tasks at the same time. Most programs we’ve written so far, however, run sequentially one step at a time, from top to bottom. If a single task takes a long time (like downloading a file or performing a complex calculation), the entire program has to wait.

Concurrency allows a program to run multiple tasks at the same time, independently of each other. Many languages achieve this with threads, which are separate paths of execution. Creating and managing threads can be slow and resource-intensive. Go, however, makes concurrency simple and efficient with goroutines and channels, allowing us to run many tasks concurrently without the overhead of traditional threads.
### Goroutines
A goroutine is a lightweight, independent path of execution in Go. It allows a function to run at the same time as other functions, enabling true concurrency. Unlike traditional threads, goroutines are extremely cheap to create we can run thousands or even millions without heavy memory use.

To start a goroutine, simply put the `go` keyword before a function call. Go will run that function concurrently in the background while the rest of your program continues.
```go
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// Start a goroutine
	go say("world")
	
	// Run this in the main "thread"
	say("hello")
}
```
**Output:**
```
hello
world
hello
world
hello
world
```
Notice how "hello" and "world" are interleaved? The `main` function and the `say("world")` goroutine are running at the same time!
### Using `sync.WaitGroup`
When you launch a goroutine, the `main` function (which runs in its own goroutine) doesn't automatically wait for it to finish.
```go
func main() {
	go say("I am a goroutine")
	
	fmt.Println("Main function is done.")
}
```
In this code, the program will likely exit immediately. The "I am a goroutine" message may never print because the `main` goroutine finishes and terminates the program before the new goroutine gets a chance to run. To fix this, we need a way to make main wait for other goroutines to complete.

A `sync.WaitGroup` is a simple and idiomatic way to wait for a collection of goroutines to finish.
It's essentially a concurrent counter. We use it to tell `main` how many "jobs" it needs to wait for.

It has three methods:
- **`Add(n)`:** Increments the counter by `n`. We call this before starting the goroutine(s) to signal that `n` tasks are about to begin.
- **`Done()`:** Decrements the counter by one. This is called from within the goroutine to signal that it has finished its work.
- **`Wait()`:** Blocks the current goroutine (e.g., `main`) and waits until the counter becomes zero.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	fmt.Println("Main: Waiting for workers...")
	wg.Wait()
	fmt.Println("Main: All workers done.")
}
```
Here, we used a **`sync.WaitGroup`** (named `wg`) to coordinate our goroutines.   
In the `main` function, before we start each `worker` goroutine, we call **`wg.Add(1)`** to tell the `WaitGroup` "one more task is starting."   
After launching all the goroutines, `main` calls **`wg.Wait()`**, which simply pauses `main` and waits until the counter goes back down to zero.   
Inside each `worker` function, the **`defer wg.Done()`** line is crucial; it ensures that the counter is decreased by one right before the function finishes.   

This pattern guarantees that the `main` function will only continue and print "All workers done." after all three `worker` goroutines have completed.
