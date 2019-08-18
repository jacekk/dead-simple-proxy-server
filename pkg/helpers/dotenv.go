package helpers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

const (
	customEnvsFile  = ".env"
	defaultEnvsFile = "dist.env"
)

// GetProjectDir - this very project root directory
func GetProjectDir() string {
	_, fileName, _, isOk := runtime.Caller(0)

	if !isOk {
		log.Fatal("Runtime Caller could not be recovered.")
	}

	projectDir := filepath.Join(fileName, "../../../")

	return projectDir
}

// LoadEnvs - loads contents of '.env' and 'dist.env' files,
// and overrides each system ENV which has not been defined.
func LoadEnvs(projectDir string) {
	// Gitignored. Has precedence over `dist.env`, but NOT over system ENVs.
	err := godotenv.Load(filepath.Join(projectDir, customEnvsFile))
	if err != nil {
		fmt.Fprintf(os.Stdout, "File '%s' has NOT been loaded. \n", customEnvsFile)
	} else {
		fmt.Fprintf(os.Stdout, "File '%s' has been loaded. \n", customEnvsFile)
	}

	// Defaults. Under VCS.
	err = godotenv.Load(filepath.Join(projectDir, defaultEnvsFile))
	if err != nil {
		log.Fatalf("File '%s' is not present. It should contain all used ENVs, which could be overwritten by system ENVs.", defaultEnvsFile)
	} else {
		fmt.Fprintf(os.Stdout, "File '%s' has been loaded. \n", defaultEnvsFile)
	}
}
