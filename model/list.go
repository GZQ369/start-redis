package model

import (
	"errors"
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
		var sd  = (*redisObject)(res.ptr)
		d := (*ChainList)(sd.ptr)
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
		var d = (*redisObject)(res.ptr)
		n := (*ChainList)(d.ptr)
		s,err:= n.listDelNode(n.listFirst())

		res.lru = time.Now().Unix()
		return string(s.Buf), err
	}
}

func (r *RedisDb) Rpop(k string) (resp string, err error) {
	res, ok := r.dict[k]
	if !ok {
		return "", errors.New("key nil")
	} else {

		d := (*redisObject)(res.ptr)
		n := (*ChainList)(d.ptr)
		s,err:= n.listDelNode(n.listLast())

		res.lru = time.Now().Unix()
		return string(s.Buf), err
	}
}

func (r *RedisDb) Lrange(k string) ([]string,error) {
	res, ok := r.dict[k]
	if !ok {
		return []string{}, errors.New("key nil")
	} else {

		n := (*redisObject)(res.ptr)
		t := (*ChainList)(n.ptr)
		res,err:=t.listLRange(k)
		return res,err
	}
}
func (r *RedisDb) Lindex(k string, index int) (string,error) {
	res, ok := r.dict[k]
	if !ok {
		return "", errors.New("key nil")
	} else {

		n := (*redisObject)(res.ptr)
		t := (*ChainList)(n.ptr)
		s,err := t.listIndex(index)
		return string(s.Value.Buf),err
	}
}
func (r *RedisDb) Linsert(k,v string, index int) ([]string,error) {
	res, ok := r.dict[k]
	if !ok {
		return []string{}, errors.New("key nil")
	}else {
		n := (*redisObject)(res.ptr)
		t := (*ChainList)(n.ptr)
		newNode := new(ListNode)
		newNode.Value = Sdshdr{
			Buf: []byte(v),
		}

		t.listInsertNode(newNode,index)
		s,err:= t.listLRange(k)
		return s,err
	}

}

func (r *RedisDb) LSet(k,v string, index int) (string,error) {
	res, ok := r.dict[k]
	if !ok {
		return "", errors.New("key nil")
	} else {

		n := (*redisObject)(res.ptr)
		t := (*ChainList)(n.ptr)
		s,err := t.listLset(v,index)
		return s,err
	}
}