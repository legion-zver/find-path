package findpath3d

import (
    "sync"
    "math"
    "errors"    
)

// Node - интерфейс ноды графа
type Node interface {
    ID() int    
    Tag() string    
    Position() (float64, float64, float64)   
    Distance(to Node) (float64, error) 
}

// Edge - интерфейс ребер графа 
type Edge interface {
    Source() Node
    Target() Node
    Weight(method int) float64
    SetMethodFactor(method int, factor float64)
}

// Нода
type node struct {
    id  int
    tag string
    x   float64
    y   float64
    z   float64
}

func (n *node) ID() int {
    return n.id
}

func (n *node) Tag() string {
    return n.tag
}

func (n *node) Position() (float64, float64, float64) {
    return n.x, n.y, n.z
}

func (n *node) Distance(to Node) (float64, error) {
    if to != nil {
        xTo, yTo, zTo := to.Position()
        return math.Abs(math.Sqrt(math.Pow(xTo-n.x,2)+math.Pow(yTo-n.y,2)+math.Pow(zTo-n.z,2))), nil 
    }
    return math.MaxFloat64, errors.New("to node is nil")
}

// Ребро
type edge struct {
    source  Node
    target  Node
    factors map[int]float64
}

func (e *edge) Source() Node {
    return e.source
}

func (e *edge) Target() Node {
    return e.target
}

func (e *edge) Weight(method int) float64 {
    f := float64(1.0)
    if v, ok := e.factors[method]; ok {
        f = v
    }
    if f <= 0 { f = 0.00000001 }
    if e.source != nil {
        if w, err := e.source.Distance(e.target); err == nil {
            return w*f
        }
    }
    return math.MaxFloat64
}

func (e *edge) SetMethodFactor(method int, factor float64) {
    if e.factors == nil {
        e.factors = make(map[int]float64)
    }
    e.factors[method] = factor
}

// Graph - интерфейс графа
type Graph interface {

}

// Граф
type graph struct {
    mutex sync.RWMutex

    nodes map[int]Node
    edges map[int]map[int]Edge
}