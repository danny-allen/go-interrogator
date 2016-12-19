package interrogator

import (
	"fmt"
	"log"
)

// Data struct for Ask.
type Question struct {

	// Public properties.
	Content 	string
	Open 		bool
	Answers 	map[string][]string
	Response	string
}

// Create a new question.
func NewQuestion(content string) *Question {
	q := new(Question)
	q.Content = content
	q.Answers = make(map[string][]string)
	return q
}

// Store answers for a question.
func (q *Question) SetAnswer(answer string, values []string) *Question {

	// Set the answers on the question.
	q.Answers[answer] = values

	// Return the question.
	return q
}

// Ask a question.
func (q *Question) Ask() *Question {

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
	if(q.Open) {

		// Set the key as response if found.
		q.Response = response

	} else {

		// Loop the answer groups.
		for key, arr := range q.Answers {

			// Loop each answer.
			for _, value := range arr {

				// Check user answer against accepted answers.
				if (value == response) {

					// Set the key as response if found.
					q.Response = key
				}
			}
		}
	}

	// Return the question.
	return q
}

// Check the response.
func (q *Question) IsResponse(key string) bool {

	// Check the response matches.
	if(q.Response == key) {
		return true
	} else {
		return false
	}
}