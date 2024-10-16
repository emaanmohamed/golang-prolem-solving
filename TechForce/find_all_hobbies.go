package main

import "fmt"

// return a slice containing the names of the people who enjoy the hobby
func findAllHobbyists(hobby string, hobbies map[string][]string) []string {
	var hobbyists []string
	for key, value := range hobbies {
		for _, val := range value {
			if val == hobby {
				hobbyists = append(hobbyists, key)
			}
		}
	}
	return hobbyists
}

func main() {
	hobbies := map[string][]string{
		"Steve": []string{"Fashion", "Piano", "Reading"},
		"Patty": []string{"Drama", "Magic", "Pets"},
		"Chad":  []string{"Puzzles", "Pets", "Yoga"},
	}

	fmt.Println(findAllHobbyists("Yoga", hobbies))
}
