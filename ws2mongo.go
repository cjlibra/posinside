package main



import (

	"golang.org/x/net/websocket"

	"fmt"

	"log"

)



var origin = "http://192.168.0.100/"

var url = "ws://192.168.0.100/api/v1/wsstream"



func main() {

	ws, err := websocket.Dial(url, "", origin)

	if err != nil {

		log.Fatal(err)

	}

	message := []byte(`{
		
		jsonrpc: "2.0",
		
		method : "get",
		
		params : {
			
			entity: "channels",
			
			project_uid: project_uid,
			
			userEmail: userEmail,
			
			authToken: authToken
			
		}
		
	}`)

	_, err = ws.Write(message)

	if err != nil {

		log.Fatal(err)

	}

	fmt.Printf("Send: %s\n", message)



	var msg = make([]byte, 512)

	m, err := ws.Read(msg)

	if err != nil {

		log.Fatal(err)

	}

	fmt.Printf("Receive: %s\n", msg[:m])



	ws.Close()//关闭连接

}
