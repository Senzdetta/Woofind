// https://github.com/Senzdetta/Woofind

package console

import (
    "github.com/Senzdetta/Woofind/module/sysinfo"
)

type Sysinfo struct{}
func (c Sysinfo) Execute(args []string) {
    sysinfo.MachineInfo()
}

// Copyright (c) 2026 Senzdetta