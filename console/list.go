// https://github.com/Senzdetta/Woofind

package console

import (
    "os"
    "github.com/Senzdetta/Woofind/module/list"
    "github.com/Senzdetta/Woofind/utils/invinput"
)

type List struct{}
func (c List) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    list.Show(os.Args[2])
}

// Copyright (c) 2026 Senzdetta