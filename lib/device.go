/*
Copyright © 2022 Prafulla P. Ranadive <pprandive@rediffmail.com>

*/

package lib

import(
    "log"
    _ "github.com/google/gopacket"
    "github.com/google/gopacket/pcap"
)

// GetDevices returns the list of devices.
func GetDevices() ([]string, error){
    devices, err := pcap.FindAllDevs()
    if err != nil {
        log.Printf("Failed to get the device list: '%v'",err)
        return []string{}, err
    }

    devNames := []string{}
    for _, device := range devices{
        devNames = append(devNames, device.Name)
    }
    return devNames, nil
}
