package pancakes

import (
	"fmt"
	"regexp"
	"strings"
)

func GetFlipCounts(input string) string {
	parts := strings.Split(input, "\n")
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

	var newStack string
	if strings.HasPrefix(stack, "+") {
		re := regexp.MustCompile(`^\++`)
		numberToFlip := len(re.FindString(stack))
		newStack = re.ReplaceAllString(stack, strings.Repeat("-", numberToFlip))
	} else {
		re := regexp.MustCompile(`^\-+`)
		numberToFlip := len(re.FindString(stack))
		newStack = re.ReplaceAllString(stack, strings.Repeat("+", numberToFlip))
	}

	return getFlipCount(newStack, numberOfFlips+1)
}
