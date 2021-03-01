package model

type ZSkipList struct {
	header *ZSkipListNode
	tail   *ZSkipListNode
	length int     //节点的数量

}
type ZSkipListNode struct {
	object   Sdshdr
	score    int64
	step     int             //跨度
	level  []*ZSkipListNode //记录目前跳跃表内，层数最大的那个节点的层数

	forward  *ZSkipListNode //前进指针
	backward *ZSkipListNode //后退指针

}
