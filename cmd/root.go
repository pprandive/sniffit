/*
Copyright © 2022 Prafulla P. Ranadive <pprandive@rediffmail.com>

*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sniffit",
	Short: "An application/tool to capture network traffic.",
	Long: `Use this tool to capture the network traffic. 
You may first want to list the available devices and then choose the device 
to capture traffic. By default, interface value 'any' and port '11111' will be 
used to capture the network packets and they will be stored in file 
'/tmp/sniffit.pcap' unless different path is specified.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sniffit.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


