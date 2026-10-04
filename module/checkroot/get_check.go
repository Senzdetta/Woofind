// https://github.com/Senzdetta/Woofind

package checkroot

import (
    "fmt"
    "os"
    "github.com/Senzdetta/Woofind/utils/color"
)

func GetCheck() {
    if os.Geteuid() == 0 {
        fmt.Printf(
            "%s[*] %sRunning as %sroot%s\n",
            color.B, color.N, color.GG, color.N,
        )
    } else {
        fmt.Printf(
            "%s[*] %sNot running as %sroot%s\n",
            color.B, color.N, color.GG, color.N,
        )
    }
}

// Copyright (c) 2026 Senzdetta