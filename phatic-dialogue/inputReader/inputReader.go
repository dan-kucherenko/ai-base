package inputReader

import (
	"bufio"
	"os"
	"strings"
)

func ReadUserInput() string {
	input := ""
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	input = strings.ToLower(input)
	return input
}
