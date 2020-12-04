package main

import (
	"net"
	"strconv"
	"strings"
)

func setDataconn(reqStr string) (net.Conn, error) {
	address, err := parseAddressReqs(reqStr)
	if err != nil {
		return nil, err
	}
	dataConn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	return dataConn, nil
}

func parseAddressReqs(reqStr string) (string, error) {
	addressReqs := strings.Split(reqStr, ",")
	ipAdd := strings.Join(addressReqs[0:4], ".")
	tmp1, err := strconv.Atoi(addressReqs[4])
	if err != nil {
		return "", err
	}
	tmp2, err := strconv.Atoi(addressReqs[5])
	if err != nil {
		return "", err
	}
	portNum := tmp1*256 + tmp2
	address := ipAdd + ":" + strconv.Itoa(portNum)
	return address, nil
}
