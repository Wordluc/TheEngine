package physics

import (
	"game/core/base"
)

type RigidBody struct {
	o          base.Object
	quadtree   *QuadTree
	toSimulate bool
	isStatic   bool
}

func NewRigidBody(object base.Object) RigidBody {
	return RigidBody{
		o:          object,
		toSimulate: true,
	}
}

func (r *RigidBody) SetQuadTree(q *QuadTree) {
	r.quadtree = q
}

func (r *RigidBody) GetHitbox() *base.Hitbox {
	return r.o.GetHitbox()
}

func (r *RigidBody) Move(v base.Vec[float32]) {
	r.o.MoveBy(v)
}

func (a *RigidBody) ResolveCollision(b *RigidBody, v base.Vec[float32]) error {
	a.Move(v)
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

	posA := hitboxA.GetPos()
	posB := hitboxB.GetPos()

	sizeA := hitboxA.GetBox()
	sizeB := hitboxB.GetBox()

	centerA := base.Vec[float32]{X: posA.X + sizeA.X/2, Y: posA.Y + sizeA.Y/2}
	centerB := base.Vec[float32]{X: posB.X + sizeB.X/2, Y: posB.Y + sizeB.Y/2}

	centerDistance := base.SubVecs(centerA, centerB)
	totalSize := base.AddVecs(sizeA, sizeB)
	distance := base.NewVec(centerDistance.X-totalSize.X/2, centerDistance.Y-totalSize.Y/2)

	if distance.X < 0 || distance.Y < 0 {
		return a.ResolveCollision(b, distance)
	}
	return nil
}
