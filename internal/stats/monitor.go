package stats

import (
	"os/exec"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

type SystemMood struct {
	CPUUsage   float64
	RAMUsage   float64
	NetActive  bool
	OnBattery  bool
	HasCommits bool
	Uptime     uint64
}

func GetCurrentStats() SystemMood {
	c, _ := cpu.Percent(time.Millisecond*200, false)
	m, _ := mem.VirtualMemory()
	n, _ := net.IOCounters(false)
	h, _ := host.Info()

	// Check Git
	cmd := exec.Command("git", "log", "--all", "--since='3 days ago'", "--oneline")
	out, _ := cmd.Output()

	return SystemMood{
		CPUUsage:   c[0],
		RAMUsage:   m.UsedPercent,
		NetActive:  n[0].BytesSent > 0, // Simplified check
		HasCommits: len(strings.TrimSpace(string(out))) > 0,
		Uptime:     h.Uptime / 3600, // Uptime in hours
	}
}
