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

type SimpleEntity[opt any] struct {
	o      base.Object
	opt    opt
	update func(float32) error
	start  func(opt opt) base.Object
}

func (s *SimpleEntity[opt]) Start() error {
	s.o = s.start(s.opt)
	return nil
}

func (s *SimpleEntity[opt]) getObject() base.Object {
	return s.o
}

func (s *SimpleEntity[opt]) Update(dt float32) error {
	return s.Update(dt)
}

func NewSimpleEntity[opt any](o base.Object, option opt, start func(opt opt) base.Object, update func(float32) error) *SimpleEntity[opt] {
	return &SimpleEntity[opt]{
		o:      o,
		update: update,
		opt:    option,
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
	dt := rl.GetFrameTime()
	for i := range e.entities {
		err := e.entities[i].Update(dt)
		if err != nil {
			return err
		}
	}
	var (
		obj base.Object
		d   base.Drawable
		ok  bool
	)

	for i := range e.entities {
		obj = e.entities[i].getObject()
		if d, ok = obj.(base.Drawable); ok {
			d.Draw()
		}
	}
	return nil
}
