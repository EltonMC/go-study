package main

import "net"
import "fmt"
import "bufio"

func OpenQueue(conn net.Conn, name string) {
	text := "{\"Header\": \"declare\", \"Queue\":\"" + name + "\"}"
	// send to socket
	fmt.Fprintf(conn, text+"\n")
	// listen for reply
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Status: " + message + "\n")
}

func SendMessage(conn net.Conn, name string, message string) {
	text := "{\"Header\": \"publish\", \"Queue\":\"" + name + "\", \"Message\":\"" + message + "\"}"

	// send to socket
	fmt.Fprintf(conn, text+"\n")
	// listen for reply
	message, _ = bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Status: " + message + "\n")
}

func ConsumeQueue(conn net.Conn, name string) {
	for {
		text := "{\"Header\": \"consume\", \"Queue\":\"" + name + "\"}"

		// send to socket
		fmt.Fprintf(conn, text+"\n")

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("\n\tFila: " + name + "\n\tMensagem: " + message)
	}
}

func main() {

	// connect to this socket
	conn, _ := net.Dial("tcp", "localhost:8081")
	for {
		fmt.Print("1-Criar Fila\n2-Escrever na Fila\n3-Consumir Fila\nSelecione uma opção: ")
		var input string
		fmt.Scanln(&input)
		switch input {
		case "1":
			fmt.Print("Digite o nome da fila: ")
			var queue string
			fmt.Scanln(&queue)
			OpenQueue(conn, queue)
			break
		case "2":
			fmt.Print("Digite o nome da fila: ")
			var queue string
			fmt.Scanln(&queue)

			fmt.Print("Digite a mensagem: ")
			var message string
			fmt.Scanln(&message)

			SendMessage(conn, queue, message)
			break
		case "3":
			fmt.Print("Digite o nome da fila: ")
			var queue string
			fmt.Scanln(&queue)
			aux, _ := net.Dial("tcp", "localhost:8081")
			go ConsumeQueue(aux, queue)
			break
		default:
			fmt.Print("Opção incorreta\n\n")
			break
		}
	}
}
