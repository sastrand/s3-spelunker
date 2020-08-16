package search

import "math"

// prime64: the largest prime number that will fit in a uint64
// $ echo '2^64-59' | bc
var prime64 uint64 = 18446744073709551557

// RabinKarpCLRS: a limited implementation of the Rabin-Karp matching algorithm
// as presented in Cormen Leiserson Rivest & Stein 3rd ed. page 993
func RabinKarpCLRS(text []uint64, pattern []uint64, alphabetSize uint64) (match bool, startIndex int){
	n := len(text)
	m := len(pattern)
	h := uint64(math.Floor(math.Pow(float64(alphabetSize), float64(m-1)))) % prime64
	var p uint64 = 0
	var t uint64 = 0
	// precompute the hash of the pattern and start of text
	for i := 0; i < m; i++ {
		p = (alphabetSize * p + pattern[i]) % prime64
		t = (alphabetSize * t + text[i]) % prime64
	}
	// check for success, continue hash roll
	for s := 0; s <= n-m; s++ {
		if p == t {
			if checkSubstringInts(text, pattern, s, 0) {
				return true, s
			}
		}
		if s < n-m {
			t = (alphabetSize * (t - (text[s] * h)) + text[s + m]) % prime64
		}
	}
	return false, 0
}

func checkSubstringInts(text []uint64, pattern []uint64, anchorIndexCorpus, currIndexSubStr int) (match bool) {
	if currIndexSubStr == len(pattern) {
		return true
	}
	if text[anchorIndexCorpus + currIndexSubStr] == pattern[currIndexSubStr] {
		return checkSubstringInts(text, pattern, anchorIndexCorpus, currIndexSubStr + 1)
	}
	return false
}
