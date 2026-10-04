// https://github.com/Senzdetta/Woofind

package main

import (
    "os"
    "strings"
    "github.com/Senzdetta/Woofind/console"
)

func main() {
    args := os.Args[1:]
    input := strings.Join(args, " ")
    console.WoofindConsole(input)
}

// Copyright (c) 2026 Senzdetta