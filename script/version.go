
package script

import (
	_ "embed"
	"encoding/json"

	"github.com/kmcsr/semver"
)


type DependMap map[string]semver.ComparatorSet

type packageJsonT struct {
	Name string `json:"name"`
	Version semver.Version `json:"version"`
	Dependencies DependMap `json:"dependencies"`
	DevDependencies DependMap `json:"devDependencies"`
}

//go:embed types/package.json
var _pkgJson []byte
var pkgJson = func()(v *packageJsonT){
	v = new(packageJsonT)
	if err := json.Unmarshal(_pkgJson, v); err != nil {
		panic("liter.script: Cannot parse embedded types/package.json: " + err.Error())
	}
	return
}()

// Version returns the version of liter plugin engine
// This is the version of "go-liter-plugin" when building the package
func Version()(semver.Version){
	return pkgJson.Version
}

func Dependencies()(DependMap){
	return pkgJson.Dependencies
}
