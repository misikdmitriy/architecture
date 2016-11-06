package searcher

import (
	"strings"
	"sync"
)

//SearchResult structure
type SearchResult struct {
	Position int
	Result   string
}

var (
	locker sync.Mutex
)

//Search function
func Search(str, subStr string) []SearchResult {
	array := make([]SearchResult, 0)
	startIndex := 0

	for index := 0; index < len(str); index++ {
		if str[index] == ' ' {
			locker.Lock()
			go cutWords(&startIndex, index, str, subStr, &array)
		}
	}

	var subArray = searchForAllSubstrings(str[startIndex:len(str)], subStr)
	subArray = shiftAllPositions(subArray, startIndex)
	return appendArray(array, subArray)
}

func cutWords(startIndex *int, index int, str, subStr string, array *[]SearchResult) {
	var subArray = searchForAllSubstrings(str[*startIndex:index], subStr)
	subArray = shiftAllPositions(subArray, *startIndex)
	*array = appendArray(*array, subArray)
	*startIndex = index + 1
	locker.Unlock()
}

func searchForAllSubstrings(str, subStr string) []SearchResult {
	startPosition := 0
	array := make([]SearchResult, 0)

	for {
		result := getFullSubstring(str[startPosition:len(str)], subStr)

		if result != "" {
			position := strings.Index(str[startPosition:len(str)], result)

			array = append(array, createSearchResult(position+startPosition, result))

			startPosition += position + 1
		} else {
			break
		}
	}

	return array
}

func createSearchResult(position int, result string) SearchResult {
	newSearchResult := new(SearchResult)
	newSearchResult.Position = position
	newSearchResult.Result = result

	return *newSearchResult
}

func shiftAllPositions(searchResult []SearchResult, shift int) []SearchResult {
	for index := 0; index < len(searchResult); index++ {
		searchResult[index].Position += shift
	}

	return searchResult
}

func getFullSubstring(str, subStr string) string {
	allSubStrs := splitWithAsterisks(subStr)
	startIndex := 0
	finishIndex := 0
	lastFoundIndex := 0

	for currentIndex, searchStr := range allSubStrs {
		if searchStr == "*" {
			if currentIndex == 0 {
				startIndex = 0
			}
			if currentIndex == len(allSubStrs)-1 {
				finishIndex = len(str)
			}
			continue
		} else {
			index := strings.Index(str[lastFoundIndex:len(str)], searchStr)
			lastFoundIndex += index + 1
			if currentIndex == 0 {
				startIndex = index
			}
			if currentIndex == len(allSubStrs)-1 {
				finishIndex = lastFoundIndex + len(searchStr) - 1
			}

			if index == -1 {
				return ""
			}
		}
	}

	return str[startIndex:finishIndex]
}

func appendArray(array1, array2 []SearchResult) []SearchResult {
	for index := 0; index < len(array2); index++ {
		array1 = append(array1, array2[index])
	}

	return array1
}

func splitWithAsterisks(str string) []string {
	allSubStrs := make([]string, 0)
	startIndex := 0

	for {
		if startIndex >= len(str) {
			break
		}
		index := strings.Index(str[startIndex:len(str)], "*") + startIndex
		if index-startIndex != -1 {
			if index > 0 {
				allSubStrs = append(allSubStrs, str[startIndex:index])
			}
			allSubStrs = append(allSubStrs, "*")
			startIndex = index + 1
		} else {
			allSubStrs = append(allSubStrs, str[startIndex:len(str)])
			startIndex = len(str)
		}
	}

	return allSubStrs
}
