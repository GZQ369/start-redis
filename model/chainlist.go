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
	head *ListNode
	tail *ListNode
	len  int64
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
func (c *ChainList) listInsertNode(newNode *ListNode, index int) (*ChainList, error) {
	if int64(index) >= c.len {
		return nil, errors.New("the index out of list")
	}
	ls, _ := c.listIndex(index) //出入至该ls节点之前

	if bytes.Equal(ls.Value.Buf, c.head.Value.Buf) {
		c.listAddNodeHead(newNode.Value)

		return c, nil

	} else if bytes.Equal(ls.Value.Buf, c.tail.Value.Buf) {
		c.listAddNodeTail(newNode.Value)
		return c, nil
	}
	newNode.Next = ls
	newNode.prev = ls.prev
	ls.prev = newNode
	newNode.prev.Next = newNode
	c.len++
	return c, nil
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
//lset更新指定索引上的值
func (c *ChainList)listLset(v string, index int) (string,error) {
	ls,err :=c.listIndex(index)
	if err !=nil{
		return "", err
	}
	ls.Value.Buf = []byte(v)
	return "OK",err
}
////返回链表在给定索引上的节点
func (c *ChainList) listIndex(index int) (*ListNode, error) {
	var cur int = 0
	res := c.head
	for ; cur < index && res != nil; cur++ {
		res = res.Next
	}
	if res != nil {
		return res, nil
	} else {
		return nil, errors.New("the index out of list")
	}
}

//从链表中删除给定节点
func (c *ChainList) listDelNode(v *ListNode) (Sdshdr, error) {
	//head := c.head
	//if bytes.Equal(head.Value.Buf, v.Value.Buf){
	//	c.head = head.Next
	//	return v.Value,nil
	//}else {
	//	for head != nil {
	//		if bytes.Equal(head.Value.Buf, v.Value.Buf) {
	//			head.prev.Next = head.Next
	//			if head.Next != nil {
	//				head.Next.prev = head.prev
	//			}
	//			return v.Value, nil
	//		} else {
	//			head = head.Next
	//		}
	//	}
	//}
	if v == nil {
		return Sdshdr{}, errors.New("list out of index")
	}

	prev := (*v).prev
	next := (*v).Next

	if c.IsHead(v) {
		c.head = next
	} else {
		(*prev).Next = next
	}

	if c.IsTail(v) {
		c.tail = prev
	} else {
		(*next).prev = prev
	}
	c.len--

	return v.Value, nil
}
func (c *ChainList) listFirst() *ListNode {
	return c.head
}
func (c *ChainList) listLast() *ListNode {
	return c.tail
}

func (c *ChainList) listLRange(k string) ([]string, error) {
	temp := c.head
	var res []string
	for temp != nil {
		res = append(res, string(temp.Value.Buf))
		temp = temp.Next
	}
	return res, nil
}

//将链表的表尾节点弹出，然后将被弹出的节点插入链表的表头，成为新的表头节点
//复制一个给定链表的副本
//复制一个给定链表的副本
//释放给定链表，以及链表中的所有节点
func (c *ChainList) IsHead(l *ListNode) bool {
	if c.head == l {
		return true
	}
	return false
}

func (c *ChainList) IsTail(l *ListNode) bool {
	if c.tail == l {
		return true
	}
	return false
}
