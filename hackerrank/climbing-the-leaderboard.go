package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'climbingLeaderboard' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY ranked
 *  2. INTEGER_ARRAY player
 */

func binarySearch(a []int32, key int32) int32 {
	lo := int32(0)
	hi := int32(len(a) - 1)

	for lo <= hi {
		mid := int32(lo + (hi-lo)/2)
		if a[mid] == key {
			return mid
		} else if a[mid] < key && key < a[mid-1] {
			return mid
		} else if a[mid] > key && key >= a[mid+1] {
			return mid + 1
		} else if a[mid] < key {
			hi = mid - 1
		} else if a[mid] > key {
			lo = mid + 1
		}
	}
	return -1
}

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	// rankedNew := rank(ranked)
	// fmt.Println(rankedNew)
	// fmt.Println(ranked)
	rank := make([]int32, len(ranked), len(ranked))
	rank[0] = 1
	for i := 1; i < len(ranked); i++ {
		if ranked[i] == ranked[i-1] {
			rank[i] = rank[i-1]
		} else {
			rank[i] = rank[i-1] + 1
		}
	}
	fmt.Println(rank)
	nalice := make([]int32, len(player), len(player))
	for j, v := range player {
		if v > ranked[0] {
			nalice[j] = 1
		} else if v < ranked[len(ranked)-1] {
			nalice[j] = rank[len(ranked)-1] + 1
		} else {
			nalice[j] = rank[binarySearch(ranked, v)]
		}
	}

	return nalice
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	rankedCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	rankedTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var ranked []int32

	for i := 0; i < int(rankedCount); i++ {
		rankedItemTemp, err := strconv.ParseInt(rankedTemp[i], 10, 64)
		checkError(err)
		rankedItem := int32(rankedItemTemp)
		ranked = append(ranked, rankedItem)
	}

	playerCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	playerTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var player []int32

	for i := 0; i < int(playerCount); i++ {
		playerItemTemp, err := strconv.ParseInt(playerTemp[i], 10, 64)
		checkError(err)
		playerItem := int32(playerItemTemp)
		player = append(player, playerItem)
	}

	result := climbingLeaderboard(ranked, player)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
