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
    calcRLF(starttime,3600,5,db)
    



}

func calcRLF(starttime time.Time, timelast time.Duration , sects int ,db mysql.Conn){
	type POSITION struct {
		ID_name int
		Pos_x  int
		Pos_y int
		Pos_z int
		Pos_time time.Time
	}
	var  pos_one POSITION
	var  pos_ones []POSITION
	endtime := starttime.Add(timelast*time.Second)
	sql := `select * from mobile42 where last_heard >= "%s" and last_heard <= "%s"`
        endtimestr := endtime.Format(timelayout)
	starttimestr :=starttime.Format(timelayout)
	sql = fmt.Sprintf(sql,starttimestr,endtimestr)
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
	fmt.Println(pos_ones)
}

