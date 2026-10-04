// https://github.com/Senzdetta/Woofind

package console

import (
    "github.com/Senzdetta/Woofind/module/version"
)

type Version struct{}
func (c Version) Execute(args []string) {
    version.WoofindVersion()
}

// Copyright (c) 2026 Senzdetta