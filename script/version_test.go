package script_test

import (
	. "github.com/kmcsr/go-liter/script"
	"testing"
)

func TestVersion(t *testing.T) {
	t.Logf("The current plugin engine version is %s", Version())
}
