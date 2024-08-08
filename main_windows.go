package main

import "embed"

//go:embed build/windows.txt
var myFiles embed.FS
