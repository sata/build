package build

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

func Build(path string) error {
	return sh.RunV("go", "build", fmt.Sprintf("%s/main.go", path))
}
