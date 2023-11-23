package liter_test

import (
	"testing"

	"github.com/google/uuid"
	. "github.com/kmcsr/go-liter"
)

func TestAuthClient_GetPlayerUUID(t *testing.T) {
	names := []string{"north", "jeb_", "ckupen"}
	for _, n := range names {
		info, err := DefaultAuthClient.GetPlayerInfo(n)
		if err != nil {
			t.Logf("Cannot get uuid of player '%s': %v", n, err)
			continue
		}
		t.Logf("Uuid of player %s is %s[%v]", n, info.Name, info.Id)
	}
}

func TestAuthClient_GetPlayerProfile(t *testing.T) {
	uuids := []string{"79c839d8-f837-40a6-a519-e46aa1b029ee", "853c80ef-3c37-49fd-aa49-938b674adae6", "7a0ba4fe-e6ec-4bfe-99fc-56bf677a15a5"}
	for _, n := range uuids {
		profile, err := DefaultAuthClient.GetPlayerProfile(uuid.MustParse(n))
		if err != nil {
			t.Logf("Cannot get profile of player '%s': %v", n, err)
			continue
		}
		t.Logf("Profile of player %s is %s[%v]", n, profile.Name, profile.Id)
	}
}
