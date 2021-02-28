package redis

import (
	"fmt"
	"unsafe"
)

type dict struct {
	dType    dictType
	privdata *unsafe.Pointer
	dicter   [2]dictht
	rehash   int64
}
type dictht struct {
	table    []dictEntry
	size     int64
	sizeMask int64
	used     int64
}
type dictEntry struct {
	key   string
	value []*unsafe.Pointer
	next  *dictEntry //将多个hash值相同的键值对连接在一起，解决hash冲突问题
}
type dictType struct {
	//计算hash值的函数
}

func dictCreate() dict {
	return dict{}
}
func (d dict)dictAdd(key string, v ...interface{}) dict {
	d.dicter[0].table[0].key = key
	for _, arg := range v { //迭代不定参数
		ag:=arg
		switch arg.(type) {
		case int:
			res:= append(d .dicter[0].table[0].value, unsafe.Pointer(&ag))
		case string:
			fmt.Println(arg, "is string")
		case float64:
			fmt.Println(arg, "is float64")
		case bool:
			fmt.Println(arg, " is bool")
		case int64:
			fmt.Println(arg, " is bool")
		default:
			fmt.Println("未知的类型")
		}
	}
}