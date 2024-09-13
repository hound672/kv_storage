package parser

import (
	"strings"

	"github.com/hound672/kv_storage/internal/models"
)

type Parser struct {
}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(text string) (models.Command, error) {
	tokens := strings.Split(text, " ")
	if len(tokens) <= 1 {
		return models.Command{}, models.ErrInvalidCommand
	}

	args := tokens[1:]

	switch tokens[0] {
	case "SET":
		return parseSET(args)
	case "GET":
		return parseGET(args)
	case "DEL":
		return parseDEL(args)
	}

	return models.Command{}, nil
}

func parseSET(tokens []string) (models.Command, error) {
	if len(tokens) != 2 {
		return models.Command{}, models.ErrInvalidCommand
	}

	return models.Command{
		Type: models.CommandTypeSet,
		Args: tokens,
	}, nil
}

func parseGET(tokens []string) (models.Command, error) {
	if len(tokens) != 1 {
		return models.Command{}, models.ErrInvalidCommand
	}

	return models.Command{
		Type: models.CommandTypeGet,
		Args: tokens,
	}, nil
}

func parseDEL(tokens []string) (models.Command, error) {
	if len(tokens) != 1 {
		return models.Command{}, models.ErrInvalidCommand
	}

	return models.Command{
		Type: models.CommandTypeDel,
		Args: tokens,
	}, nil
}
