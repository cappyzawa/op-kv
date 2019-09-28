package helper_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/cappyzawa/op-kv/pkg/helper"
)

func TestPrinterPair(t *testing.T) {
	cases := []struct {
		name     string
		args     []string
		expect   string
		existErr bool
	}{
		{
			name:     "with username & password",
			args:     []string{"username", "password"},
			expect:   fmt.Sprintf("| %-20s| %-60s|\n", "username", "password"),
			existErr: false,
		},
		{
			name:     "with looooong username & password",
			args:     []string{"looooooooooooooooooooooooooooooooooogusername", "password"},
			expect:   fmt.Sprintf("| %-20s| %-60s|\n", "looooooooooooooooooooooooooooooooooogusername", "password"),
			existErr: false,
		},
	}
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			outStream := new(bytes.Buffer)
			p := helper.NewPrinter(helper.PrinterOut(outStream))
			actualErr := p.Pair(c.args[0], c.args[1])
			if actualErr != nil && !c.existErr {
				t.Errorf("error should not be occurred, but actual is %v", actualErr)
			}
			if actualErr == nil && c.existErr {
				t.Error("error should be occurred")
			}
			if outStream.String() != c.expect {
				t.Errorf("output should be %s, but actual is %v", c.expect, outStream.String())
			}
		})
	}
}

func TestPrinterHeader(t *testing.T) {
	outStream := new(bytes.Buffer)
	p := helper.NewPrinter(helper.PrinterOut(outStream))
	expectDeli := "-------------------------------------------------------------------------------------"
	expectStr := fmt.Sprintf("| %-20s| %-60s|\n%s\n", "USERNAME", "PASSWORD", expectDeli)
	actual := p.Header()
	var expectErr error = nil
	if actual != expectErr {
		t.Errorf("error should not be occurred, but actual is %v", actual)
	}
	if outStream.String() != expectStr {
		t.Errorf("output should be %s, but actual is %s", expectStr, outStream.String())
	}
}
