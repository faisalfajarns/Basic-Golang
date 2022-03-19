package helper

import (
	"container/list"
	"fmt"
)

func DoubleLinkedList() {
	data := list.New()

	data.PushBack("Faisal")
	data.PushBack("Fajar")
	data.PushBack("Nursaid")

	fmt.Println(*data.Front())
	for e := data.Front(); e != nil; e = e.Next(){

		fmt.Println(e.Value)
	
	}

}