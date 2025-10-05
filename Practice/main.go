package main 




func isHappy(n int) bool {
	seen := make(map[int]bool)


		for n != 1{
			sum := 0	
			for n > 0{
				digit := n % 10
				sum += digit * digit
				n /= 10
			}
			if seen[sum]{
				return false
			}			
			seen[sum] = true
			n = sum
		}		
		return true
	}

func main() {
	println(isHappy(17))
}
