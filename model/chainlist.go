package model

import (
	"bytes"
	"errors"
	"unsafe"
)

type ListNode struct {
	prev  *ListNode
	Next  *ListNode
	Value Sdshdr
}
type ChainList struct {
	head  *ListNode
	tail  *ListNode
	len   int64
	//dup   unsafe.Pointer //复制节点所报存的值
	//free  unsafe.Pointer //释放节点所报存的值
	//match unsafe.Pointer //比对链表节点的值和另一个输入值是否相等
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
func ChainListCreate() *ChainList {
	res := new(ChainList)
	res.len = 0
	return res
}

//将一个包含给定值的新节点添加到给定链表的表头
func (c *ChainList) listAddNodeHead(data Sdshdr) {
	newNode := new(ListNode)
	newNode.Value = data

	if c.GetSize() == 0 { // head
		c.head = newNode
		c.tail = newNode
		newNode.prev = nil
		newNode.Next = nil
	} else { //  tail
		newNode.Next = c.head
		c.head.prev = newNode
		c.head = newNode

	}
	c.len++
}

func (c *ChainList) GetSize() int64 {
	return c.len
}

//将一个包含给点值的新节点添加到给定链表的表尾
func (c *ChainList) listAddNodeTail(data Sdshdr) {
	//res := c.tail
	//c.tail = &node
	//node.prev = res
	//node.Next = nil
	//return c
	newNode := new(ListNode)
	newNode.Value = data

	if c.GetSize() == 0 { // head
		c.head = newNode
		c.tail = newNode
		newNode.prev = nil
		newNode.Next = nil
	} else { //  tail
		newNode.prev = c.tail
		c.tail.Next = newNode
		c.tail = newNode
	}
	c.len++
}

//将一个包含给定值的新节点添加到给定节点的之前或之后
func (c *ChainList) listInsertNode(node ListNode) *ChainList {

	return c
}

//查找并返回链表中包含给定值的节点
func (c *ChainList) listSearchKey(v string) (*ListNode, error) {
	var res *ListNode
	res = c.head

	for res != nil {
		if bytes.Equal(res.Value.Buf, []byte(v)) {
			return res, nil
		}
		res = res.Next
	}
	return nil, errors.New("key node not exits")
}

////返回链表在给定索引上的节点
//func (c *ChainList) listIndex(index int) Sdshdr {
//	var cur int = 0
//	res := c.head
//	for cur <= index && res {
//
//	}
//}

//从链表中删除给定节点
func (c *ChainList) listDelNode(v ListNode) (Sdshdr,error) {
	//head := c.head
	//if string(head.Value.Buf) == string(v.Value.Buf){
	//	c.head = head.Next
	//}else {
	//	for head != nil {
	//		if string(head.Value.Buf) == string(v.Value.Buf) {
	//			head.prev.Next = head.Next
	//			if head.Next != nil{
	//				head.Next.prev = head.prev
	//			}
	//			break
	//		} else {
	//			head = head.Next
	//		}
	//	}
	//}
	res:= c.head
	for res != nil{
		if bytes.Equal(res.Value.Buf, v.Value.Buf){
			res.prev.Next = res.Next
			if res.Next !=nil{
				res.Next.prev = res.prev
			}

			return v.Value,nil
		}
		res = res.Next
	}
	return Sdshdr{},errors.New("list out of element")
}
//输出节点的值
func (c *ChainList)listLRange(k string, start,stop int) ([]string,error) {
	res:= c.head
	var ls []string
	for res !=nil{
		ls= append(ls, string(res.Value.Buf))
		res = res.Next
	}
	return ls,nil
}

func (c *ChainList) listFirst() ListNode {
	return *c.head
}
func (c *ChainList) listLast() ListNode {
	return *c.tail
}

//将链表的表尾节点弹出，然后将被弹出的节点插入链表的表头，成为新的表头节点
//复制一个给定链表的副本
//复制一个给定链表的副本
//释放给定链表，以及链表中的所有节点
