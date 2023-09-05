
package nbt

import (
	"strings"
)

func addIndent(str string)(string){
	const indent = "  "
	end := strings.HasPrefix(str, "\n")
	if end {
		str = str[:len(str) - 1]
	}
	str = indent + strings.ReplaceAll(str, "\n", "\n" + indent)
	if end {
		str += "\n"
	}
	return str
}
