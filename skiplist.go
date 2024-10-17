package main

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Byte-keyed skip list example

var max_tower int

type node struct {
	key   []byte
	val   []byte
	links []*node
}

type SkipList struct {
	head          *node
	height        int
	probabilities map[int]uint32
	rnd           *rand.Rand
}

func NewSkipList(max int) *SkipList {
	probs := make(map[int]uint32)

	p := 1.0

	for i := 0; i < max; i++ {
		probs[i] = uint32(p * float64(math.MaxUint32))
		p *= 0.5
	}

	sl := &SkipList{
		height:        1,
		probabilities: probs,
		rnd:           rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	sl.head = &node{
		links: make([]*node, max),
	}
	return sl
}

func (sl *SkipList) randLevel() int {
	dice := sl.rnd.Uint32()

	max := len(sl.probabilities)
	height := 1
	for height < max && dice <= sl.probabilities[height] {
		height++
	}

	return height
}

func (sl *SkipList) Get(key []byte) ([]byte, error) {
	found, _ := sl.lookup(key)

	if found == nil {
		return nil, errors.New("key not found")
	}

	return found.val, nil
}

func (sl *SkipList) Insert(key []byte, val []byte) {
	found, path := sl.lookup(key)

	if found != nil {
		found.val = val
		return
	}

	height := sl.randLevel()
	nd := &node{key: key, val: val, links: make([]*node, len(sl.probabilities))}

	for level := 0; level < height; level++ {
		prev := path[level]

		if prev == nil { // если превысили высоту текущей башни, т.е. на предложенном уровне еще нет списка
			prev = sl.head
		}
		nd.links[level] = prev.links[level]
		prev.links[level] = nd
	}

	if height > sl.height {
		sl.height = height
	}
}

func (sl *SkipList) Delete(key []byte) error {
	found, path := sl.lookup(key)

	if found == nil {
		return errors.New("key not found")
	}

	for h := sl.height - 1; h >= 0; h-- {
		if path[h].links[h] != found {
			continue
		}
		path[h].links[h] = found.links[h]
		found.links[h] = nil
		if sl.head.links[h] == nil {
			sl.height--
		}

	}

	found = nil

	return nil
}

func (sl *SkipList) lookup(key []byte) (*node, []*node) {
	var next *node
	path := make([]*node, len(sl.probabilities)) // TODO: refactor this

	prev := sl.head
	for level := sl.height - 1; level >= 0; level-- {
		for next = prev.links[level]; next != nil; next = prev.links[level] {
			if bytes.Compare(key, next.key) <= 0 {
				break
			}
			prev = next
		}
		path[level] = prev
	}

	if next != nil && bytes.Equal(key, next.key) {
		return next, path
	}
	return nil, path
}

func (sl *SkipList) Print() {
	var res string

	for level := sl.height - 1; level >= 0; level-- {
		res += fmt.Sprintf("%d ", level)
		for i, next := 0, sl.head.links[level]; next != nil; i, next = i+1, next.links[level] {
			res += fmt.Sprintf("-> %v ", string(next.key))
		}
		res += "\n"
	}

	fmt.Println(res)
}
