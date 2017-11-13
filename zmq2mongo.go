/*
package main



import "fmt"

import zmq "github.com/alecthomas/gozmq"



func main() {

	context, _ := zmq.NewContext()

	socket, _ := context.NewSocket(zmq.SUB)
	socket.SetSockOptString(zmq.SUBSCRIBE,"#pos")
	socket.Connect("tcp://127.0.0.1:7001")

	for {

		data,err := socket.Recv(0)
		fmt.Println(string(data))

	}


}
*/



package main



import (

	zmq "github.com/pebbe/zmq4"



	"fmt"

)



func main() {

	//  Socket to talk to server

	fmt.Println("Connecting to hello world server...")

	subscriber, _ := zmq.NewSocket(zmq.SUB)

	defer subscriber.Close()
	subscriber.SetSubscribe("#pos")

	subscriber.Connect("tcp://192.168.0.100:7001")


	for {

		reply, _ := subscriber.Recv(0)

		fmt.Println("Received ", reply)
	}


}
