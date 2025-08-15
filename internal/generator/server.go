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

//go:embed templates/server-setup.tmpl
var serverSetupTmpl embed.FS

// generateServer uses server-setup.tmpl and writes to internal/mcp/{{.Package}}/server.go
func generateServer(basePath string, templates embed.FS, cfg config.ProjectConfig) error {
	fmt.Println("generating server setup")
	outPath := filepath.Join(basePath, "server.go")

	fmt.Println("parsing embedded template")
	tmplContent, err := serverSetupTmpl.ReadFile("templates/server-setup.tmpl")
	if err != nil {
		return err
	}
	tmpl, err := template.New("server-setup").Parse(string(tmplContent))
	if err != nil {
		return errors.Join(ErrCreatingFile, err)
	}

	// Create output file
	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Println("writing file")
	// Execute template with config
	return tmpl.Execute(f, cfg)
}
