package model

import (
	"errors"
	"fmt"
	"time"
	"unsafe"
)

//压缩列表
//
////双端链表
//type doubleSide struct {
//	head int
//	tail int
//	len int
//	free int
//}
//type chainNode struct {
//	head *chainNode
//	tail *chainNode
//	value redisObject
//}

func (r *RedisDb) Lpush(k string, v ...string) (int64, error) {
	res, OK := r.dict[k]
	if !OK {
		cl := ChainListCreate()
		for _, va := range v {
			cl.listAddNodeHead(*sdsHdrNew(va))
		}
		tmp, _ := listObjectNew()
		tmp.ptr = unsafe.Pointer(cl)
		tmp.lru = time.Now().Unix()
		res = redisObject{Type: List{}, Enconding: linkedList, lru: time.Now().Unix(), Refound: int(cl.GetSize()), ptr: unsafe.Pointer(&tmp)}
		return cl.GetSize(), nil
	} else {

		for _, va := range v {
			(*ChainList)(res.ptr).listAddNodeHead(*sdsHdrNew(va))
			(*ChainList)(res.ptr).len++
		}
		res.lru = time.Now().Unix()
		return (*ChainList)(res.ptr).GetSize(), nil
	}
}

func (r *RedisDb) Rpush(k string, v ...string) (int64, error) {
	res, ok := r.dict[k]
	if !ok {
		cl := ChainListCreate()
		for _, va := range v {
			cl.listAddNodeTail(*sdsHdrNew(va))
		}
		tmp, _ := listObjectNew()
		tmp.ptr = unsafe.Pointer(cl)
		tmp.lru = time.Now().Unix()
		r.dict[k] = redisObject{Type: List{}, Enconding: linkedList, lru: time.Now().Unix(), Refound: int(cl.GetSize()), ptr: unsafe.Pointer(&tmp)}
		return cl.GetSize(), nil
	} else {
		var d *ChainList = (*ChainList)(res.ptr)
		for _, va := range v {
			d.listAddNodeTail(*sdsHdrNew(va))
			d.len++
		}
		res.lru = time.Now().Unix()
		return (*ChainList)(res.ptr).GetSize(), nil
	}
}

func (r *RedisDb) Lpop(k string) (resp string, err error) {
	res, ok := r.dict[k]
	if !ok {
		return "", errors.New("key nil")
	} else {
		var d  = *(*ChainList)(res.ptr)
		n := d.listFirst()
		//s := d.listDelNode(n)
		fmt.Println(34535)

		res.lru = time.Now().Unix()

		return string(n.Value.Buf), nil
	}
}

func (r *RedisDb) Rpop(k string) (resp string, err error) {
	res, ok := r.dict[k]
	if !ok {
		return "", errors.New("key nil")
	} else {

		n := *(*redisObject)(res.ptr)

		t:=(*ChainList)(n.ptr)
		s,_ := t.listDelNode(t.listLast())

		res.lru = time.Now().Unix()
		return string(s.Buf), nil
	}
}
func (r *RedisDb)Lrange(k string) ([]string,error) {
	res, ok := r.dict[k]
	if !ok {
		return []string{}, errors.New("key nil")
	} else {
		var tmp = (*redisObject)(res.ptr)
		var t = (*ChainList)(tmp.ptr)
		ls,_:=t.listLRange(k,0,int(t.len))
		return ls,nil
	}
}
func (r *RedisDb)Lindex(k string,index int )  {

}