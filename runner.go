package opkv

type Runner interface {
	Run(path string, args []string) (string, error)
}
