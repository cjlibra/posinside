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
       "labix.org/v2/mgo/bson"

	// "github.com/kardianos/service"
)
/*
const  (
  URL = "127.0.0.1:27017"
  timelayer= "2006-01-02 15:04:05"
)
*/
func GetNewWarning() (string,int) {

        session ,err := mgo.Dial(URL)
	if err != nil {
		fmt.Println("数据库无法连接",err)
		return err.Error() , -1
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic,true)

	db := session.DB("posdata")
	collection := db.C("inoutwarn")
	var idnums []int
	err = collection.Find(nil).Distinct("id",&idnums)
	if err != nil {
		fmt.Println("setting can not  find distinct :",err)
		return err.Error(), -2
	}
	var  allwarns []string
	for _,idnum := range(idnums) {
		//query_str :=fmt.Sprintf("\"id\" :  %d ",idnum)
		var obj_warn interface{} 
		collection.Find(bson.M{"id" : idnum}).Sort("\"_id\" : -1").Limit(1).One(&obj_warn)
		
		str_warn, err := json.Marshal(obj_warn)
		if err != nil {
			return err.Error() , -3
		}
		allwarns = append(allwarns,string(str_warn))


	}
	str_warns,_ := json.Marshal(allwarns)
	fmt.Println(string(str_warns))
	return string(str_warns), 0
}

