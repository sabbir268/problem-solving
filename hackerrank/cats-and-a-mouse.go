package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func convertN(n int32) int32 {
	if n < 0 {
		return n * -1
	}
	return n
}

// Complete the catAndMouse function below.
func catAndMouse(x int32, y int32, z int32) string {
	// fmt.Println(x)
	var catA int32 = convertN(x - z)

	var catB int32 = convertN(y - z)
	var result string
	if catA == catB {
		result = "Mouse C"
	} else if catA < catB {
		result = "Cat A"
	} else {
		result = "Cat B"
	}

	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		xyz := strings.Split(readLine(reader), " ")

		xTemp, err := strconv.ParseInt(xyz[0], 10, 64)
		checkError(err)
		x := int32(xTemp)

		yTemp, err := strconv.ParseInt(xyz[1], 10, 64)
		checkError(err)
		y := int32(yTemp)

		zTemp, err := strconv.ParseInt(xyz[2], 10, 64)
		checkError(err)
		z := int32(zTemp)

		result := catAndMouse(x, y, z)

		fmt.Fprintf(writer, "%s\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}