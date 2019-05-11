package easy

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func groupEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for index, _ := range a {
		if a[index] != b[index] {
			return false
		}
	}
	return true
}


func TestTwoSum(t *testing.T) {
	for _, unit := range []struct {
		m        []int
		n          int
		expected []int
	}{
		{[]int{2, 7, 11, 5}, 9, []int{0, 1} },
		{[]int{2, 7, 11, 5}, 7, []int{0, 3} },
		{[]int{2, 7, 11, 5}, 16, []int{2, 3} },
		{[]int{2, 7, 11, 5}, 18, []int{1, 2} },
		{[]int{3, 2, 4}, 6, []int{1, 2} },
	}{
		if actually := twoSum(unit.m, unit.n); !groupEqual(actually, unit.expected){
			t.Errorf("twoSum: [%v], actually: [%v]", unit, actually)
		}
		if actually := twoSumOnePass(unit.m, unit.n); !groupEqual(actually, unit.expected){
			t.Errorf("twoSum: [%v], actually: [%v]", unit, actually)
		}
	}
}

func TestMap(t *testing.T) {
	m := map[int]int{1:1, 2:2, 3:3, 4:4}
	for k,v := range m {
		t.Log(k, v)
	}
}

func TestDefer(t *testing.T) {
	fmt.Println(1)
	defer fmt.Println(111)
	fmt.Println(2)
	for i:=0;i<3;i++ {
		defer fmt.Println(i)
	}
	fmt.Println(3)
	defer fmt.Println(222)
}

type Employee struct {
	Id string
	Name string
	Age int
}

func TestCreateEmployee(t *testing.T) {
	e1 := Employee{"1","aa",1}
	e2 := new(Employee)
	e1.Age = 10
	e2.Age = 11
	t.Log(e1, e2)
}

type interface1 interface {
	hello()
}

type obj1 struct {
}

func (o *obj1)hello() {
	fmt.Println(111)
}

func TestInterface(t *testing.T) {
	var i interface1 = &obj1{}
	i.hello()
}

func TestGoroutine(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	var count = 0
	for i := 0; i < 5000;i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			count++
			wg.Done()
		}()
	}
	//time.Sleep(time.Second*1)
	wg.Wait()
	fmt.Println(count)
}

func service() string {
	time.Sleep(time.Millisecond*50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond*100)
	fmt.Println("Task is done.")
}

func AsyncService() chan string {
	//retCh := make(chan string)
	retCh := make(chan string , 1) // buffer channel
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestService(t *testing.T) {

	//fmt.Println(service())
	//otherTask()

	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
}













