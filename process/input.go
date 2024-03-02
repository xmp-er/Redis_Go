package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/xmp-er/Redis_Go/validatior"
)

var Map = make(map[string]string)
var Backup_Map = make(map[string]string)

func main() {

	var port string //taking port as flag
	flag.StringVar(&port, "p", ":8000", "Port number on which the TCP Server will be exposed")
	flag.Parse()

	if !validatior.Is_Valid_Port(port) { //validating port
		fmt.Println("Please pass valid port number as :<your_desired_port>")
	}
	listener, err := net.Listen("tcp", port) //making listener

	if err != nil {
		log.Fatal(err)
	}

	sigTermUser := make(chan os.Signal, 1)

	signal.Notify(sigTermUser, os.Interrupt, syscall.SIGTERM)

	go func() {
		for {
			connection, err := listener.Accept() //listening for connections
			select {
			case <-sigTermUser:
				sigTermUser <- os.Kill
				fmt.Println(1)
				return
			default:
				go func() {
					if err != nil {
						log.Fatal(err)
					}
					process(connection)
					connection.Close()
				}()
			}
		}
	}()
	<-sigTermUser
	fmt.Println("\n Closing the listener")
	listener.Close() //closing the listener once its done
}

func process(conn net.Conn) {
	for {
		temp_inp := make([]byte, 8192)
		n, err := conn.Read(temp_inp) //reading the input and taking it to string
		// conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\n\r\nConnection established and response recieved\r\n")))

		if err != nil {
			fmt.Println(err)
			conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\n\r\n %s \r\n", err.Error())))
			break
		}

		str := string(temp_inp[:n]) //processing begins

		str = strings.TrimSpace(str)

		temp_str := strings.Split(str, "\n") //the relevant data part is on the last line

		str = temp_str[len(temp_str)-1]
		str = strings.TrimSpace(str)

		//checking if the input is correct
		_, err = validatior.Validate_input(str)
		if err != nil {
			temp_err := err.Error()
			conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\n\r\n %s \r\n", temp_err)))
			fmt.Println(temp_err)
			break
		}
		//splitting the string into array
		st := strings.Split(str, " ")

		//string holding our final result
		var res string = ""

		//if the command is GET,SET or DEL, we handle it via the crud_commands handler
		switch st[0] {
		case "SET", "GET", "DEL":
			res = crud(st)
		case "INCR", "INCRBY":
			res = incr_cmds(st)
		case "MULTI": //taking Multi as we will be operating from the function after we get all these cmds
			res = transactional_cmds(st)
		case "COMPACT":
			res = additional_commands(st)
		case "DISCONNECT":
			conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\n\r\n %s \r\n", "Closing connection")))
			conn.Close()
			return
		}
		if res != "" {
			fmt.Println(res)
			conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\n\r\n %s \r\n", res)))
		}
	}
	conn.Close() //closing the connection once the processing is done
}
