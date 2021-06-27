package object

import (
	"errors"
	"sync"
	"time"
)

type History interface {
	Insert(time.Time, Object) error
	Prune(time.Time)
	Identifier() ObjectIdentifier
	Get(t time.Time) (Object, bool)
}

type history struct {
	history    map[time.Time]Object
	identifier ObjectIdentifier
	mux        sync.RWMutex
}

func NewHistory(identifier ObjectIdentifier) History {
	return &history{
		history:    make(map[time.Time]Object),
		identifier: identifier,
	}
}

func (hist *history) Insert(time time.Time, obj Object) error {
	if hist.identifier != obj.Identifier() {
		return errors.New("identifier mismatch inserting into history")
	}
	hist.mux.Lock()
	defer hist.mux.Unlock()
	hist.history[time] = obj
	return nil
}

func (hist *history) Prune(prune time.Time) {
	hist.mux.Lock()
	defer hist.mux.Unlock()
	for t := range hist.history {
		if t.Before(prune) {
			delete(hist.history, t)
		}
	}
}

func (hist *history) Get(t time.Time) (Object, bool) {
	hist.mux.RLock()
	defer hist.mux.RUnlock()
	obj, ok := hist.history[t]
	return obj, ok
}

func (hist *history) Identifier() ObjectIdentifier {
	return hist.identifier
}
