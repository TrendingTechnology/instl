package internal

import (
	"log"
	"os"
	"path/filepath"

	"github.com/pterm/pterm"
)

func GetInstallPath(username, programName string) string {
	basePath := pterm.Sprintf("/usr/local/instl/%s/%s", username, programName)
	basePath = filepath.Clean(basePath)
	os.MkdirAll(basePath, 0755)

	return basePath
}

// AddToPath adds a value to the global system path environment variable.
func AddToPath(path, filename string) {
	err := os.Symlink(path+"/"+filename, "/usr/local/bin/"+filepath.Base(path))
	if err != nil {
		log.Fatal(err)
	}
}