package redis

import (
	"errors"
	"unsafe"
)

type dict struct {
	key      map[string][]dictEntry
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
	field string
	value Sdshdr
	next  *dictEntry //将多个hash值相同的键值对连接在一起，解决hash冲突问题
}
type dictType struct {
	//计算hash值的函数
}

func dictCreate() dict {
	return dict{}
}
func (d dict) dictAdd(key string, v ...string) (dict, error) {
	if len(v)%2 != 0 {
		return dict{}, errors.New("ERR wrong number of arguments for HMSET")
	}
	d.dicter[0].size = int64(len(v) / 2)
	var tb []dictEntry
	for i := 0; i < len(v)-1; i++ {
		tb = append(tb, dictEntry{v[i], Sdshdr{Buf: []byte(v[i+1])}, nil})
	}
	d.dicter[0].table = tb
	d.key = map[string][]dictEntry{key:tb}
	return d, nil
}

//将给定的值加入到字典中，如果键值已经存在于字典，那么新值取代原有的值
func (d dict) dictReplace(key string, v ...string) (dict, error) {

	if len(v)%2 != 0 {
		return dict{}, errors.New("ERR wrong number of arguments for HMSET")
	}
	d.dicter[0].size = int64(len(v) / 2)
	var tb []dictEntry
	for i := 0; i < len(v)-1; i++ {
		tb = append(tb, dictEntry{v[i], Sdshdr{Buf: []byte(v[i+1])}, nil})
	}
	d.dicter[0].table = tb
	if _, ok :=d.key[key];ok {
		d.key = map[string][]dictEntry{key:tb}
	}
	return d, nil
}

//返回给定键的值
func (d dict)dictFetchValue(key string) ([]dictEntry, error) {
	if v, ok :=d.key[key];ok {
		return v,nil
	}else {
		return []dictEntry{},errors.New("keys not exits")
	}
}
//从字典随机返回一个键值对

//从字典中删除给定键所对应的键值对
//释放给定字典，以及字典包含的所有的键值对O（n）
