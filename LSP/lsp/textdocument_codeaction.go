package lsp

// CodeActionRequest CodeAction represents a code action request.
type CodeActionRequest struct {
	Request
	Params TextDocumentCodeActionParams `json:"params"`
}

// TextDocumentCodeActionParams represents the parameters of a `textDocument/codeAction` request.
type TextDocumentCodeActionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Range        Range                  `json:"range"`
	Context      CodeActionContext      `json:"context"`
}

type CodeActionContext struct {
	// Add fields here as needed
}

// TextDocumentCodeActionResponse represents a code action response.
type TextDocumentCodeActionResponse struct {
	Response
	Result []CodeAction `json:"result"`
}

type CodeAction struct {
	Title   string         `json:"title"`
	Edit    *WorkspaceEdit `json:"edit,omitempty"`
	Command *Command       `json:"command,omitempty"`
}

type Command struct {
	Title     string        `json:"title"`
	Command   string        `json:"command"`
	Arguments []interface{} `json:"arguments,omitempty"`
}
