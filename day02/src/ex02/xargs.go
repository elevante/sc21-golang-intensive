package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func myXargs() {
	scanner := bufio.NewScanner(os.Stdin)

	args := []string{"ls", "-la"}

	for scanner.Scan() {
		arg := scanner.Text()
		args = append(args, filepath.Join(".", arg))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading standard input:", err)
		return
	}
	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = "."

	if err := cmd.Run(); err != nil {
		fmt.Println("Command execution error:", err)
		return
	}
}

func main() {
	myXargs()
}
