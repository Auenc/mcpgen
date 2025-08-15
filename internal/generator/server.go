package generator

import (
	"embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/auenc/mcpgen/internal/config"
)

// generateServer uses server-setup.tmpl and writes to internal/mcp/{{.Package}}/server.go
func generateServer(basePath string, templates embed.FS, cfg config.Config) error {
	outPath := filepath.Join(basePath, "server.go")

	tmplContent, err := templates.ReadFile("templates/server-setup.tmpl")
	if err != nil {
		return err
	}
	tmpl := template.New("server-setup").Funcs(templateFuncMap)
	tmpl, err = tmpl.Parse(string(tmplContent))
	if err != nil {
		return errors.Join(ErrCreatingFile, err)
	}

	fmt.Printf("CREATE %s\n", outPath)
	if cfg.DryRun {
		return nil
	}

	// Create output file
	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()
	// Execute template with config
	return tmpl.Execute(f, cfg.MCPGen)
}
