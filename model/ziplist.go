package model

import "unsafe"

type ZipList struct {
	zlbytes int         //整个压缩列表的内存字节数
	zltail  uintptr     //记录压缩列表尾节点距离列表的起始地址有多少字节，通过偏移量可确定表尾节点的地址
	zllen   int64       //节点的数量
	entry   []entryNode //压缩列表的各个节点，长度由保存的内容决定
	zlend   int64       //用于标记压缩列表的末端
}
type entryNode struct { //压缩列表节点
	content             *Sdshdr
	previousEntryLength uintptr //上个节点的长度
}

func ZipListNew(v ...string) ZipList {
	res := ZipList{}
	entry := entryNode{}
	for i, v := range v {
		if i == 0 {
			entry.previousEntryLength = 0
		} else {
			entry.previousEntryLength = uintptr(entry.content.Len)
		}
		entry.content = sdsHdrNew(v)

		res.entry = append(res.entry, entry)
		res.zlbytes += entry.content.Len
		res.zllen += 1
		res.zltail += unsafe.Offsetof(entry.content)
	}
}
