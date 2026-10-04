// https://github.com/Senzdetta/Woofind

package console

import (
    "os"
    "github.com/Senzdetta/Woofind/utils/invinput"
    "github.com/Senzdetta/Woofind/module/dumpstring"
)

type Dumpstring struct{}
func (c Dumpstring) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    dumpstring.Analyzer(args[2])
}

// Copyright (c) 2026 Senzdetta