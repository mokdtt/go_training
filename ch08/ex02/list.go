package main

import (
	"io"
	"net"
	"os/exec"
	"strings"
)

func execLs(reqs []string, currentDir string, dataConn net.Conn) error {
	defer dataConn.Close()
	params := []string{"-la", currentDir}
	/*
		if len(reqs) != 1 {
			params = append(params, reqs[1:]...)
		}
	*/
	lscmd := exec.Command("ls", params...)
	out, err := lscmd.CombinedOutput()
	if err != nil {
		return err
	}
	s := strings.ReplaceAll(string(out), "\n", "\r\n")
	io.WriteString(dataConn, s)
	return nil
}
