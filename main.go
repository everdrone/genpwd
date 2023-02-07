package main

import (
	"fmt"
	"os"

	"github.com/everdrone/genpasswd/cmd"
	"github.com/everdrone/genpasswd/internal"
)

func main() {
	// from: https://github.com/spf13/cobra/issues/914#issuecomment-548411337
	if err := cmd.RootCmd.Execute(); err != nil {
		// if we have ErrSilent, we don't want to print the error
		if err != internal.ErrSilent {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}
}
