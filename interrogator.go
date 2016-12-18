package interrogator

import (
	"fmt"
	"log"
)

// Data struct for Ask.
type Question struct {

	// Define interface for public methods.
	AInterface

	// Public properties.
	Content 	string
	Open 		bool

	// Private properties.
	response 	string
	answers		map[string][]string
}

// Ask interface
type AInterface interface {

	// Public methods.
	Question()
	Response()
	Answers()
}

// Store answers for a question.
func (q Question) Answers(question Question, answers map[string][]string) Question {

	// Set the answers on the question.
	question.answers = answers

	// Return the question.
	return question
}

// Ask a question.
func (q Question) Ask(question Question) Question {

	// Ask the question.
	fmt.Println(q.Content)

	// Declare Response string.
	var response string

	// Get the response.
	_, err := fmt.Scanln(&response)

	// Check for errors.
	if err != nil {
		log.Fatal(err)
	}

	// If open question, just add the response.
	if(question.Open) {

		// Set the key as response if found.
		question.response = response

	} else {

		// Loop the answer groups.
		for key, arr := range question.answers {

			// Loop each answer.
			for _, value := range arr {

				// Check user answer against accepted answers.
				if (value == response) {

					// Set the key as response if found.
					question.response = key
				}
			}
		}
	}

	// Return the question.
	return question
}

// Check the response.
func (q Question) Response(question Question, key string) bool {

	// Check the response matches.
	if(question.response == key) {
		return true
	} else {
		return false
	}
}


// Tests.
func main() {

	//// Declare question.
	//q := interrogator.Question{
	//	Content: "Would you like to upgrade",
	//}
	//
	//// Declare answers.
	//answers := map[string][]string{}
	//
	//// Define answers.
	//answers["yes"] = []string{"yes", "y"}
	//answers["no"] = []string{"no", "n"}
	//
	//// Ask question.
	//q = q.Ask(q)
	//q = q.Answers(q, answers)
	//
	//// On yes response.
	//if(q.Response(q, "yes")) {
	//	fmt.Println("YES!")
	//}
	//
	//// On no response.
	//if(q.Response(q, "no")) {
	//	fmt.Println("Nope!")
	//}
}