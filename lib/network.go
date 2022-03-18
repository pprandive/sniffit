/*
Copyright © 2022 Prafulla P. Ranadive <pprandive@rediffmail.com>

*/

package lib

import (
    "log"
    "fmt"
    "os"
    "os/signal"
    "syscall"

    "github.com/google/gopacket"
    "github.com/google/gopacket/pcap"
    "github.com/google/gopacket/pcapgo"
    "github.com/google/gopacket/layers"
    _ "net/http"
    _ "net/http/pprof"
)

type Sniffer interface{
    Capture(iface, filter, fileName string) error
}

func NewSniffer() Sniffer{
    return &SnifferImpl{}
}

var (
    buffer = int32(65535)
)

type SnifferImpl struct{
}

// Capture Starts to capture the packets on the specified interface and stores it in the file
// specified.
func (snif *SnifferImpl) Capture(iface, filter, fileName string) error{

   if !deviceCheck(iface) {
        log.Fatal("Failed to open device :: ", iface)
        return fmt.Errorf("Failed to open device :: %s", iface)
    }

    handler, err := pcap.OpenLive(iface, buffer, false, pcap.BlockForever)

    if err != nil {
        log.Fatal(err)
        return err
    }
    defer handler.Close()

    if err := handler.SetBPFFilter(filter); err != nil {
        log.Fatal(err)
        return err
    }

    log.Println("iface=", iface, "filter=", filter, "fileName", fileName)

    count := 1
    termChan := make(chan os.Signal, 1)
    signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT)

    go func(){

        // destination path you want your pcap saved to
        outputFilename := fileName
        outputFile, err := os.Create(outputFilename)
        if err != nil {
            log.Fatal("Failed to create packet capture output file", "err:", err, "filename:", outputFilename)
            return
        }
        defer func(outputFile *os.File) {
            if err := outputFile.Close(); err != nil {
                log.Fatal("Failed to close file", err)
            }
        }(outputFile)

        outputWriter := pcapgo.NewWriter(outputFile)
        if err := outputWriter.WriteFileHeader(8192, layers.LinkTypeEthernet); err != nil {
            log.Fatal("Failed to write file header", err)
            return
        }

        source := gopacket.NewPacketSource(handler, handler.LinkType())
        for  packet := range source.Packets() {
            err = outputWriter.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
            if err != nil {
               log.Fatal("Failed to write packet data")
            }
            count++
        }
   }()

    select {
        case <-termChan:
            log.Print("\n\nPackets captured: ", count)
            os.Exit(0)
    }
    return nil
}

func deviceCheck(name string) bool {
    devices, err := pcap.FindAllDevs()

    if err != nil {
        log.Panic(err)
    }

    for _, device := range devices {
        if device.Name == name {
            return true
        }
    }
    return false
}
