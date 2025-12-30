package game

import (
	"fmt"
	"text-adventure/ast"
)

type CommandHandler struct {
	World *World
}

func NewCmdHandler(world *World) *CommandHandler {
	return &CommandHandler{
		World: world,
	}
}

func (ch *CommandHandler) Execute(cmd ast.Command) string {
	switch cmd.Verb.TokenLiteral() {
	case "go":
		return ch.handlerMove(cmd)
	case "look":
		return ch.HandleLook(cmd)
	default:
		return ch.handleUnknown(cmd)
	}
}

func (ch *CommandHandler) handlerMove(cmd ast.Command) string {
	room := ch.World.Player.CurrentRoom
	dir := getDirectionInt(cmd.Object.Noun.Literal)
	nextRoom, ok := room.Exits[dir]
	if !ok {
		return "I can't go that direction"
	}
	ch.World.Player.CurrentRoom = nextRoom.To
	return nextRoom.To.Description
}

func (ch *CommandHandler) HandleLook(cmd ast.Command) string {

	return ""
}

func (ch *CommandHandler) handleUnknown(cmd ast.Command) string {
	return fmt.Sprintf("I can't do %s", cmd.Verb.TokenLiteral())
}
