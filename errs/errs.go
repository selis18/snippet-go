package errs

import "fmt"

func CE(err error, s string) {
	if err != nil {
		fmt.Println("s", err)
	}
}
