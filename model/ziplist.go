package model

type ZipList struct {
	zlbytes int64       //整个压缩列表的内存字节数
	zltail  int64       //记录压缩列表尾节点距离列表的起始地址有多少字节，通过偏移量可确定表尾节点的地址
	zllen   int64       //节点的数量
	entry   []entryNode //压缩列表的各个节点，长度由保存的内容决定
	zlend   int64       //用于标记压缩列表的末端
}
type entryNode struct { //压缩列表节点
	content                 []byte
	previousEntryLength uintptr //上个节点的长度
}


