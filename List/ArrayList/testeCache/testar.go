package main

import (
	"ArrayList"
	"time"
)

func testar(n int) int64 {
	var listaTeste ArrayList.ArrayList[int]
	for i := 1; i < n; i++ {
		listaTeste.Add(i)
	}
	t := time.Now()
	listaTeste.AddPos(0, 0)
	return time.Since(t).Microseconds()
}
