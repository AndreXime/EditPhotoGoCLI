// main.go
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	imagePaths := findImageFiles(".")
	if len(imagePaths) == 0 {
		fmt.Println("Nenhuma imagem encontrada no diretório atual.")
		return
	}

	model := newModel(imagePaths)
	program := tea.NewProgram(model)

	if _, err := program.Run(); err != nil {
		fmt.Printf("Ocorreu um erro ao executar o programa: %v\n", err)
		os.Exit(1)
	}
}

func findImageFiles(rootDirectory string) []string {
	var filePaths []string
	_ = filepath.Walk(rootDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // Ignora arquivos que não podem ser lidos.
		}
		if !info.IsDir() {
			extension := strings.ToLower(filepath.Ext(path))
			switch extension {
			case ".jpg", ".jpeg", ".png", ".webp":
				filePaths = append(filePaths, path)
			}
		}
		return nil
	})
	return filePaths
}
