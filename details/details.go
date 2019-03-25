package details

import (
	info "github.com/Helcaraxan/moddep-direct-dep"
	printer "github.com/Helcaraxan/moddep-indirect-dep"
)

const data = "test-content"

func PrintDefault() {
	printer.Print(&info.Info{Content: data})
}
