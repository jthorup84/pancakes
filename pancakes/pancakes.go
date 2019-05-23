package pancakes

import (
	"fmt"
	"regexp"
	"strings"
)

func GetFlipCounts(input string) string {
	parts := strings.Split(input, "\n")
	// Didn't see the need to use the first line to determine
	// number of test cases. I just loop over all the lines that are given.
	// numberOfTestCases, _ := strconv.Atoi(parts[0])
	stacks := parts[1:]
	var flipCounts []string

	for i, stack := range stacks {
		flipCount := getFlipCount(stack, 0)
		flipCounts = append(flipCounts, fmt.Sprintf("Case #%d: %d", i+1, flipCount))
	}

	return strings.Join(flipCounts, "\n")
}

func getFlipCount(stack string, numberOfFlips int) int {
	// check to see if all pancakes are happy side up
	allHappy, _ := regexp.MatchString(`^\++$`, stack)
	if allHappy {
		return numberOfFlips
	}

	// flip consecutive pancakes of same orientation
	newStack := flipPancakes(stack)

	return getFlipCount(newStack, numberOfFlips+1)
}

func flipPancakes(stack string) string {
	var regex *regexp.Regexp
	// determine if flipping happy or blank side pancakes
	if strings.HasPrefix(stack, "+") {
		regex = regexp.MustCompile(`^\++`)
	} else {
		regex = regexp.MustCompile(`^\-+`)
	}

	numberToFlip := len(regex.FindString(stack))

	return regex.ReplaceAllString(stack, strings.Repeat(stack[:1], numberToFlip))
}
