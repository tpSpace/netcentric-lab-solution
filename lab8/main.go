package main

import "fmt"

func main() {
	// TODO: Implement the main function
	// 1. USser able to share a file with another user, use UDP to broadcast to annouce
	// 2. User able to download a file from another user, use TCP to download the file
	// 3. User able to list all the files that are shared by other users and search for a file

	// now let's create a terminal application

	for {
		// 1. Print the menu
		// 2. Ask the user to choose an option
		// 3. Perform the action based on the option
		// 4. Repeat
		fmt.Println("Welcome to the P2P File Sharing System")
		fmt.Println("1. Share a file")
		fmt.Println("2. Download a file")
		fmt.Println("3. List all files")
		fmt.Println("4. Search for a file")
		fmt.Println("5. Exit")

		var option int
		fmt.Print("Choose an option: ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			fmt.Println("You choose to share a file")
			fmt.Println("Enter the file path: ")
			var filepath string
			fmt.Scanln(&filepath)
			fmt.Println("Enter the file size: ")
			var filesize int64
			fmt.Scanln(&filesize)
			fmt.Println("File path:", filepath)
			fmt.Println("File size:", filesize)
			sendBroadcast(filepath, filesize)

		case 2:
			fmt.Println("You choose to download a file")
		case 3:
			fmt.Println("You choose to list all files")
			for {
				message := listenForBroadcast()
				fmt.Println("Received message:", message)
				fmt.Println("R for reload, Q for quit or index for download:")
				var option string
				fmt.Scanln(&option)
				if option == "Q" {
					break
				} else if option == "R" {
					continue
				} else {
					// download the file
				}

			}

		case 4:
			fmt.Println("You choose to search for a file")
		case 5:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option")
		}

	}
}
