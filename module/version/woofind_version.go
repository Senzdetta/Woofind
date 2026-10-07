// https://github.com/Senzdetta/Woofind

package version

import (
    "fmt"
    "github.com/Senzdetta/Woofind/utils/color"
)

const (
    name = "Woofind"
    version = "v0.1.20261008"
    developer = "Senzdetta"
    homepage = "https://github.com/Senzdetta/Woofind"
)

func WoofindVersion() {
    fmt.Printf(
        "%s- %s%s %s-%s\n",
        color.DG, color.GG, name, color.DG, color.N,
    )

    fmt.Printf(
        "%sVersion: %s%s%s\n",
        color.N, color.GG, version, color.N,
    )

    fmt.Printf(
        "%sDeveloper: %s%s%s\n",
        color.N, color.GG, developer, color.N,
    )

    fmt.Printf(
        "%sHomepage: %s%s%s\n",
        color.N, color.GG, homepage, color.N,
    )
}

// Copyright (c) 2026 Senzdetta