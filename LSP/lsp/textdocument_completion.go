package lsp

// CompletionRequest Hover represents a hover for a symbol at a specific location.
type CompletionRequest struct {
	Request
	Params CompletionParams `json:"params"`
}

// CompletionParams represents the parameters of a `textDocument/hover` request.
type CompletionParams struct {
	TextDocumentPositionParams
}

type CompletionResponse struct {
	Response
	Result []CompletionItem `json:"result"`
}

type CompletionItem struct {
	Label         string `json:"label"`
	Detail        string `json:"detail"`
	Documentation string `json:"documentation"`
}
