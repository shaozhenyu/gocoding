package main

import (
	"fmt"
	"strings"
)

func main() {
	emails := []string{"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"}
	fmt.Println(numUniqueEmails(emails))
}

func numUniqueEmails(emails []string) int {
	m := make(map[string]struct{})
	for _, e := range emails {
		s := strings.Split(e, "@")
		left, right := s[0], s[1]
		left = strings.Replace(strings.Split(left, "+")[0], ".", "", -1)
		m[left + right] = struct{}{}
	}
	return len(m)
}
