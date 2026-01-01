package generalOutput

import (
	"bufio"
	"fmt"
	"os"
)

func HelloWorld() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your name: ")
	input, _ := reader.ReadString('\n')

	fmt.Println("Hello ", input)

}
