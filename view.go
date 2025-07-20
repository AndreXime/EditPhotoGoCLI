// view.go (Refatorado)
package main

import (
	"fmt"
	"strings"
)

func (m model) View() string {
	var builder strings.Builder

	switch m.currentState {
	case choosingFileState:
		builder.WriteString(m.styles.Header.Render("📂 Selecione uma imagem para processar:") + "\n")
		for i, path := range m.imagePaths {
			// Determina o prefixo: cursor e checkbox
			cursor := "  " // 2 espaços para alinhamento
			if i == m.cursorIndex {
				cursor = "> "
			}

			checkbox := "[ ]" // Checkbox não marcado
			if _, ok := m.selectedFiles[path]; ok {
				checkbox = "[x]" // Checkbox marcado
			}

			// Monta a linha
			line := fmt.Sprintf("%s%s %s", cursor, checkbox, path)

			// Aplica estilo se a linha estiver selecionada (marcada)
			if _, ok := m.selectedFiles[path]; ok {
				builder.WriteString(m.styles.Selected.Render(line) + "\n")
			} else {
				builder.WriteString(line + "\n")
			}
		}
		builder.WriteString(m.styles.Help.Render("\n↑ ↓ para navegar | espaço para marcar/desmarca r | Enter para confirmar | q para sair"))

	case processingState:
		builder.WriteString(m.styles.Header.Render(fmt.Sprintf("⚙️ Processando os %d itens, por favor aguarde...", len(m.selectedFiles))))
		builder.WriteString(m.styles.Help.Render("\nIsso pode levar alguns instantes dependendo do tamanho da imagem."))

	case choosingActionState:
		header := fmt.Sprintf("🖼️ %d arquivos selecionados", len(m.selectedFiles))
		builder.WriteString(m.styles.Header.Render(header))
		builder.WriteString("\n\nEscolha uma ação:\n")

		actions := []string{"Converter formato", "Comprimir para um tamanho máximo (MB)"}
		for i, action := range actions {
			if i == m.cursorIndex {
				builder.WriteString(m.styles.Selected.Render("> "+action) + "\n")
			} else {
				builder.WriteString("  " + action + "\n")
			}
		}
		builder.WriteString(m.styles.Help.Render("\n↑ ↓ para navegar | Enter para confirmar"))

	case choosingFormatState:
		builder.WriteString(m.styles.Header.Render("🎯 Escolha o formato de destino:") + "\n")
		for i, format := range supportedFormats {
			if i == m.cursorIndex {
				builder.WriteString(m.styles.Selected.Render("> "+format) + "\n")
			} else {
				builder.WriteString("  " + format + "\n")
			}
		}
		builder.WriteString(m.styles.Help.Render("\n↑ ↓ para navegar | Enter para confirmar"))

	case enteringSizeLimitState:
		if m.sizeInputError != "" {
			builder.WriteString(m.styles.ErrorHeader.Render(m.sizeInputError) + "\n\n")
		}
		prompt := "📉 Digite o tamanho máximo em MB (ex: 1.5) e pressione Enter:\n"
		builder.WriteString(prompt)
		builder.WriteString(m.styles.Selected.Render("> " + m.sizeInputBuffer))

	case doneState:
		style := m.styles.SuccessMsg
		if strings.HasPrefix(m.resultMessage, "❌") {
			style = m.styles.ErrorMsg
		}
		builder.WriteString(style.Render(m.resultMessage))
		builder.WriteString(m.styles.Help.Render("\n\nPressione Enter para sair."))

	default:
		builder.WriteString(m.styles.ErrorMsg.Render("Erro: Estado desconhecido da aplicação."))
	}

	return builder.String()
}
