package main

import (
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
)

func getFile(filename string, currentDir string, dataConn net.Conn) error {
	defer dataConn.Close()
	filePath := filepath.Join(currentDir, filename)
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	_, err = dataConn.Write(buf)
	if err != nil {
		return err
	}
	return nil
}
