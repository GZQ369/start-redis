package model

import (
	"math/rand"
	"sync"
	"time"
)

type ZSkipList struct {
	head   *ZSkipListNode
	tail   *ZSkipListNode
	length int //节点的数量
	level  int
	mut    *sync.RWMutex
	rand   *rand.Rand
}
type ZSkipListNode struct {
	object Sdshdr
	score  int64
	next   []*ZSkipListNode //指针

}

func (list *ZSkipList) randomLevel() int {
	level := 1
	for ; level < list.level && list.rand.Uint32()&0x1 == 1; level++ {
	}
	return level
}
func NewZSkipList(level int) *ZSkipList {

	list := &ZSkipList{}
	if level <= 0 {
		level = 32
	}
	list.level = level
	list.head = &ZSkipListNode{next: make([]*ZSkipListNode, level, level)}
	list.tail = &ZSkipListNode{}
	list.mut = &sync.RWMutex{}
	list.rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	for index := range list.head.next {
		list.head.next[index] = list.tail
	}
	return list
}
func (list *ZSkipList) ZSkipListAdd(score int64, ob Sdshdr) {

	list.mut.Lock()
	defer list.mut.Unlock()
	//1.确定插入深度
	level := list.randomLevel()
	//2.查找插入部位
	update := make([]*ZSkipListNode, level, level)
	node := list.head

	for index := level - 1; index >= 0; index-- {
		for {
			node1 := node.next[index]
			if node1 == list.tail || node1.score > score {
				update[index] = node //找到一个插入部位
				break
			} else if node1.score == score {
				node1.object = ob
				return
			} else {
				node = node1
			}
		}

	}
	//3.执行插入
	newNode := &ZSkipListNode{ob, score, make([]*ZSkipListNode, level, level)}
	for index, node := range update {
		node.next[index], newNode.next[index] = newNode, node.next[index]
	}
	list.length++
}

func (list *ZSkipList) ZSkipLength() int {

	list.mut.RLock()
	defer list.mut.RUnlock()
	return list.length
}
func (list *ZSkipList) ZSkipRemove(key int64) bool {

	list.mut.Lock()
	defer list.mut.Unlock()
	//1.查找删除节点

	node := list.head
	remove := make([]*ZSkipListNode, list.level, list.level)
	var target *ZSkipListNode
	for index := len(node.next) - 1; index >= 0; index-- {
		for {
			node1 := node.next[index]
			if node1 == list.tail || node1.score > key {
				break
			} else if node1.score == key {

				remove[index] = node //找到啦
				target = node1
				break
			} else {
				node = node1
			}
		}
	}

	//2.执行删除
	if target != nil {
		for index, node1 := range remove {
			if node1 != nil {
				node1.next[index] = target.next[index]
			}
		}
		list.length--
		return true
	}
	return false
}
func (list *ZSkipList) ZSkipFind(key int64) interface{} {

	list.mut.RLock()
	defer list.mut.RUnlock()
	node := list.head
	for index := len(node.next) - 1; index >= 0; index-- {
		for {
			node1 := node.next[index]
			if node1 == list.tail || node1.score > key {
				break
			} else if node1.score == key {
				return node1.object
			} else {
				node = node1
			}
		}
	}
	return nil
}
