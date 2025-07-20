// ffmpeg.go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// A função foi renomeada para refletir o processamento em lote.
func (m model) executeFFmpeg() tea.Cmd {
	return func() tea.Msg {
		// Pega os caminhos do mapa de arquivos selecionados.
		filesToProcess := make([]string, 0, len(m.selectedFiles))
		for path := range m.selectedFiles {
			filesToProcess = append(filesToProcess, path)
		}

		var successCount, errorCount int
		var processedFiles, errorDetails []string

		// Itera sobre cada arquivo selecionado para processamento.
		for _, sourcePath := range filesToProcess {
			var outputPath string
			var err error

			switch m.selectedAction {
			case "convert":
				baseName := strings.TrimSuffix(sourcePath, filepath.Ext(sourcePath))
				outputPath = fmt.Sprintf("%s_converted.%s", baseName, m.targetFormat)
				cmd := exec.Command("ffmpeg", "-i", sourcePath, outputPath, "-y")
				err = cmd.Run()

			case "compress":
				fileName := filepath.Base(sourcePath)
				outputPath = "compressed_" + fileName
				targetSizeInBytes := int64(m.compressionSizeLimitMB * 1024 * 1024)

				// Verifica se o arquivo original já atende ao critério.
				var sourceInfo os.FileInfo
				sourceInfo, err = os.Stat(sourcePath)
				if err == nil && sourceInfo.Size() <= targetSizeInBytes {
					var input []byte
					input, err = os.ReadFile(sourcePath)
					if err == nil {
						err = os.WriteFile(outputPath, input, 0644)
					}
				} else if err == nil {
					// Se o arquivo for maior, tenta a compressão iterativa.
					var compressionSuccess bool
					// Itera da maior qualidade (q=2) para a menor (q=31).
					for q := 2; q <= 31; q++ {
						cmd := exec.Command("ffmpeg", "-i", sourcePath, "-q:v", fmt.Sprint(q), outputPath, "-y")
						runErr := cmd.Run()
						if runErr != nil {
							err = runErr
							break
						}

						var fileInfo os.FileInfo
						fileInfo, err = os.Stat(outputPath)
						if err != nil {
							break
						}

						if fileInfo.Size() <= targetSizeInBytes {
							compressionSuccess = true
							break
						}
					}
					if err == nil && !compressionSuccess {
						err = fmt.Errorf("não foi possível comprimir ao tamanho desejado")
					}
				}
			}

			// Avalia o resultado da operação para este arquivo.
			if err != nil {
				errorCount++
				errorDetails = append(errorDetails, fmt.Sprintf("    - %s: %v", filepath.Base(sourcePath), err))
				_ = os.Remove(outputPath) // Remove arquivo de saída se a operação falhou.
			} else {
				successCount++
				processedFiles = append(processedFiles, outputPath)
			}
		}

		// Monta a mensagem de resumo final.
		var builder strings.Builder
		builder.WriteString(fmt.Sprintf("✅ Processamento em lote concluído! %d com sucesso, %d com erro.\n", successCount, errorCount))
		if successCount > 0 {
			builder.WriteString("\nArquivos gerados:\n")
			for _, file := range processedFiles {
				builder.WriteString(fmt.Sprintf("  - %s\n", file))
			}
		}
		if errorCount > 0 {
			builder.WriteString("\nDetalhes dos erros:\n")
			builder.WriteString(strings.Join(errorDetails, "\n"))
		}

		return processCompleteMsg(builder.String())
	}
}
