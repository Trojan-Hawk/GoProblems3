// Student Name: Timothy Cassidy
// Student Number: G00333333

// Resources
// https://golang.org/pkg/math/rand/#Rand.Intn
// 

package main

// format package import
import "fmt"
import "time"
import "math/rand"

// generic responses string array
var responseItems = []string {
	"I'm not sure what you're trying to say. Could you explain it to me?",
	"How does that make you feel?",
	"Why do you say that?",
}// responseItems

// function that takes a string and returns a response string
func ElizaResponse(inputString string) string {
	// Intn returns, as an int, a non-negative pseudo-random number
	// using this number to randomly pick a response from the array
	return responseItems[rand.Intn(len(responseItems))]
}

// Hello World main function
func main() {
	// using the system clock to set a random seed value
	rand.Seed(time.Now().UTC().UnixNano())
	
	// input and random responses
	fmt.Println("People say I look like both my mother and father.")
	fmt.Println(ElizaResponse("People say I look like both my mother and father."))
	fmt.Println("Father was a teacher.")
	fmt.Println(ElizaResponse("Father was a teacher."))
	fmt.Println("I was my father's favourite.")
	fmt.Println(ElizaResponse("I was my father's favourite."))
	fmt.Println("I'm looking forward to the weekend.")
	fmt.Println(ElizaResponse("I'm looking forward to the weekend."))
	fmt.Println("My grandfather was French!")
	fmt.Println(ElizaResponse("My grandfather was French!"))
	
}// main