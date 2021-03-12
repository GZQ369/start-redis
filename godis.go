package main

import (
	"./model"
	"fmt"
	"sort"
)

type n struct {
	d *int
}
type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}
func main() {
	db :=model.RedisNew()

	dd,_:=db.Rpush("fdsg", "qqqqadddddddddddddd","fd23","dff","fdfsa","llllll")
	db.Linsert("fdsg","d11111111232432",4)
	sdsd,df := db.LSet("fdsg","3344567",0)
	ls,_ :=db.Lrange("fdsg")
	db.Hset("234","df","FD","gfd","5426","www","999999")
	db.Hset("234","df","F34D","gfdsfd","5426","www","999999")
	db.HDel("234","gfd")
	qw,er:= db.HgetAll("234")
	dds, _ := db.Hlen("234")

	l:=[]string{"331","123","5","35"}
	sort.Strings(l)
	a:= []int64{12,1245,45456,1,2415245612,12,132346376,146574979,14,6597946378,65746546544,12524,123}
	sort.Slice(a, func(i, j int) bool {
		return a[i]>a[j]
	})
	fmt.Println(dd,sdsd,df,ls,qw,er,dds, a)
	var peo People = &Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
