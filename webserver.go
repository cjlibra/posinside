package main
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"strings"
	"mypack"

	//"github.com/ziutek/mymysql/autorc"
	//_ "github.com/ziutek/mymysql/thrsafe" // You may also use the native engine


	// "github.com/kardianos/service"
)
func GetPosition(w http.ResponseWriter, r *http.Request) {
	bb,_ := json.Marshal(pos_all)
	w.Write(bb)

}
func GetNewWarning(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	ret,retint := mypack.GetNewWarning()
	if  retint !=  0 {
		w.Write([]byte(ret + fmt.Sprintf("+++%d++",retint)))
		return 
	}
 	w.Write([]byte(ret))

}
func SetZoneAlarm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	settingstr := r.Form["set"]
	ret := mypack.Setzone(settingstr[0])
	w.Write([]byte(ret))

}
func webserver(){

	http.HandleFunc("/GetNewWarning",GetNewWarning)
	http.HandleFunc("/GetPosition", GetPosition)
	http.HandleFunc("/SetZoneAlarm",SetZoneAlarm)

	err := http.ListenAndServe(":10080", nil)
	if err != nil {

		fmt.Println("ListenAndServer: ", err)

	}
}

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
type POS struct {
	Id  int
	Name string
	Position_x int
	Position_y int
	Position_z int
	Timestr  string
}
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
	Last_heard                time.Time
	Announce_timestamp        time.Time
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

var pos_all []POS
func main() {
	var pos_one POS

	go webserver()

	for {
		out := httpGet()
		//fmt.Println(out)
		//var rec_pos RECORD_POS
		var recs_pos []RECORD_POS

		json.Unmarshal([]byte(out), &recs_pos)
		//fmt.Println(recs_pos)

		pos_all=[]POS{}
		for _, rec_pos := range recs_pos {
			if strings.Contains(strings.ToUpper(rec_pos.Name),strings.ToUpper("mobile")) != true {
				continue
			}
			//sql = fmt.Sprintf(sql, rec_pos.Id, rec_pos.Name, rec_pos.Position_x, rec_pos.Position_y, rec_pos.Position_z, rec_pos.Mac_address, rec_pos.Parent_mac_address, rec_pos.Bridge_mac_address, rec_pos.Sw_version)
			pos_one.Id = rec_pos.Id
			pos_one.Name = rec_pos.Name
			pos_one.Position_x = rec_pos.Position_x
			pos_one.Position_y = rec_pos.Position_y
			pos_one.Position_z = rec_pos.Position_z
			pos_one.Timestr = rec_pos.Last_heard.String()
			pos_all = append(pos_all,pos_one)
	//		fmt.Println(rec_pos)
		}
		time.Sleep(time.Second*1)
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
