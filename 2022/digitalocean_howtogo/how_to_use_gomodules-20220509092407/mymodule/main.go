package main

import (
	"github.com/spf13/cobra"
	"mymodule/mypackage"
)

func main() { 
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			mypackage.PrintHello()
		},
	}
	cmd.Execute()
}
