package base

import (
	"fmt"
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
	Id        string
}

func NewRigidBody(toSimulate, isStatic bool, mass float32) RigidBody {
	return RigidBody{
		toSimulate: toSimulate,
		isStatic:   isStatic,
		Mass:       mass,
		Friction:   5,
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
	r.force.Add(*v.MultScalar(r.Mass))
}

func (r *RigidBody) setObject(o Object) {
	r.o = o
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
		v := r.GetVelocity()
		if v.X > 0 {
			sign = 1
		} else if v.X < 0 {
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
		} else {
			norm := v.Normalize()
			magn := v.Magnitude()
			r.ApplyForce(*norm.MultScalar(magn * magn * (-0.3)))
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

	if (v.X != 0 && math.Abs(float64(v.X)) < math.Abs(float64(v.Y))) || v.Y == 0 {
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

func getAxis(h *Hitbox, posObject Vec[float32]) (res []Vec[float32]) {
	verts := h.GetVertex()
	var a, b Vec[float32]
	for i := range verts {
		pos := AddVecs(posObject, *h.Pos)
		a = *verts[i].Clone().Add(pos)
		b = *verts[(i+1)%len(verts)].Clone().Add(pos)
		edge := FromAtoBVec(a, b)
		// perpendicular normal — rotate edge 90°
		normal := Vec[float32]{X: -edge.Y, Y: edge.X}
		res = append(res, *normal.Normalize())
	}
	return res
}
func getOverlap(minA, maxA, minB, maxB float32) float32 {
	var AB, BA float32
	AB = maxA - minB
	BA = maxB - minA
	if AB < 0 || BA < 0 {
		return 0
	}

	return min(AB, BA)

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

	posA := a.o.GetPos()
	posB := b.o.GetPos()

	sizeA := hitboxA.GetOuterBox()
	sizeB := hitboxB.GetOuterBox()

	if a.touch {
		a.Collision = nil
		a.touch = false
	}
	if posA.X+sizeA.X < posB.X || posB.X+sizeB.X < posA.X ||
		posA.Y+sizeA.Y < posB.Y || posB.Y+sizeB.Y < posA.Y {
		return nil
	}

	axisA := getAxis(hitboxA, posA)
	var minA, maxA, minB, maxB, dist float32
	var err error
	var minDist float32 = math.MaxFloat32
	var axesToMove Vec[float32]
	for _, axA := range axisA {
		minA, maxA = hitboxA.ProjectionOn(posA, axA)
		minB, maxB = hitboxB.ProjectionOn(posB, axA)
		dist = getOverlap(minA, maxA, minB, maxB)
		if dist == 0 {
			return nil
		}
		if a.Id == "ball" && b.Id == "block" {
			fmt.Printf("A %v %v %v %v\n", axA, minA, maxA, dist)
		}
		if math.Abs(float64(dist)) < math.Abs(float64(minDist)) {
			axesToMove = *axA.MultScalar(dist)

			minDist = dist
		}

	}
	err = a.ResolveCollision(b, axesToMove)
	if err != nil {
		return err
	}

	return nil
}
