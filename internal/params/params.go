package params

import (
	"github.com/spf13/pflag"
)

type Parameters struct {
	SourceDirectory string
	TargetDirectory string
}

func (p *Parameters) DeclareFlags() {
	pflag.StringVar(&p.SourceDirectory, "source-directory", "", "path to the extracted directory of Google takeout container")
	pflag.StringVar(&p.TargetDirectory, "target-directory", "", "path to the target directory where the edited photos has to be saved to")
}

func (p *Parameters) Obtain() {
	pflag.Parse()
}
