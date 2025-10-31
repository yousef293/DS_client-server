package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
)

type Message struct {
	Author  string
	Content string
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer client.Close()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("You can start chatting! Type 'exit' to quit.\n")

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if strings.ToLower(text) == "exit" {
			fmt.Println("👋 Goodbye!")
			break
		}

		var history []Message
		msg := Message{Author: name, Content: text}

		// Send message to server
		err = client.Call("ChatServer.SendMessage", msg, &history)
		if err != nil {
			fmt.Printf("⚠️ Error sending message: %v\n", err)
			break
		}

		fmt.Println("\n--- Chat History ---")
		for _, m := range history {
			fmt.Printf("[%s]: %s\n", m.Author, m.Content)
		}
		fmt.Println("--------------------\n")
	}
}
