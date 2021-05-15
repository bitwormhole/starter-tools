package demo

import (
	"testing"

	"github.com/bitwormhole/starter-tools/tools/configen"
)

func TestConfigen(t *testing.T) {
	args := []string{}
	err := configen.Run(args)
	if err != nil {
		t.Fatal(err)
	}
}
