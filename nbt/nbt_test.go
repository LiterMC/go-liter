package nbt_test

import (
	"os"
	"testing"

	"github.com/kmcsr/go-liter/nbt"
)

func TestReadNBT(t *testing.T) {
	type T struct {
		file string
	}
	datas := []T{
		{"testdata/level.dat.1"},
		{"testdata/servers.dat.1"},
		{"testdata/bigtest.nbt"},
	}
	for _, d := range datas {
		resname := d.file + ".txt"
		bts, err := os.ReadFile(d.file)
		if err != nil {
			t.Fatalf("Error when reading file %q: %v", d.file, err)
		}
		nbt, err := nbt.ParseFromBytes(bts)
		if err != nil {
			t.Errorf("Couldn't parse nbt file %q: %v", d.file, err)
			continue
		}
		// os.WriteFile(resname, ([]byte)(nbt.String()), 0644)
		expect, err := os.ReadFile(resname)
		if (string)(expect) != nbt.String() {
			t.Errorf("Couldn't parse nbt file %q correctly", d.file)
		}
	}
}
