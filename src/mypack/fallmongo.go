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
const  (
  URL1 = "127.0.0.1:57017"  //202.127.26.254 
  //timelayer = "2006-01-02 15:04:05"
)
/*
C. 查询最新一次行动的定位数据
先查询出最新一次的行动ID：
db.actions.find({}, {actionId: 1}).sort({_id: -1}).limit(1)
可以看到打印出的数据中有 actionId 字段，复制它的值，例如 action-eddfaff8-278d-479a-83c7-cc6355e575c0
执行命令：
use action-eddfaff8-278d-479a-83c7-cc6355e575c0
执行命令：
db.history.find()
可以看到打印出了本次行动中所有的人员数据信息，包括位置、姿态、警报等
*/

func GetFallInfo() string {

        session ,err := mgo.Dial(URL1)
	if err != nil {
		fmt.Println("数据库无法连接",err)
		return "error"
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic,true)

	db := session.DB("test")

	collection := db.C("actions")
	type ACTIONSTRU struct {
		ID  bson.ObjectId `bson:"_id"`
		ActionID string  `bson:"actionId"`


	}
        var  actionId  ACTIONSTRU 
	err = collection.Find(nil ).Select(bson.M{"actionId":1}).Sort("-_id").One(&actionId)
	if err != nil {
		fmt.Println(err.Error())
		return "error1"
	}
//	fmt.Println(actionId.ActionID)
	db_action := session.DB(actionId.ActionID)
	col_action := db_action.C("history")
	var  actions  interface{}
	err = col_action.Find(nil).One(&actions)
//	fmt.Println(actions)
	if err != nil {
		fmt.Println(err.Error())
		return "error2"
	}
	bb_actions, err := json.Marshal(&actions)
	if err != nil {
		fmt.Println(err.Error())
		return "error3"
	}
	return string(bb_actions)



}

