package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Welcome message
	fmt.Println("Git commiter")

	// Quick way to check if the arguments are there
	if len(os.Args) == 4 {
		processByArgs(os.Args)
	} else {
		processByHand()
	}
}

// Takes values from the command line arguments
func processByArgs(args []string) {
	title := args[1]
	what := args[2]
	why := args[3]

	commit(title, what, why)
}

// Takes values by asking them manually
func processByHand() {
	title := read("Title")
	what := read("What did you do?")
	why := read("Why did you do it?")

	commit(title, what, why)
}

// Reads a value from the standard input
func read(label string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(label + ": ")
	text, _ := reader.ReadString('\n')

	// Convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)

	return text
}

// Makes a git commit with the given values
func commit(title, what, why string) {
	command := fmt.Sprintf(`"%v | What was made: %v | Why was made: %v"`, title, what, why)

	// Adds all
	output, err := exec.Command("git", "add", ".").CombinedOutput()

	// Checks for errors
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}

	// Commits
	output, err = exec.Command("git", "commit", "-m", command).CombinedOutput()

	// Checks again
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}

	fmt.Println(string(output))
}
