package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Input(label string) string {
	var output string 
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stderr, label + " ")
		output, _ = r.ReadString('\n')

		if output != "" {
			break
		}
	}
	return strings.TrimSpace(output)
}