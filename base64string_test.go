package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

var testStrings = [...]string{
	"sudo",
	"password",
	"turkeyTrot677!",
	"the ^little^ brown fox()&",
}

// ConvertToBase64 converts the input string to a Base64String
// using the base64 command.
// This is necessary to reproduce the strings generate script's procedure.
func ConvertToBase64String(s string) (Base64String, error) {
	cmd := exec.Command("base64")
	cmd.Stdin = strings.NewReader(s)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return Base64String(out.String()), nil
}

func TestBase64StringDecode(t *testing.T) {
	for _, s := range testStrings {
		b64, err := ConvertToBase64String(s)
		if err != nil {
			t.Fatal(err)
		}
		if b64.String() != s {
			t.Fatalf("b64.String() = '%s', but it should be '%s'\n", b64.String(), s)
		}
	}
}
