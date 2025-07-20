// styles.go
package main

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	Header      lipgloss.Style
	Selected    lipgloss.Style
	Unselected  lipgloss.Style
	Help        lipgloss.Style
	ErrorHeader lipgloss.Style
	ErrorMsg    lipgloss.Style
	SuccessMsg  lipgloss.Style
}

// Renomeada para seguir a convenção de construtores em Go.
func newDefaultStyles() *Styles {
	s := new(Styles)

	s.Header = lipgloss.NewStyle().
		Foreground(lipgloss.Color("36")). // Ciano
		Bold(true)

	s.Selected = lipgloss.NewStyle().
		Foreground(lipgloss.Color("32")) // Verde

	s.Unselected = lipgloss.NewStyle().
		Foreground(lipgloss.Color("250")) // Cinza claro

	s.Help = lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")) // Cinza escuro

	s.ErrorHeader = lipgloss.NewStyle().
		Foreground(lipgloss.Color("9")). // Vermelho
		Bold(true)

	s.ErrorMsg = lipgloss.NewStyle().
		Foreground(lipgloss.Color("9"))

	s.SuccessMsg = lipgloss.NewStyle().
		Foreground(lipgloss.Color("10")) // Verde claro

	return s
}
