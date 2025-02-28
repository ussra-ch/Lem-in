package lem

type Room struct {
	name    string
	x, y    string
}
type Colony struct {
	NumAnts int
	rooms   map[string]*Room
	links   map[string][]string
	start   *Room
	end     *Room
}
