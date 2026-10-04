// https://github.com/Senzdetta/Woofind

package owner

import (
    "sync"
)

func worker(
    jobs <-chan string,
    wg *sync.WaitGroup,
    targetUIDs map[uint32]string,
    mu *sync.Mutex,
) {
    defer wg.Done()
    for filePath := range jobs {
        checkOwner(
            filePath, targetUIDs,
            mu,
        )
    }
}

// Copyright (c) 2026 Senzdetta