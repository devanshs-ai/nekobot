package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/devanshs-ai/nekobotic/internal/ascii"
	"github.com/devanshs-ai/nekobotic/internal/stats"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	s := stats.GetCurrentStats()

	messages := []string{
		"Watching the threads...",
		"Scanning for vulnerabilities...",
		"Memory looks clean, Devansh.",
		"Ready for the next commit?",
		"The Net is vast and infinite...",
		"Cyber-brain sync: 98%.",
	}

	msg := messages[rand.Intn(len(messages))]

	// The "Fancy" Layout
	fmt.Println(ascii.GetNeko(s.CPUUsage, s.RAMUsage, s.HasCommits, s.Uptime))
	fmt.Printf("%s───[ %s ]───────────────────────────────────%s\n", ascii.Teal, msg, ascii.Reset)
	fmt.Printf("%sCPU: %.1f%% | RAM: %.1f%% | UPTIME: %dh%s\n", ascii.Dim, s.CPUUsage, s.RAMUsage, s.Uptime, ascii.Reset)
}
