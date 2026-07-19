package lsp

type TextDocumentItem struct {
	// The text document's URI.
	URI string `json:"uri"`

	//  The text document's language identifier.
	LanguageID string `json:"languageId"`

	// The version number of this document (it will increase after each
	// change, including undo/redo).
	Version int `json:"version"`

	// The content of the opened text document.
	Text string `json:"text"`
}

type TextDocumentIdentifier struct {
	URI string `json:"uri"`
}

type VersionTextDocumentIdentifier struct {
	TextDocumentIdentifier
	Version int `json:"version"`
}

// textDocument/didOpen
type DidOpenTextDocumentNotification struct {
	Notification
	Params DidOpenTextDocumentParams `json:"params"`
}

type DidOpenTextDocumentParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}

// textDocument/didChange
type TextDocumentDidChangeNotification struct {
	Notification
	Params DidChangeTextDocumentParams `json:"params"`
}

type DidChangeTextDocumentParams struct {
	TextDocument   VersionTextDocumentIdentifier    `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

// An event describing the change to a text document. If only a text is provided
// its considered to be the full content of the document
type TextDocumentContentChangeEvent struct {
	Text string `json:"text"`

	// Range
	// RangeLength
}
