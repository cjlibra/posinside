package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"strings"
	zmq "github.com/pebbe/zmq4"

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
	Last_heard               time.Time 
	Announce_timestamp       time.Time 
	Position_update_timestamp time.Time 
	Root_formed               bool
	Root_formed_timestamp     time.Time 
	Node_type_id              int
	Sequence_number           int
	Sublocation_id            int
	Battery_voltage           int
	Battery_remaining_charge  int
	Created_at                time.Time 
	Updated_at                time.Time 
	Mobile_dimension_mode     int
	Current_system_timestamp  time.Time 
	Hardware_bt_present       int
	Device_groups             []string
	Category_list             []string
}
const  (
  URL = "127.0.0.1:27017"
  ZMQ_URL = "tcp://192.168.0.100:7001"
  NODE_URL = "http://192.168.0.100/nodes.json"
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
	out := httpGet()
		//fmt.Println(out)
		//var rec_pos RECORD_POS
	var recs_pos []RECORD_POS

	json.Unmarshal([]byte(out), &recs_pos)
		//fmt.Println(recs_pos)

	fmt.Println("Connecting to zmq server...")

	subscriber, _ := zmq.NewSocket(zmq.SUB)

	defer subscriber.Close()
	subscriber.SetSubscribe("#pos")

	subscriber.Connect(ZMQ_URL)


	for {

		reply, _ := subscriber.Recv(0)

		fmt.Println(reply)
		rec_pos := UpdateRecPos(reply,recs_pos)
		err = collection.Insert(&rec_pos)
		fmt.Println(rec_pos)
		if err != nil {
			fmt.Println("数据库无法插入" + err.Error())
			continue
		}
	}
	
}
func UpdateRecPos(sub_reply string , recs_pos []RECORD_POS) RECORD_POS  {
	type ZMQPOS struct {
		Ts  int64 `bson:"ts"`
		Id  int `bson:"id"`
		X   int `bson:"x"`
		Y   int `bson:"y"`
		Z   int `bson:"z"` 
	}
	reply := strings.Replace(sub_reply,"#pos","",-1)
	var zmqpos ZMQPOS
	err := json.Unmarshal([]byte(reply),&zmqpos)
	if err != nil {
		fmt.Println(reply,err.Error())
		return  RECORD_POS{}
	}
	for seq,rec_pos := range recs_pos {
		if rec_pos.Id ==  zmqpos.Id {
			recs_pos[seq].Position_x = zmqpos.X
			recs_pos[seq].Position_y = zmqpos.Y
			recs_pos[seq].Position_z  = zmqpos.Z
			recs_pos[seq].Last_heard = time.Unix(zmqpos.Ts,0).In(time.Local)
			return recs_pos[seq]
		}

	}
	return  RECORD_POS{} 


	//  Socket to talk to server



}

func httpGet() string {
	resp, err := http.Get(NODE_URL)
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
