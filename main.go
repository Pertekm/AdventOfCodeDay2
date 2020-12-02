package main
import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
)

func main() {
    fmt.Println("Starting Advent of Code Day 2")

    passwordsMatchPolicyCount := 0
    input := readInput()
    inputLines := strings.Split(input, "\n")
    indexLast := 0

    for index, line := range inputLines {
      //fmt.Println("Doing line: ", index, line)
      lineParts := strings.Split(line, " ")

      policyTimesWhole := lineParts[0]
      policyTimesRanges := strings.Split(policyTimesWhole, "-")
      rangeStart, _ := strconv.Atoi(policyTimesRanges[0])
      rangeEnd, _ := strconv.Atoi(policyTimesRanges[1])

      policyChar := string(lineParts[1][0]) // Second Part, only first Char relevant (convert from byte)
      password := lineParts[2]

      //fmt.Println("rangeStart: ", rangeStart, " rangeEnd: ", rangeEnd, " policyChar: ", policyChar, " password:", password)

      countOfPolicyChar := strings.Count(password, policyChar)
      
      if countOfPolicyChar >= rangeStart && countOfPolicyChar <= rangeEnd {
        // passwordMatchesPolicy
        passwordsMatchPolicyCount++
        //fmt.Println("password Match Policy")
      } else {
        //fmt.Println("password does not Match Policy")
      }

      indexLast = index
    }

    fmt.Println("passwords match policy count: ", passwordsMatchPolicyCount, " of total ", indexLast + 1)
}

func readInput() string {
    content, err := ioutil.ReadFile("puzzleInput.txt")

    if err != nil {
        return ""
    }

    // Convert []byte to string and print to screen
    text := string(content)
    return text
}