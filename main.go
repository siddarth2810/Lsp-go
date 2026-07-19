package main

import (
	"bufio"
	"encoding/json"
	"log"
	"lsp/analyser"
	"lsp/lsp"
	"lsp/rpc"
	"os"
)

func main() {
	logger := getLogger("/home/sid/personal/lsp/log.txt")
	logger.Println("Hey lsp started!")
	scanner := bufio.NewScanner(os.Stdin)

	//rpc.Split is implemented in the format which scanner.Split needs
	scanner.Split(rpc.Split)

	state := analyser.NewState()

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error: %s", err)
			continue
		}

		handleMessage(logger, state, method, contents)
	}
}

func handleMessage(logger *log.Logger, state analyser.State, method string, contents []byte) {
	logger.Printf("Received message with method: %s", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Hey, couldn't parse this: %s", err)
		}
		logger.Printf("Connected to: %s %s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version)

		// send this reply back !
		msg := lsp.NewInitializeResponse(request.ID)
		reply := rpc.EncodeMessage(msg)
		writer := os.Stdout
		writer.Write([]byte(reply))
		logger.Print("Sent the reply")

	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/didOpne: %s", err)
			return
		}
		logger.Printf("Opened : %s", request.Params.TextDocument.URI)
		state.OpenDocument(
			request.Params.TextDocument.URI,
			request.Params.TextDocument.Text)

	case "textDocument/didChange":
		var request lsp.TextDocumentDidChangeNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("textDocument/didChange: %s", err)
			return
		}
		logger.Printf("Changed: %s", request.Params.TextDocument.URI)
		for _, change := range request.Params.ContentChanges {
			state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
		}

	}
}

// can't print to stdout because its used to communicate with client
func getLogger(filename string) *log.Logger {
	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("hey, not a good file")
	}
	return log.New(logFile, "[LSP]", log.Ldate)
}
