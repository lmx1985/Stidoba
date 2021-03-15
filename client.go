package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Эта часть отвечает за верификацию пароля (если пароль не верный, то програма отрубается)

	fmt.Println("Server:")
	buff := make([]byte, 1024)
	n, err := conn.Read(buff)
	if err != nil {
		return
	}
	fmt.Print(string(buff[0:n]))

	var sourcee string
	_, err = fmt.Scanln(&sourcee)
	if err != nil {
		fmt.Println("Некорректный ввод", err)
		return
	}
	// отправляем сообщение серверу
	if n, err := conn.Write([]byte(sourcee)); n == 0 || err != nil {
		fmt.Println(err)
		if err != nil {
			return
		}
	}
	otvet := make([]byte, 1024)
	n, err = conn.Read(otvet)
	fmt.Println(string(otvet[0:n]))
	z := string(otvet[0:n])

	if z != "Вход разрешен" {
		os.Exit(1)
	}

	//Начинаем обмен данными с сервером, если пароль введен верно
	for {
		var source string
		fmt.Print(">>> ")

		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			source = sc.Text()
			break
		}
		// отправляем сообщение серверу
		if n, err := conn.Write([]byte(source)); n == 0 || err != nil {
			fmt.Println(err)
			return
		}
		// получем ответ
		fmt.Println("Ответ:")
		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		if err != nil {
			break
		}
		fmt.Print(string(buff[0:n]))
		buff = nil
		n = 0
		fmt.Println()

	}
}
