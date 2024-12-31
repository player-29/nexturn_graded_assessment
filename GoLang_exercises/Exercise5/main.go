// Exercise 5: Climate Data Analysis 
// Topics Covered: Go Arrays, Go Strings, Go Type Casting, Go Functions, Go Conditions, 
// Go Loops 
// Case Study: 
// You are tasked with analyzing climate data from multiple cities. The data includes the 
// city name, average temperature (°C), and rainfall (mm). 
// 1. Data Input: Create a slice of structs to store data for each city. Input data can be 
// hardcoded or taken from the user. 
// 2. Highest and Lowest Temperature: Write functions to find the city with the 
// highest and lowest average temperatures. Use conditions for comparison. 
// 3. Average Rainfall: Calculate the average rainfall across all cities using loops. Use 
// type casting if necessary. 
// 4. Filter Cities by Rainfall: Use loops to display cities with rainfall above a certain 
// threshold. Prompt the user to enter the threshold value. 
// 5. Search by City Name: Allow users to search for a city by name and display its 
// data. 
// Bonus: 
// • Add error handling for invalid city names and invalid input for thresholds.

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type City struct {
	Name          string
	Temperature   float64
	Rainfall      float64
}

func main() {
	// Hardcoded data for cities
	cities := []City{
		{"New York", 20.5, 120.7},
		{"Los Angeles", 25.3, 50.2},
		{"Chicago", 15.0, 150.0},
		{"Houston", 30.1, 200.3},
		{"San Francisco", 18.6, 80.4},
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nClimate Data Analysis Menu:")
		fmt.Println("1. Find city with the highest and lowest temperature")
		fmt.Println("2. Calculate average rainfall across all cities")
		fmt.Println("3. Filter cities by rainfall threshold")
		fmt.Println("4. Search for city by name")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			highest, lowest := findTemperatureExtremes(cities)
			fmt.Printf("City with highest temperature: %s (%.2f°C)\n", highest.Name, highest.Temperature)
			fmt.Printf("City with lowest temperature: %s (%.2f°C)\n", lowest.Name, lowest.Temperature)
		case "2":
			averageRainfall := calculateAverageRainfall(cities)
			fmt.Printf("Average rainfall across all cities: %.2f mm\n", averageRainfall)
		case "3":
			fmt.Print("Enter rainfall threshold: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			threshold, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				continue
			}
			filterCitiesByRainfall(cities, threshold)
		case "4":
			fmt.Print("Enter city name: ")
			cityName, _ := reader.ReadString('\n')
			cityName = strings.TrimSpace(cityName)
			city, err := searchCityByName(cities, cityName)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("City: %s, Temperature: %.2f°C, Rainfall: %.2f mm\n", city.Name, city.Temperature, city.Rainfall)
			}
		case "5":
			fmt.Println("Exiting program. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}

func findTemperatureExtremes(cities []City) (City, City) {
	highest := cities[0]
	lowest := cities[0]
	for _, city := range cities {
		if city.Temperature > highest.Temperature {
			highest = city
		}
		if city.Temperature < lowest.Temperature {
			lowest = city
		}
	}
	return highest, lowest
}

func calculateAverageRainfall(cities []City) float64 {
	totalRainfall := 0.0
	for _, city := range cities {
		totalRainfall += city.Rainfall
	}
	return totalRainfall / float64(len(cities))
}

func filterCitiesByRainfall(cities []City, threshold float64) {
	fmt.Printf("Cities with rainfall above %.2f mm:\n", threshold)
	found := false
	for _, city := range cities {
		if city.Rainfall > threshold {
			fmt.Printf("City: %s, Rainfall: %.2f mm\n", city.Name, city.Rainfall)
			found = true
		}
	}
	if !found {
		fmt.Println("No cities found with rainfall above the given threshold.")
	}
}

func searchCityByName(cities []City, name string) (City, error) {
	for _, city := range cities {
		if strings.EqualFold(city.Name, name) {
			return city, nil
		}
	}
	return City{}, errors.New("city not found")
}
