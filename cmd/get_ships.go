package cmd

import (
	"fmt"
	"github.com/roryj/neptunes-pride/lib"
	"github.com/spf13/cobra"
)

var TestCommand = &cobra.Command{
	Use:   "test",
	Short: "Get the available ships for you",
	Long:  `Get the available ships for you`,
	Run: func(cmd *cobra.Command, args []string) {
		config := lib.ParseConfigFile()

		fmt.Println("Creating client")
		client := lib.NewNeptuneClient("test", config)
		fmt.Println("Client created")

		fmt.Println("Logging in")
		client.Login(config.Username, config.Password)
		fmt.Println("Logged in")

		data := client.GetShips()

		users := lib.GetAllPlayers()

		//		for
	},
}
