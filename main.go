package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/dynamic"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <proto-file> <message-name>")
		return
	}

	protoFile := os.Args[1]   // например "hello.proto"
	messageName := os.Args[2] // например "hello.Person"

	// 1. Парсим .proto
	parser := protoparse.Parser{}
	fds, err := parser.ParseFiles(protoFile)
	if err != nil {
		log.Fatalf("failed to parse proto: %v", err)
	}

	// 2. Находим дескриптор сообщения
	md := fds[0].FindMessage(messageName)
	if md == nil {
		log.Fatalf("message %s not found in %s", messageName, protoFile)
	}

	// 3. Создаём динамическое сообщение
	msg := dynamic.NewMessage(md)

	// 4. Пример JSON
	jsonData := []byte(`{"name":"Alice","age":30}`)

	// 5. JSON -> Protobuf
	if err := msg.UnmarshalJSON(jsonData); err != nil {
		log.Fatalf("failed to unmarshal json: %v", err)
	}

	fmt.Println("✅ Parsed JSON into Protobuf message")
	fmt.Println(msg)

	// 6. Protobuf -> JSON
	out, err := msg.MarshalJSON()
	if err != nil {
		log.Fatalf("failed to marshal to json: %v", err)
	}
	fmt.Println("✅ Converted back to JSON:", string(out))
}