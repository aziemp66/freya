# Websocket Chat Application with Concurrency (Go-routines)

This is a simple WebSocket chat application built using Go (Golang) on the back-end to enable real-time communication between users. The application utilizes Go-routines for concurrency, allowing multiple users to interact with the chat system simultaneously.

## Features

- Real-time chat functionality using WebSocket protocol
- Support for multiple users to chat concurrently
- Simple and easy-to-use interface

## Getting Started

To run this chat application on your local machine, follow these steps:

1. **Clone the repository**

```bash
git clone https://github.com/your-username/websocket-chat.git
cd websocket-chat
```
2. **Copy the .env.example to .env and modify the environtment variables as you fits**
```bash
cp .env.example .env
```

3. **Build and run the server**

```bash
go build ./cmd/main.go -o ./build/main
./main
```

The server should now be running defaultly on `http://localhost:5000`.

4. **Start chatting!**

Enter your username, type your message in the input field, and hit Enter to send the message. Your message will be visible to all other users connected to the chat.

## Concurrency with Go-routines

This chat application utilizes Go-routines to achieve concurrency. Go-routines are lightweight threads that allow multiple tasks to run concurrently without blocking each other. The use of Go-routines ensures that multiple users can send and receive messages simultaneously, providing a seamless chat experience.

## Technologies Used

- Go (Golang) - The back-end language used to handle WebSocket connections and manage chat functionality.
- Websocket - Computer communications protocol, providing full-duplex communication channels over a single TCP connection

## Project Structure
```
|-- cmd/main.go # Main server code
|-- internal # Application Logic
|-- common/websocket/
| |-- ... # WebSocket utility code
```
## Contributing

If you would like to contribute to this project, feel free to open issues or submit pull requests. Any contributions are welcome!

## License

This project is licensed under the [MIT License](LICENSE). You are free to use, modify, and distribute the code as you see fit.
