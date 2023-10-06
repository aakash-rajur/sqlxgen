package utils

import (
	"fmt"
	"os"
)

func ExitWithError(err error) {
	fmt.Printf("error: %+v\n", err)

	os.Exit(1)
}
