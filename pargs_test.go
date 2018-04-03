package pargs_test

import (
	"testing"

	"github.com/mcandre/pargs"
)

func TestVersion(t *testing.T) {
	if pargs.Version == "" {
		t.Errorf("Expected pargs version to be non-blank")
	}
}
