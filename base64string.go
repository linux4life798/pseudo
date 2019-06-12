package main

import "encoding/base64"

// Base64String holds a string constant in base64 format.
// It allows easy access to the decoded plain text through
// the String() function.
// This format allows for a small degree of binary obfuscation.
type Base64String string

func NewBase64String(s string) Base64String {
	s = base64.StdEncoding.EncodeToString([]byte(s))
	return Base64String(s)
}

// String returns the decoded plain text
func (s Base64String) String() string {
	plain, _ := base64.StdEncoding.DecodeString(string(s))
	return string(plain)
}
