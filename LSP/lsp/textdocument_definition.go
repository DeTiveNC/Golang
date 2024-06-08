package lsp

// DefinitionRequest Hover represents a hover for a symbol at a specific location.
type DefinitionRequest struct {
	Request
	Params DefinitionParams `json:"params"`
}

// DefinitionParams represents the parameters of a `textDocument/hover` request.
type DefinitionParams struct {
	TextDocumentPositionParams
}

type DefinitionResponse struct {
	Response
	Result Location `json:"result"`
}
