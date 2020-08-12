package script

import (
	"errors"
	"os/exec"
)

func NewScript(service string) (Script, error) {
	return WindowsScript{}, errors.New("Unsupported OS.")
}

// AppleScript is script
type WindowsScript struct {
	service string
	path    string
}

// Exec call osascript
func (a WindowsScript) Exec(command string) ([]byte, error) {
	return exec.Command("cscript -e", a.path, command).Output()
}
