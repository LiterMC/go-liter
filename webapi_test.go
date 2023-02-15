
package liter_test

import (
	"testing"

	. "github.com/kmcsr/go-liter"
)

func TestAuthClient_GetPlayerUUID(t *testing.T){
	names := []string{"north", "jeb_", "ckupen"}
	for _, n := range names {
		n, uid, err := DefaultAuthClient.GetPlayerUUID(n)
		if err != nil {
			t.Logf("Cannout get uuid of player '%s': %v", n, err)
			continue
		}
		t.Logf("Uuid of player %s is %v", n, uid)
	}
}
