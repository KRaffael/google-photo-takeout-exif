package params

type ParameterService interface {
	DeclareFlags()
	Obtain()
	Validate() error

	GetSourceDirectory() string
	GetTargetDirectory() string
}

func New() ParameterService {
	return &Parameters{}
}
