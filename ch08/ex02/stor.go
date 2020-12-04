package main

import (
	"io/ioutil"
	"net"
	"path/filepath"
)

func putFile(filename string, currentDir string, dataConn net.Conn) error {
	defer dataConn.Close()
	buf, err := ioutil.ReadAll(dataConn)
	if err != nil {
		return err
	}

	filePath := filepath.Join(currentDir, filename)
	ioutil.WriteFile(filePath, buf, 0775)
	return nil
}
