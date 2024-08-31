package main

import (
	"fmt"
	"os"
)

// Global variables
// L is capital, so it is supposed to be a public variable
const LoginToken string = "someValue" // Public variable

func main() {
	// var username string = "John Doe"
	// var age1 int = 25
	// var married bool = false
	// var salary float32 = 5000.50

	// fmt.Println("Hello,", username, "Your age is", age1, "and your married status is", married, "and your salary is", salary)

	// // Default values and aliases
	// var num int      // 0
	// var str string   // ""
	// var boolean bool // false
	// fmt.Println("Default value of number is:", num)
	// fmt.Println("Default value of string is:", str)
	// fmt.Println("Default value of boolean is:", boolean)

	// // Go also support implicit type declaration
	// var name = "John Doe"
	// fmt.Println("Hello, Implicitly Typed: ", name)

	// // Var syntax is also not required, cannot be used outside of a function as it is a shorthand syntax, cannot be used as a global variable
	// age2 := 25
	// fmt.Println("Hello, Implicitly Typed: ", age2)

	// fmt.Println("Login Token: ", LoginToken)

	// User input
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter rating between 1 to 5: ")
	// // Comma ok || error error
	// // Used as try catch block
	// input, _ := reader.ReadString('\n')

	// // String to int conversion
	// rating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// } else {
	// 	fmt.Println("Rating is:", rating+1)
	// }

	// Time in Go
	// currentTime := time.Now()
	// fmt.Println("Current Time: ", currentTime.Format("01-02-2006 15:04:05 Monday"))

	// // Pointers in Go
	// var ptr *int
	// fmt.Println("Value of ptr: ", ptr) // <nil>

	// // Memory address of a variable
	// myNumber := 23
	// var myNumberPtr *int = &myNumber
	// fmt.Println("Memory address of myNumber: ", myNumberPtr) // Memory address of myNumber
	// fmt.Println("Value of myNumber: ", *myNumberPtr)         // Value of myNumber

	// *myNumberPtr = *myNumberPtr + 2
	// fmt.Println("Value of myNumber: ", *myNumberPtr) // Value of myNumber after increment

	// // Arrays in Go
	// var numbers [5]int
	// numbers[0] = 1
	// numbers[1] = 2
	// numbers[2] = 3
	// numbers[3] = 4
	// numbers[4] = 5
	// // numbers[5] = 6 // Error: index out of range [5] with length 5
	// fmt.Println("Numbers Array: ", numbers)

	// Slices in Go
	// var fruits = []string{"Apple", "Banana", "Mango"}
	// fmt.Println("Fruits Slice: ", fruits) // [Apple Banana Mango]

	// // Append to a slice
	// fruits = append(fruits, "Orange")                     // [Apple Banana Mango Orange]
	// fruits = append(fruits[1:4])                          // [Banana Mango Orange]
	// fmt.Println("Fruits Slice After Appending: ", fruits) // [Banana Mango Orange]

	// var scores = make([]int, 3)
	// scores[0] = 345
	// scores[1] = 678
	// scores[2] = 123
	// // scores[3] = 40 // Will cause an error
	// scores = append(scores, 235)          // [345 678 123 235]
	// fmt.Println("Scores Slice: ", scores) //  [345 678 123 235]

	// fmt.Println("Scores sorted: ", sort.IntsAreSorted(scores)) // false
	// sort.Ints(scores)
	// fmt.Println("Sorted Scores Slice: ", scores) // [123 235 345 678]

	// // Remove an element from a slice based on index
	// index := 2
	// scores = append(scores[:index], scores[index+1:]...) // [123 235 678]
	// fmt.Println("Scores Slice After Removing: ", scores) // [123 235 678]

	// // Maps in Go (Key-Value pair)
	// // map[keyType]valueType
	// languages := make(map[string]string)
	// languages["JS"] = "JavaScript"
	// languages["PY"] = "Python"
	// languages["RB"] = "Ruby"
	// languages["GO"] = "Golang"
	// fmt.Println("Languages Map: ", languages)       // map[GO:Golang JS:JavaScript PY:Python RB:Ruby]
	// fmt.Println("JS stands for: ", languages["JS"]) // JavaScript

	// // Delete a key-value pair from a map
	// delete(languages, "RB")
	// fmt.Println("Languages Map After Deleting: ", languages) // map[GO:Golang JS:JavaScript PY:Python]

	// for key, value := range languages {
	// 	fmt.Println("Key:", key, "Value:", value)
	// }

	// // Structs in Go (Custom Data Types) (Object in JS) (Class in Java)
	// danish := User{Name: "Danish", Email: "danishsidd@gmail.com", isVerified: true, age: 25}
	// fmt.Println("User: ", danish) // {Danish danishsidd@gmail.com true 25}
	// fmt.Println("User Name: ", danish.Name)

	// Conditional Statements in Go
	// loginCount := 12
	// var result string
	// if else
	// if loginCount < 10 {
	// 	result = "Regular User"
	// } else if loginCount > 10 {
	// 	result = "Excellent User"
	// } else {
	// 	result = "Awesome User"
	// }
	// fmt.Println("User Result: ", result)

	// Switch Case
	// switch loginCount {
	// case 1:
	// 	result = "One Time User"
	// case 10:
	// 	result = "Ten Time User"
	// default:
	// 	result = "Awesome User"
	// }
	// fmt.Println("User Result: ", result)

	// Loops in Go
	// days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	// For Loop
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println("Day", d+1, "is", days[d])
	// }
	// Commonly used for loop
	// for index, day := range days {
	// 	fmt.Println("Day", index, "is", day)
	// }

	// While Loop
	// i := 0
	// for i <= 6 {
	// 	fmt.Println("Day:", i+1, "is", days[i])
	// 	i++
	// }

	// // Functions in Go
	// greet("Danish")
	// result := add(2, 3)
	// fmt.Println("Result of addition is:", result)

	// totalPrice, message := calculateBill(23, 45, 67, 89)
	// fmt.Println("Total Price is:", totalPrice)
	// fmt.Println("Message is:", message)

	// File Handling in Go
	// content := "Hello from Go"
	// file, err := os.Create("./info.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// length, err := io.WriteString(file, content)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Length of the content is:", length)
	// defer file.Close()

	// // Read from a file
	// readFile("./info.txt")
}

// type User struct {
// 	Name       string
// 	Email      string
// 	IsVerified bool
// 	Age        int
// }

// func greet(name string) {
// 	fmt.Println("Hello,", name)
// }

// func add(a int, b int) int {
// 	return a + b
// }

// func calculateBill(values ...int) (int, string) {
// 	total := 0
// 	for _, value := range values {
// 		total += value
// 	}
// 	return total, "Hello"
// }

// func readFile(filename string) {
// 	dataByte, err := os.ReadFile(filename)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Data from file is:", string(dataByte))
// }
