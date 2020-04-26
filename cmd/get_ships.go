package cmd

import (
	"fmt"

	"github.com/roryj/neptunes-pride/lib"
	"github.com/spf13/cobra"
)

var GetShipsCommand = &cobra.Command{
	Use:   "get-ships",
	Short: "Get the available ships for you",
	Long:  `Get the available ships for you`,
	Run: func(cmd *cobra.Command, args []string) {

		_ = cmd.ParseFlags(args)
		gameID, _ := cmd.Flags().GetString("game-id")

		config := lib.ParseConfigFile()

		fmt.Println("Creating client")
		client := lib.NewNeptuneClient(gameID, "test", config)
		fmt.Println("Client created")

		fmt.Println("Logging in")
		client.Login()
		fmt.Println("Logged in")

		data := client.GetShips()

		fmt.Printf("\nresult of ships: %+v\n", data)

		users := lib.GetAllPlayers()

		fmt.Printf("\nresult of players: %+v\n", users)

		//		for
	},
}
