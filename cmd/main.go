package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newRootCommand() error {
	var rootCmd = &cobra.Command{
		Use:   "t1dash",
		Short: "Type 1 Diabetes Data Dashboard",
	}
	rootCmd.AddCommand(newServerCommand())

	err := rootCmd.Execute()
	return err
}

func main() {
	err := newRootCommand()
	if err != nil {
		fmt.Println(err.Error())
	}

	os.Exit(0)
}
