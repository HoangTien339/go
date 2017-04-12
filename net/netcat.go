package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	// "time"
)

const (
	Host  = "bullandcow-challenge.framgia.vn"
	Port  = "2015"
	Sleep = 1
)

func streamTCP(host string, port string, m map[string]int) bool {
	var output string
	var c, b int
	buf := make([]byte, 1024)

	destinationPort := fmt.Sprintf(":%v", port)
	con, err := net.Dial("tcp", host+destinationPort)
	if err != nil {
		fmt.Printf("Connect error: %s\n", err)
	}
	fmt.Println("Connected to:", Host+Port)

	preDigitNum := 0
	digitNum := 0
	correct := false
	for {
		var nBytes int
		var err error
		nBytes, err = con.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Printf("Read error: %s\n", err)
			}
			break
		}
		output = string(buf[0:nBytes])
		fmt.Printf(output)
		isWrite := strings.Index(output, "Your answer")
		digitNum = getRequireDigitNum(output)
		fmt.Printf("Guessing... ! Digitnum: %v | isWrite: %v\n", digitNum, isWrite)
		if isWrite > 0 {
			if preDigitNum != digitNum || correct {
				fmt.Println("Map initialized")
				m = make(map[string]int)
				var digits []byte = []byte("1234567890")
				if digitNum == 6 {
					digits = []byte("1234567890")
				}
				fmt.Printf("Map number of previous: %v | current: %v\n", preDigitNum, digitNum)
				generate4DigitsNumber(m, digits, 0, digitNum)
				preDigitNum = digitNum
				correct = false
			}

			numStr := pickAGuess4Digits(m, digitNum)

			num := []byte(numStr + "\n")
			_, err = con.Write(num)
			if err != nil {
				fmt.Printf("Write error: %s\n", err)
			}

			// Submit 4 digits number & Get output result
			nBytes, err = con.Read(buf)
			output = string(buf[0:nBytes])

			fmt.Println(string(numStr))
			fmt.Print(output)

			// Parse cow & bull
			c, b = parseCowsBullsFromString(output)

			removeNonMatching4Digits(numStr, m, c, b)

			if checkForWin(b, output) {
				correct = true
			}
		}
		// time.Sleep(Sleep * time.Millisecond)
	}
	return false
}

func generate4DigitsNumber(m map[string]int, digits []byte, fixed int, digitNum int) {
	if fixed == digitNum {
		m[string(digits[:digitNum])] = 0
		return
	}
	for i := fixed; i < len(digits); i++ {
		digit, _ := strconv.ParseInt(string(digits[i]), 10, 0)
		if (fixed == 0) && (digit == 0) {
			continue
		}
		digits[fixed], digits[i] = digits[i], digits[fixed]
		generate4DigitsNumber(m, digits, fixed+1, digitNum)
		digits[fixed], digits[i] = digits[i], digits[fixed]
	}
}

func removeNonMatching4Digits(guess string, m map[string]int, c int, b int) {
	for matching := range m {
		var cows, bulls int = 0, 0
		for ig, vg := range guess {
			switch strings.IndexRune(matching, vg) {
			case -1: // Skip case not found
			case ig:
				bulls++
			default:
				cows++
			}
		}
		if cows != c || bulls != b {
			delete(m, matching)
		}
	}
}

func parseCowsBullsFromString(s string) (c, b int) {
	c = strings.Count(s, "cow")
	b = strings.Count(s, "bull")
	return
}

func pickAGuess4Digits(m map[string]int, digitNum int) string {
	var guess string
	switch digitNum {
	case 4:
		_, ok := m["1234"]
		if ok {
			guess = "1234"
		}
	case 5:
		_, ok := m["12345"]
		if ok {
			guess = "12345"
		} else {
			_, ok := m["67890"]
			if ok {
				guess = "67890"
			}
		}
	case 6:
		_, ok := m["123456"]
		if ok {
			guess = "123456"
		} else {
			_, ok := m["789012"]
			if ok {
				guess = "789012"
			}
		}
	}
	if len(guess) > 0 {
		delete(m, guess)
		return guess
	}
	for guess = range m {
		delete(m, guess)
		break
	}
	return guess
}

func checkForWin(b int, output string) bool {
	if (b == 4) || (strings.Index(output, "Correct") > 0) {
		return true
	}
	return false
}

func getRequireDigitNum(s string) int {
	trimToNum := func(r rune) bool {
		if n := r - '0'; n > 0 && n <= 9 {
			return false
		}
		return true
	}
	n := strings.TrimFunc(s, trimToNum)
	i, _ := strconv.Atoi(n)
	fmt.Println("Parse Number: ", i)
	return i
}

func main() {
	m := make(map[string]int)
	r := false
	for !r {
		r = streamTCP(Host, Port, m)
	}
}
