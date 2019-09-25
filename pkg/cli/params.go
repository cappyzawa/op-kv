package cli

// OpKvParams describes paramters
type OpKvParams struct {
}

var _ Params = (*OpKvParams)(nil)

// Runner runs ex commands
func (p *OpKvParams) Runner() {

}
