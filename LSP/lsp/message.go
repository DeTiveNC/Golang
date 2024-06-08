package lsp

type Request struct {
	RPC    string `json:"jsonrpc"`
	ID     int    `json:"id"`
	Method string `json:"method"`

	// Params is a union type
}

type Response struct {
	RPC string `json:"jsonrpc"`
	ID  *int   `json:"id,omitempty"`

	// Result is a union type
	// Error is a union type
}

type Notification struct {
	RPC    string `json:"jsonrpc"`
	Method string `json:"method"`
}
