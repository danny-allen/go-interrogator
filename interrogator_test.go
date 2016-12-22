package interrogator

import (
	"testing"
	"github.com/danny-allen/go-interrogator"
	"strconv"
	"fmt"
)

// Test values struct.
var InterrogatorTests = []struct {
	Content 	string
	Open		bool
	Answers 	map[string][]string
	Response	string
}{
	{
		"Did it work?",
		false,
		map[string][]string {
			"yes": {"yes", "y"},
			"no": {"no", "y"},
		},
		"",
	},
}

func TestNewQuestion(t *testing.T) {

	// Loop the test questions.
	for i, testQ := range InterrogatorTests {

		// Create a new question
		q := interrogator.NewQuestion(testQ.Content)

		// Check content.
		CheckContent(t, q, i)

		// Check answers.
		CheckAnswers(t, q, i)
	}
}


// Check the Content property of the question.
func CheckContent(t *testing.T, q *interrogator.Question, i int){

	// Check the content against th
	if(q.Content != InterrogatorTests[i].Content) {
		t.Error("CheckContent: There is something wrong with the content. Error occured in index " + strconv.Itoa(i) + ".")
	}
}

// Check the Answers property of the question.
func CheckAnswers(t *testing.T, q *interrogator.Question, i int) {

	fmt.Println(InterrogatorTests[i].Answers)

	if(len(InterrogatorTests[i].Answers) == 0 ){
		t.Error("There is something wrong with the answers. Error occured in index " + strconv.Itoa(i) + ".")
	}

	//
	if(len(InterrogatorTests[i].Answers) > 0) {

		// Add the answers.
		for key, answer := range InterrogatorTests[i].Answers {
			q.SetAnswer(key, answer)
		}
	}
}
