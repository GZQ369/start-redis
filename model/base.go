package model

import "bytes"

type Sdshdr struct {
	//记录buf数组中已经使用字节的数量
	//等于sds所报存字符串的长度
	Len int
	//no use
	Free int
	//byte array for save
	Buf  []byte
}

func sdsHdrNew(v string) *Sdshdr {
	res:=[]byte(v)
	return &Sdshdr{Len: len(res),Buf: res}
}
//返回一个空的SDS
func SdsEmpty(key string) Sdshdr {
	return Sdshdr{}
}
func (s *Sdshdr)SdsLen() int {
	s.Len = len(s.Buf)
	return s.Len
}
//创建一个给定SDS的副本
func (s *Sdshdr)SdsUp() Sdshdr {
	return Sdshdr{
		Len: s.Len,
		Buf: s.Buf,
	}
}
//清空SDS保存的字符串内容
func (s *Sdshdr)SdsClear() *Sdshdr {
	s.Buf=[]byte{}
	return s
}
//将给定的字符串拼接到SDS字符串的末尾
func (s *Sdshdr)SdsCat(str string) *Sdshdr {
	s.Buf = append(s.Buf, []byte(str)...)
	return s
}
//将给定的sds字符串拼接到另一个SDS字符串的末尾
func (s *Sdshdr)SdsCatsDs(sds1 Sdshdr) *Sdshdr {
	s.Buf = append(s.Buf, sds1.Buf...)
	return s
}
//将给定的C字符串复制到Sds里面，覆盖SDS原有的字符串
func (s *Sdshdr)SdsCpy(str string) *Sdshdr {
	s.Buf = []byte(str)
	return s
}
//用空字符将SDS扩展到指定的长度
func (s *Sdshdr)SdsGrowZero(long int) *Sdshdr {
	for i:=0;i<long;i++{
		s.Buf = append(s.Buf, 32)
		s.Len++
	}
	return s
}
//保留给定区间内的数据，不在区间内的数据会被清除或者覆盖
func (s *Sdshdr)SdsRange()  {}
//接受一个SDS和一个C字符串作为参数，从Sds移除所有在C字符串中出现的字符
func (s *Sdshdr)SdsTrim()  {

}
//对比两个SDS字符串是否相同
func (s *Sdshdr)SdsCmp(sds Sdshdr)  bool {
	return bytes.Equal(s.Buf, sds.Buf)
}