package issueUtil

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	//"strings"
)

func openEditor() (string, error) {
	filePath := os.TempDir() + "/gotraining_issue"
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	file.Close()

	cmd := exec.Command("nvim", filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	//s = strings.Replace(string(content), "\n", " ", -1)
	s := string(content)
	return s, nil
}

func InputText() (string, error) {
	message := `# |<----  タイトルを入力します ---->|
# |<----  nvimと入力するとnvimが開き, それ以外の場合その文字列がタイトルとして処理されます ---->|
> `
	fmt.Printf(message)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	if s == "nvim" {
		r, err := openEditor()
		if err != nil {
			return "", err
		}
		s = string(r)
	}
	return s, nil
}
