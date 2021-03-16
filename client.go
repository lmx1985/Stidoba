package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Эта часть отвечает за верификацию пароля (если пароль не верный, то програма отрубается)

	fmt.Println("Server:")
	buff := make([]byte, 4096)
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
	n, _ = conn.Read(otvet)
	fmt.Println(string(otvet[0:n]))
	z := string(otvet[0:n])

	if z != "Вход разрешен" {
		os.Exit(1)
	}
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
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
		buff := make([]byte, 4096)
		n, err := conn.Read(buff)
		if err != nil {
			break
		}
		fmt.Print(string(buff[0:n]))
		if string(buff[0:n]) == "exit" {
			break
		}
		if string(buff[0:n]) == "copy\n" { // дописываем функционал по приему файла
			//dir, _ := os.Getwd()
			//fmt.Println(dir)

			fmt.Println("Начинаю скачивание...")
			const BUFFERSIZE = 4096
			bufferFileName := make([]byte, 64)
			bufferFileSize := make([]byte, 10)

			conn.Read(bufferFileSize)
			fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)

			conn.Read(bufferFileName)
			fileName := strings.Trim(string(bufferFileName), ":")

			newFile, err := os.Create(fileName)

			if err != nil {
				fmt.Println(err)
			}
			defer newFile.Close()
			var receivedBytes int64

			for (fileSize - receivedBytes) > BUFFERSIZE {
				io.CopyN(newFile, conn, (fileSize - receivedBytes))
				conn.Read(make([]byte, (receivedBytes+BUFFERSIZE)-fileSize))
			}

			io.CopyN(newFile, conn, BUFFERSIZE)
			receivedBytes += BUFFERSIZE
			time.Sleep(2)

			fmt.Println("Received file completely!")
			newFile.Close()

		}

		buff = nil
		n = 0
		fmt.Println()
	}
}
