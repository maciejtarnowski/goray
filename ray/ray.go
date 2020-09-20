package ray

import (
	"goray/matrix"
	"goray/tuple"
	"math"
)

type Ray struct {
	Origin    *tuple.Tuple
	Direction *tuple.Tuple
}

func NewRay(origin, dir *tuple.Tuple) *Ray {
	return &Ray{Origin: origin, Direction: dir}
}

func (r *Ray) Position(t float64) *tuple.Tuple {
	return r.Origin.Add(r.Direction.Multiply(t))
}

func (r *Ray) Transform(m *matrix.Matrix) *Ray {
	return &Ray{Origin: m.MultiplyTuple(r.Origin), Direction: m.MultiplyTuple(r.Direction)}
}

type Intersection struct {
	T      float64
	Object Object
}

func NewIntersection(t float64, o Object) *Intersection {
	return &Intersection{T: t, Object: o}
}

type Intersections struct {
	elements []*Intersection
}

func NewIntersections(elements ...*Intersection) *Intersections {
	return &Intersections{elements: elements}
}

func (is *Intersections) Add(i *Intersection) {
	is.elements = append(is.elements, i)
}

func (is *Intersections) Len() int {
	return len(is.elements)
}

func (is *Intersections) Swap(i, j int) {
	is.elements[i], is.elements[j] = is.elements[j], is.elements[i]
}

func (is *Intersections) Less(i, j int) bool {
	return is.ValueAt(i) < is.ValueAt(j)
}

func (is *Intersections) ValueAt(index int) float64 {
	return is.elements[index].T
}

func (is *Intersections) Get(index int) *Intersection {
	return is.elements[index]
}

func (is *Intersections) GetAll() []*Intersection {
	return is.elements
}

func (is *Intersections) ObjectAt(index int) Object {
	return is.Get(index).Object
}

func (is *Intersections) Hit() *Intersection {
	lowest := math.MaxFloat64
	lowestIndex := -1

	for index, x := range is.elements {
		if x.T > 0 && x.T < lowest {
			lowest = x.T
			lowestIndex = index
		}
	}

	if lowestIndex == -1 {
		return nil
	}
	return is.Get(lowestIndex)
}
