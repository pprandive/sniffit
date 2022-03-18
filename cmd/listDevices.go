/*
Copyright © 2022 Prafulla P. Ranadive <pprandive@rediffmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

    "sniffit/lib"
)

// listDevicesCmd represents the listDevices command
var listDevicesCmd = &cobra.Command{
	Use:   "listDevices",
	Short: "A sub-command to list available interfaces",
	Long: `Using this command, one can know the available interfaces on this machine.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("listDevices called")
        devices, err := lib.GetDevices()
        if err != nil{
            fmt.Println("Something went wrong while fetching devices list.")
            return
        }
        counter := 1
        for _, device := range devices{
            fmt.Println(counter, ":", device)
            counter++
        }
	},
}

func init() {
	rootCmd.AddCommand(listDevicesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listDevicesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listDevicesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
