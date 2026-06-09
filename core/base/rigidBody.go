package base

import (
	"math"
	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type CollisionDetail struct {
	Vec[float32]
	r *RigidBody
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

	Mass      float32
	velocity  Vec[float32]
	force     Vec[float32]
	touch     bool
	Collision CollisionDetails
	Friction  float32
}

func NewRigidBody(toSimulate, isStatic bool, mass float32) RigidBody {
	return RigidBody{
		toSimulate: toSimulate,
		isStatic:   isStatic,
		Mass:       mass,
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
	r.force.Add(*v.MultScalar(1 / rl.GetFrameTime()))
}

func (r *RigidBody) ApplyAcceleration(v Vec[float32]) {
	v.MultScalar(r.Mass)
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

func (r *RigidBody) ComputeVelocity(dt float32) *Vec[float32] {
	if !r.toSimulate || r.isStatic {
		return nil
	}

	acceleration := r.force.Clone()
	acceleration.MultScalar(1.0 / r.Mass)
	acceleration.MultScalar(dt)
	return r.velocity.Clone().Add(*acceleration)
}

func (r *RigidBody) Integrate(dt float32) {
	if !r.toSimulate || r.isStatic {
		return
	}

	friction := func(r *RigidBody) {
		var sign float32
		if r.GetVelocity().X > 0 {
			sign = 1
		} else if r.GetVelocity().X < 0 {
			sign = -1
		} else {
			return
		}
		f := r.GetForce().Clone()
		if f.Y < 0 {
			return
		}
		if f.X != 0 {
			return
		}
		if r.Collision.CheckIf(func(cd CollisionDetail) bool {
			if cd.Y < 0 {
				f.X = f.Y * cd.r.Friction * (-sign)
				return true
			}
			return false
		}) {
			f.Y = 0
			r.ApplyForce(*f)
		}

	}

	vBefore := r.ComputeVelocity(dt)
	friction(r)
	vAfter := r.ComputeVelocity(dt)
	if (vBefore.X == 0) || (vAfter.X > 0 && vBefore.X > 0) || (vAfter.X < 0 && vBefore.X < 0) {
		r.velocity = *vAfter
	} else {
		r.velocity = Vec[float32]{}
	}
	r.o.MoveBy(*r.velocity.Clone().MultScalar(10 * dt))
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
		b,
	})
	a.o.MoveBy(res)
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
