// Package question holds functions for questions to the user.
package question

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// YesNo takes keyboard input for a decision question and returns true or false.
func YesNo(message string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s [Y|n] ", message)
	decision, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	decision = strings.TrimSpace(decision)

	if decision == "y" || decision == "yes" || decision == "Y" {
		return true
	}
	return false
}
