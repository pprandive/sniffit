# sniffit
A simple naive tool to capture network packets.

# Build instructions
For the first time build, run  
**go mod tidy**  
Followed by,  
**go build -o sniffit main.go**

# Help about the tool  
**sudo ./sniffit**  
Use this tool to capture the network traffic.  
You may first want to list the available devices and then choose the device  
to capture traffic. By default, interface value 'any' and port '11111' will be  
used to capture the network packets and they will be stored in file  
'/tmp/sniffit.pcap' unless different path is specified.  

**Usage:**  
  sniffit [command]  
  
Available Commands:  
  completion  Generate the autocompletion script for the specified shell  
  help        Help about any command  
  **listDevices** A sub-command to list available interfaces  
  **start**       Command to start the capture.  
  
Flags:  
  -h, --help   help for sniffit  

**Use "sniffit [command] --help" for more information about a command.**  

**Example run**  
**sudo ./sniffit start**  
2022/03/18 08:23:43 iface= any filter= port 11111 fileName /tmp/sniffit.pcap  
