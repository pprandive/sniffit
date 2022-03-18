/*
Copyright © 2022 Prafulla P. Ranadive <pprandive@rediffmail.com>

*/

package lib

import (
    "testing"
)

// TestSniffer Tests the function Capture
func TestSniffer(t *testing.T){
    sniffer := NewSniffer()
    if err := sniffer.Capture("enp0s3", "", "/tmp/sniffit.pcap"); err != nil{
        t.Errorf("error isn't expected. Received: %v", err)
    }

}
