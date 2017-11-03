package main


import (
    "math"
    "fmt"
    "mypack"
    "os"
    "io/ioutil"


)


func main(){
	fn := "test.txt"

	s := []byte("Hello World!")

	ioutil.WriteFile(fn, s, os.ModeAppend)
	ioutil.WriteFile(fn, s, os.ModeAppend)
	ioutil.WriteFile(fn, s, os.ModeAppend)
	ioutil.WriteFile(fn, s, os.ModeAppend)
	return
	mypack.GetNewWarning()
	return
	x := make(map[string] map[string] int )
	x1 := make ( map[int] int )
	x["a"] =  make( map[string] int )
	x["a"]["b"]= 1
	x["a"]["c"]= 2
	x["a"]["c"]= x["a"]["c"]+ 1
	x1[1]=2
        fmt.Println(x1)
	
	return
	a := math.Sqrt(4)
	fmt.Println(math.Abs(a))
	var bb []int
	u := 1
	bb=append(bb,u)
	u = 2
	bb=append(bb,u)
	fmt.Println(bb)
	aaa := 9.0
	bbb := 7 
	bbb = bbb*1.0
	aaa = aaa*1.0
	ccc := float64(aaa)/float64(bbb)
	fmt.Println(ccc)
	fmt.Println(math.Sqrt(-1))



}
