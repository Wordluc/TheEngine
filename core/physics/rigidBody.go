package physics

import "game/core/base"

type RigidBody struct {
	o        base.Object
	quadtree *QuadTree
}

func NewRigidBody(object base.Object) RigidBody {
	return RigidBody{
		o: object,
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
