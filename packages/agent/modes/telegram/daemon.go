package telegram

import (
	"path/filepath"
	"time"

	"github.com/nlf/ncode/packages/agent/modes/bot"
)

// PIDPath returns the location of the bot's pid file.
func PIDPath(ncodeHome string) string { return filepath.Join(ncodeHome, "bot.pid") }

// LogPath returns the location of the bot's log file (stdout+stderr
// from a detached `ncode bot start`).
func LogPath(ncodeHome string) string { return filepath.Join(ncodeHome, "logs", "bot.log") }

// WritePID persists pid to bot.pid. Overwrites any existing file.
func WritePID(ncodeHome string, pid int) error { return bot.WritePIDFile(PIDPath(ncodeHome), pid) }

// ReadPID returns the pid stored in bot.pid, or 0 if the file doesn't
// exist. Returns an error for any other read/parse failure.
func ReadPID(ncodeHome string) (int, error) { return bot.ReadPIDFile(PIDPath(ncodeHome)) }

// RemovePID deletes the pid file if it exists.
func RemovePID(ncodeHome string) error { return bot.RemovePIDFile(PIDPath(ncodeHome)) }

// IsRunning returns (pid, true) if a live process with the recorded
// pid exists, or (pid, false) if the pid file points to a dead process.
// Stale pid files are left in place; the caller may remove them.
func IsRunning(ncodeHome string) (int, bool, error) { return bot.IsRunningAt(PIDPath(ncodeHome)) }

// StopProcess asks pid to exit and waits up to graceful for it to stop,
// then escalates to a forced kill. Returns nil if the process is gone.
func StopProcess(pid int, graceful time.Duration) error { return bot.StopProcess(pid, graceful) }
