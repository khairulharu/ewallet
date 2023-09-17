package main

import (
	"github.com/khairulharu/ewallet/internal/component"
	"github.com/khairulharu/ewallet/internal/config"
)

func main() {
	config := config.New()
	component.NewDatabase(config)

}
