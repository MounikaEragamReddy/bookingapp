package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50

// var bookings = make([]map[string]string,0) ; list of map
var bookings = make([]UserData, 0) //list of struct

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()
	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T \n", conferenceName, conferenceTickets, remainingTickets)
	//fmt.Println("Welcome to", conferenceName, "booking-application")
	//fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available. ")
	//fmt.Printf("Welcome to %v booking-application\n", conferenceName)
	//fmt.Printf("We have total of %v tickets and still %v tickets are available\n", conferenceTickets, remainingTickets)

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		//isValidCity := city == "Singapore" || city == "London"
		if isValidName && isValidEmail && isValidTicketNumber {
			bookTickets(userTickets, firstName, lastName, email)
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)
			// fmt.Printf("The whole array: %v\n", bookings)
			// fmt.Printf("The first Value: %v\n", bookings[0])
			// fmt.Printf("The type of array:%T\n", bookings)
			// fmt.Printf("The length of slice:%v\n", len(bookings))

			//Call PrintFirstNames
			firstNames := getFirstNames()
			fmt.Printf("The first names of the bookings in our app: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our Conference was completely booked.Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name u entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email adress u entered doesnt contain @")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you enteted is invalid ")
			}
			//fmt.Println("User Input is invalid, Please try again")
			//fmt.Printf("we have only %v tickets reamaining, so u cant book %v tickets\n", remainingTickets, userTickets)
		}

		// city := "London"

		// switch city {
		// case "NewYork":
		// 	//execute code for booking Newyork conforence tickets
		// case "Singapore", "Hong kong":
		// 	//execute code for booking Singapore or Hongkong conforence tickets
		// case "London", "Berlin":
		// 	//execute code for booking Londona nd Berlin conforence tickets
		// case "Mexico City":
		// 	//execute code for booking Newyork conforence tickets
		// default:
		// 	fmt.Println("No valid city")
		// }
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and still %v tickets are available\n", conferenceTickets, remainingTickets)
}
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("Please Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Please Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Please Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Please Enter number if tickets neeeded:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//create a map for a user
	//var userData = make(map[string]string)
	//create struct
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	// userData["firstName"]= firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n ", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets, you will get confirmation message on %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("Number of tickets available are %v for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var Ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("=======================================================")
	fmt.Printf("Sending Ticket:\n %v \n to email address %v\n", Ticket, email)
	fmt.Println("=======================================================")
	wg.Done()
}
