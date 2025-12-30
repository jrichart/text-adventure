package game

import "strings"

// Game State
type World struct {
	Rooms  []*Room
	Player *Player
}

type Room struct {
	ID          string
	Description string
	Exits       map[Direction]*Exit // north, south, etc.
	Contents    []*GameObject
	IsLit       bool
}

type Exit struct {
	To          *Room
	IsLocked    bool
	RequiresKey string // ID of key item needed
	IsHidden    bool   // for secret bookshelf door
}

type GameObject struct {
	ID          string
	Names       []string // "book", "red book", "journal"
	Description string

	// Properties
	IsContainer bool
	Contents    []GameObject
	IsTakeable  bool
	IsFixed     bool // like a desk

	// For special interactive items
	OnInteract func(*World, *GameObject) string
}

type Player struct {
	CurrentRoom *Room
	Inventory   []*GameObject
}

type Direction int

const (
	UNKNOWN = iota
	North
	Northwest
	West
	Southwest
	South
	Southeast
	East
	Northeast
)

func getDirectionInt(dir string) Direction {
	lower := strings.ToLower(dir)
	switch lower {
	case "north":
		return North
	case "northwest":
		return Northwest
	case "west":
		return West
	case "southwest":
		return Southwest
	case "south":
		return South
	case "southeast":
		return Southeast
	case "east":
		return East
	case "northeast":
		return Northeast
	default:
		return UNKNOWN
	}
}

func oppositeDirection(dir Direction) Direction {
	switch dir {
	case North:
		return South
	case Northwest:
		return Southeast
	case West:
		return East
	case Southwest:
		return Northeast
	case South:
		return North
	case Southeast:
		return Northwest
	case East:
		return West
	case Northeast:
		return Southwest
	default:
		return UNKNOWN
	}
}

const (
	Tower = iota
	LibraryEntrance
	ReferenceDesk
	ReshelvingCart
	SkeletonRoom
	CardCatalog
	Study
	NonFiction
	Fiction
	Murder
	TheStacks
	ExitRoom
)

// Optional: helpful for debugging
var RoomNames = map[int]string{
	Tower:           "Tower",
	LibraryEntrance: "Library Entrance",
	ReferenceDesk:   "Reference Desk",
	ReshelvingCart:  "Reshelving Cart",
	SkeletonRoom:    "Skeleton Room",
	CardCatalog:     "Card Catalog",
	Study:           "Study",
	NonFiction:      "Non-Fiction",
	Fiction:         "Fiction",
	Murder:          "Murder",
	TheStacks:       "The Stacks",
	ExitRoom:        "Exit Room",
}
