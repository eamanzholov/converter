Converter: JSON ↔ Protobuf

Converter — это простой инструмент на Go для конвертации JSON в Protobuf и обратно динамически, без необходимости генерировать .pb.go файлы.

Проект использует библиотеку github.com/jhump/protoreflect и позволяет работать с любыми .proto файлами на лету.



🔹 Возможности
	•	Динамическое создание Protobuf сообщений на основе .proto файлов.
	•	Конвертация JSON → Protobuf.
	•	Конвертация Protobuf → JSON.
	•	Не требуется сгенерированный Go код (.pb.go), можно использовать любые сообщения.



🔹 Установка

git clone https://github.com/eamanzholov/converter.git
cd converter
go mod tidy



🔹 Использование

go run main.go <proto-file> <message-name>

Пример:

go run main.go hello.proto hello.Person

	•	<proto-file> — путь к твоему .proto файлу.
	•	<message-name> — полное имя сообщения (включая пакет, например hello.Person).

Программа автоматически:
	1.	Парсит .proto файл.
	2.	Создаёт динамическое Protobuf-сообщение.
	3.	Конвертирует пример JSON в Protobuf.
	4.	Сериализует Protobuf обратно в JSON.



🔹 Пример вывода

✅ Parsed JSON into Protobuf message
name:"Alice" age:30

✅ Converted back to JSON: {"name":"Alice","age":30}



🔹 Пример .proto файла

syntax = "proto3";
package hello;

message Person {
  string name = 1;
  int32 age = 2;
}



🔹 Требования
	•	Go 1.20+
	•	Библиотеки:
	•	github.com/jhump/protoreflect/desc/protoparse
	•	github.com/jhump/protoreflect/dynamic

