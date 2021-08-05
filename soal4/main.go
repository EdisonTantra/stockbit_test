package anagram

import (
	"anagram"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text (delimeter ','): ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	input := strings.Split(text, ",")

	result := anagram.Detect(input)
	fmt.Println(result)
}
