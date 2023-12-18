# Building-A-TCP-Server-in-GOlang

Here is a simple program to make a [**TCP**](https://www.geeksforgeeks.org/tcp-connection-establishment/) connection in [GO](https://go.dev/). This program demonstrates a simple [**TCP**](https://www.geeksforgeeks.org/tcp-connection-establishment/) connection between server and a client. You can add as many clients as possible in the following. This program will simply let you send a message from the client to the server and the server will respond whether it received the message from the client or not.

# How to run the program?

You can write the following command in the terminal in order to start the server:

```console
go run main.go
```
# To make multiple connections
To make a connection with the server open another terminal and run the following command:

```console
telnet localhost 3000
```
This will create a connection between the server and a cliend.
You can add as many clients as you can using this.
