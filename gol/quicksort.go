package gol

import (
	"reflect"
)

type SortData struct {
	Less func(i, j int) bool
	Swap func(i, j int)
}

func Sort(data interface{}, lessThan func(i, j int) bool) {
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Slice {
		panic("Sort data is not slice")
	}
	swapper := reflect.Swapper(data)

	quicksort(SortData{Less: lessThan, Swap: swapper}, 0, v.Len())
}

func quicksort(data SortData, left, length int) {
	// fmt.Printf("quicksort:[%+v, %+v)\n", left, left+length)
	if length == 2 {
		if data.Less(left+1, left) {
			data.Swap(left, left+1)
		}
		return
	}
	if length <= 1 {
		return
	}
	i, j, pivot := left-1, left+length-1, left+length-1
	ch := make([]chan int, 3)
	for i := range ch {
		ch[i] = make(chan int)
	}
	go func() {
		<-ch[0]
		for {
			if i >= j {
				ch[2] <- 1
				break
			}
			i++
			if data.Less(pivot, i) {
				ch[2] <- 1
				<-ch[0]
			}
		}
	}()
	go func() {
		<-ch[1]
		for {
			if j <= i {
				ch[2] <- 1
				break
			}
			j--
			if data.Less(j, pivot) {
				ch[2] <- 1
				<-ch[1]
			}
		}
	}()

	for {
		ch[0] <- 1
		<-ch[2]
		ch[1] <- 1
		<-ch[2]
		// fmt.Printf("i, j, pivot = %+v, %+v, %+v\n", i, j, pivot)
		if i >= j {
			if data.Less(i, j) {
				data.Swap(j, pivot)
				pivot = j
			} else {
				data.Swap(i, pivot)
				pivot = i
			}
			break
		} else {
			data.Swap(i, j)
		}
	}

	quicksort(data, left, pivot-left)
	quicksort(data, pivot+1, length-(pivot-left+1))
}
