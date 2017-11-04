package mypack 

import (
	"encoding/json"
	"fmt"
//       "io/ioutil"
//	"net/http"
//	"time"
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

func GetInOutWarn(userid int, in int) string {

        session ,err := mgo.Dial(URL)
	if err != nil {
		fmt.Println("数据库无法连接",err)
		return "error"
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic,true)

	db := session.DB("posdata")

	collection := db.C("inoutwarn")
	var inoutwarn interface{}
	err = collection.Find(bson.M{"id": userid , "setting.in" : in}).Sort("-_id").One(&inoutwarn)
	if err != nil {
		fmt.Println(err.Error())
		return "error,can not find inwarn userid= "+fmt.Sprintf("%d",userid)
	}
	bb_inoutwarn, err := json.Marshal(&inoutwarn)
	if err != nil {
		fmt.Println(err.Error())
		return "error,can not json to obj by"
	}
	return string(bb_inoutwarn)



}


