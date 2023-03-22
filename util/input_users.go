package util

import (
	"bufio"
	"fmt"
	"os"
)

func InputUserInteger() int32 {
	var inputUserInt int32
	fmt.Scanln(&inputUserInt)
	return inputUserInt
}

func InputUserString() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return input, nil
}
