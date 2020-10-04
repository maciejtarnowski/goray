package ray

import "goray/tuple"

type Computation struct {
	T float64
	Object Object
	Point *tuple.Tuple
	EyeV *tuple.Tuple
	NormalV *tuple.Tuple
	Inside bool
}

func (i *Intersection) PrepareComputations(r *Ray) *Computation {
	c := &Computation{}

	c.T = i.T
	c.Object = i.Object
	c.Point = r.Position(c.T)
	c.EyeV = r.Direction.Negate()
	c.NormalV = c.Object.NormalAt(c.Point)

	if c.NormalV.Dot(c.EyeV) < 0 {
		c.Inside = true
		c.NormalV = c.NormalV.Negate()
	} else {
		c.Inside = false
	}

	return c
}
