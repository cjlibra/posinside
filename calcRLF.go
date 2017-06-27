package main

import (
	"math"
	"fmt"
	"time"
	"github.com/ziutek/mymysql/mysql"
	"github.com/ziutek/mymysql/godrv"
)


var timelayout string 
func main(){
    timelayout ="2006-01-02 15:04:05"
    godrv.Register("SET NAMES utf8")

    db := mysql.New("tcp", "", "127.0.0.1:3306", "root", "root", "posdata")

    err := db.Connect()
    if err != nil {
	fmt.Println("数据库无法连接")
	return
    }
    defer db.Close()
    starttime,r := time.ParseInLocation(timelayout,"2017-06-15 13:00:00",time.Local)
    fmt.Println(starttime,r)
    calcRLF(starttime,3600,"mobile42",db)
    



}

type POSITION struct {
	ID_name int
	Pos_x  int
	Pos_y int
	Pos_z int
	Pos_time time.Time
	WhichGrid int
}
func calcRLF(starttime time.Time, timelast time.Duration , labels string ,db mysql.Conn){
	var  pos_one POSITION
	var  pos_ones []POSITION
	endtime := starttime.Add(timelast*time.Second)
	sql := `select * from %s where last_heard >= "%s" and last_heard <= "%s"`
        endtimestr := endtime.Format(timelayout)
	starttimestr :=starttime.Format(timelayout)
	sql = fmt.Sprintf(sql,labels,starttimestr,endtimestr)
	fmt.Println(sql)
	rows,res,err := db.Query(sql)
	if err != nil {
		fmt.Println("查询失败")
		return
	}
	for _,row := range rows{
	    pos_one.ID_name = row.Int(res.Map("id"))
	    pos_one.Pos_x = row.Int(res.Map("position_x"))
	    pos_one.Pos_y = row.Int(res.Map("position_y"))
	    pos_one.Pos_z = row.Int(res.Map("position_z"))
	    timestr := row.Str(res.Map("last_heard"))
	    pos_one.Pos_time,_ =time.ParseInLocation(timelayout,timestr,time.Local)
	    pos_ones = append(pos_ones,pos_one)



	}
	//fmt.Println(pos_ones)
	CutIntoPiece(pos_ones)
	fmt.Println(len(pos_many))
	if len(pos_many) == 0 {
		pos_many = append(pos_many,pos_ones)
		fmt.Println(labels,"人员逗留")
		return
	}
	Lengthmax := SearchLongestLength(pos_many)
	rlf := CalcRlf(Lengthmax)
	fmt.Println("RLF is:",rlf,Lengthmax)

}


func CalcRlf(LengthMax float64) float64{
	T_Length := 200.0
	result := math.Sqrt((LengthMax - 1 ) / ( T_Length - 1))
	return result

	
}

func WhichLengthIsMax(Lengths []float64) float64{
	A0 := Lengths[0]
	for _,Length := range Lengths {
		if Length > A0 {
			A0 = Length
		}
	}
	return A0 
}
func SearchLongestLength(pos_many [][]POSITION) float64{
	var Length float64 
	var Lengths []float64
	for _,pos_ones := range pos_many {
		Length = 0.0
		for id_pos_one,pos_one := range pos_ones {
			if id_pos_one <  len(pos_ones)  - 1 {
				Length = Length + CalcDistance(pos_one.Pos_x,pos_one.Pos_y,pos_ones[id_pos_one+1].Pos_x,pos_ones[id_pos_one+1].Pos_y)
			}
		}
		Lengths = append(Lengths,Length)
	}

	return WhichLengthIsMax(Lengths)


}

var pos_many [][]POSITION
func CutIntoPiece(pos_ones []POSITION){
	A :=InWhichGrid(pos_ones[0].Pos_x, pos_ones[0].Pos_y)
	B := 0
	for seqid,pos_one := range pos_ones {
		if InWhichGrid(pos_one.Pos_x,pos_one.Pos_y) != A {
			pos_many = append(pos_many,pos_ones[B:seqid])
			B = seqid
			A = InWhichGrid(pos_one.Pos_x,pos_one.Pos_y) 
		}
	}
}

func  InWhichGrid(x int, y int) int {
	if x <= 2000  && y <= 2000 {
		return 1
	}
	if x >= 2000  && y <= 2000 {
		return 2
	}
	if x <= 2000  && y >= 2000 {
		return 3
	}
	if x >= 2000  && y >= 2000 {
		return 4
	}
	return 0



}

func CalcDistance(x1 int,y1 int,x2 int,y2 int) float64 {
	x := math.Abs(float64(x1 - x2))
	y := math.Abs(float64(y1 - y2))
	z := math.Sqrt(x*x+y*y)
	return z
}
