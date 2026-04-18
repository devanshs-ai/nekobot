package main

import (
	"fmt"

	"github.com/devanshs-ai/nekobotic/internal/ascii"
	"github.com/devanshs-ai/nekobotic/internal/stats"
)

func main() {
	s := stats.GetCurrentStats()
	fmt.Println(ascii.GetNeko(s.CPUUsage, s.HasCommits))
}
