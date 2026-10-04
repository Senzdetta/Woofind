// https://github.com/Senzdetta/Woofind

package decode

func looksRot13(s string) bool {
    if len(s) == 0 {
        return false
    }

    for _, c := range s {
        if (c >= 'a' &&
            c <= 'z') || (c >= 'A' &&
            c <= 'Z') {
                return true
        }
    }

    return false
}

// Copyright (c) 2026 Senzdetta