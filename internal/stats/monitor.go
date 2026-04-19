package stats

import (
	"os/exec"
	"strings"
	"time"

	"sort"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
)

type ProcessInfo struct {
	Name  string
	Usage float64
}

func GetTopProcess(byMemory bool) ProcessInfo {
	processes, _ := process.Processes()
	var processList []ProcessInfo

	for _, p := range processes {
		name, _ := p.Name()
		var usage float64
		if byMemory {
			m, _ := p.MemoryPercent()
			usage = float64(m)
		} else {
			usage, _ = p.CPUPercent()
		}

		if name != "" && name != "idle" && !strings.Contains(name, "nekobotic") {
			processList = append(processList, ProcessInfo{Name: name, Usage: usage})
		}
	}

	sort.Slice(processList, func(i, j int) bool {
		return processList[i].Usage > processList[j].Usage
	})

	if len(processList) > 0 {
		return processList[0]
	}
	return ProcessInfo{Name: "System", Usage: 0}
}

type SystemMood struct {
	CPUUsage   float64
	RAMUsage   float64
	NetActive  bool
	OnBattery  bool
	HasCommits bool
	Uptime     uint64
}

func GetCurrentStats() SystemMood {
	c, _ := cpu.Percent(time.Millisecond*50, false)
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
