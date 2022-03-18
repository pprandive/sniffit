/*
Copyright © 2022 Prafulla P. Ranadive <pprandive@rediffmail.com>

*/

package cmd

import (
    "log"

    "sniffit/lib"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Command to start the capture.",
	Long: `Use this sub-command to start the packet capture. 
Interface, port and the file name are the options to specify while calling this. 
By default, the packets captured are stored in the file '/tmp/sniffit.pcap'`,
	Run: func(cmd *cobra.Command, args []string) {
        port, _ := cmd.Flags().GetString("port")
        iface, _ := cmd.Flags().GetString("iface")
        fileName, _ := cmd.Flags().GetString("fileName")
        sniffer := lib.NewSniffer()
        if err := sniffer.Capture(iface, "port "+port, fileName); err != nil{
            log.Printf("error isn't expected. Received: %v", err)
        }
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	startCmd.Flags().StringP("port", "p", "11111", "Port number. Default is 11111, which is default port of 'sockperf', too.")
	startCmd.Flags().StringP("iface", "i", "any", "Interface name.")
	startCmd.Flags().StringP("fileName", "f", "/tmp/sniffit.pcap", "File to save the captured packets.")
}
