package model

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"unsafe"
)

type SdsString struct {
	buf string
	len int64
}
//使用整数值实现的字符串对象 :8个字节的长度，整形
type SdsInt struct {
	buf float64
}

func SdsIntNew(v float64) *SdsInt {
	return &SdsInt{buf: v}
}

//使用embestr编码实现的简单动态支付串对象,只能读取，长度小于39个字节，修改后转化为raw

//使用简单字符串实现的字符串对象,长度大于39个字

func (r *RedisDb) Set(key, value string) (res string, err error) {
	r.dict[key], err = stringObjectNew(value)
	res = "OK"
	return res, err
}

func (r *RedisDb) Get(key string) (s string, err error) {
	res, ok := r.dict[key]
	if !ok {
		err = errors.New("the key not exits")
		return "", err
	}
	if res.Enconding == sdsInt {
		var i *SdsInt = (*SdsInt)(res.ptr)
		s = strconv.FormatFloat(i.buf, 'f', -1, 64)
	} else {
		var j *Sdshdr = (*Sdshdr)(res.ptr)
		s = string(j.Buf)
	}
	return s, err
}
func (r *RedisDb) Append(key, v string) (out int, err error) {
	res, ok := r.dict[key]
	if !ok {
		sds := sdsHdrNew(v)
		r.dict[key] = redisObject{Type: String{}, Enconding: sdsHdr, lru: time.Now().Unix(), Refound: 1, ptr: unsafe.Pointer(sds)}
		out = sds.SdsLen()
	} else if res.Enconding == sdsInt {
		var i *SdsInt = (*SdsInt)(res.ptr)
		s := strconv.FormatFloat(i.buf, 'f', -1, 64)
		newSds := sdsHdrNew(s + v)
		res.ptr = unsafe.Pointer(newSds)
		res.Enconding = sdsHdr
		res.lru = time.Now().Unix()
		res.Refound++
		r.dict[key] = res
		out = newSds.SdsLen()
	} else {
		var i *Sdshdr = (*Sdshdr)(r.dict[key].ptr)
		i.SdsCatsDs(Sdshdr{Buf: []byte(v)})
		out = i.SdsLen()
	}
	return out, nil
}
func (r *RedisDb) Incrby(k string, v int64) (out string, err error) {
	res, ok := r.dict[k]
	if !ok {
		sds := SdsIntNew(float64(v))
		r.dict[k] = redisObject{Type: String{}, Enconding: sdsInt, lru: time.Now().Unix(), Refound: 1, ptr: unsafe.Pointer(sds)}
		out = string(v)
		return out, nil
	}
	if res.Enconding == sdsHdr {
		res.lru = time.Now().Unix()
		return "", errors.New("ERR value is not an integer or out of range")
	}
	var i *SdsInt = (*SdsInt)(res.ptr)
	i.buf = i.buf + float64(v)
	res.lru = time.Now().Unix()
	out = strconv.Itoa(int(i.buf))
	return out, nil
}
func (r *RedisDb) Decrby(k string, v int64) (out string, err error) {
	res, ok := r.dict[k]
	if !ok {
		sds := SdsIntNew(-float64(v))
		r.dict[k] = redisObject{Type: String{}, Enconding: sdsInt, lru: time.Now().Unix(), Refound: 1, ptr: unsafe.Pointer(sds)}
		out = string(v)
		return out, nil
	}
	if res.Enconding == sdsHdr {
		return "", errors.New("ERR value is not an integer or out of range")
	}
	var i *SdsInt = (*SdsInt)(res.ptr)
	i.buf = i.buf - float64(v)
	res.lru = time.Now().Unix()
	out = strconv.Itoa(int(i.buf))
	return out, nil
}
func (r *RedisDb) StrLen(k string) (l int, err error) {
	res, ok := r.dict[k]
	if !ok {
		err = errors.New("the key not exits")
		return 0, err
	}
	if res.Enconding == sdsInt {
		var i *SdsInt = (*SdsInt)(res.ptr)
		s := strconv.FormatFloat(i.buf, 'f', -1, 64)
		return len(s), nil
	}
	var i *Sdshdr = (*Sdshdr)(res.ptr)
	res.lru = time.Now().Unix()
	l = i.SdsLen()
	return l, nil
}

//将字符串特定索引上的值设置为给定的字符
func (r *RedisDb) SetRange(k,v string, index int,) (s string, err error) {
	res, ok := r.dict[k]
	if !ok {
		err = errors.New("the key not exits")
		return "", err
	}
	if res.Enconding==sdsInt{
		var i *SdsInt = (*SdsInt)(res.ptr)
		s = strconv.FormatFloat(i.buf, 'f', -1, 64)
		if index <len([]byte(s)){
			var byteTemp []byte
			temp := []byte(s)[index+1:]
			byteTemp = append(byteTemp,[]byte(s)[:index]...)
			byteTemp= append(byteTemp, []byte(v)...)
			byteTemp = append(byteTemp, temp...)
			s = string(byteTemp)
			re:=Sdshdr{Buf: byteTemp}
			re.Len = re.SdsLen()
			res.ptr = unsafe.Pointer(&re)
			res.Enconding = sdsHdr
			res.lru = time.Now().Unix()
			res.Refound++
			//r.dict[k] = res
		}else {
			err = errors.New("index integer or out of range")
			return
		}
		return s,nil
	}
	var i *Sdshdr = (*Sdshdr)(res.ptr)
	if index <i.SdsLen(){
		temp := i.Buf[index+1:]
		i.Buf = append([]byte{},i.Buf[:index]...)
		i.Buf = append(i.Buf, []byte(v)...)
		i.Buf = append(i.Buf, temp...)
		s = string(i.Buf)
	}else {
		err = errors.New("index integer or out of range")
		return
	}
	return
}

//拷贝对象所保存的整数值后转换成字符串值，然后取出指定索引的字符
func  (r *RedisDb) GetRange(k string, index int) (s string, err error) {
	res, ok := r.dict[k]
	if !ok {
		err = errors.New("the key not exits")
		return "", err
	}
	if res.Enconding == sdsInt {
		var i *SdsInt = (*SdsInt)(res.ptr)
		s = strconv.FormatFloat(i.buf, 'f', -1, 64)
		res.lru = time.Now().Unix()
		fmt.Println(s)
		if index <= len(s)-1 {
			s = string(s[index])
		}else {
			err = errors.New("index integer or out of range")
		}
		return s,err
	}
	var i *Sdshdr = (*Sdshdr)(res.ptr)
	res.lru = time.Now().Unix()
	if index <= i.SdsLen() -1 {
		s  = string(i.Buf[index])
	}else {
		err = errors.New("index integer or out of range")
	}
	return s,err
}
