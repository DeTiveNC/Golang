package analysis

import (
	"fmt"
	"github.com/detivenc/lsp-go/lsp"
	"strings"
)

type State struct {
	// Depends on the editor would be more fields
	Documents map[string]string
}

func NewState() State {
	return State{
		Documents: make(map[string]string),
	}
}

func getDiagnosticsForFile(text string) []lsp.Diagnostic {
	var diagnostics []lsp.Diagnostic
	for row, line := range strings.Split(text, "\n") {
		if strings.Contains(line, "VS Code") {
			idx := strings.Index(line, "VS Code")
			diagnostics = append(diagnostics, lsp.Diagnostic{
				Range:    LineRange(row, idx, idx+len("VS Code")),
				Severity: 1,
				Source:   "Common Sense",
				Message:  "Please make sure we use good language",
			})
		}

		if strings.Contains(line, "Neovim") {
			idx := strings.Index(line, "Neovim")
			diagnostics = append(diagnostics, lsp.Diagnostic{
				Range:    LineRange(row, idx, idx+len("Neovim")),
				Severity: 2,
				Source:   "Common Sense",
				Message:  "Great choice :)",
			})

		}
	}

	return diagnostics
}

func (s *State) OpenDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text
	return getDiagnosticsForFile(text)
}

func (s *State) UpdateDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text
	return getDiagnosticsForFile(text)
}

func (s *State) Hover(id int, uri string, position lsp.Position) lsp.HoverResponse {
	// In a real language server, this look up the type in our type analysis

	document := s.Documents[uri]

	return lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.HoverResult{
			Contents: fmt.Sprintf("File: %s, Character: %d", uri, len(document)),
		},
	}
}

func (s *State) Definition(id int, uri string, position lsp.Position) lsp.DefinitionResponse {
	// In a real language server, this look up the type in our type analysis

	return lsp.DefinitionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: lsp.Location{
			URI: uri,
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
				End: lsp.Position{
					Line:      position.Line - 1,
					Character: 0,
				},
			},
		},
	}
}

func (s *State) TextDocumentCodeAction(id int, uri string) lsp.TextDocumentCodeActionResponse {
	// In a real language server, this look up the type in our type analysis

	text := s.Documents[uri]

	var actions []lsp.CodeAction
	for row, line := range strings.Split(text, "\n") {
		idx := strings.Index(line, "VS Code")
		if idx >= 0 {
			replaceChange := map[string][]lsp.TextEdit{}
			replaceChange[uri] = []lsp.TextEdit{
				{
					Range:   LineRange(row, idx, idx+len("VS Code")),
					NewText: "Neovim",
				},
			}

			actions = append(actions, lsp.CodeAction{
				Title: "Replace VS C*de with a superior editor",
				Edit:  &lsp.WorkspaceEdit{Changes: replaceChange},
			})

			censorChange := map[string][]lsp.TextEdit{}
			censorChange[uri] = []lsp.TextEdit{
				{
					Range:   LineRange(row, idx, idx+len("VS Code")),
					NewText: "VS C*de",
				},
			}

			actions = append(actions, lsp.CodeAction{
				Title: "Censor to VS C*de",
				Edit:  &lsp.WorkspaceEdit{Changes: censorChange},
			})
		}
	}

	response := lsp.TextDocumentCodeActionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: actions,
	}

	return response
}

func LineRange(line, start, end int) lsp.Range {
	return lsp.Range{
		Start: lsp.Position{
			Line:      line,
			Character: start,
		},
		End: lsp.Position{
			Line:      line,
			Character: end,
		},
	}
}

func (s *State) TextDocumentCompletion(id int, uri string) lsp.CompletionResponse {

	// Ask your static analysis tools to figure out good completions
	items := []lsp.CompletionItem{
		{
			Label:         "Neovim (BTW)",
			Detail:        "Very cool editor",
			Documentation: "Fun to watch in videos.",
		},
	}

	response := lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: items,
	}

	return response
}
