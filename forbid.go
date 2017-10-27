package main

import (
//	"encoding/json"
	"fmt"
//       "io/ioutil"
//	"net/http"
	"time"
//	"strings"
//
	//"github.com/ziutek/mymysql/autorc"
	//_ "github.com/ziutek/mymysql/thrsafe" // You may also use the native engine

//	"github.com/ziutek/mymysql/mysql"
//	_ "github.com/ziutek/mymysql/native" // Native engine
       "labix.org/v2/mgo"
       "labix.org/v2/mgo/bson"

	// "github.com/kardianos/service"
)

/*
"id":28,
"name":"Bridge5",
"position_x":-1,
"position_y":-1,
"position_z":2100,
"mac_address":"e4956efffea50e97",
"parent_mac_address":"e4956efffea50e97",
"bridge_mac_address":"e4956efffea50e97",
"sw_version":"e912a9ec",
"reporting_frequency":10,
"ranging_frequency":30,
"last_heard":"2017-05-08T15:39:19.000Z",
"announce_timestamp":"2017-05-08T11:04:03.000Z",
"position_update_timestamp":"1970-01-01T00:00:00.000Z",
"root_formed":true,
"root_formed_timestamp":"2017-05-08T11:04:03.000Z",
"node_type_id":4,
"sequence_number":0,
"sublocation_id":1,
"battery_voltage":null,
"battery_remaining_charge":null,
"created_at":"2017-04-10T15:13:21.000Z",
"updated_at":"2017-05-08T11:04:03.000Z",
"mobile_dimension_mode":0,
"current_system_timestamp":"2017-05-08T15:39:33.487+00:00",
"hardware_bt_present":null,
"device_groups":[],
"category_list":[]
*/
type RECORD_POS struct {
	Id                        int
	Name                      string
	Position_x                int
	Position_y                int
	Position_z                int
	Mac_address               string
	Parent_mac_address        string
	Bridge_mac_address        string
	Sw_version                string
	Reporting_frequency       int
	Ranging_frequency         int
	Last_heard                string
	Announce_timestamp        string
	Position_update_timestamp string
	Root_formed               bool
	Root_formed_timestamp     string
	Node_type_id              int
	Sequence_number           int
	Sublocation_id            int
	Battery_voltage           int
	Battery_remaining_charge  int
	Created_at                string
	Updated_at                string
	Mobile_dimension_mode     int
	Current_system_timestamp  string
	Hardware_bt_present       int
	Device_groups             []string
	Category_list             []string
}
const  (
  URL = "127.0.0.1:27017"
  timelayer = "2006-01-02 15:04:05"
)
type SETTING struct {
	Id int
	S_time string
	E_time string
	X1 int
	Y1 int
	X2 int
	Y2 int
	In int
}
type INOUTWARN struct {
	Id int 
	Name string
	Time string
	Pos_x int
	Pos_y int
	Setting SETTING 
}

var setting SETTING
var settings []SETTING
func main() {

        session ,err := mgo.Dial(URL)
	if err != nil {
		fmt.Println("数据库无法连接",err)
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic,true)

	db := session.DB("posdata")

	col_setting := db.C("setting")
	load_setting(col_setting)
	fmt.Println(settings)

	collection := db.C("positon")
	col_warn := db.C("inwarn")
	iter := collection.Find(bson.M{"id" : bson.M{"$gt" : 37}}).Iter()
	var result RECORD_POS
	num := 0
	for iter.Next(&result) {
		num += 1
		fmt.Printf("Result: %v. sqis : %d\n", result.Name,num)
		warningots :=checkposbysetting(result)
	        var inoutwarn INOUTWARN	
		fmt.Println("this is in_warningots",warningots)
		for _, warningot := range(warningots){
			inoutwarn.Id = result.Id
			inoutwarn.Name = result.Name
			inoutwarn.Time =  result.Last_heard 
			inoutwarn.Pos_x = result.Position_x
			inoutwarn.Pos_y = result.Position_y
			inoutwarn.Setting = settings[warningot]
			err := col_warn.Insert(&inoutwarn)
			if err != nil {
				fmt.Println("inwarn can not insert:",err)
                                
			}

		}
	}
}

func checkposbysetting(rec_pos RECORD_POS ) []int {

	var ret []int
	t1,_ := time.ParseInLocation(timelayer,rec_pos.Last_heard,time.Local)
	it1 := t1.Unix()
	for inx ,setting := range(settings) {
		ts,_ := time.ParseInLocation(timelayer,setting.S_time,time.Local)
		its := ts.Unix()
		te,_ := time.ParseInLocation(timelayer,setting.E_time,time.Local)
		ite := te.Unix()
		if it1 < its || it1 > ite && false {
			continue
		}else{
			if rec_pos.Id != setting.Id {
				continue
			}else{
				if  setting.In == 0 {
					if !(rec_pos.Position_x > setting.X1 && rec_pos.Position_y > setting.Y1 && rec_pos.Position_x < setting.X2 && rec_pos.Position_y < setting.Y2) {
						continue
					}else{
						ret = append(ret,inx)
						continue

					}
				}else{
					if rec_pos.Position_x > setting.X1 && rec_pos.Position_x < setting.X2 && rec_pos.Position_y > setting.Y1 && rec_pos.Position_y < setting.Y2 {
						continue
					}else{
						ret = append(ret,inx)
						continue

					}

				}
			}


		}
	}
	return ret


}

func load_setting(col *mgo.Collection){
	iter := col.Find(nil).Iter()
	for iter.Next(&setting){
		settings = append(settings,setting)
	}

}
