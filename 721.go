package main

import "sort"

// Given a list of accounts where each element accounts[i] is a list of strings, where the first element accounts[i][0] is a name, and the rest of the elements are emails representing emails of the account.
//
// Now, we would like to merge these accounts. Two accounts definitely belong to the same person if there is some common email to both accounts. Note that even if two accounts have the same name, they may belong to different people as people could have the same name. A person can have any number of accounts initially, but all of their accounts definitely have the same name.
//
// After merging the accounts, return the accounts in the following format: the first element of each account is the name, and the rest of the elements are emails in sorted order. The accounts themselves can be returned in any order.
//
//
//
// Example 1:
//
// Input: accounts = [["John","johnsmith@mail.com","john_newyork@mail.com"],["John","johnsmith@mail.com","john00@mail.com"],["Mary","mary@mail.com"],["John","johnnybravo@mail.com"]]
// Output: [["John","john00@mail.com","john_newyork@mail.com","johnsmith@mail.com"],["Mary","mary@mail.com"],["John","johnnybravo@mail.com"]]
// Explanation:
// The first and third John's are the same person as they have the common email "johnsmith@mail.com".
// The second John and Mary are different people as none of their email addresses are used by other accounts.
// We could return these lists in any order, for example the answer [['Mary', 'mary@mail.com'], ['John', 'johnnybravo@mail.com'],
// ['John', 'john00@mail.com', 'john_newyork@mail.com', 'johnsmith@mail.com']] would still be accepted.
//
// Example 2:
//
// Input: accounts = [["Gabe","Gabe0@m.co","Gabe3@m.co","Gabe1@m.co"],["Kevin","Kevin3@m.co","Kevin5@m.co","Kevin0@m.co"],["Ethan","Ethan5@m.co","Ethan4@m.co","Ethan0@m.co"],["Hanzo","Hanzo3@m.co","Hanzo1@m.co","Hanzo0@m.co"],["Fern","Fern5@m.co","Fern1@m.co","Fern0@m.co"]]
// Output: [["Ethan","Ethan0@m.co","Ethan4@m.co","Ethan5@m.co"],["Gabe","Gabe0@m.co","Gabe1@m.co","Gabe3@m.co"],["Hanzo","Hanzo0@m.co","Hanzo1@m.co","Hanzo3@m.co"],["Kevin","Kevin0@m.co","Kevin3@m.co","Kevin5@m.co"],["Fern","Fern0@m.co","Fern1@m.co","Fern5@m.co"]]
//
//
//
// Constraints:
//
// 1 <= accounts.length <= 1000
// 2 <= accounts[i].length <= 10
// 1 <= accounts[i][j] <= 30
// accounts[i][0] consists of English letters.
// accounts[i][j] (for j > 0) is a valid email.

// tc: O(nm log(nm)), n: # of accounts, m: average emails count for an account
func accountsMerge(accounts [][]string) [][]string {
	edges := make(map[string][]int)

	for i, account := range accounts {
		for j := 1; j < len(account); j++ {
			edges[account[j]] = append(edges[account[j]], i)
		}
	}

	size := len(accounts)
	groups := make([]int, size)
	ranks := make([]int, size)

	for i := range groups {
		groups[i] = i
		ranks[i] = 1
	}

	for _, edge := range edges {
		for i := 1; i < len(edge); i++ {
			p1, p2 := find(groups, edge[0]), find(groups, edge[i])

			if p1 != p2 {
				if ranks[p1] >= ranks[p2] {
					groups[p2] = p1
					ranks[p1] += ranks[p2]
				} else {
					groups[p1] = p2
					ranks[p2] += ranks[p1]
				}
			}
		}
	}

	accountEmails := make(map[int]map[string]bool)

	for i, account := range accounts {
		idx := find(groups, i)

		if _, ok := accountEmails[idx]; !ok {
			accountEmails[idx] = make(map[string]bool)
		}

		for j := 1; j < len(account); j++ {
			accountEmails[idx][account[j]] = true
		}
	}

	data := make([][]string, 0)

	for idx, emailTable := range accountEmails {
		name := accounts[idx][0]

		tmp := []string{name}

		emails := make([]string, 0)
		for email := range emailTable {
			emails = append(emails, email)
		}
		sort.Strings(emails)

		tmp = append(tmp, emails...)

		data = append(data, tmp)
	}

	return data
}

func find(groups []int, idx int) int {
	if groups[idx] != idx {
		groups[idx] = find(groups, groups[idx])
	}

	return groups[idx]
}

//	Notes
//	1.	inspired from sample code, not need to have a special edges table,
//		just store all previous emails -> idx, as long as record exist in
//		map, do union
