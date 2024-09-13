package models

type CommandType int

const (
	CommandTypeSet CommandType = iota + 1
	CommandTypeGet
	CommandTypeDel
)

type Command struct {
	Type CommandType
	Args []string
}
