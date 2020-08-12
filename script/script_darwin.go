package script

import (
	"os/exec"
)

// NewScript is hoge
func NewScript(service string) (Script, error) {
	in := "/mac/app.js"
	out := "cached_app.js"
	path, err := createScriptFile(in, out)
	if err != nil {
		return nil, err
	}

	return AppleScript{
		service: service,
		path:    path,
	}, nil
}

// AppleScript is script
type AppleScript struct {
	service string
	path    string
}

// Exec call osascript
func (a AppleScript) Exec(command string) ([]byte, error) {
	return exec.Command("osascript", a.path, a.service, command).Output()
}
