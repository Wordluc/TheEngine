package core

import (
	"game/core/base"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Entity interface {
	Start() error
	getObject() base.Object
	Update(dt float32) error
}

type SimpleEntity struct {
	o      base.Object
	update func(float32) error
}

func (s *SimpleEntity) Start() error {
	return nil
}

func (s *SimpleEntity) getObject() base.Object {
	return s.o
}

func (s *SimpleEntity) Update(dt float32) error {
	return s.Update(dt)
}

func NewSimpleEntity(o base.Object, update func(float32) error) *SimpleEntity {
	return &SimpleEntity{
		o:      o,
		update: update,
	}
}

type EntityEngine struct {
	toStart  []Entity
	entities []Entity
}

func (e *EntityEngine) AppendEntity(entity Entity) {
	e.toStart = append(e.toStart, entity)
}

func (e *EntityEngine) RunEntitiesStarts() error {
	if e.toStart == nil {
		return nil
	}
	for i := range e.toStart {
		err := e.toStart[i].Start()
		if err != nil {
			return err
		}
	}
	e.entities = append(e.entities, e.toStart...)
	e.toStart = nil
	return nil
}

func (e *EntityEngine) RunEntitiesUpdates() error {
	var (
		obj base.Object
		d   base.Drawable
		ok  bool
		dt  float32 = rl.GetFrameTime()
	)
	for i := range e.entities {
		err := e.entities[i].Update(dt)
		if err != nil {
			return err
		}
	}

	for i := range e.entities {
		obj = e.entities[i].getObject()
		if d, ok = obj.(base.Drawable); ok {
			d.Draw()
		}
	}
	return nil
}
