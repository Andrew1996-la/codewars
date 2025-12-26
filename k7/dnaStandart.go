package k7

func DNAStrand(dna string) string {
	symbols := map[rune]rune{
		'A': 'T',
		'T': 'A',
		'G': 'C',
		'C': 'G',
	}

	res := ""

	for _, char := range dna {
		res += string(symbols[char])
	}
	return res
}
