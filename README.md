# Description

This repo houses the sudo impersonating utility called _pseudo_.
This tool serves to expose a potential attack on a user's password
through the use of bash abstractions, such as `alias sudo=~/bin/pseudo` or `sudo() { ~/bin/pseudo "$@"; }`.

The [pseudo-poc](pseudo-poc) directory contains the most simple
Proof of Concept implementation.

The local directory holds the more obfuscated and convincing version.