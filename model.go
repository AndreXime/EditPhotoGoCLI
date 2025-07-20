// model.go
package main

type appState int

const (
	choosingFileState appState = iota
	choosingActionState
	choosingFormatState
	enteringSizeLimitState
	processingState
	doneState
)

// O tipo da mensagem de status foi renomeado para ser mais específico.
type processCompleteMsg string

// A struct 'model' teve seus campos renomeados para maior clareza.
type model struct {
	imagePaths             []string
	cursorIndex            int
	currentState           appState
	selectedFiles          map[string]struct{}
	selectedAction         string
	targetFormat           string
	compressionSizeLimitMB float64
	resultMessage          string
	styles                 *Styles

	sizeInputError  string
	sizeInputBuffer string
}

// A lista de formatos agora tem um nome mais descritivo.
var supportedFormats = []string{"jpg", "png", "webp", "jpeg"}

// A função de inicialização do modelo foi renomeada para seguir convenções.
func newModel(imagePaths []string) model {
	return model{
		imagePaths:    imagePaths,
		styles:        newDefaultStyles(),
		selectedFiles: make(map[string]struct{}),
		currentState:  choosingFileState, // O estado inicial é definido explicitamente.
	}
}
