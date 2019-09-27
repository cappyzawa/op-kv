package cli

import "github.com/cappyzawa/op-kv/pkg/helper"

// OpKvParams describes parameters
type OpKvParams struct {
}

var _ Params = (*OpKvParams)(nil)

// Runner runs ex commands
func (p *OpKvParams) Runner(opts ...helper.Opts) helper.Runner {
	return helper.NewRunner(opts...)
}
