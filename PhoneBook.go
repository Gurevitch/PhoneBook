package main

import (
	"fmt"
)

var pl = fmt.Println
var phonebook = make(map[string]string)

func insertNewUser(username string, phonenumber string) {
	// If the key exists
	if _, ok := phonebook[username]; ok {
		pl("user already excised...")
	} else {
		pl("New user will be added to the phonebook")
		phonebook[username] = phonenumber
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
func main() {

	phonebook["Roii"] = "0545300668"
	phonebook["Lior"] = "0542223763"
	/*reader := bufio.NewReader(os.Stdin)
	NewUserName, err := reader.ReadString('\n')
	if err == nil {
		NewPhoneNumber, err := reader.ReadString('\n')
		if err == nil {
			insertNewUser(NewUserName, NewPhoneNumber)
		}
	}
	*/
	insertNewUser("Abc", "0543156781")
	printPhoneBook()
	deleteUserFromPhoneBook("Abc")
	printPhoneBook()
}

/*
-add search by name
-add search by phonenumber

*/
