package main

import "fmt"

func main() {

	m := make(map[string]string)

	m = map[string]string{
		"AUD": "Australian dollar",
		"USD": "US dollar",
		"INR": "Indian Rupees",
		"JPY": "Japan Yen",
	}

	for key, value := range m {
		fmt.Println(key, value)
	}

	delete(m, "GBP")
	fmt.Println("Currency with GBP deleted")

	// Comma idiom
	if val, ok := m["USD"]; ok {
		fmt.Printf("The value %s is present", val)
	} else {
		fmt.Println("The key is not found")
	}

	nameAndHobby := map[string][]string{
		"steven": []string{"F1", "Table Tennis", "Coding", "Reading"},
		"gorge":  []string{"Sleeping", "watching tv", "Eating"},
	}

	nameAndHobby["Tom"] = []string{"Reading", "dancing", "walking", "jogging"}

	for k, v := range nameAndHobby {
		fmt.Printf("%v likes \n", k)
		for j, val := range v {
			fmt.Printf("\t %v %v\n", j, val)
		}
	}

	currency := map[string]map[string]int{
		"Great Britain Pound": {"GBP": 1},
		"Euro":                {"EUR": 2},
		"USA Dollar":          {"USD": 3},
	}

	for key, value := range currency {
		fmt.Printf("Currency Name: %v\n", key)
		for k, v := range value {
			fmt.Printf("\t Currency Code: %v\n\t\t Ranking: %v\n", k, v)
		}
	}
}
