package server

import "net"
import "fmt"
import "bufio"
import "strings"
import "encoding/json"

//Declare map
var queue map[string]chan string

func Connect(conn net.Conn) {

	fmt.Println("Connect!")
	type Json struct {
		Header  string
		Queue   string
		Message string
	}
	// run loop forever (or until ctrl-c)
	for {

		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')

		// read json message
		dec := json.NewDecoder(strings.NewReader(message))

		var j Json
		dec.Decode(&j)

		switch j.Header {
		case "declare":
			exit := QueueDeclare(j.Queue)
			conn.Write([]byte(exit + "\n"))
			break
		case "publish":
			exit := Publish(j.Queue, j.Message)
			conn.Write([]byte(exit + "\n"))
			break
		case "consume":
			exit := Consume(j.Queue)
			conn.Write([]byte(exit + "\n"))
			break
		case "data":
		case "revoke":
		default:
			conn.Write([]byte("Error" + "\n"))
			break
		}

	}
}

func QueueDeclare(name string) string {
	//Declare queue
	queue[name] = make(chan string, 100)
	return "success!"
}

func Publish(name, message string) string {
	queue[name] <- message
	return "success!"
}

func Consume(name string) string {
	return <-queue[name]
}

func main() {
	queue = make(map[string]chan string)
	queue["start"] = make(chan string, 100)
	ln, _ := net.Listen("tcp", ":8081")
	for {
		conn, _ := ln.Accept()
		go Connect(conn)
	}
	forever := make(chan bool)
	<-forever

}
