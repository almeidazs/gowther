package prompts

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

var IS_COLOR_ENABLED = os.Getenv("NO_COLOR") == ""

func c(s, code string) string {
	if !IS_COLOR_ENABLED {
		return s
	}

	return code + s + reset
}

const (
	reset = "\033[0m"
	bold  = "\033[1m"

	purplePastel = "\033[38;5;141m"
	gray         = "\033[90m"
	cyan         = "\033[36m"
	red          = "\033[31m"
)

func Input(label, def string) (string, error) {
	fmt.Printf(
		"%s %s (%s)\n",
		c("?", purplePastel),
		c(label, bold),
		def,
	)

	input, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	value := strings.TrimSpace(input)

	if value == "" {
		return def, nil
	}

	return value, nil
}

func Confirm(label string) (bool, error) {
	fmt.Printf(
		"%s %s %s\n",
		c("?", purplePastel),
		c(label, bold),
		c("(y/N)", gray),
	)

	fmt.Printf("%s ", c(">", cyan))

	for {
		input, err := reader.ReadString('\n')

		if err != nil {
			return false, err
		}

		value := strings.ToLower(strings.TrimSpace(input))

		switch value {
		case "y", "yes":
			return true, nil
		case "", "n", "no":
			return false, nil
		default:
			fmt.Print(c("Please answer y or n: ", red))
		}
	}
}
