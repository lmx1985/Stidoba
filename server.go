package main

import (
	"fmt"
	"net"
	"strings"
)

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
		text := strings.Split(source, " ")
		if len(text) > 1 {
			c = text[0]
			p = text[1]
		} else {
			c = text[0]
			p = ""
		}
		fmt.Println(c, p)
		otv := (Dir(string(c), string(p)))
		if len(otv) != 0 {
			conn.Write([]byte(otv))
		} else {
			otv = "Успешно"
			conn.Write([]byte(otv))
		}
	}

}
