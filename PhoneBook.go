package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var pl = fmt.Println
var phonebook = make(map[string]string)

func insertNewUser() {
	pl("please inset the name and phone number for the user")
	var Newusername string
	inputReader := bufio.NewReader(os.Stdin)
	Newusername, _ = inputReader.ReadString('\n')
	if len([]rune(Newusername)) >= 3 {
		pl("now enter the phone number")
		var Newphonenumber string
		inputReader := bufio.NewReader(os.Stdin)
		Newphonenumber, _ = inputReader.ReadString('\n')
		if len([]rune(Newphonenumber)) < 10 {
			pl("the phone you entered is too short..")
		} else {
			// If the key exists
			if _, ok := phonebook[Newusername]; ok {
				pl("user already excised... \n")
			} else {
				pl("New user will be added to the phonebook \n")
				nameWithCapitalLetter := []rune(Newusername)
				nameWithCapitalLetter[0] = unicode.ToUpper(nameWithCapitalLetter[0])
				s := string(nameWithCapitalLetter)
				phonebook[s] = Newphonenumber
			}
		}
	}

}
func printPhoneBook() {
	for k, v := range phonebook {
		pl(k, v)
	}
}
func deleteUserFromPhoneBook(username string) {
	delete(phonebook, username)
	fmt.Printf("User %s has been deleted \n", username)
}

func checkViaUserName(username string) {
	pl("checking via username")
	value, ok := phonebook[username]
	if ok {
		fmt.Printf("the phone number of %v is : %v \n", username, value)
	} else {
		fmt.Println("user not found")
	}
}

func checkViaPhoneNumber(phonenumber string) {
	pl("checking via phone number ")
	for name := range phonebook {
		if phonebook[name] == phonenumber {
			fmt.Printf("the user of this phonenumber is %v \n ", name)
			return
		}
	}
	fmt.Println("user not found")
}
func main() {

	phonebook["Roii"] = "0545300668"
	phonebook["Lior"] = "0542223763"

	pl("hello and welcome to my phonebook")
	for {
		pl("press 1 for : insert new user \npress 2 for : check user by name \npress 3 for : check user by phone number\npress 4 for : delete user from phonebook \npress 5 for : print the phone book")
		var task int
		fmt.Scanf("%d", &task)
		switch task {
		case 1: //insert new user
			{
				insertNewUser()
			}

		case 2:
			{
				pl("please inset the name and to check")
				reader := bufio.NewReader(os.Stdin)
				usernameforchecking, err := reader.ReadString('\n')
				if err == nil {
					checkViaUserName(usernameforchecking)
				}
			}
		case 3:
			{
				pl("please inset the phone number and to check")
				reader := bufio.NewReader(os.Stdin)
				phonenumberforchecking, err := reader.ReadString('\n')
				if err == nil {
					checkViaPhoneNumber(phonenumberforchecking)
				}
			}
		case 4:
			{
				pl("please inset the name you want to delete")
				reader := bufio.NewReader(os.Stdin)
				usernametodelete, err := reader.ReadString('\n')
				if err == nil {
					deleteUserFromPhoneBook(usernametodelete)
				}
			}
		case 5:
			printPhoneBook()
		default:
		}
	}

}
