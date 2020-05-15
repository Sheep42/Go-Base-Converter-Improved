package base_converter

import "fmt"

type inputErr struct {
	input string
	msg   string
}

func (e *inputErr) Error() string {
	return fmt.Sprintf("%s %s", e.input, e.msg)
}
