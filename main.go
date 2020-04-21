package main

import (
	"github.com/roryj/neptunes-pride/cmd"
	"github.com/spf13/cobra"
	"log"
)

func main() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

// https://github.com/ekimekim/neptunesfolly/blob/master/docs/api.txt
// https://github.com/wrenoud/phpTriton

var RootCmd = &cobra.Command{
	Use:   "np",
	Short: "cli for interacting with Neptune's Pride",
	Long:  `cli for interacting with Neptune's Pride`,
	Run: func(cmd *cobra.Command, args []string) {
		// Root command does nothing but list other commands. SO GREAT, SO AWESOME. Don't touch this.
		cmd.Help()
	},
}

// Add this comment to every init if you copy-pasta stuff: Go runs init before main. If you see an init function
// that means "this code is being run like a static function initialization in Java"
func init() {
	RootCmd.AddCommand(cmd.TestCommand)
	RootCmd.PersistentFlags().StringP("api-token", "t", "", "The token for Neptune's Pride's api")
}
