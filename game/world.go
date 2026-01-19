package game

func NewWorld() *World {
	rooms := make([]*Room, 12)

	// Create all rooms
	rooms[Tower] = buildRoom(towerDescription, true)
	rooms[LibraryEntrance] = buildRoom(libraryEntranceDescription, true)
	rooms[CirculationDesk] = buildRoom(circulationDeskDescription, true)
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
	connect(LibraryEntrance, CirculationDesk, East)
	connect(CirculationDesk, ReshelvingCart, South)
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

var towerDescription = `You are in a brick lined room that is dimly lit. 
Above you there is a trap door with a ladder extending down, but out of reach. Light filters through the door.

Ahead, the brick surrounds large windows and a double door all made of glass and aluminum. The glass reflects back the room you are in.
To the left of the door there is a small rectangular door at chest height made out of metal.
`
var libraryEntranceDescription = `As you enter, a couple of flourecent bulbs half-heartedly flicker on providing some light to the room and a new buzzing sound.

There is a welcome sign (TODO) ahead and a book return chute to your left that is empty, but full of undisturbed dust.

Further ahead you can see a desk.
`

var circulationDeskDescription = `You stand at the circulation desk. Weak fluorescent light continues to flickers overhead, 
casting uneven shadows across the scratched wooden surface.

On the desk: a coffee-stained notepad, a rubber stamp reading "MISFILED"

The air smells of old paper and disappointment.
`

var cardCatalogDescription = `You are now at the card catalog, its brass pulls tarnished and labeled 
in faded ink. Some drawers sit slightly ajar.
`
