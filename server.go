package main

import (
	"fmt"
	"net"
)

var dict = map[string]string{
	"red":    "красный",
	"green":  "зеленый",
	"blue":   "синий",
	"yellow": "желтый",
}

func main() {
	listener, err := net.Listen("tcp", ":4545")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		conn.Write([]byte("Введите пароль:"))

		passw := make([]byte, (1024 * 4))
		n, err := conn.Read(passw)
		if n == 0 || err != nil {
			fmt.Println("Read error:", err)
			return
		}
		ps := string(passw[0:n])

		if ps == "123456" { // если данные не найдены в словаре
			conn.Write([]byte("Вход разрешен"))
			go handleConnection(conn) // запускаем горутину для обработки запроса
		} else {
			conn.Write([]byte("Введен не верный пароль"))
			conn.Close()
		}
	}
}

// обработка подключения
func handleConnection(conn net.Conn) {
	var c, p string

	// на основании полученных данных получаем из словаря перевод
	fmt.Print(Dir("", ""))

	defer conn.Close()
	for {
		// считываем полученные в запросе данные
		input := make([]byte, (1024 * 4))
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println("Read error:", err)
			input = nil
			break
		}
		source := string(input[0:n])
		// !!! Разбираем ответ сервера на команды (Команда  -  Путь)

		for i, v := range source {
			if string(v) == string(" ") {
				c = string(source[0:i])
				p = string(source[i+1:])
				break
			} else {
				c = source
				p = ""
			}
			conn.Write([]byte(Dir(c, p)))

			// на основании полученных данных получаем из словаря перевод
			//target, ok := dict[source]
			//if ok == false { // если данные не найдены в словаре
			//	target = "undefined"
			//}
			// выводим на консоль сервера диагностическую информацию
			//fmt.Println(source, "-", target)
			// отправляем данные клиенту
			//conn.Write([]byte(target))
		}
	}
}