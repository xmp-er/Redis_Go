# REDIS 0.1 EXPOSED VIA TCP SERVER IN GO

Redis 0.1 build entirely on Go exposed via TCP server with in-memory storage

The program will accept DB commands as inputs from the command line from client connection and process them by creating DB structures in memory. The output will be displayed on the standard output stream and to the client with connection

## How to Run and Hit Server

 - To run locally, clone and use `go run ./process` to run the program as the main file resides in process folder

 - The default server is 8000, it can be hit via telnet as `telnet localhost 8000`

For custom port, it can be passed as a flag varaible as `go run ./process -p :<custom_port>`



## Commands

the basic features of Redis 0.1 are implemented such as

 - `SET <k> <v>` to set the key to a certain value
 - `GET <k>` to get value of a key that has been set
 - `DELETE <k>` to delete a key

Apart from basic crud commands, arithmetic operation commands are also added as 

 - `INCR <k>` increments value of specified key by 1 if the value is integer
 - `INCRBY <k> <v>` increments value of specified key by v if the value is integer

Commands to execute mass instructions at once have been added such as 
 - `MULTI` starts queueing the commands 
 - `EXEC` executes the queued commands
 - `DISCARD` discards the queued commands

Some additional commands implemented are
 - `COMPACT` shows the final SET value of a key, if the value was a integer.

## Database structure

<b><u>Redis has databses indexed 0-15</u></b>, post connection establishment we will have to select the database to operate on, these databases can be accessed via any of the connection, new or established, and values can be modified accordingly.

 - `SELECT <db_number>` selects datbase ranged from 0-15 based on the number provided

## DISCONNECT CONNECTION

To terminate a already established connection, `DISCONNECT` command can be used which will terminate the connection from client-side.

## Demonstration of application

A video demonstration of the application with all the commands over multiple TCP connections communication with the same in-memory database at the same time can be found attached.

![](Redis_0.1_Go_Demo.mkv)
