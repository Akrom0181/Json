	package main

	import (
	 "bufio"
	 "fmt"
	 "os"
	 "strconv"
	 "strings"
	)
	
	func main() {
	
	 	file, err := os.Open("file.txt")
	 	if err != nil {
	  		fmt.Println("Error: ", err)
	  	return
	 	}
	 	defer file.Close()
	
	 	scanner := bufio.NewScanner(file)
	 	var nums []int
	 	var words []string
	 	for scanner.Scan() {
	  		line := scanner.Text()
	
	  	for _, word := range strings.Fields(line) {
	   		if num, err := strconv.Atoi(word); err == nil {
				nums = append(nums, num)
	   		} else {
				words = append(words, word)
	   		}
	  	}
		}
	
	 	if err := scanner.Err(); err != nil {
	  		fmt.Println("Error", err)
	  		return
		}
	
	 	fmt.Println("Number:", nums)
	 	fmt.Println("Words:", words)
}
	
