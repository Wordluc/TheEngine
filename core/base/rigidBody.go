package base

type RigidBody struct {
	o          Object
	quadtree   *QuadTree
	toSimulate bool
	isStatic   bool
}

func NewRigidBody(toSimulate, isStatic bool) RigidBody {
	return RigidBody{
		toSimulate: toSimulate,
		isStatic:   isStatic,
	}
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

func (a *RigidBody) ResolveCollision(b *RigidBody, v Vec[float32]) error {
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

	centerA := Vec[float32]{X: posA.X + sizeA.X/2, Y: posA.Y + sizeA.Y/2}
	centerB := Vec[float32]{X: posB.X + sizeB.X/2, Y: posB.Y + sizeB.Y/2}

	centerDistance := SubVecs(centerA, centerB)
	totalSize := AddVecs(sizeA, sizeB)
	distance := NewVec(centerDistance.X-totalSize.X/2, centerDistance.Y-totalSize.Y/2)

	if distance.X < 0 || distance.Y < 0 {
		return a.ResolveCollision(b, distance)
	}
	return nil
}
