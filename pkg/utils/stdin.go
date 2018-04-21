package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StdinQuery(question string) string {

	fmt.Println(question)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	return text
}