package main

var testCases = []struct {
	description string
	input       string
	expected    string
}{
	{
		description: "no bracket",
		input:       "hello",
		expected:    "",
	},
	{
		description: "only open bracket",
		input:       "(hello",
		expected:    "",
	},
	{
		description: "only close bracket",
		input:       "test)",
		expected:    "",
	},
	{
		description: "one word in bracket",
		input:       "(hello)",
		expected:    "hello",
	},
	{
		description: "two words in different bracket",
		input:       "(hello)(world)",
		expected:    "hello",
	},
	{
		description: "have many open bracket",
		input:       "(((hello",
		expected:    "",
	},
	{
		description: "only bracket",
		input:       "()",
		expected:    "",
	},
	{
		description: "have many close bracket",
		input:       "hello))))",
		expected:    "",
	},
	{
		description: "have many bracket",
		input:       "(((hello))))",
		expected:    "hello",
	},
	{
		description: "have many words",
		input:       "(hello world apple)",
		expected:    "hello",
	},
	{
		description: "have space around",
		input:       "( hello )",
		expected:    "hello",
	},
	{
		description: "have many words",
		input:       "( hello world apple )",
		expected:    "hello",
	},
	{
		description: "have many words",
		input:       "(   hello   )",
		expected:    "hello",
	},
}
