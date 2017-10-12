package grid

import (
	"github.com/VivaLaPanda/endless-territories/entity"
	"github.com/VivaLaPanda/endless-territories/event"
)

// Grid is a representation of a 2D grid of cells that contain entities
// Grids themselves are an entity.
// A grid is the core datastructure used to provide the state of the game
type Grid struct {
	boardArray [][][]entity.Entity
	x, y       int              // The grid's location relative to it's parent, assuming it has one
	eventQueue chan event.Event // A channel used to buffer event processing
	spriteLoc  string           // A string used to serve sprites
	alive      bool             // Alive represents the state of the eventProcessor
}

// BuildGrid is the constructor for the Grid struct
// It expects no arguments and returns a pointer to a grid
// It also spins up a function to process events and marks the grid as alive
func BuildGrid() *Grid {
	g := &Grid{}
	g.alive = true
	g.eventQueue = make(chan event.Event, 10)
	g.eventProcessor()

	return g
}

// GetGridStack takes a pointer to a grid and a location on that grid and returns
//   the entities at that location
func (g *Grid) GetGridStack(x int, y int) []entity.Entity {
	return g.boardArray[x][y]
}

// GetLocation takes a pointer to a grid and returns its location relative
//   to its parent.
func (g *Grid) GetLocation() (x int, y int) {
	return g.x, g.y
}

// GetSprite takes a pointer to a grid and returns its location url
func (g *Grid) GetSprite() (filename string) {
	return g.spriteLoc
}

// QueueEvent makes sure an eventProcessor is running and then adds the passwed
//   event struct to the queue
func (g *Grid) QueueEvent(ev event.Event) (succeeded bool) {
	if g.alive != true {
		go g.eventProcessor()
	}
	g.eventQueue <- ev

	return true
}

func (g *Grid) eventProcessor() {
	for ev := range g.eventQueue {
		switch ev.EvType {
		case ("SUSPEND"):
			g.suspend()
		}
	}
}

func (g *Grid) suspend() {
	close(g.eventQueue)
	g.alive = false
}
