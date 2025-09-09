package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// newRootCommand creates the root command for cli
func newRootCommand() error {
	var rootCmd = &cobra.Command{
		Use:   "t1dash",
		Short: "Type 1 Diabetes Data Dashboard",
	}
	rootCmd.AddCommand(newServerCommand())

	err := rootCmd.Execute()
	return err
}

// main entry point
func main() {
	err := newRootCommand()
	if err != nil {
		fmt.Println(err.Error())
	}

	os.Exit(0)
}
