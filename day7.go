package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func day7() {
	content, err := os.ReadFile("inputs/7.txt")
	if err != nil {
		log.Fatal(err)
	}
	cardSets := string(content)
	setSlice := strings.Split(cardSets, "\r\n")
	handMap := make(map[string]int)
	tierSlice := make([][]string, 7)
	for _, set := range setSlice {
		splitSet := strings.Split(set, " ")
		num, err := strconv.Atoi(splitSet[1])
		if err != nil {
			log.Fatal(err)
		}
		handMap[splitSet[0]] = num
	}
	for c := range handMap {
		cCount := make(map[rune]int)
		longest := 0
		jCount := 0
		for _, r := range c {
			if r == 'J' {
				jCount++
			}
			cCount[r]++
			if cCount[r] > 3 {
				longest = 4
			} else if cCount[r] > 2 {
				longest = 3
			}
		}
		lc := len(cCount)
		if jCount > 0 && jCount != 5 {
			lc--
		}
		if lc == 1 {
			tierSlice[0] = append(tierSlice[0], c)
		} else if lc == 2 {
			if longest == 4 || jCount > 1 || jCount == 1 && longest == 3 {
				tierSlice[1] = append(tierSlice[1], c)
			} else {
				tierSlice[2] = append(tierSlice[2], c)
			}

		} else if lc == 3 {
			if longest == 3 || jCount > 0 {
				tierSlice[3] = append(tierSlice[3], c)
			} else {
				tierSlice[4] = append(tierSlice[4], c)
			}
		} else if lc == 4 {
			tierSlice[5] = append(tierSlice[5], c)
		} else if lc == 5 {
			tierSlice[6] = append(tierSlice[6], c)
		}
	}
	rank := len(setSlice)
	total := 0
	for i := range tierSlice {
		slices.SortFunc(tierSlice[i], func(i, j string) int {
			n := 0
			for i[n] == j[n] {
				n++
			}
			oT := []byte{
				'A': 'A',
				'K': 'B',
				'Q': 'C',
				'T': 'D',
				'9': 'E',
				'8': 'F',
				'7': 'G',
				'6': 'H',
				'5': 'I',
				'4': 'J',
				'3': 'K',
				'2': 'L',
				'J': 'M',
			} // lmao the ones that's greater has the weaker card value abit confusion
			if oT[i[n]] > oT[j[n]] {
				return 1 // i gets pushed forward
			} else {
				return -1 // order stays the same
			}
		})
		for _, group := range tierSlice[i] {
			total += rank * handMap[group]
			rank--
		}
	}
	// for i, t := range tierSlice {
	// 	fmt.Println(i + 1)
	// 	fmt.Println(t)
	// 	fmt.Printf("\n")
	// }
	// fmt.Println(rank)
	fmt.Println(total)
	// too low 249495725
}

// func day7() {
// 	content, err := os.ReadFile("inputs/7.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	cardSets := string(content)
// 	setSlice := strings.Split(cardSets, "\r\n")
// 	handMap := make(map[string]int)
// 	tierSlice := make([][]string, 7)
// 	for _, set := range setSlice {
// 		splitSet := strings.Split(set, " ")
// 		num, err := strconv.Atoi(splitSet[1])
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		handMap[splitSet[0]] = num
// 	}
// 	for c := range handMap {
// 		cCount := make(map[rune]int)
// 		longest := 0
// 		for _, r := range c {
// 			cCount[r]++
// 			if cCount[r] > 3 {
// 				longest = 4
// 			} else if cCount[r] > 2 {
// 				longest = 3
// 			}
// 		}
// 		if len(cCount) == 1 {
// 			tierSlice[0] = append(tierSlice[0], c)
// 		} else if len(cCount) == 2 {
// 			if longest == 4 {
// 				tierSlice[1] = append(tierSlice[1], c)
// 			} else {
// 				tierSlice[2] = append(tierSlice[2], c)
// 			}

// 		} else if len(cCount) == 3 {
// 			if longest == 3 {
// 				tierSlice[3] = append(tierSlice[3], c)
// 			} else {
// 				tierSlice[4] = append(tierSlice[4], c)
// 			}
// 		} else if len(cCount) == 4 {
// 			tierSlice[5] = append(tierSlice[5], c)
// 		} else {
// 			tierSlice[6] = append(tierSlice[6], c)
// 		}
// 	}
// 	rank := len(setSlice)
// 	total := 0
// 	for i := range tierSlice {
// 		slices.SortFunc(tierSlice[i], func(i, j string) int {
// 			n := 0
// 			for i[n] == j[n] {
// 				n++
// 			}
// 			oT := []byte{
// 				'A': 'A',
// 				'K': 'B',
// 				'Q': 'C',
// 				'J': 'D',
// 				'T': 'E',
// 			}
// 			if i[n] > '9' && j[n] > '9' {
// 				if oT[i[n]] > oT[j[n]] {
// 					return 1
// 				} else {
// 					return -1
// 				}
// 			} else { // ascii: Letters > big numbers > smaller numbers
// 				if i[n] > j[n] {
// 					return -1 // 1 means infront, -1 means behind
// 				} else {
// 					return 1
// 				}
// 			}
// 		})
// 		for _, group := range tierSlice[i] {
// 			total += rank * handMap[group]
// 			// fmt.Printf("%s %d\n", group, handMap[group])
// 			rank--
// 		}
// 	}
// 	fmt.Println(total)
// 	// too low 46428
// 	// correct 249204891
// }
