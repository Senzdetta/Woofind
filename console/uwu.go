// https://github.com/Senzdetta/Woofind

package console

import (
    "time"
    "fmt"
    "github.com/Senzdetta/Woofind/utils/cursor"
    "github.com/Senzdetta/Woofind/module/uwu"
)

type UWU struct{}
func (c UWU) Execute(args []string) {
    cursor.Hide()
    uwu.Nyanners(5 * time.Second)
    cursor.Visible()

    fmt.Println()
}

// Copyright (c) 2026 Senzdetta