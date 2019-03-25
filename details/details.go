package details

import (
	info "github.com/Helcaraxan/moddep-direct-dep"
	printer "github.com/Helcaraxan/moddep-indirect-dep"
)

func PrintDefaultInfo() {
	printer.Print(info.DefaultInfo())
}

func PrintInfo(i *info.Info) {
	printer.Print(i)
}
