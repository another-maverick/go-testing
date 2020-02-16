package main

import (
	"fmt"
	"io"
	"testing"
)

type MyWriter struct {
	//
}

// This is supposed to implement io.Writer interface but it does not since Write method does not return int
func (m *MyWriter) Write([]byte) error {
	return nil
}

func main() {
	testWriter := map[string]interface{} {
		"test": &MyWriter{},
	}

	//Test the testWriter

	if _, ok := testWriter["test"].(io.Writer); !ok {
		fmt.Errorf("There is a problem. %s does not implement io.Writer", testWriter["test"])
	}


}

//Here compiler itself would fail
func TestWriter(t *testing.T) {
	var _ io.Writer = &MyWriter{}
}

/*
Here is the error that we will see

./assertion_test.go:34:6: cannot use &MyWriter literal (type *MyWriter) as type io.Writer in assignment:
*MyWriter does not implement io.Writer (wrong type for Write method)
have Write([]byte) error
want Write([]byte) (int, error)
FAIL	command-line-arguments [build failed]

*/
