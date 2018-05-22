package cmd

import (
	"fmt"
	"os"
)

// Execute command tool
func Execute() {
	rootCmd.AddCommand(
		startAppCmd,
		startProjectCmd,
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
