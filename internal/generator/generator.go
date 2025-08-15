package generator

import (
	"errors"
	"os"
	"path/filepath"

	"embed"

	"github.com/auenc/mcpgen/internal/config"
)

var (
	ErrCreatingOutputDirectory = errors.New("error creating output directory")
	ErrGeneratingServer        = errors.New("error generating server")
	ErrCreatingFile            = errors.New("error creating file")

	//go:embed templates/*
	templates embed.FS
)

func Generate(cfg config.Config) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	basePath := filepath.Join(cwd, cfg.MCPGen.OutputDir, cfg.MCPGen.Package)

	// Ensure output directory exists
	if err := os.MkdirAll(basePath, 0755); err != nil {
		return errors.Join(ErrCreatingOutputDirectory, err)
	}
	err = generateServer(basePath, templates, cfg.MCPGen)

	if err != nil {
		return errors.Join(ErrGeneratingServer, err)
	}
	return nil
}
