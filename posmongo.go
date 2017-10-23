package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"strings"

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
const  (
  URL = "127.0.0.1:27017"
)
func Cuttimestr(str string) string {

	a1 := strings.Split(str,".")[0]
	a2 := strings.Replace(a1,"T"," ",-1)
	return a2



}
func main() {

        session ,err := mgo.Dial(URL)
	if err != nil {
		fmt.Println("数据库无法连接",err)
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic,true)

	db := session.DB("posdata")
	collection := db.C("positon")
	for {
		out := httpGet()
		//fmt.Println(out)
		//var rec_pos RECORD_POS
		var recs_pos []RECORD_POS

		json.Unmarshal([]byte(out), &recs_pos)
		//fmt.Println(recs_pos)

		for _, rec_pos := range recs_pos {
			err = collection.Insert(&rec_pos)
			fmt.Println(rec_pos)
			if err != nil {
				fmt.Println("数据库无法插入" + err.Error())
				time.Sleep(time.Second * 5)
				continue
			}
		}
		time.Sleep(time.Second * 1)
	}
}

func httpGet() string {
	resp, err := http.Get("http://192.168.0.100/nodes.json")
	if err != nil {
		return "http get error"

	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "ioutil.ReadAll error"

	}

	return string(body)
}
