package error

import (
    "errors"
    "testing"
)

// type error Interface {
//     Error() string
// }

type MyError struct {
    Msg string
}

func (e *MyError) Error() string {
    return e.Msg
}

func printError(err error) {
    if err != nil {
        println(err.Error())
    }
}

func TestBuiltinError(t *testing.T) {
    err := errors.New("error message")
    printError(err) // error message

    myErr := &MyError{"my error message"}
    printError(myErr) // my error message
}
