// https://github.com/Senzdetta/Woofind

package decode

func init() {
    for i := range b91DecTable {
        b91DecTable[i] = -1
    }

    for i, b := range b91Alphabet {
        b91DecTable[b] = i
    }
}

// Copyright (c) 2026 Senzdetta