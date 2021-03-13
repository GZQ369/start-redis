package model

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//type Dict struct {
//	//key     KvObject map[string]redisObject
//	dType    dictType
//	privdata *unsafe.Pointer
//	dicter   [2]dictht
//	rehash   int64
//}
type dictht struct {
	table    []dictEntry
	size     int64
	sizeMask map[string]int64
	used     int64
}
type dictEntry struct {
	field map[string]Sdshdr
	next  *dictEntry //将多个hash值相同的键值对连接在一起，解决hash冲突问题
}
type dictType struct {
	//计算hash值的函数
}

func dictHtNew() *dictht {
	return &dictht{}
}
func (d *dictht) dictAdd(v []string) error {
	if len(v)%2 != 0 {
		return errors.New("ERR wrong number of arguments for HMSET")
	}
	var tb []dictEntry
	d.sizeMask = make(map[string]int64)
	for i := 0; i < len(v)-1; i += 2 {

		d.size++
		d.sizeMask[v[i]] = d.size-1

		tb = append(tb, dictEntry{field: map[string]Sdshdr{v[i]: Sdshdr{Buf: []byte(v[i+1])}}})

	}
	d.table = append(d.table, tb...)
	return nil
}

//将给定的值加入到字典中，如果键值已经存在于字典，那么新值取代原有的值
func (d *dictht) dictReplace(field []string) error {

	if len(field)%2 != 0 {
		return  errors.New("ERR wrong number of arguments for HMSET")
	}
	var tb []dictEntry

	for i := 0; i < len(field)-1; i += 2 {
		fmt.Println(d.sizeMask)
		if index, OK := d.sizeMask[field[i]]; OK {
			fmt.Println(field[i],index)
			d.table[index] = dictEntry{field: map[string]Sdshdr{field[i]: {Buf: []byte(field[i+1])}}}
		} else {

			d.size++
			d.sizeMask[field[i]] =  d.size

			tb = append(tb, dictEntry{field: map[string]Sdshdr{field[i]: {Buf: []byte(field[i+1])}}})

		}
	}
	d.table = append(d.table, tb...)
	return nil
}

//返回给定键的值,字段
func (d *dictht) dictFetchValue(field string) (dictEntry, error) {
	if v, ok := d.sizeMask[field]; ok {
		return d.table[v], nil
	} else {

		return dictEntry{}, errors.New("field not exits")
	}
}

//从字典随机返回一个键值对
func (d *dictht) dictGetRandomKey() (res dictEntry, err error) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(int(d.size))
	for k, v := range d.table {
		if k == n {
			return v, nil
		}
	}
	return res, nil
}

//从字典中删除给定键所对应的键值对
func (d *dictht) dictDelete(key string) (string,error) {
	if v, ok := d.sizeMask[key]; ok {
		d.table = append(d.table[:v], d.table[v+1:]...)
		delete(d.sizeMask, key)
		d.size--
		return "OK",nil
	} else {
		return "ERROR",errors.New("field not exits")
	}
}

//释放给定字典，以及字典包含的所有的键值对O（n）
func (d *dictht) dictRelease() {
	d.sizeMask = map[string]int64{}
	d.size = 0
	d.table = []dictEntry{}
	d.used = 0
}
