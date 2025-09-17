package data

type Stuff struct {
	Head  string
	Chest string
	Feet  string
}

type Person struct {
	Name   string
	Class  string
	Level  int
	Hp     int
	Hpmax  int
	Bag    []string
	Gold   int
	Spells []string
	Eq     Stuff
	Size   int
	Up     int
}

type Mob struct {
	Name  string
	Hp    int
	Hpmax int
	Atk   int
}
