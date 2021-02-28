package redis

import (
	"unsafe"
)

type ListNode struct {
	prev  *ListNode
	Next  *ListNode
	Value *interface{}
}
type ChainList struct {
	head  *ListNode
	tail  *ListNode
	len   int
	dup   *unsafe.Pointer //复制节点所报存的值
	free  *unsafe.Pointer //释放节点所报存的值
	match *unsafe.Pointer //比对链表节点的值和另一个输入值是否相等
}
func ifyType(x interface{}) unsafe.Pointer {
	switch v := x.(type) {
	case string:
		return unsafe.Pointer(&v)
	case int:
		return unsafe.Pointer(&v)
	case bool:
		return unsafe.Pointer(&v)
	default:
		return unsafe.Pointer(&v)

	}
}
//初始化
func listCreate() ChainList {
	return ChainList{}
}

//将一个包含给定值的新节点添加到给定链表的表头
func (c *ChainList)listAddNodeHead(node ListNode) *ChainList{
	c.head = &node
	return c
}

//将一个包含给点值的新节点添加到给定链表的表尾
func (c *ChainList)listAddNodeTail(node ListNode) *ChainList{
	c.tail = &node
	return c
}
//将一个包含给定值的新节点添加到给定节点的之前或之后
func (c *ChainList)listInsertNode(node ListNode) *ChainList{

	return c
}
//查找并返回链表中包含给定值的节点
func (c *ChainList)listSearchKey(v *interface{}) ListNode {
	var res *ListNode
	res = c.head

	for res != nil{
		res = res.Next
	}
}
//返回链表在给定索引上的节点
//从链表中删除给定节点
//将链表的表尾节点弹出，然后将被弹出的节点插入链表的表头，成为新的表头节点
//复制一个给定链表的副本
//复制一个给定链表的副本
//释放给定链表，以及链表中的所有节点
