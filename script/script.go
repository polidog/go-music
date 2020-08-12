package script

import (
	"io/ioutil"
	"os"
	"path/filepath"

	_ "go-music/script/statik"

	"github.com/rakyll/statik/fs"
)

// Script interface
type Script interface {
	Exec(command string) ([]byte, error)
}

func createScriptFile(in, out string) (string, error) {
	dir := os.TempDir()
	path := filepath.Join(dir, out)
	if exists(path) {
		return path, nil
	}

	fileSystem, err := fs.New()
	if err != nil {
		return "", err
	}

	file, err2 := fileSystem.Open(in)
	if err2 != nil {
		return "", err2
	}
	defer file.Close()

	contents, readErr := ioutil.ReadAll(file)
	if readErr != nil {
		return "", readErr
	}

	err = ioutil.WriteFile(path, contents, 0777)
	if err != nil {
		return "", err
	}
	return path, nil
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
