package search

// rabinKarp : implements Rabin-Karp substring match algorithm
func rabinKarp(substring string, corpus string) (lineNum int){
	return lineNum
}

func traverseCorpusRK(corpus []rune, substring []rune, currIndexCorpus int, candidateHash, targetHash uint64) (match bool, startIndex int) {
	return false, 0
}

// unicodePolyHash: implements a polynomial rolling hash function
// hash = sum from i = 0 to n - 1 of s[i] * p^i mod m
// for input string s of length n, prime number p roughly equal to cardinality of utf8
// and m, a prime number close to the max value of a uint64
func unicodePolyHash(input []rune) uint64 {
	return 0
}
