package main

import (
	"fmt"
	"time"
	"sync"
)


const conferenceTickets = 50
var conferenceName  = "Go Conference"
var remainingTickets int = conferenceTickets
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets int
}

var wg = sync.WaitGroup{}


func main(){

greetUsers()


	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)


	if isValidTicketNumber && isValidName && isValidEmail {
		
		bookTicket( remainingTickets, userTickets,  firstName, lastName, email )

		wg.Add(1)
		go sendTicket( userTickets,  firstName, lastName, email )

		firstNames := getFirstNames()

		fmt.Printf("The first names of bookings are %v\n", firstNames)

		if remainingTickets == 0 {
			//end program
			fmt.Println("Our conference is sold out. Come back next year.")
		}
	} else {
		if !isValidName {
			fmt.Println("First or last name you entered is too short!")
		}
		if !isValidEmail {
			fmt.Println("Email address entered does not contain '@' sign")
		}
		if !isValidName {
			fmt.Println("Number of tickets you entered is invalid")
		}

	}

	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets, and %v tickets available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
 
func getFirstNames() []string{
		firstNames := []string{}
	
		for _, booking := range bookings {
			firstNames = append(firstNames, booking.firstName)
		}

		return firstNames

}


func getUserInput()(string, string, string, int){
	var firstName string
	var lastName string
	var email string
	var userTickets int
	// ask user for their name
	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Please enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets int, userTickets int, firstName string, lastName string, email string){
	remainingTickets = remainingTickets - userTickets

	//create a map for a user
	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookins is %v", bookings)


	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation emailt at %v\n", firstName, lastName, userTickets, email)	
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#########")
	fmt.Printf("Sending ticket: \n  %v  \nto email address %v\n", ticket, email)
	fmt.Println("#########")
	wg.Done()
}