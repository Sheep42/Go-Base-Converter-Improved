package errors

import "fmt"

type InputErr struct {
	input string
	msg   string
}

func (e *InputErr) Error() string {
	return fmt.Sprintf("Error: Value: %s - %s", e.input, e.msg)
}

func ThrowInputError(input string, msg string) error {
	return &InputErr{input, msg}
}
