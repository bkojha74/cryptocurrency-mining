package main

import (
	"bufio"
	"fmt"
	"mining/internal/utils"
	"os"
	"strconv"
	"strings"
)

func main() {
	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	if err != nil {
		panic(err)
	}
	defer stdout.Close()
	writer := bufio.NewWriterSize(stdout, 16*1024*1024)
	defer writer.Flush()

	inputStr, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	var input []byte
	for _, s := range strings.Fields(string(inputStr)) {
		i, err := strconv.ParseInt(s, 16, 64)
		if err != nil {
			fmt.Fprintf(writer, "Invalid input %s\n", s)
		}
		input = append(input, byte(i))
	}
	fmt.Fprintf(writer, "INPUT => %X\n", input)
	nonce := utils.FindNonce(input)
	fmt.Fprintf(writer, "OUTPUT => %X\n", nonce)
}
