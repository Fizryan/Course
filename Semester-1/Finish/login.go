package main

import "fmt"

func main() {
	users := map[string]string{
		"admin": "admin",
		"user1": "pass1",
		"user2": "pass2",
	}

	var username, password string
	var failed int

	for {
		fmt.Print("Username: ")
		fmt.Scan(&username)

		fmt.Print("Password: ")
		fmt.Scan(&password)

		if storedPassword, exists := users[username]; exists && password == storedPassword {
			fmt.Println("Login berhasil!")
			break
		} else {
			failed++
			fmt.Println("Login gagal! Silahkan coba lagi!")
		}
	}

	fmt.Println("Total login gagal:", failed)
}
