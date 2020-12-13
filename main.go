package main
import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
  "unicode/utf8"
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
      policyCharPosition1, _ := strconv.Atoi(policyTimesRanges[0])
      policyCharPosition2, _ := strconv.Atoi(policyTimesRanges[1])

      policyChar := string(lineParts[1][0]) // Second Part, only first Char relevant (convert from byte)
      password := lineParts[2]

      //fmt.Println("password: ", password, " len: ", len(password))
      position1Matches := utf8.RuneCountInString(string(password)) >= policyCharPosition1 && string(password[policyCharPosition1-1]) == policyChar

      position2Matches := utf8.RuneCountInString(string(password)) >= policyCharPosition2 && string(password[policyCharPosition2-1]) == policyChar

      if(position1Matches && !position2Matches) || (!position1Matches && position2Matches) {
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