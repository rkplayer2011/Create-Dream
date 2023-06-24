package main

import "fmt"

func main() {

	// Creating a dream
	fmt.Println("Beginning the process of creating a dream...")
	
    // Setting the goal of the dream
	fmt.Println("Setting a goal for the dream...")
	var goal string 
	fmt.Println("What is the goal of the dream?")
	fmt.Scan(&goal)
	
    // Choosing a location for the dream
	fmt.Println("Choosing a location for the dream...")
	var location string
	fmt.Println("Where is the dream located?")
	fmt.Scan(&location)
	
    // Gathering supplies for the dream
	fmt.Println("Gathering supplies for the dream...")
	var supplies string
	fmt.Println("What supplies will be needed for the dream to take shape?")
	fmt.Scan(&supplies)
	
    // Inviting people to the dream
	fmt.Println("Inviting people to the dream...")
	var people string
	fmt.Println("Who should be invited to the dream?")
	fmt.Scan(&people)
	
    // Planning the dream
	fmt.Println("Planning the dream...")
	var plan string
	fmt.Println("What steps need to be taken to make the dream come to life?")
	fmt.Scan(&plan)
	
    // Gathering resources for the dream
	fmt.Println("Gathering resources for the dream...")
	var resources string
	fmt.Println("What resources are needed for the dream to be successful?")
	fmt.Scan(&resources)
	
    // Preparing for the dream
	fmt.Println("Preparing for the dream...")
	var preparation string
	fmt.Println("What needs to be done to prepare for the dream to come true?")
	fmt.Scan(&preparation)
	
    // Crafting details of the dream
	fmt.Println("Crafting details of the dream...")
	var details string
	fmt.Println("What details need to be included in the dream?")
	fmt.Scan(&details)
	
    // Making the dream come true
	fmt.Println("Making the dream come true...")
	var action string
	fmt.Println("What action needs to be taken to make the dream come true?")
	fmt.Scan(&action)
	
    // Celebrating the dream
	fmt.Println("Celebrating the dream...")
	var celebration string
	fmt.Println("How will the dream be celebrated?")
	fmt.Scan(&celebration)

	// Dream is complete
	fmt.Println("The dream is now complete!")
		
	fmt.Printf("The dream is located in %s, with a goal of %s and requires %s as supplies. %s were invited to be part of the dream, and it required %s to make it come to life. The details of the dream included %s in order to make it real. %s was done to make it true, and it was celebrated with %s.\n", location, goal, supplies, people, plan, details, action, celebration)
	
}