package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"np/lib"
)

var TestCommand = &cobra.Command{
	Use:   "test",
	Short: "Get the available ships for you",
	Long:  `Get the available ships for you`,
	Run: func(cmd *cobra.Command, args []string) {
		config := lib.ParseConfigFile()

		fmt.Println("Creating client")
		client := lib.NewNeptuneClient("test", config)
		fmt.Println( "Client created")

		fmt.Println( "Logging in")
		client.Login(config.Username, config.Password)
		fmt.Println( "Logged in")

		data := client.GetShips()

		users := lib.GetAllPlayers()

		for
	},
}