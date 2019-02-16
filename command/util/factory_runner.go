package util

import (
	"github.com/cappyzawa/op-kv"
)

type factoryImpl struct {
}

func NewFactory() Factory {
	return &factoryImpl{}
}

func (f *factoryImpl) CommandRunner() opkv.Runner {
	return opkv.NewRunner()
}
