package base

import (
	"math"
	"slices"
)

type CollisionDetail struct {
	Vec[float32]
	o Object
}
type CollisionDetails []CollisionDetail

func (c CollisionDetails) CheckIf(check func(CollisionDetail) bool) bool {
	return slices.ContainsFunc(c, check)
}

type RigidBody struct {
	o          Object
	quadtree   *QuadTree
	toSimulate bool
	isStatic   bool

	mass      float32
	velocity  Vec[float32]
	force     Vec[float32]
	touch     bool
	Collision CollisionDetails
}

func NewRigidBody(toSimulate, isStatic bool, mass float32) RigidBody {
	return RigidBody{
		toSimulate: toSimulate,
		isStatic:   isStatic,
		mass:       mass,
	}
}

func (r *RigidBody) Touch() {
	r.touch = true
}

// Return force reference
func (r *RigidBody) GetForce() (res *Vec[float32]) {
	return &r.force
}

// Return velocity reference
func (r *RigidBody) GetVelocity() (res *Vec[float32]) {
	return &r.velocity
}

func (r *RigidBody) ApplyForce(v Vec[float32]) {
	r.force.Add(v)
}

func (r *RigidBody) ApplyImpulse(v Vec[float32]) {
	r.velocity.Add(v)
}

func (r *RigidBody) ApplyAcceleration(v Vec[float32]) {
	v.MultScalar(r.mass)
	r.force.Add(v)
}

func (r *RigidBody) SetObject(o Object) {
	r.o = o
}

func (r *RigidBody) SetQuadTree(q *QuadTree) {
	r.quadtree = q
}

func (r *RigidBody) GetHitbox() *Hitbox {
	return r.o.GetHitbox()
}

func (r *RigidBody) Move(v Vec[float32]) {
	r.o.MoveBy(v)
}

func (r *RigidBody) Integrate(dt float32) {
	if !r.toSimulate || r.isStatic {
		return
	}

	acceleration := r.force.Clone()
	acceleration.MultScalar(1.0 / r.mass)
	acceleration.MultScalar(dt)

	r.Move(*r.velocity.Add(*acceleration).Clone().MultScalar(dt * 10))

	r.force = NewVec[float32](0, 0)
}

func (a *RigidBody) ResolveCollision(b *RigidBody, v Vec[float32]) error {
	res := NewVec[float32](0, 0)

	aToB := FromAtoBVec(a.o.GetPos(), b.o.GetPos())

	if math.Abs(float64(v.X)) < math.Abs(float64(v.Y)) {
		res.AddScalars(v.X, 0)
		if math.Signbit(float64(res.X)) == math.Signbit(float64(aToB.X)) {
			res.MultScalar(-1)
		}
		if math.Signbit(float64(a.velocity.X)) == math.Signbit(float64(aToB.X)) {
			a.velocity.X = 0
		}
	} else {
		res.AddScalars(0, v.Y)
		if math.Signbit(float64(res.Y)) == math.Signbit(float64(aToB.Y)) {
			res.MultScalar(-1)
		}
		if math.Signbit(float64(a.velocity.Y)) == math.Signbit(float64(aToB.Y)) {
			a.velocity.Y = 0
		}
	}
	a.Collision = append(a.Collision, CollisionDetail{
		res,
		b.o,
	})
	a.Move(res)
	return nil
}

func (a *RigidBody) Collide(b *RigidBody) error {
	if !a.toSimulate || !b.toSimulate {
		return nil
	}
	if a.isStatic {
		return nil
	}

	hitboxA := a.o.GetHitbox()
	hitboxB := b.o.GetHitbox()

	posA := AddVecs(hitboxA.GetPos(), a.o.GetPos())
	posB := AddVecs(hitboxB.GetPos(), b.o.GetPos())

	sizeA := hitboxA.GetBox()
	sizeB := hitboxB.GetBox()

	centerA := Vec[float32]{X: posA.X + sizeA.X/2, Y: posA.Y + sizeA.Y/2}
	centerB := Vec[float32]{X: posB.X + sizeB.X/2, Y: posB.Y + sizeB.Y/2}

	centerDistance := SubVecs(centerA, centerB)
	centerDistance = NewVec(float32(math.Abs(float64(centerDistance.X))), float32(math.Abs(float64(centerDistance.Y))))
	totalSize := AddVecs(sizeA, sizeB)
	distance := NewVec(centerDistance.X-totalSize.X/2, centerDistance.Y-totalSize.Y/2)
	if a.touch {
		a.Collision = nil
		a.touch = false
	}
	if distance.X < 0 && distance.Y < 0 {
		return a.ResolveCollision(b, distance)
	}
	return nil
}
