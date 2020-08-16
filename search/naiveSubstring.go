package search

// naiveSubstringSearch: traverses the string and checks the entire substring against any potential matches
func NaiveSubstringSearch(corpus, substring string) (match bool, startIndex int) {
	if len(corpus) < len(substring) {
		return false, 0
	}
	return traverseCorpus([]rune(corpus), []rune(substring), 0)
}

func traverseCorpus(corpus []rune, substring []rune, currIndexCorpus int) (match bool, startIndex int) {
	if currIndexCorpus == len(corpus) - 1 {
		return false, 0
	}
	if corpus[currIndexCorpus] == substring[0] {
		if checkSubstring(corpus, substring, currIndexCorpus, 1) {
			return true, currIndexCorpus
		}
	}
	return traverseCorpus(corpus, substring, currIndexCorpus + 1)
}

func checkSubstring(corpus []rune, substring []rune, anchorIndexCorpus, currIndexSubStr int) (match bool) {
	if currIndexSubStr == len(substring) {
		return true
	}
	if corpus[anchorIndexCorpus + currIndexSubStr] == substring[currIndexSubStr] {
		return checkSubstring(corpus, substring, anchorIndexCorpus, currIndexSubStr + 1)
	}
	return false
}
