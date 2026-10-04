// https://github.com/Senzdetta/Woofind

package decode

func looksMorse(s string) bool {
    if len(s) == 0 {
        return false
    }

    hasMorseChars := false
    for _, c := range s {
        if c == '.' || c == '-' {
            hasMorseChars = true
        } else if c != ' ' && c != '/' {
            return false
        }
    }

    return hasMorseChars
}

// Copyright (c) 2026 Senzdetta