package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var databseCmd = &cobra.Command{
	Use:   "db",
	Short: "Start database command of Golang",
	Long:  `Start database command of Golang`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("database command starting...", args)
		switch args[0] {
		case "create":
			fmt.Println("Creating Database Done!")
		case "migrate":
			fmt.Println("Auto Migration Done!")
		}
	},
}

func init() {
	rootCmd.AddCommand(databseCmd)
}
