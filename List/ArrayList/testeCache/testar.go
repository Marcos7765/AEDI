package main

import (
	"lista"
	"time"
)

func testar(n int) int64 {
	var listaTeste lista.ArrayList[int]
	for i := 1; i < n; i++ {
		listaTeste.Add(i)
	}
	t := time.Now()
	listaTeste.AddPos(0, 0)
	return time.Since(t).Microseconds()
}
