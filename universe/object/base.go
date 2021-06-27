package object

import (
	"github.com/mkuchenbecker/storage/universe/vector"
)

type ObjectIdentifier string

type Object interface {
	Position() vector.Vector
	Identifier() ObjectIdentifier
}

type object struct {
	position   vector.Vector
	identifier ObjectIdentifier
	// orientation vector.Vector
	// velocity    vector.Vector
}

func New(identifier ObjectIdentifier, position vector.Vector) Object {
	return &object{
		position:   position,
		identifier: identifier,
	}
}

func (obj *object) Position() vector.Vector {
	return obj.position
}

func (obj *object) Identifier() ObjectIdentifier {
	return obj.identifier
}
