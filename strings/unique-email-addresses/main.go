package main

import (
	"fmt"
	"strings"
)

func main() {
	input1 := []string{
		"test.email+alex@leetcode.com",
		"test.e.mail+bob.cathy@leetcode.com",
		"test.e.mail+bob.cathy@leetcode.com",
		"testemail+david@lee.tcode.com",
	}

	fmt.Println(numUniqueEmails(input1))

	input2 := []string{
		"a@leetcode.com",
		"b@leetcode.com",
		"c@leetcode.com",
	}
	fmt.Println(numUniqueEmails(input2))
}

func numUniqueEmails(emails []string) int {
	var emailCountMap = make(map[string]int)
	for _, email := range emails {
		identifier := uniqueIdentifier(email)
		_, found := emailCountMap[identifier]
		if found {
			emailCountMap[identifier]++
		} else {
			emailCountMap[identifier] = 1
		}
	}
	return len(emailCountMap)
}

func uniqueIdentifier(email string) string {
	parts := strings.Split(email, "@")
	localName := parts[0]
	localName = strings.ReplaceAll(localName, ".", "")
	localName = strings.Split(localName, "+")[0]
	domainName := parts[1]
	identifier := localName + "@" + domainName
	return identifier
}
