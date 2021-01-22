package main
import "fmt"
import "sort"

func main() {
	barVal := map[string]int{
		"alpha": 34,
		"bravo": 56,
		"charlie": 23,
		"delta": 87,
		"echo": 56,
		"foxtrot": 12,
		"golf": 34,
		"hotel": 16,
		"indio": 87,
		"juliet": 65,
		"kilo": 43,
		"lima": 98,
	}

	fmt.Printf("Unsorted:\n")
	for k, v := range barVal {
		fmt.Printf("key: %v\tvalue: %v\n", k, v)
	}

	keys := make([]string, len(barVal))
	i := 0
	for k := range barVal {
		keys[i] = k
		i++
	}

	sort.Strings(keys)
	fmt.Printf("\nSorted:\n")
	for _, k := range keys {
		fmt.Printf("key: %v\tvalue: %v\n", k, barVal[k])		
	}
}