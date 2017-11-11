package twofer

// Outputs the sentence "One for X, one for me." where X is the name, or "you" if it is not provided

func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
