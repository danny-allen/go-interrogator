# Go Interrogator

This package was built to ask questions. You can create two types of questions. Open and Closed.

## Installation

Run the following command to install the package:

```sh
go get github.com/danny-allen/go-interrogator
```

## Examples

### Closed Question

This is an example of a closed question. It has specific acceptable answers.

```golang
// Create a question.
q := interrogator.NewQuestion(
    "You are currently using an old version. Would you like to upgrade? [y/n]",
)

// Define answers.
q.SetAnswer("yes", []string{"yes", "y"})
q.SetAnswer("no", []string{"no", "n"})

// Ask the question.
q.Ask()

// On yes response.
if(q.IsResponse("yes")) {
    fmt.Println("Starting update...")
}

// On no response.
if(q.IsResponse("no")) {
    fmt.Println("Okay, but you may be missing out on some cool features!")
}
```

### Open Question

You can also create an open question. To do this, you can change the `Open` property on the `Question` struct to true.

```golang
    // Create a question.
    q = interrogator.NewQuestion(
        "What is your name?",
    )

    // Set the question to open.
    q.Open = true

    // Ask the question.
    q.Ask()

    // Output the response.
    fmt.Println(q.Response)
```