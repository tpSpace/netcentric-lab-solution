// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"os"
// )

// type User struct {
// 	Id       int      `json:"id"`
// 	Username string   `json:"username"`
// 	Password string   `json:"password"`
// 	Fullname string   `json:"fullname"`
// 	Email    []string `json:"email"`
// 	Address  []string `json:"address"`
// }

// func main() {
// 	// Open our jsonFile
// 	jsonFile, err := os.Open("user.json")

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	byteValue, _ := io.ReadAll(jsonFile)
// 	var users []User

// 	json.Unmarshal(byteValue, &users)

// 	fmt.Print(users[0].Username)
// 	// edit the json file
// 	users[0].Username = "newusername"

// 	fmt.Println("\nSuccessfully Opened users.json")
// 	defer jsonFile.Close()

// }
