package game

func NewWorld() *World {
	rooms := make([]*Room, 12)

	// Create all rooms
	rooms[Tower] = buildRoom(towerDescription, true)
	rooms[LibraryEntrance] = buildRoom(libraryEntranceDescription, true)
	rooms[ReferenceDesk] = buildRoom("", true)
	rooms[ReshelvingCart] = buildRoom("", true)
	rooms[SkeletonRoom] = buildRoom("", true)
	rooms[CardCatalog] = buildRoom("", true)
	rooms[Study] = buildRoom("", false)
	rooms[NonFiction] = buildRoom("", true)
	rooms[Fiction] = buildRoom("", true)
	rooms[Murder] = buildRoom("", true)
	rooms[TheStacks] = buildRoom("", true)
	rooms[ExitRoom] = buildRoom("", true)

	// Helper to connect rooms bidirectionally
	connect := func(r1, r2 int, dir Direction) {
		rooms[r1].Exits[dir] = &Exit{To: rooms[r2]}
		rooms[r2].Exits[oppositeDirection(dir)] = &Exit{To: rooms[r1]}
	}

	// Connect the world
	connect(Tower, LibraryEntrance, East)
	connect(LibraryEntrance, ReferenceDesk, East)
	connect(ReferenceDesk, ReshelvingCart, South)
	connect(ReshelvingCart, SkeletonRoom, South)
	connect(SkeletonRoom, CardCatalog, East)
	connect(CardCatalog, Study, East)
	connect(CardCatalog, NonFiction, North)
	connect(NonFiction, Fiction, East)
	connect(Fiction, Murder, North)
	connect(Murder, TheStacks, West)

	// Special one-way/hidden exits
	rooms[TheStacks].Exits[North] = &Exit{
		To:       rooms[ExitRoom],
		IsHidden: true,
	}

	return &World{
		Rooms: rooms,
		Player: &Player{
			CurrentRoom: rooms[Tower],
			Inventory:   []*GameObject{},
		},
	}
}

func buildRoom(desc string, lit bool) *Room {
	return &Room{
		Description: desc,
		Exits:       make(map[Direction]*Exit),
		Contents:    []*GameObject{},
		IsLit:       lit,
	}
}

var towerDescription = `You are in a brick lined room. 
Above there is a trap door with a ladder above your head. 
Ahead, there appears to be an entrace to another room.
`
var libraryEntranceDescription = `This appears to be the entrance to a library.
There is a welcome sign (TODO) and a book return schute`
