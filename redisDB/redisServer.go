package redisDB

import "../model"
type redisServer struct {
	db []*model.RedisDb   //存有16个数组
	dbNum   int  //数据库的数量
}

