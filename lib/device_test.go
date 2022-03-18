/*
Copyright © 2022 Prafulla P. Ranadive <pprandive@rediffmail.com>

*/

package lib

import (
    "testing"
)

// TestGetDevices Tests the function GetDevices()
func TestGetDevices(t *testing.T){
    devices, _ := GetDevices()
    if devices == nil{
        t.Errorf("expected atleast one device, got nothing")
    }
    t.Log("Devices list is:",devices)
}
