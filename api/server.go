package api

import (
	"fmt"
	"net"
)

func StartSysLogUDPServer(port int) error {

	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: 514})
	if err != nil {
		return fmt.Errorf("Error starting UDP server: %w\n", err)
	}
	defer conn.Close()
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return fmt.Errorf("Error resolving TCP address: %w\n", err)
	}
	forwardPort, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return fmt.Errorf("Error dialing TCP: %w\n", err)
	}
	defer forwardPort.Close()

	buffer := make([]byte, 1024)

	for {
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			return fmt.Errorf("Error reading from UDP: %w\n", err)
		}
		_, err = forwardPort.Write(buffer[:n])
		if err != nil {
			return fmt.Errorf("Error writing to TCP: %w\n", err)
		}
	}
}
