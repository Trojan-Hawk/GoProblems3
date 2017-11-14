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
import "strings"

// generic responses string array
var responseItems = []string {
	"I'm not sure what you're trying to say. Could you explain it to me?",
	"How does that make you feel?",
	"Why do you say that?",
}// responseItems

// generic feeling responses string array
var feelings = []string {
	"very good.",
	"not very good.",
	"pretty good.",
	"could be better.",
}

// function that takes a string and returns a response string
func ElizaResponse(inputString string) string {
	// MustCompile is like Compile but panics if the expression cannot be parsed. 
	// It simplifies safe initialization of global variables holding compiled 
	// regular expressions. 
	regXp := regexp.MustCompile("(?i)\\bfather\\b")

	// FindStringIndex returns a two-element slice of integers defining the 
	// location of the leftmost match in s of the regular expression. The match 
	// itself is at s[loc[0]:loc[1]]. A return value of nil indicates no match. 
	// in this case we are checking to see if the regxp "father" is contained within the input string
	if len(regXp.FindStringIndex(inputString)) > 0 {
		return "Why don't you tell me more about your father?"
	}// if
	
	// if the input begins with how are you
	regXp = regexp.MustCompile("(?i)how are you(.*)")
	subString := regXp.FindStringSubmatch(inputString)
	
	if len(subString) > 1 {
		return feelings[rand.Intn(len(responseItems))]
	}// if
	
	// if the input begins with how
	regXp = regexp.MustCompile("(?i)how(.*)")
	subString = regXp.FindStringSubmatch(inputString)
	
	if len(subString) > 1 {
		str := tidyOutput(subString[1])
		return "I don't know, how" + str	
	}// if
	
	// if the input begins with “i am “, “I AM “, “I’m “, “Im “, “i’m “
	regXp = regexp.MustCompile("(?i)i(?:'| a|)?m(.*)")
	subString = regXp.FindStringSubmatch(inputString)
	
	if len(subString) > 1 {
		str := tidyOutput(subString[1])
		return "How do you know you are" + str
	}// if
	
	// if the input begins with what
	regXp = regexp.MustCompile("(?i)what(.*)")
	subString = regXp.FindStringSubmatch(inputString)
	
	if len(subString) > 1 {
		str := tidyOutput(subString[1])
		return "What do you mean what" + str
	}// if
	
	// Intn returns, as an int, a non-negative pseudo-random number
	// using this number to randomly pick a response from the array
	return responseItems[rand.Intn(len(responseItems))]
}// ElizaResponse

// function to tidy up the string to output
func tidyOutput(str string) string {
	// replacing the full stop with a question mark
	str = strings.Replace(str, ".", "?", 1)
	// reflecting the pronouns in the captured groups
	str = strings.Replace(str, "you are", "I'm", 1)
	str = strings.Replace(str, "you’re", "I'm", 1)
	str = strings.Replace(str, "your", "my", 1)
	str = strings.Replace(str, "you", "I", 1)
	str = strings.Replace(str, "me", "you", 1)
	str = strings.Replace(str, "i am", "you are", 1)
	return str
}// tidyOutput

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
	fmt.Println()
	
	fmt.Println("I am happy.")
	fmt.Println(ElizaResponse("I am happy."))
	fmt.Println("I am not happy with your responses.")
	fmt.Println(ElizaResponse("I am not happy with your responses."))
	fmt.Println("I am not sure that you understand the effect that your questions are having on me.")
	fmt.Println(ElizaResponse("I am not sure that you understand the effect that your questions are having on me."))
	fmt.Println("I am supposed to just take what you’re saying at face value?")
	fmt.Println(ElizaResponse("I am supposed to just take what you’re saying at face value?"))
	fmt.Println()
	
	// More input patterns
	// what
	fmt.Println("What do you think you are doing?")
	fmt.Println(ElizaResponse("What do you think you are doing?"))
	fmt.Println("What is life?")
	fmt.Println(ElizaResponse("What is life?"))
	fmt.Println("What is wubbalubbadubdub?")
	fmt.Println(ElizaResponse("What is wubbalubbadubdub?"))
	fmt.Println()
	
	// how are you
	fmt.Println("How are you?")
	fmt.Println(ElizaResponse("How are you?"))
	fmt.Println("How are you feeling today?")
	fmt.Println(ElizaResponse("How are you feeling today?"))
	fmt.Println("How are your family doing?")
	fmt.Println(ElizaResponse("How are your family doing?"))
	fmt.Println()
	
	// how
	fmt.Println("How do you build a car?")
	fmt.Println(ElizaResponse("How do you build a car?"))
	fmt.Println("How do you know what I am saying?")
	fmt.Println(ElizaResponse("How do you know what I am saying?"))
	fmt.Println("How can this be possiable?")
	fmt.Println(ElizaResponse("How can this be possiable?"))

	
}// main