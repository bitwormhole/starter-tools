package demo

import (
	"testing"

	"github.com/bitwormhole/starter-tools/tools/configen1"
)

func TestConfigen(t *testing.T) {
	args := []string{}
	err := configen1.Run(args)
	if err != nil {
		t.Fatal(err)
	}
}
