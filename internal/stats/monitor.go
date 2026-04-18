package stats

import (
	"os/exec"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

type SystemMood struct {
	CPUUsage   float64
	HasCommits bool
}

func GetCurrentStats() SystemMood {
	c, _ := cpu.Percent(time.Millisecond*500, false)

	// Fancy Check: Look for commits in the last 3 days across ANY local repo
	// (or just the current one if preferred)
	cmd := exec.Command("git", "log", "--all", "--since='3 days ago'", "--oneline")
	out, _ := cmd.Output()

	return SystemMood{
		CPUUsage:   c[0],
		HasCommits: len(strings.TrimSpace(string(out))) > 0,
	}
}
