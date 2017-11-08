package handlers

import (
	"math/rand"
	"reflect"
	"time"
)

// NewUIntHandler creates a hander for faking the uint type
func NewUIntHandler() *Liar {
	liar := Liar{
		Kind: reflect.Uint,
		Type: "uint",
	}

	liar.Fill = func(field reflect.Value, args Tag) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		field.Set(reflect.ValueOf(uint(r.Uint32())))
	}

	return &liar
}
