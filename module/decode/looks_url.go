// https://github.com/Senzdetta/Woofind

package decode

import (
    "strings"
)

func looksURL(s string) bool {
    return strings.Contains(s, "%")
}

// Copyright (c) 2026 Senzdetta