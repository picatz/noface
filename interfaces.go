package noface

import (
  "os"
  "fmt"
  "errors"
  "github.com/google/gopacket/pcap"
)

// Simply lookup the first available interface.
func FirstIface() (string, error) {
  // Find all the devices.
  devices, err := pcap.FindAllDevs()

  // Check there was no errors.
  if err != nil {
    fmt.Println(err)
  }

  // Do we have any devices to use?
  if len(devices) > 0 {
    device  := devices[0]
    // Check if address list for device is empty.
    if len(device.Addresses) > 0 {
      address := device.Addresses[0]
      ip      := address.IP
      // If the IP address is availalbe, use that.
      if len(ip) > 0 {
        return device.Name, nil
      }
    }
  }

  // Error out when things don't go as planned.
  return "", errors.New("noface")
}
