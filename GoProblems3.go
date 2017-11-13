// Student Name: Timothy Cassidy
// Student Number: G00333333

// Resources
// https://golang.org/pkg/math/rand/#Rand.Intn
// https://golang.org/pkg/regexp/
// 

package main

// package imports
import "fmt"
import "time"
import "math/rand"
import "regexp"

// generic responses string array
var responseItems = []string {
	"I'm not sure what you're trying to say. Could you explain it to me?",
	"How does that make you feel?",
	"Why do you say that?",
}// responseItems

// function that takes a string and returns a response string
func ElizaResponse(inputString string) string {
	// MustCompile is like Compile but panics if the expression cannot be parsed. 
	// It simplifies safe initialization of global variables holding compiled 
	// regular expressions. 
	fatherXp := regexp.MustCompile("(?i)\\bfather\\b")
	

	// FindStringIndex returns a two-element slice of integers defining the 
	// location of the leftmost match in s of the regular expression. The match 
	// itself is at s[loc[0]:loc[1]]. A return value of nil indicates no match. 
	// in this case we are checking to see if the regxp "father" is contained within the input string
	if len(fatherXp.FindStringIndex(inputString)) > 0 {
		return "Why don't you tell me more about your father?"
	}// if
	
	// Intn returns, as an int, a non-negative pseudo-random number
	// using this number to randomly pick a response from the array
	return responseItems[rand.Intn(len(responseItems))]
}// ElizaResponse

// Main function
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
	fmt.Println("I am looking forward to the weekend.")
	fmt.Println(ElizaResponse("I'm looking forward to the weekend."))
	fmt.Println("My grandfather was French!")
	fmt.Println(ElizaResponse("My grandfather was French!"))
	
}// main