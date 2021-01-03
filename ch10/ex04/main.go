package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type depsList struct {
	Name string
	Deps []string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: go run main.go [package]")
		os.Exit(1)
	}
	pkgName := os.Args[1]
	deps := getDeps(pkgName)
	//初期化(指定した物の依存パッケージlist)
	depsSet := make(map[string]struct{})
	for _, pkg := range deps {
		depsSet[pkg] = struct{}{}
	}
	//同じ依存パッケージのものがあるか照合
	pkgAll := getAll()
	scanner := bufio.NewScanner(bytes.NewReader(pkgAll))
	for scanner.Scan() {
		pkgs := strings.Fields(scanner.Text())
		for _, pkg := range pkgs[1:] {
			if _, ok := depsSet[pkg]; ok {
				fmt.Println(pkgs[0])
				break
			}
		}
	}
}

func getDeps(pkgName string) []string {
	args := append([]string{"list", "-json"}, pkgName)
	s, err := exec.Command("go", args...).Output()
	if err != nil {
		os.Exit(1)
	}
	var data depsList
	if err := json.Unmarshal(s, &data); err != nil {
		os.Exit(1)
	}
	return data.Deps
}

func getAll() []byte {
	args := []string{"list", `-f={{.ImportPath}} {{join .Deps " "}}`, "..."}
	out, err := exec.Command("go", args...).Output()
	if err != nil {
		os.Exit(1)
	}
	return out
}
