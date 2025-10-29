## Objectives
- Working with Strings, Arrays, Slices, and Maps
- Comparison and Logical Operators
- Conditional Statements
- Working with Loops
## Working with Go's Data Structures
### Strings
In Go, we use strings to store text values. Strings are created by wrapping text in double quotes (`"`). We can also use backticks (`` ` ``) to create **raw string literals**, which can span multiple lines and ignore escape characters.   
Strings in Go are **immutable** and **UTF-8 encoded**.
```go
// Short declaration operator, the most common way to declare
name := "Alice"
// Raw string literal
greeting := `Hello,
world!`
```
#### Accessing Bytes and Runes in a String
A string in Go is a read-only slice of bytes. Indexing a string accesses its individual **bytes**, not characters.
```go
word := "Go"
fmt.Println(word[0]) // Output: 71 (the ASCII/UTF-8 value for 'G')
fmt.Println(word[1]) // Output: 111 (the ASCII/UTF-8 value for 'o')
```
A "character" in Go is called a **rune** (which is an alias for `int32`). To access the individual runes, we  use a `for...range` loop:
```go
for index, runeValue := range "Go" {
    fmt.Printf("At index %d, rune is %c\n", index, runeValue)
}
// Output:
// At index 0, rune is G
// At index 1, rune is o
```
Trying to access an index that’s out of bounds will cause a **panic** (which stops the program).
#### Accessing Substrings (Slicing)
We get substrings using Go's **slice** syntax `[start:end]`. This is a fundamental concept in Go.
```go
word := "Programming"
fmt.Println(word[0:5]) // Output: Progr (from index 0 up to, but not including, 5)
fmt.Println(word[7:])  // Output: ming (from index 7 to the end)
fmt.Println(word[:5])  // Output: Progr (from the start up to index 5)
```
We select the starting index and the end index, the character at the end index isn't included.
#### Common String Functions (from the `strings` package)
Go provides most string manipulation functions in the standard library `strings` package.   
##### `len()`
Returns the number of **bytes** in the string, not the number of runes (characters).
```go
name := "Alice"
fmt.Println(len(name)) // Output: 5

// A string with multi-byte characters
name_jp := "日本語"
fmt.Println(len(name_jp)) // Output: 9 (3 runes, 3 bytes each)
```
Not all stringcomposed of character of single bytes some symbole take more then one byte, to get the count of runes, we  use the `utf8` package:
```go
import "unicode/utf8"
fmt.Println(utf8.RuneCountInString(name_jp)) // Output: 3
```
##### `strings.ToUpper()` and `strings.ToLower()`
Convert a string to all uppercase or lowercase.
```go
import "strings"
fmt.Println(strings.ToUpper("hello")) // Output: HELLO
fmt.Println(strings.ToLower("WORLD")) // Output: world
```
##### `strings.TrimSpace()`
Removes extra whitespace from the beginning and end.
```go
input := "   hello   "
fmt.Println(strings.TrimSpace(input)) // Output: hello
```
##### `strings.Contains()`
Checks if a string contains a certain substring.
```go
sentence := "Go is fun"
fmt.Println(strings.Contains(sentence, "fun"))    // Output: true
fmt.Println(strings.Contains(sentence, "boring")) // Output: false
```
##### `+` or `fmt.Sprintf()`
Concatenates strings. The `+` operator is simple. For more complex formatting (Go's version of "interpolation"), use `fmt.Sprintf`.
```go
greeting := "Hello, " + "Alice"
fmt.Println(greeting) // Output: Hello, Alice

name := "Bob"
// %s is the "verb" for a string
greetingFmt := fmt.Sprintf("Hello, %s", name)
fmt.Println(greetingFmt) // Output: Hello, Bob
```
##### `strings.ReplaceAll()`
Replaces parts of a string.
```go
text := "I like Java"
fmt.Println(strings.ReplaceAll(text, "Java", "Go")) // Output: I like Go
```
##### `strings.Split()`
Splits a string into a **slice** of strings, based on a separator.
```go
words := "apple,banana,orange"
fmt.Println(strings.Split(words, ",")) // Output: [apple banana orange]
```

### Arrays
An **array** in Go is a fixed-length sequence of elements of the same type. The length is part of its type.
```go
// An array of 3 strings. The size [3] is fixed.
var colors [3]string
colors[0] = "red"
colors[1] = "green"
colors[2] = "blue"

// Declare and initialize with an array literal
numbers := [5]int{1, 2, 3, 4, 5}
```
Arrays are rigid and not used as often as slices in Go.
### Slices
A **slice** is the workhorse of Go. It's a dynamic, flexible view into the elements of an underlying array. Slices are far more common than arrays.
```go
// A slice literal (looks like an array, but without a size)
colors := []string{"red", "green", "blue"}

// A slice of integers
numbers := []int{10, 20, 30, 40, 50}
```
We can also create slice using special function `make()` it take 3 argument:
- **type** → type of elements (e.g., `[]string`, `[]int`)
- **length** → number of elements the slice currently has
- **capacity** → total space reserved in memory (optional)
```go
// Create a slice of strings with length 5 and capacity 5
s := make([]string, 5)
```
If we try to append a new element, Go will automatically allocate a bigger underlying array and increase the slice’s capacity to fit the new data.
#### Accessing Elements in a Slice
Just like arrays, we access slice elements using index numbers starting from 0:
```go
colors := []string{"red", "green", "blue"}
fmt.Println(colors[0]) // Output: red
```
Accessing an index that doesn’t exist will cause a **panic**.
#### Accessing Sub-Slices (Slicing)
We use the `[start:end]` syntax to get a new slice that references a portion of the original.
```go
numbers := []int{10, 20, 30, 40, 50}
fmt.Println(numbers[1:4]) // Output: [20 30 40] (index 1 up to 4)
fmt.Println(numbers[0:2]) // Output: [10 20]
```
#### Slice Functions and Idioms
##### `len()` and `cap()`
- `len()`: Returns the number of elements currently in the slice.
- `cap()`: Returns the **capacity** of the slice (how many elements it can hold before needing to reallocate memory).
```go
colors := []string{"red", "green", "blue"}
fmt.Println(len(colors)) // Output: 3
fmt.Println(cap(colors)) // Output: 3
```
##### `append()`
This is the **built-in** function to add elements to a slice. `append` returns a new slice, so you must re-assign it.
```go
colors := []string{"red", "green", "blue"}
// Note the re-assignment: colors = ...
colors = append(colors, "yellow")
fmt.Println(colors) // Output: [red green blue yellow]
```
##### Inserting an Element 
When we want to insert an element at a specific position in a slice,  we use a small trick with `append` and slicing. The idea is to split the slice into two parts: everything before the index where we want to insert, and everything after it. Then we use `append` twice  first to add the new value, and then to join the second half of the slice back. 
```go
// Insert "black" at index 2
colors = append(colors[:2], append([]string{"black"}, colors[2:]...)...)
//                 ^-- part before    ^-- new element(s)   ^-- part after
fmt.Println(colors) // Output: [red green black blue yellow]
```
Here we add `"black"` into the `colors` slice at position `2`. we taking the elements before that position (`colors[:2]`), then adding `"black"` as a small slice, and then adding the rest of the original slice (`colors[2:]`).   
The `...` is used so the elements of the second slice are expanded and appended individually. This builds a new slice with `"black"` placed in the middle.
##### Removing the Last Element
We remove last Element by using slicing we slice from start to last element index  `len(colors)-1`.
```go
colors := []string{"red", "green", "blue"}
lastColor := colors[len(colors)-1]
colors = colors[:len(colors)-1] // Slice off the last element
fmt.Println(lastColor) // Output: blue
fmt.Println(colors)    // Output: [red green]
```
##### Removing an Element by Index (Idiom)
We remove an element from the middle by connecting the part before it with the part after it using `append`.
```go
// Remove element at index 1 ("green")
arr := []int{1, 2, 3, 4}
arr = append(arr[:1], arr[2:]...)
fmt.Println(arr) // Output: [1 3 4]
```
##### Sorting Slices 
We sort element of slice by using  the sort package
```go
import "sort"

numbers := []int{3, 1, 5, 2}
sort.Ints(numbers)
fmt.Println(numbers) // Output: [1 2 3 5]

names := []string{"Zoe", "Alice", "Bob"}
sort.Strings(names)
fmt.Println(names) // Output: [Alice Bob Zoe]
```
##### Joining a Slice 
We can join element of slice into string by using ``strings.Join()`
```go
colors := []string{"red", "green", "blue"}
fmt.Println(strings.Join(colors, ", ")) // Output: red, green, blue
```
### Maps
**Maps** are Go's built-in hash table type. They store collections of key-value pairs.
```go
// A map with string keys and int values
// Using a map literal:
person := map[string]int{
    "age":       30,
    "zipcode":   90210,
}

// Using the `make` function:
capitals := make(map[string]string)
capitals["France"] = "Paris"
capitals["Japan"] = "Tokyo"
```
We can also use mixed types by using the  `interface{}` , or the `any` type:
```go
person := map[string]interface{}{
    "name":      "Alice",
    "age":       30,
    "isStudent": false,
}
```
#### Accessing Values
We use the key in square brackets to get a value:
```go
fmt.Println(person["name"]) // Output: Alice
fmt.Println(person["age"])  // Output: 30
```
Accessing a non-existent key returns the **zero value** for the value type (e.g., `0` for `int`, `""` for `string`, `false` for `bool`), and it return `nil`if map the we used ``interface{}`` or `any` as type
```go
fmt.Println(person["height"]) // Output: <nil> (if map is map[string]any)
                             // or 0 (if map is map[string]int)
```
To properly check if value exist , we use the **comma-ok** idiom.
With this pattern, Go returns two values:
- the value stored in the map (or a zero value if not found)
- a boolean (`ok`) that tells us whether the key actually exists

If `ok` is `true`, the key is present. If `false`, the key is not in the map.
```go
value, ok := person["height"]
if ok {
    fmt.Println("Height exists:", value)
} else {
    fmt.Println("Height key does not exist")
}
// Output: Height key does not exist

age, ok := person["age"]
if ok {
    fmt.Println("Age exists:", age)
}
// Output: Age exists: 30
```

#### Adding and Updating Values
We can add key value by just assign if the value key don't exist it will be added to the map, as key value pair if it exist it will edit the old value that was stored in the key:
```go
person["email"] = "alice@example.com" // Add
person["age"] = 31                      // Update
```
#### Removing Entries
We use the built-in `delete` function:
```go
delete(person, "isStudent")
fmt.Println(person) // Output: map[age:31 email:alice@example.com name:Alice]
```
#### Map Functions 
##### Iterating over Keys and Values
To get keys and values, you must iterate. The `for...range` loop is perfect for this. **Note: Maps in Go are unordered.** The iteration order is random.
```go
person := map[string]any{"name": "Alice", "age": 31}
for key, value := range person {
    fmt.Printf("Key: %s, Value: %v\n", key, value)
}

// To get *only* keys:
for key := range person {
    fmt.Println("Key:", key)
}
```
##### `len()`
Returns the number of key-value pairs.
```go
fmt.Println(len(person)) // Output: 2
```
##### Checking for a Key
Use the comma-ok idiom.
```go
_, ok := person["name"]
fmt.Println(ok) // Output: true

_, ok = person["salary"]
fmt.Println(ok) // Output: false
```
## Comparison and Logical Operators
### Comparison
These operators compare values and always return a boolean (`true` or `false`). 
- `>` (Greater than)
- `<` (Less than)
- `==` (Equal to)
- `!=` (Not equal to)
- `>=` (Greater than or equal to)
- `<=` (Less than or equal to)

```go
fmt.Println(5 > 4)    // => true
fmt.Println(4 == 4)   // => true
fmt.Println("hi" != "hello") // => true
```
### Logical Operators
Used to combine multiple boolean conditions.
#### `||` (OR)
Returns `true` if **at least one** condition is true.
```go
fmt.Println(true || false)  // => true
fmt.Println(false || false) // => false
```
#### `&&` (AND)
Returns `true` only if **all** conditions are true.
```go
fmt.Println(true && true)   // => true
fmt.Println(true && false)  // => false
```
#### `!` (NOT)
Reverses the logical state.
```go
fmt.Println(!true)  // => false
fmt.Println(!false) // => true
```
#### Truth Table
| A     | B     | A && B | A \| B | !A    |
| ----- | ----- | ------ | ------ | ----- |
| true  | true  | true   | true   | false |
| true  | false | false  | true   | false |
| false | true  | false  | true   | true  |
| false | false | false  | false  | true  |
## Conditional Statements
Up to this point, our programs have been running in a straight line from the first instruction at the top to the last one at the bottom. However, real-world programs often need more flexibility. Sometimes we want to run certain code **only when a specific condition is true**, or choose between different paths depending on what is happening in the program. To handle situations like this, Go provides **conditional statements**, such as `if`, `else if`, and `else`. These allow our code to make decisions and react to different scenarios instead of always following the same flow.
### Single Condition with `if`
Sometimes we want a piece of code to run **only when a condition is true**. If the condition is false, Go simply skips that block and continues, we can create this by single if block
```go
age := 18
if age >= 18 {
    fmt.Println("You are an adult.")
}
```
We begin with the keyword `if`, followed by the condition we want to check. After that, we write a block of code that will run only if the condition is true.
### Alternative Path with `if-else`
We can improve this by adding an **alternative path**. If the condition is true, one block runs; otherwise, the `else` block runs.
```go
isRaining := false
if isRaining {
    fmt.Println("Bring an umbrella.")
} else {
    fmt.Println("No umbrella needed today.")
}
```
The `else` block does not need a condition. It simply runs as the alternative code when the `if` condition is not met.
### Multiple Conditions with `else if`
When we have **more than two possible outcomes**, we can chain conditions using `else if`. Go will check them in order and execute the first one that is true.
```go
score := 85
if score >= 90 {
    fmt.Println("Excellent performance")
} else if score >= 80 {
    fmt.Println("Good job")
} else {
    fmt.Println("Room for improvement")
}
```
### `if` with a Short Statement
Go allows you to write a small statement (like creating or assigning a variable) **before the condition**, separated by a semicolon.  
The variable only exists inside the `if` / `else` block.   
This is useful when we want to prepare a value and immediately check it.
```go
if n := 3; n > 5 {
    fmt.Println("n is greater than 5")
} else {
    fmt.Println("n is 5 or less")
}
```
### The `switch`, `case` Statement
When we need to check a value against multiple possible options, Go provides the `switch` statement. It allows us to write cleaner and more organized code compared to chaining many `else if` conditions. With `switch`, we evaluate a single value and match it against different cases, running the code for the first case that matches.
```go
day := "Monday"

switch day {
case "Monday":
    fmt.Println("Start of the week")
case "Friday":
    fmt.Println("Almost the weekend!")
case "Saturday", "Sunday": // Multiple values in one case
    fmt.Println("It’s the weekend!")
default:
    fmt.Println("Just another weekday")
}
```
To use a `switch` statement, we start with the `switch` keyword followed by the variable we want to check. Inside the block, we write our cases each one begins with `case` followed by a value to match. If a case matches the value, its code runs. When a case should match multiple values, we separate them with commas. Finally, a `default` case can be added to handle any situation where none of the previous cases match.
### Expressionless `switch`
Sometimes we don’t just want to compare a single value  we want to check more flexible or complex conditions. Go allows a `switch` without a value, which means each `case` can contain a boolean condition. The first condition that evaluates to `true` will run, making this a cleaner alternative to multiple `if-else if` statements.
```go
score := 85

switch { // No variable here!
case score >= 90:
    fmt.Println("Excellent performance")
case score >= 80:
    fmt.Println("Good job")
default:
    fmt.Println("Room for improvement")
}
```

## Loops
Sometimes we need to repeat an action many times for example, processing all values in a slice, performing a calculation repeatedly, or running a block of code until a condition is met. Instead of writing the same code over and over, we use **loops**. A loop lets us run a piece of code multiple times efficiently. Go provides different loop patterns to help us handle various repeating tasks, with `for` being the main looping construct in the language.
### The Traditional `for` Loop
The traditional `for` loop lets us repeat code a **specific number of times**. This is useful when we know exactly how many iterations we want. It consists of three parts:
1. **Initialization** – usually creates a counter variable (e.g., `i := 0`)
2. **Condition** – the loop runs while this is `true` (e.g., `i < 5`)
3. **Post / Increment** – updates the counter each loop (e.g., `i++`)
```go
for i := 0; i < 5; i++ {
    fmt.Printf("Count: %d\n", i)
}
// Prints 0, 1, 2, 3, 4
```
In this example:
- The loop starts at `i = 0`
- Runs while `i` is less than 5
- Increases `i` by 1 each time (`i++`)
### The `for-range` Loop
The `for-range` loop is Go's way of iterating over collections like slices, arrays, maps, and strings. Instead of using a counter manually, Go gives us the index and value automatically.    
`range` automatically provides two values each iteration:
- the **index** (position in the collection)
- the **value** at that index
```go
fruits := []string{"apple", "banana", "cherry"}

for i, fruit := range fruits {
    fmt.Printf("Index %d, Fruit %s\n", i, fruit)
}
```
- `i` is the index (0, 1, 2…)
- `fruit` is the value at that index

If we only care about the value and not the index, we can use the **blank identifier** `_` to ignore it:
```go
for _, fruit := range fruits {
    fmt.Println(fruit)
}
```
This makes the code cleaner when you don't need the index and only care about the items.
#### Iterating Over a Map
We can also use for to loop over map, When looping over a map, range gives us:
- key
- value

```go
prices := map[string]float64{
    "apple":  1.2,
    "banana": 0.8,
    "cherry": 2.5,
}

for fruit, price := range prices {
    fmt.Printf("%s costs %.2f\n", fruit, price)
}
```
If we only need the keys:
```go
for fruit := range prices {
    fmt.Println(fruit)
}
```
### Looping While a Condition is True
In Go, a `for` loop can be used not only with an initializer and increment  it can also run **purely based on a condition**.  This form behaves similarly to a traditional **while loop** in other languages.
```go
i := 0
for i < 3 { // This is Go's "while" loop
    fmt.Printf("i is %d\n", i)
    i++ // i = i + 1
}
```
Here, the code keeps repeating **as long as `i` is less than 3**, The loop stops once `i` becomes equal to 3 (the condition becomes false).
### `break` and `continue`
Go gives us more control over our loops by using `break` and `continue`.  These keywords help us decide **when to stop a loop early** or **skip parts of it**.     
**`break`:** this keyword help us to  immediately stops the loop and exits it.
```go
for i := 1; i <= 5; i++ {
    if i == 3 {
        break // stop the loop when i reaches 3
    }
    fmt.Println(i)
}
```
In the other hand `continue` skips the current iteration and moves to the next one.
```go
for i := 1; i <= 5; i++ {
    if i == 3 {
        continue // skip printing when i is 3
    }
    fmt.Println(i)
}
```
## Tasks
### Task 1: Strings
- Create a string variable `word` with the value `"Developer"`.
    - Print the first byte and the last byte.
    - Convert the word to uppercase and print it.
- Create a string variable `phrase` with this text: `"I love Python"`.
    - Replace "Python" with "Go" and print the result.
- Create a string variable `sentence` with this text: `"Learning Go is fun"`.
    - Print `true` if "Go" is in the sentence.
    - Print the number of **bytes** in the sentence.
    - Print the number of **runes** (characters) in the sentence.

### Task 2: Slices
- Create a slice of colors: `[]string{"red", "green", "blue"}`.
    - Print the first and last element.
    - Add `"yellow"` to the end of the slice.
    - Insert `"black"` at the beginning of the slice.
    - Remove and print the last element.
    - Check if `"green"` exists in the slice (print `true` or `false`).
- Given the slice `[]int{5, 3, 8, 1}`:
    - Remove the second item (3).
    - Insert `25` at index 2 (after 30).
	- Sort the slice in ascending order.

### Task 3: Maps
- Create a map representing a person with name (string), age (int), and city (string).
    - Print the name.
    - Add a new key-value pair for `"occupation": "Developer"`.
    - Print all the keys of the map.
    - Check if the map contains the key `"age"` (print `true` or `false`).
- Merge this map with another map: `{"hobbies": []string{"reading", "coding"}}`.
- Delete the city information from the map.
### Task 4: Conditionals
- Write an `if-else` statement that checks if a number is even or odd.
- Create a grading system using the `switch` statement (with no expression) for the following:

|**Class interval**|**Letter grade**|
|---|---|
|80-100|A+|
|75-79|A|
|70-74|A−|
|65-69|B+|
|60-64|B|
|55-59|B−|
|50-54|C+|
|45-49|C|
|40-44|D|
|0-39|F|


### Task 5: Quadratic Equation Solver
Write a Go program that reads three `float64` numbers, **a**, **b**, and **c**, from the user, and then solves the quadratic equation: $ax^2 + bx + c = 0$.

The program should handle the following cases:
1. When the discriminant ($b^2 - 4ac$) is positive, print the two real and distinct roots.
2. When the discriminant is zero, print the one real root.
3. When the discriminant is negative, print "No real solutions".

(Hint: You will need the `fmt` package to scan for input and the `math` package for `math.Sqrt` and `math.Pow`.)

### Task 6: Loops
- Create a program that asks the user to enter a positive integer and calculates its factorial using a `for` loop (in the style of a `while` loop).
