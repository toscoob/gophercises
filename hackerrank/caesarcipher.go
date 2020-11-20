package main

import (
	"bufio"
	"fmt"
	//"io"
	"os"
	"strconv"
	"strings"
)

var asciiLower = "abcdefghijklmnopqrstuvwxyz"
var asciiUpper = strings.ToUpper(asciiLower)

func getShiftedChar(baseStr string, lookup string, k int32) (string, bool) {
	idx := strings.Index(baseStr, lookup)
	if idx < 0 {
		return lookup, false
	} else {
		newIdx := (idx + int(k)) % len(baseStr)
		return string(baseStr[newIdx]), true
	}
}
// Complete the caesarCipher function below.
func caesarCipher(s string, k int32) string {
	var ret string

	for _, ch := range s {
		newCh, ok := getShiftedChar(asciiLower, string(ch), k)
		if !ok {
			newCh, _ = getShiftedChar(asciiUpper, string(ch), k)
		}
		ret = ret + newCh
	}

	return ret
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)
	_ = n

	s := readLine(reader)

	kTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := caesarCipher(s, k)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

//func readLine(reader *bufio.Reader) string {
//	str, _, err := reader.ReadLine()
//	if err == io.EOF {
//		return ""
//	}
//
//	return strings.TrimRight(string(str), "\r\n")
//}
//
//func checkError(err error) {
//	if err != nil {
//		panic(err)
//	}
//}

