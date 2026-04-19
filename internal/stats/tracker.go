package stats

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type TrackingData struct {
	ProjectName  string    `json:"project_name"`
	StartTime    time.Time `json:"start_time"`
	LastSeen     time.Time `json:"last_seen"`
	TotalMinutes int       `json:"total_minutes"`
}

func GetTrackingFilePath() string {
	home, _ := os.UserHomeDir()
	dir := filepath.Join(home, ".config", "nekobotic")
	os.MkdirAll(dir, 0755)
	return filepath.Join(dir, "activity.json")
}

func UpdateActivity(currentProject string) TrackingData {
	path := GetTrackingFilePath()
	data := TrackingData{ProjectName: currentProject, StartTime: time.Now(), LastSeen: time.Now()}

	// Load existing data if it exists
	file, err := os.ReadFile(path)
	if err == nil {
		json.Unmarshal(file, &data)
	}

	now := time.Now()

	// If it's the same project and the last seen was less than 30 mins ago,
	// we consider it the same session.
	if data.ProjectName == currentProject && now.Sub(data.LastSeen).Minutes() < 30 {
		duration := now.Sub(data.LastSeen).Minutes()
		data.TotalMinutes += int(duration)
	} else {
		// New session or new project
		data.ProjectName = currentProject
		data.StartTime = now
	}

	data.LastSeen = now

	// Save back to file
	newData, _ := json.Marshal(data)
	os.WriteFile(path, newData, 0644)

	return data
}
