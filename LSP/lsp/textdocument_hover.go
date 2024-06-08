package lsp

// DefinitionRequest Hover represents a hover for a symbol at a specific location.
type HoverRequest struct {
	Request
	Params HoverParams `json:"params"`
}

// HoverParams represents the parameters of a `textDocument/hover` request.
type HoverParams struct {
	TextDocumentPositionParams
}

type HoverResponse struct {
	Response
	Result HoverResult `json:"result"`
}

type HoverResult struct {
	Contents string `json:"contents"`
}
