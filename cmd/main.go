package main

import (
	"bufio"
	"errors"
	"fmt"
	"log/slog"
	"os"

	parserPkg "github.com/hound672/kv_storage/internal/compute/parser"
	"github.com/hound672/kv_storage/internal/models"
	"github.com/hound672/kv_storage/internal/storage"
	inmemory "github.com/hound672/kv_storage/internal/storage/in_memory"
)

func main() {
	slog.Info("Running")

	parser := parserPkg.New()
	store := inmemory.New()
	engine := storage.New(store)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("> ")
		rawMessage, _ := reader.ReadString('\n')
		fmt.Println(rawMessage)

		cmd, err := parser.Parse(rawMessage)
		if err != nil {
			if errors.Is(err, models.ErrInvalidCommand) {
				fmt.Println("Invalid command")
				continue
			}

			fmt.Println("Unknown error: %s", err)
			continue
		}

		res, err := engine.Execute(cmd)
		if err != nil {
			if errors.Is(err, models.ErrKeyNotFound) {
				fmt.Println("key not found")
				continue
			}

			fmt.Println("Unknown error: %s", err)
			continue
		}

		if res != "" {
			fmt.Printf("Result: %s\n", res)
		}
		fmt.Println("OK")
	}
}
