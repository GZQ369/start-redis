package model

import "strconv"

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

func (r *RedisDb)Lpush(k string, v ...string) (le int, err error) {
	r.dict[k]
}

