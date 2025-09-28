package mypackages

import (
	"bufio"
	"os"
	"strings"
)

func UserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	userinput := strings.TrimSpace(input)
	return userinput
}
