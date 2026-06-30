package core

import (
	"github.com/Wordluc/TheEngine/core/base"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Entity interface {
	Start() error
	GetObject() base.Object
	Update(dt float32) error
}

type SimpleEntity struct {
	o      base.Object
	update func(float32) error
}

func (s *SimpleEntity) Start() error {
	return nil
}

func (s *SimpleEntity) GetObject() base.Object {
	return s.o
}

func (s *SimpleEntity) Update(dt float32) error {
	return s.update(dt)
}

func NewSimpleEntity(o base.Object, update func(float32) error) *SimpleEntity {
	if update == nil {
		update = func(f float32) error { return nil }
	}
	return &SimpleEntity{
		o:      o,
		update: update,
	}
}

type EntityEngine struct {
	toStart  []Entity
	entities []Entity
	quadtree base.QuadTree
}

func NewEntityEngine(w, h float32) *EntityEngine {
	quadtree := base.NewQuadTree(base.Vec[float32]{}, base.NewVec(w, h), nil)
	return &EntityEngine{
		quadtree: *quadtree,
	}
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
		e.quadtree.Insert(e.toStart[i].GetObject())
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
	if dt == 0 {
		return nil
	}
	for i := range e.entities {
		err := e.entities[i].Update(dt)
		if err != nil {
			return err
		}
	}
	for i := range e.entities {
		obj = e.entities[i].GetObject()
		base.UseModifierRef(obj, func(r *base.RigidBody) {
			r.Touch()
			r.Integrate(dt)
			r.GetVelocity().CapAt(base.Vec[float32]{X: 20, Y: 20})
		})
	}
	e.quadtree.Foreach(func(elements []base.QuadTreeElement) {
		for i := range elements {
			for j := range elements {
				if i == j {
					continue
				}
				a, okA := elements[i].(base.Object)
				b, okB := elements[j].(base.Object)
				if !okA || !okB {
					continue
				}
				ra := base.GetModifierRef[*base.RigidBody](a)
				rb := base.GetModifierRef[*base.RigidBody](b)
				if ra == nil || rb == nil {
					continue
				}
				ra.Collide(rb)
			}
		}
	})
	for i := range e.entities {
		obj = e.entities[i].GetObject()
		if d, ok = obj.(base.Drawable); ok {
			d.Draw()
		}
	}
	return nil
}
