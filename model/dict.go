package model

import (
	"errors"
	"time"
	"unsafe"
)

func (r *RedisDb) Hset(k string, v ...string) error {
	if len(v)%2 != 0 {
		return errors.New("ERR wrong number of arguments for HMSET")
	}
	res, OK := r.dict[k]
	if !OK {
		dc := dictHtNew()
		err := dc.dictAdd(v)
		res = redisObject{Type: Hash{}, Enconding: dictht{}, lru: time.Now().Unix(), Refound: 1, ptr: unsafe.Pointer(dc)}
		r.dict[k] = res
		return err
	} else {
		d := (*dictht)(res.ptr)
		err := d.dictReplace(v)
		return err
	}
}

func (r *RedisDb) HgetAll(k string) ([]string, error) {
	res, OK := r.dict[k]
	if !OK {
		return []string{}, errors.New("empty array")

	} else {
		d := (*dictht)(res.ptr)
		var res []string
		for _, de := range d.table {
			for k, v := range de.field {
				res = append(res, k)
				res = append(res, string(v.Buf))

			}
		}
		return res, nil
	}
}

func (r *RedisDb) Hget(k, field string) (string, error) {
	res, OK := r.dict[k]
	if !OK {
		return "", errors.New("empty array")

	} else {
		d := (*dictht)(res.ptr)
		res, err := d.dictFetchValue(field)
		resp := res.field[field].Buf
		return string(resp), err
	}
}
//字段是否存在
func (r *RedisDb)Hexists(k, field string) bool {
	res, OK := r.dict[k]
	if !OK {
		return false

	} else {
		d := (*dictht)(res.ptr)
		_, err := d.dictFetchValue(field)
		if err !=nil{
			return false
		}
		return true
	}
}
func (r *RedisDb)HDel(k, field string) error {
	res, OK := r.dict[k]
	if !OK {
		return errors.New("key not exits")

	} else {
		d := (*dictht)(res.ptr)
		_, err := d.dictDelete(field)

		return err
	}
}

func (r *RedisDb)Hlen(k string) (int64,error) {
	res, OK := r.dict[k]
	if !OK {
		return 0,errors.New("key not exits")

	} else {
		d := (*dictht)(res.ptr)

		return d.size,nil
	}
}