package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/z6wdc/go-cobra/internal/controller"
)

var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Greet someone",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a name")
			return
		}
		controller.RunGreet(args[0])
	},
}

func init() {
	rootCmd.AddCommand(greetCmd)
}
