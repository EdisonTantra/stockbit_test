package anagram

var testCases = []struct {
	description string
	candidates  []string
	expected    [][]string
}{
	{
		description: "no matches",
		candidates: []string{
			"hello",
			"world",
			"apple",
			"orange"},
		expected: [][]string{
			{"apple"},
			{"hello"},
			{"orange"},
			{"world"},
		},
	},
	{
		description: "detects one pair anagram",
		candidates: []string{
			"aku",
			"kua"},
		expected: [][]string{
			{"aku", "kua"},
		},
	},
	{
		description: "detects two anagrams",
		candidates: []string{
			"kita",
			"atik",
			"tika",
			"aku",
			"kua"},
		expected: [][]string{
			{"kita", "atik", "tika"},
			{"aku", "kua"},
		},
	},
	{
		description: "pdf test",
		candidates: []string{
			"kita",
			"atik",
			"tika",
			"aku",
			"kia",
			"makan",
			"kua"},
		expected: [][]string{
			{"kita", "atik", "tika"},
			{"aku", "kua"},
			{"kia"},
			{"makan"},
		},
	},
	{
		description: "typo double chars",
		candidates: []string{
			"lasttt",
			"last"},
		expected: [][]string{
			{"last"},
			{"lasttt"},
		},
	},
	{
		description: "detects anagrams case-insensitively",
		candidates: []string{
			"aku",
			"Kua",
			"uKa",
			"kamu"},
		expected: [][]string{
			{"aku", "Kua", "uKa"},
			{"kamu"},
		},
	},
	{
		description: "detects anagrams with space",
		candidates: []string{
			"Tom Marvolo Riddle",
			"I am Lord Voldemort",
			"you know who",
		},
		expected: [][]string{
			{"Tom Marvolo Riddle", "I am Lord Voldemort"},
			{"you know who"},
		},
	},
	{
		description: "anagrams must use all letters exactly once",
		candidates: []string{
			"tapper",
			"patter"},
		expected: [][]string{
			{"patter"},
			{"tapper"},
		},
	},
	{
		description: "words are not anagrams of themselves (case-insensitive)",
		candidates: []string{
			"BANANA",
			"Banana",
			"banana"},
		expected: [][]string{
			{"BANANA"},
		},
	},
}
