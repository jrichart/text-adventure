package vocabulary

type Vocabulary struct {
	verbs map[string]bool
	nouns map[string]bool
}

func New() *Vocabulary {
	return &Vocabulary{
		verbs: make(map[string]bool),
		nouns: make(map[string]bool),
	}
}

func (v *Vocabulary) AddVerb(word string) {
	v.verbs[word] = true
}

func (v *Vocabulary) AddNoun(word string) {
	v.nouns[word] = true
}

func (v *Vocabulary) IsVerb(word string) bool {
	return v.verbs[word]
}

func (v *Vocabulary) IsNoun(word string) bool {
	return v.nouns[word]
}

// The Vocabulary that is available at the start of the game.
// More could be added based on interactions during the game
func DefaultVocabulary() *Vocabulary {
	v := New()

	verbs := []string{"take", "get", "drop", "look", "examine", "go", "open", "close", "use", "peek", "turn", "grab", "pull", "turn-on", "pick-up"}
	for _, verb := range verbs {
		v.AddVerb(verb)
	}

	nouns := []string{"book", "shelf", "note", "bookshelf", "desk", "paper", "room", "door", "key", "lamp", "light", "north", "south", "east", "west", "left", "right", "up", "down", "sword"}
	for _, noun := range nouns {
		v.AddNoun(noun)
	}
	return v
}
