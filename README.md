
![sd](https://github.com/xmp-er/Redis_Go/assets/107166230/ce01c2ea-cdab-4017-b3c3-1c1c544c0690)


# REDIS 0.1 EXPOSED VIA TCP SERVER IN GO
 - [Introduction](#introduction)
   - [How to Run and Connect to the Server](#how-to-run-and-connect-to-the-server)
   - [Supported Commands](#supported-commands)
   - [Database Structure](#database-structure)
   - [Disconnecting from the Server](#disconnecting-from-the-server)
   - [Demonstration Video](#demonstration-video)

# Introduction
 

**Redis 0.1**, entirely built in Go and exposed via a **TCP server**, features **in-memory storage** capabilities. The program accepts database commands from the client connections and processes them, creating **in-memory database structures**. The output is then displayed both on the standard output stream and to the connected clients.

## How to Run and Connect to the Server
 

- Clone the repository and run the program using `go run ./process` from the 'process' folder.
- Connect to the default server at **port 8000** using `telnet localhost 8000`.
- For a custom port, use the following command: `go run ./process -p :<custom_port>`.

## Supported Commands
 

**Redis 0.1** supports basic **CRUD** commands:

- **`SET <k> <v>`**: Set the key to a specified value, if the value, has spaces, then it must be enclsed in quotes as `SET sample_key "spaced value"`.
- **`GET <k>`**: Retrieve the value of a previously set key.
- **`DELETE <k>`**: Delete a key if present.

Arithmetic operation commands include:

- **`INCR <k>`**: Increment the value of a specified key by 1 (if it's an integer), sets new key with default value of "1" if key not present.
- **`INCRBY <k> <v>`**: Increment the value of a specified key by v (if it's an integer), gives error if key not present.

Commands for executing multiple instructions at once:

- **`MULTI`**: Start queuing commands.
- **`EXEC`**: Execute the queued commands.
- **`DISCARD`**: Discard the queued commands.

Additional commands:

- **`COMPACT`**: Display the final **SET value** of a key if the value was an integer.

## Database Structure
 

**Redis** features databases indexed from **0 to 15**. After connection establishment, select a database to operate on using:

- **`SELECT <db_number>`**: Select a database ranging from 0 to 15 based on the provided number.

## Disconnecting from the Server
 

To terminate an established connection, use the **`DISCONNECT`** command, which will terminate the connection from the client-side.

## Demonstration Video
 

A video demonstrating the application, showcasing various commands executed over multiple **TCP connections** communicating with the same **in-memory database** simultaneously, can be found [here](https://github.com/xmp-er/Redis_Go/assets/107166230/5a11aec7-cfe0-4e5a-b951-80fb36c33942).

https://github.com/xmp-er/Redis_Go/assets/107166230/5a11aec7-cfe0-4e5a-b951-80fb36c33942

