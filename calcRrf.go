package main

import (
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
    DealDB(starttime,3600,db)
    



}

type POSITION struct {
	ID_name int
	Pos_x  int
	Pos_y int
	Pos_z int
	Pos_time time.Time
	WhichGrid int
}
var xx map[int] map[int] int
func DealDB(starttime time.Time, timelast time.Duration ,db mysql.Conn){
	var  pos_one POSITION
	endtime := starttime.Add(timelast*time.Second)
	sql := `select * from pos_record_tb  where id>=38 and id<=42 and last_heard >= "%s" and last_heard <= "%s"`
        endtimestr := endtime.Format(timelayout)
	starttimestr :=starttime.Format(timelayout)
	sql = fmt.Sprintf(sql,starttimestr,endtimestr)
	fmt.Println(sql)
	rows,res,err := db.Query(sql)
	if err != nil {
		fmt.Println("查询失败")
		return
	}
	xx = make(map[int] map[int] int)
	xx[1]=make(map[int] int)
	xx[2]=make(map[int] int)
	xx[3]=make(map[int] int)
	xx[4]=make(map[int] int)

	for _,row := range rows{
	    pos_one.ID_name = row.Int(res.Map("id"))
	    pos_one.Pos_x = row.Int(res.Map("position_x"))
	    pos_one.Pos_y = row.Int(res.Map("position_y"))
	    pos_one.Pos_z = row.Int(res.Map("position_z"))
	    timestr := row.Str(res.Map("last_heard"))
	    pos_one.Pos_time,_ =time.ParseInLocation(timelayout,timestr,time.Local)
	    zonenum := InWhichZone(pos_one.Pos_x,pos_one.Pos_y)
	    xx[zonenum][pos_one.ID_name ]= xx[zonenum][pos_one.ID_name] + 1




	}
	fmt.Println(xx)
	Rmin := (CalcRmin(xx))
	fmt.Println("Rmin is :",Rmin)
	fmt.Println("RRF is : ",CalcRrf(Rmin))
	//fmt.Println(pos_ones)

}



func  InWhichZone(x int, y int) int {
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
func CalcRrf(Rmin int) float64 {
	var result float64
	Trarity := 4
	if Rmin <= Trarity {
		result = float64(Trarity - Rmin) / float64(Trarity - 1)

	}else{
		result = 0
	}

	return result
}

func CalcRmin(xx map[int] map[int] int) int {
	bb := make(map[int] int)
	for idx,value := range xx {
		bb[idx] = len(value)
	}
	fmt.Println(bb)
	vtmp := 1
	for _,ivalue := range bb{
	     if ivalue < vtmp && ivalue != 0 {
		     vtmp = ivalue
	     }
	}
	return vtmp
}
