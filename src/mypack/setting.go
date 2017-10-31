package  mypack 

import (
	"encoding/json"
	"fmt"
//       "io/ioutil"
//	"net/http"
	//"time"
//	"strings"
//
	//"github.com/ziutek/mymysql/autorc"
	//_ "github.com/ziutek/mymysql/thrsafe" // You may also use the native engine

//	"github.com/ziutek/mymysql/mysql"
//	_ "github.com/ziutek/mymysql/native" // Native engine
       "labix.org/v2/mgo"
       //"labix.org/v2/mgo/bson"

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
type SETTING struct {
	Id int
	S_time  string
	E_time string
	X1  int
	Y1  int
	X2  int
	Y2  int
	In  int



}
const  (
  URL = "127.0.0.1:27017"
  timelayer= "2006-01-02 15:04:05"
)
func Setzone(settingstr string) string {


        session ,err := mgo.Dial(URL)
	if err != nil {
		fmt.Println("数据库无法连接",err)
		return err.Error()
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic,true)

	db := session.DB("posdata")
	collection := db.C("setting")
	var setting SETTING
	err = json.Unmarshal([]byte(settingstr),&setting)
	if err  != nil {
		return "json string can not to object"
	}
	/*
	setting.Id = 38
	stime, _ := time.ParseInLocation(timelayer, "2017-10-22 23:36:25", time.Local)
	etime := stime.Add(time.Hour*5)
	setting.S_time = stime.String()
	setting.E_time = etime.String()
	setting.X1 = 8000
	setting.Y1 = 3000
	setting.X2 = 9000
	setting.Y2 = 10000
	setting.In = 0
       */
	err = collection.Insert(&setting)
	if err != nil {
		fmt.Println("setting can not insert:",err)
		return err.Error()
	}
	return "ok"
}

