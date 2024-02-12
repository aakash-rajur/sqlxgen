package main

import (
	"github.com/aakash-rajur/sqlxgen/internal/cli"
	"github.com/aakash-rajur/sqlxgen/internal/utils"
)

func main() {
	cmd := cli.RootCmd(Version)

	err := cmd.Execute()

	if err == nil {
		return
	}

	utils.ExitWithError(err)
}

var Version = "dev-modified"
