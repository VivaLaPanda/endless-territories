package entity

import (
	"github.com/VivaLaPanda/endless-territories/event"
)

type Entity interface {
	GetLocation() (x int, y int)
	GetSprite() (filename string)
	QueueEvent(ev event.Event) (suceeded bool)
}
