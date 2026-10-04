// https://github.com/Senzdetta/Woofind

package console

import (
    "github.com/Senzdetta/Woofind/module/checkroot"
)

type Checkroot struct{}
func (c Checkroot) Execute(args []string) {
    checkroot.GetCheck()
}

// Copyright (c) 2026 Senzdetta