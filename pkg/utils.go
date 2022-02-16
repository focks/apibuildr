package pkg

import (
	"fmt"
	"os"
)

func CheckError(msg interface{}) {
	if msg != nil {
		fmt.Fprintln(os.Stderr, "Error:", msg)
		os.Exit(1)
	}
}
