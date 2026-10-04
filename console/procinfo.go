// https://github.com/Senzdetta/Woofind

package console

import (
    "github.com/Senzdetta/Woofind/module/procinfo"
)

type Procinfo struct{}
func (c Procinfo) Execute(args []string) {
    procinfo.ShowProcInfo()
}

// Copyright (c) 2026 Senzdetta