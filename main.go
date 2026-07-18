package main

import (
	"bufio"
	"log"
	"lsp/rpc"
	"os"
)

func main() {
	logger := getLogger("/home/sid/personal/lsp/log.txt")
	logger.Println("Hey lsp started!")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(logger, msg)
	}
}

func handleMessage(logger *log.Logger, msg any) {
	logger.Println(msg)
}

// can't print to stdout because its used to communicate with client
func getLogger(filename string) *log.Logger{
	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("hey, not a good file")
	}
	return log.New(logFile, "[LSP]", log.Ldate)
}
