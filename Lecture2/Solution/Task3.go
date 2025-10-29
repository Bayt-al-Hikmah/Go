package main

import "fmt"

func main() {
	// Creating the map
	person := map[string]any{
		"name": "Alice",
		"age":  30,
		"city": "New York",
	}

	// Printing name and adding occupation
	fmt.Println("Name:", person["name"])
	person["occupation"] = "Developer"

	// Printing all keys
	fmt.Println("Keys:")
	for key := range person {
		fmt.Println(key)
	}

	// Checking for "age" key
	_, hasAge := person["age"]
	fmt.Println("Contains 'age' key:", hasAge)

	// Merging with another map
	hobbies := map[string]any{
		"hobbies": []string{"reading", "coding"},
	}

	for key, value := range hobbies {
		person[key] = value
	}

	// Deleting city information
	delete(person, "city")
	fmt.Println("Final map:", person)
}