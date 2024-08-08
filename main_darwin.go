package main

import "embed"

//go:embed build/darwin.txt
var myFiles embed.FS
