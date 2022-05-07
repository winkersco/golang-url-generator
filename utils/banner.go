package utils

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

func PrintBanner(phrase string) {
	banner := figure.NewFigure(phrase, "", false)
	banner.Print()
	fmt.Println()
}
