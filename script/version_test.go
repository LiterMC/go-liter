
package script_test

import (
	"testing"
	. "github.com/kmcsr/go-liter/script"
)

func TestVersion(t *testing.T){
	t.Logf("The current plugin engine version is %s", Version())
}
