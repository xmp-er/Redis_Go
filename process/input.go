package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/xmp-er/Redis_Go/helper"
	"github.com/xmp-er/Redis_Go/validatior"
)

var Maps [16]map[string]string
var Backup_Maps [16]map[string]string

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

	sigTermUser := make(chan os.Signal, 1) //implementing graceful shutdown
	var wg sync.WaitGroup
	ctxShutdown, cancel := context.WithCancel(context.Background())
	signal.Notify(sigTermUser, os.Interrupt, syscall.SIGTERM) //will get notified if we put Ctrl+C for termination

	for i := 0; i < 16; i++ {
		Maps[i] = make(map[string]string)
	}

	for i := 0; i < 16; i++ {
		Backup_Maps[i] = make(map[string]string)
	}

	go func() {
		for {
			connection, err := listener.Accept() //listening for connections
			select {
			case <-sigTermUser:
				sigTermUser <- os.Kill
				cancel()
				return
			default:
				wg.Add(1)
				go func() {
					defer wg.Done()
					if err != nil {
						log.Fatal(err)
					}
					process(ctxShutdown, connection)
					connection.Close()
				}()
			}
		}
	}()
	<-sigTermUser
	fmt.Println("\n Closing the listener")
	listener.Close() //closing the listener once its done
}

func process(ctx context.Context, conn net.Conn) {
	db := -1
	for {
		select {
		case <-ctx.Done():
			goto end
		default:
			temp_inp := make([]byte, 8192)
			n, err := conn.Read(temp_inp) //reading the input and taking it to string
			// conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\n\r\nConnection established and response recieved\r\n")))

			if err != nil {
				fmt.Println(err)
				conn.Write([]byte(fmt.Sprintf("%s\n", err.Error())))
				continue
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
				conn.Write([]byte(fmt.Sprintf("%s\n", temp_err)))
				fmt.Println(temp_err)
				continue
			}
			//splitting the string into array
			st := strings.Split(str, " ")

			//string holding our final result
			var res string = ""

			if db == -1 && st[0] != "SELECT" { //making sure a database is selected first
				conn.Write([]byte(fmt.Sprintf("%s\n", "Please select a database first via the SELECT command")))
				continue
			}
			//if the command is GET,SET or DEL, we handle it via the crud_commands handler
			switch st[0] {
			case "SELECT":
				db = helper.Set_db(st)
				conn.Write([]byte(fmt.Sprintf("%s\n", "OK")))
			case "SET", "GET", "DEL":
				res = crud(st, Maps[db], Backup_Maps[db])
			case "INCR", "INCRBY":
				res = incr_cmds(st, Maps[db], Backup_Maps[db])
			case "MULTI": //taking Multi as we will be operating from the function after we get all these cmds
				res = transactional_cmds(st, Maps[db], Backup_Maps[db])
			case "COMPACT":
				res = additional_commands(st, Maps[db], Backup_Maps[db])
			case "DISCONNECT":
				conn.Write([]byte(fmt.Sprintf("%s\n", "Closing connection")))
				conn.Close()
				return
			}
			if res != "" {
				fmt.Println(res)
				conn.Write([]byte(fmt.Sprintf("%s\n", res)))
			}
		}
	}
end:
	conn.Close() //closing the connection once the processing is done
}
