package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/egoist/mygo"
)

// settings is what the app remembers between runs: the appearance, the
// topic filter and how many questions a quiz asks. The window's bounds are
// remembered by MyGo, through the window's StateKey.
type settings struct {
	Appearance string `json:"appearance"`
	Topic      string `json:"topic"`
	Limit      int    `json:"limit"`
}

// settingsFile is the settings file's name inside the app's data directory.
const settingsFile = "settings.json"

// loadSettings reads the remembered settings. A missing or unreadable file
// is not an error: it means the defaults, which is what a first run should
// show.
func loadSettings() settings {
	s := settings{Topic: "All", Limit: 10}
	path, err := settingsPath()
	if err != nil {
		return s
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return s
	}
	_ = json.Unmarshal(data, &s)
	if s.Limit <= 0 {
		s.Limit = 10
	}
	if s.Topic == "" {
		s.Topic = "All"
	}
	return s
}

// saveSettings writes the settings. A failure is ignored on purpose: the
// choice still applies to this run, and losing it on the next start is not
// worth interrupting the learner over.
func (a *app) saveSettings() {
	path, err := settingsPath()
	if err != nil {
		return
	}
	data, err := json.Marshal(settings{
		Appearance: a.appearance.String(),
		Topic:      a.topic,
		Limit:      a.limitCount(),
	})
	if err != nil {
		return
	}
	_ = os.MkdirAll(filepath.Dir(path), 0o755)
	_ = os.WriteFile(path, data, 0o644)
}

// settingsPath is where the settings live: the app's per-user data
// directory, beside the window state MyGo keeps.
func settingsPath() (string, error) {
	dir, err := mygo.App.Path(mygo.PathUserData)
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, settingsFile), nil
}
