package opkv

//go:generate counterfeiter . Runner
type Runner interface {
	Output(commands ...[]string) ([]byte, error)
}

type runner struct {
}

func NewRunner() Runner {
	return &runner{}
}

func (r *runner) Output(commands ...[]string) ([]byte, error) {
	return pipeline.Output(commands...)
}
