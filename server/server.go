package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"
)

type Message struct {
	Author  string
	Content string
}

type ChatServer struct {
	messages []Message
	mu       sync.Mutex
}

func (s *ChatServer) SendMessage(msg Message, reply *[]Message) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.messages = append(s.messages, msg)
	*reply = s.messages
	return nil
}

func (s *ChatServer) GetMessages(dummy int, reply *[]Message) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	*reply = s.messages
	return nil
}

func main() {
	server := new(ChatServer)
	err := rpc.Register(server)
	if err != nil {
		log.Fatalf("Error registering RPC: %v", err)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer listener.Close()

	fmt.Println("💬 Chatroom server started on port 1234...")
	rpc.Accept(listener)
}
