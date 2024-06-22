package params

import (
	"fmt"
	"log"
	"regexp"
)

func (p *Parameters) Validate() error {
	ok := validateSourceDirectory(p.SourceDirectory)
	ok = validateTargetDirectory(p.TargetDirectory) && ok

	if ok {
		return nil
	} else {
		return fmt.Errorf("there were parameter errors")
	}
}

var windowsPathRegex = regexp.MustCompile("(([a-zA-Z]:)|(..))\\\\(?:([^<>:\"\\/\\\\|?*]*[^<>:\"\\/\\\\|?*.]\\\\|..\\\\)*([^<>:\"\\/\\\\|?*]*[^<>:\"\\/\\\\|?*.]\\\\?|..\\\\))?")
var linuxPathRegex = regexp.MustCompile("^([.]+?/[^/ ]*)+/?$")

func validateTargetDirectory(value string) bool {
	if value == "" {
		log.Printf("ERROR: Parameter --target-directory=<xyz> is mandatory")
		return false
	}
	if windowsPathRegex.MatchString(value) || linuxPathRegex.MatchString(value) {
		log.Printf("using target-directory '%s'\n", value)
		return true
	}

	log.Printf("ERROR: Parameter --target-directory=<xyz> must be a windows or linux path")
	return false
}

func validateSourceDirectory(value string) bool {
	if value == "" {
		log.Printf("ERROR: Parameter --source-directory=<xyz> is mandatory")
		return false
	}
	if windowsPathRegex.MatchString(value) || linuxPathRegex.MatchString(value) {
		log.Printf("using source-directory '%s'\n", value)
		return true
	}

	log.Printf("ERROR: Parameter --source-directory=<xyz> must be a windows or linux path")
	return false
}
