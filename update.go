// update.go
package main

import (
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case processCompleteMsg:
		m.resultMessage = string(msg)
		m.currentState = doneState
		return m, tea.Quit

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	switch m.currentState {
	case choosingFileState:
		return m.handleFileSelection(msg)
	case choosingActionState:
		return m.handleActionSelection(msg)
	case choosingFormatState:
		return m.handleFormatSelection(msg)
	case enteringSizeLimitState:
		return m.handleSizeInput(msg)
	case processingState:
		// Nenhum comando de teclado é processado enquanto o FFmpeg está rodando.
		return m, nil
	case doneState:
		return m.handleCompletion(msg)
	}

	return m, nil
}

// As funções de 'update' foram renomeadas para 'handle' para melhor semântica.
func (m model) handleFileSelection(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.cursorIndex > 0 {
				m.cursorIndex--
			}
		case "down":
			if m.cursorIndex < len(m.imagePaths)-1 {
				m.cursorIndex++
			}

		case " ":
			path := m.imagePaths[m.cursorIndex]
			if _, ok := m.selectedFiles[path]; ok {
				// Se já estiver selecionado, remova-o
				delete(m.selectedFiles, path)
			} else {
				// Caso contrário, adicione-o
				m.selectedFiles[path] = struct{}{}
			}

		case "enter":
			if len(m.selectedFiles) > 0 {
				m.cursorIndex = 0
				m.currentState = choosingActionState
			}
		}
	}
	return m, nil
}

func (m model) handleActionSelection(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "down":
			m.cursorIndex = 1 - m.cursorIndex // Alterna entre 0 e 1
		case "enter":
			if m.cursorIndex == 0 {
				m.selectedAction = "convert"
				m.currentState = choosingFormatState
			} else {
				m.selectedAction = "compress"
				m.currentState = enteringSizeLimitState
				m.sizeInputBuffer = ""
			}
			m.cursorIndex = 0
		}
	}
	return m, nil
}

func (m model) handleFormatSelection(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.cursorIndex > 0 {
				m.cursorIndex--
			}
		case "down":
			if m.cursorIndex < len(supportedFormats)-1 {
				m.cursorIndex++
			}
		case "enter":
			m.targetFormat = supportedFormats[m.cursorIndex]
			m.currentState = processingState
			return m, m.executeFFmpeg()
		}
	}
	return m, nil
}

func (m model) handleSizeInput(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.sizeInputError != "" {
		m.sizeInputError = ""
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.sizeInputBuffer == "" {
				m.sizeInputError = "A entrada não pode estar vazia."
				return m, nil
			}
			mb, err := strconv.ParseFloat(m.sizeInputBuffer, 64)
			if err != nil {
				m.sizeInputError = "Entrada inválida. Use um número como '1.5' ou '10'."
				return m, nil
			}
			m.compressionSizeLimitMB = mb
			m.currentState = processingState
			return m, m.executeFFmpeg()

		case "backspace":
			if len(m.sizeInputBuffer) > 0 {
				m.sizeInputBuffer = m.sizeInputBuffer[:len(m.sizeInputBuffer)-1]
			}
		default:
			// Permite apenas dígitos e um ponto decimal.
			if len(msg.Runes) == 1 {
				char := msg.Runes[0]
				if (char >= '0' && char <= '9') || (char == '.' && !strings.Contains(m.sizeInputBuffer, ".")) {
					m.sizeInputBuffer += string(char)
				}
			}
		}
	}
	return m, nil
}

func (m model) handleCompletion(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return m, tea.Quit
		}
	}
	return m, nil
}
