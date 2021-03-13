package model

import (
	"errors"
	"reflect"
	"sort"
)

type IntSet struct {
	encoding string
	length   int64
	contents []string
}

func IntSetNew(v []string) *IntSet {
	return &IntSet{encoding: intSet, length: int64(len(v)),contents: v}
}
//将指定元素添加到整数集合里面
func (i *IntSet) intSetAdd(k string) IntSet {
	i.contents = append(i.contents, k)
	sort.Strings(i.contents)
	i.contents = Duplicate(i.contents)
	i.length++
	return *i
}
//从整数集合中移除给定元素
func (i *IntSet) intSetRemove(k string) (IntSet, error) {
	for item := int64(0); item < i.length; item++ {
		if i.contents[item] == k {
			i.contents = append(i.contents[:item], i.contents[item+1:]...)
			i.length--
			return *i, nil
		}

	}
	return *i, errors.New("key not exits")
}
//查找是否存在对应的键，是否存在于集合
func (i *IntSet) intSetFind(k string) bool {
	var res  []string =i.contents
	low,high := 0,len(res)-1

	for low<high{
		mid :=(low+high)>>1
		if res[mid] < k {
			low = mid+1
		}else if res[mid]==k{
			return true
		}else {
			high = mid -1
		}
	}
	return false
}
//取出给定索引上的元素
func (i *IntSet)intSetGet(k int64) string {
	return i.contents[k]
}
//f返回整数集合包含的元素的个数
func (i IntSet)intSetLen() int64 {
	return i.length
}
//返回整数集合占用的内存字节数
func (i IntSet)intSetBlobLen() (res int64 ){
	for _,item:=range i.contents{
		res+=int64(len([]byte(item)))
	}
	return res
}
//去重
func Duplicate(a interface{}) (ret []string) {
	va := reflect.ValueOf(a)
	for i := 0; i < va.Len(); i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i).String())
	}
	return ret
}
