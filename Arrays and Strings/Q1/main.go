package main


import ("fmt")


	func main() {
		word := "hello"

		m:= make(map[string]int)  //Create a map to store character counts

		for i := range word {
			char := string(word[i]) 
			if _, exists := m[char]; exists {
				m[char]++
			} else {
				m[char] = 1
				}
			if m[char] > 1 {
				fmt.Println("Repeated Character:", char)
			}
			}	
			fmt.Println("Character counts:", m)
	
	}