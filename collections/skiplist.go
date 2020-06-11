package collections

type SkipList struct {
	_maxLevel        int
	_currentMaxLevel int
	_comparator      func(k1, k2 Key) int

	_head *Node
}

//
func NewSkipList(comparator func(k1, k2 Key) int, maxLevel int) *SkipList {
	head := &Node{
		K:     Key{},
		_next: make([]*Node, 0, maxLevel),
	}
	for i := 0; i < maxLevel; i++ {
		head.SetNext(i, nil)
	}
	return &SkipList{
		_maxLevel:        maxLevel,
		_currentMaxLevel: 1,
		_comparator:      comparator,
		_head:            head,
	}
}

func (list *SkipList) getMaxHeight() int {
	return list._currentMaxLevel
}

func (list *SkipList) randomHeight() int {
	return 0
}

func (list *SkipList) keyIsAfterNode(k Key, n *Node) bool {
	return (n != nil) && (list._comparator(k, n.K) > 0)
}

func (list *SkipList) findGreaterOrEqual(k Key, prev []*Node) *Node {
	x := list._head
	level := list.getMaxHeight() - 1
	for true {
		next := x.Next(level)
		if list.keyIsAfterNode(k, next) {
			x = next
		} else {
			//  1       9
			//  1   5   9
			//  1 2 5 7 9
			if prev != nil {
				prev[level] = x
			}
			if level == 0 {
				return next
			} else {
				level--
			}
		}
	}
	return nil
}

func (list *SkipList) newNode(key Key, level int) *Node {
	return nil
}

func (list *SkipList) equal(k, k1 Key) bool {
	return true
}

func (list *SkipList) Insert(k Key) {
	prev := make([]*Node, 0, list._maxLevel)
	// find before the k and after the k
	afterNode := list.findGreaterOrEqual(k, prev)
	if list.equal(k, afterNode.K) {
		return
	}
	height := list.randomHeight()
	if height > list.getMaxHeight() {
		for i := list.getMaxHeight(); i < list._maxLevel; i++ {
			prev[i] = list._head
		}
		list._currentMaxLevel = height
	}
	x := list.newNode(k, height)

	for i := 0; i < height; i++ {
		x.SetNext(i, prev[i].Next(i))
		prev[i].SetNext(i, x)
	}
}

type Node struct {
	K     Key
	_next []*Node
}

func (n *Node) Next(level int) *Node {
	if level < 0 {
		return nil
	}
	return n._next[level]
}

func (n *Node) SetNext(level int, node *Node) {
	n._next[level] = node
}

// TODO
type Key struct {
}
