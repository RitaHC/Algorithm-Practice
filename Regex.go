package main

//////////////////// REGEX - Regular Expression ///////////////////
// Meaning: A sequence of characters that define a search pattern
// Go Package : regexp (cobtains actions like filtering, replacing, validating or extracting)

import (
	"fmt"
	"log"
	"net/mail"
	"regexp"
)

// ///////////////// Practice 3 /////////////////////
// Check if the email address is formatted correctly - USING STRING
func checkEmailValidation(email string) (string, bool) {
	a := true
	pattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	matches := pattern.MatchString(email)
	if matches == false {
		a = false
	}
	return email, a
}

// ///////////////// Practice 4 /////////////////////
// Check if the email address is formatted correctly - USING SLICE

func emailValidationSlice(emails []string) {

	// Create a regexp for an email address to be checked
	pattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	// Iterate over the slice of emails
	for _, email := range emails {
		// Check of the email pattern matches the email
		if pattern.MatchString(email) {
			fmt.Printf("Valid email: %s\n", email)
		} else {
			fmt.Printf("Invalid email: %s\n", email)
		}
	}

}

// ///////////////// Practice 5 /////////////////////
// Check if the email address is formatted correctly - USING MAPS

func emailValidationMap(nameEmails map[string]string) {
	// create a regular expression to check email validation
	pattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	// Iterate over map.value and check if it's true
	for key, value := range nameEmails {
		match := pattern.MatchString(value)
		if match {
			fmt.Println(key, "'s Email address is CORRECT")
		} else {
			fmt.Println(key, "'s Email address is INCORRECT")
		}
	}
	// Print the output
}

// ///////////////// Practice 6 /////////////////////
// Check if the email address is formatted correctly - USING STRUCTS

type Address struct {
	Name    string
	Address string
}

func emailValidationStruct(emailStruct Address) {
	pattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	match := pattern.MatchString(emailStruct.Address)
	if match {
		fmt.Println(emailStruct.Name, " : email CORRECT")
	} else {
		fmt.Println(emailStruct.Name, " : email InCORRECT")
	}
}

func main() {
	fmt.Println("------Practice 1------")
	//============== Practice 1 =================
	str := "greeksforgreeks"
	match1, err := regexp.MatchString("greeks", str)
	fmt.Println("Match: ", match1, "\nError", err)

	//============== Practice 2 =================
	fmt.Println("------Practice 2------")
	// Parse an email address
	email, err := mail.ParseAddress("Alice <alice@example.com>")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Email address is : ", email)

	//============== Practice 3 =================
	fmt.Println("------Practice 3------")
	fmt.Println(checkEmailValidation("hameer"))
	fmt.Println(checkEmailValidation("hameer@rita.com"))

	//============== Practice 4 =================
	fmt.Println("------Practice 4------")
	emails := []string{"Hameer@123.com", "Rajat376@gmail.com", "Rita"}
	emailValidationSlice(emails)

	//============== Practice 5 =================
	fmt.Println("------Practice 5------")
	nameEmails := map[string]string{"Hameer": "Hameer@123.com", "Rajat": "Rajat376@gmail.com", "Rita": "Rita.com"}
	emailValidationMap(nameEmails)

	//============== Practice 6 =================
	fmt.Println("------Practice 6------")
	// create an instance of the struct
	emailAddress1 := Address{Name: "Hameer Chauhan", Address: "Hameer@123.com"}
	emailAddress2 := Address{Name: "Rita H Chauhan", Address: "Rita.com"}
	emailAddress3 := Address{Name: "Rajat Kumar", Address: "rajat376@gmail.com"}
	emailValidationStruct(emailAddress1)
	emailValidationStruct(emailAddress2)
	emailValidationStruct(emailAddress3)
}
