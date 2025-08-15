package generator

import (
	"fmt"

	"github.com/auenc/mcpgen/internal/config"
)

func Generate(cfg config.Config) {
	fmt.Printf("we should do something: %v\n", !cfg.DryRun)
}
