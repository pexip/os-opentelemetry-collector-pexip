package gofipschecker

import (
	"crypto/fips140"
	"os"
	"path"
	"strings"
	"syscall"
)

func init() {
	// Should FIPS be enabled according to the system?
	if _, err := os.Stat("/etc/fips/enabled"); err == nil {
		// Need to have FIPS enabled in crypto
		if !fips140.Enabled() {
			ex, err := os.Executable()
			if err != nil {
				panic("Error finding executable path")
			}
			name := path.Base(ex)

			if !strings.HasSuffix(name, ".fips") {
				fipsbin := ex + ".fips"
				if _, err := os.Stat(fipsbin); err == nil {

					// Try execing, we'll either exec or fall through to the generic error
					_ = syscall.Exec(fipsbin, os.Args, os.Environ())
				}
			}

			// Our default is that we failed to start properly
			panic("FIPS should be enabled and is not")
		}
	}
}
