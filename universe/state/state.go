package state

import (
	"time"

	"github.com/mkuchenbecker/storage/universe/object"
)

type State interface {
	Object() object.Object
	Timestamp() time.Time
}

type state struct {
	obj       object.Object
	timestamp time.Time
}

func (this *state) Object() object.Object {
	return this.obj
}

func (this *state) Timestamp() time.Time {
	return this.timestamp
}

type globalstate struct {
	histories            map[object.ObjectIdentifier]object.History
	squareMagnitudeIndex map[float64]object.Object
}

// func (s *globalstate) Detect(timestamp time.Time, position vector.Vector) []object.Object {

// 	return make([]object.Object, 0)
// }
