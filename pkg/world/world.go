package world

// Resource can be mined on planets
type Resource struct {
	Name      string
	Frequency int64 // how often the resource mines (in ticks)
	Amount    int64 // amount of resource given mined every interval
}

// Planets have resources and storage.
type Planet struct {
	Name      string
	Resources map[string]Resource
	Storage   map[string]Resource
}

// Station is a simple struct for representing a Station.
// Stations only store things, they can not generate resources.
type Station struct {
	Name    string
	Storage map[string]Resource
}

// System holds a set of worlds and stations
type System struct {
	Worlds   []Planet
	Stations []Station
}

// Ship is the atomic unit of the game.
// * Ships travel from planets to stations to trade resources.
// * Ships occupy Planets or Stations.
type Ship struct {
	Name  string
	Cargo map[string]Resource
	Owner string
}
