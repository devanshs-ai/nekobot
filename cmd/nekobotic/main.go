package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/devanshs-ai/nekobotic/internal/ascii"
	"github.com/devanshs-ai/nekobotic/internal/stats"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 1. Get System Stats
	s := stats.GetCurrentStats()

	// 2. Intelligence Logic: Determine the "Culprit" based on system stress
	var topProc stats.ProcessInfo
	var xaiInsight string

	// XAI Thresholds (We'll move these to YAML soon!)
	if s.RAMUsage > 75 {
		// If RAM is high, we specifically look for the Memory hog
		topProc = stats.GetTopProcess(true)
		xaiInsight = fmt.Sprintf("High RAM usage detected. Culprit: %s", topProc.Name)
	} else if s.CPUUsage > 50 {
		// If CPU is high, we look for the CPU hog
		topProc = stats.GetTopProcess(false)
		xaiInsight = fmt.Sprintf("CRITICAL: %s is dominating CPU cycles.", topProc.Name)
	} else {
		// Otherwise, default to standard background monitoring
		topProc = stats.GetTopProcess(false)
		xaiInsight = fmt.Sprintf("System Balanced. Background: %s", topProc.Name)
	}

	// 3. Get Project Context & Tracking
	cwd, _ := os.Getwd()
	projectName := filepath.Base(cwd)
	activity := stats.UpdateActivity(projectName)

	// 4. Random Greeting Messages
	messages := []string{
		"NEURAL LINK ESTABLISHED",
		"MONITORING SUB-PROCESSES...",
		"GHOST IN THE SHELL ACTIVE",
		"SYNC RATIO: STABLE",
	}
	msg := messages[rand.Intn(len(messages))]

	// 5. Rendering the UI
	fmt.Print(ascii.GetNeko(s.CPUUsage, s.RAMUsage, s.HasCommits, s.Uptime))

	// UI Divider
	fmt.Printf("%s───[ %s ]───────────────────────────────────%s\n", ascii.Teal, msg, ascii.Reset)

	// Diagnostic Line (Now accurately reflecting the culprit)
	fmt.Printf("%sDIAGNOSTIC: %s%s\n", ascii.Dim, xaiInsight, ascii.Reset)

	// Research Mode Footer
	fmt.Printf("%sPROJECT: %s%s | %sSESSION: %dm%s | %sTOTAL SYNC: %dm%s\n",
		ascii.Pink, projectName, ascii.Reset,
		ascii.Teal, int(time.Since(activity.StartTime).Minutes()), ascii.Reset,
		ascii.Pink, activity.TotalMinutes, ascii.Reset)
}
