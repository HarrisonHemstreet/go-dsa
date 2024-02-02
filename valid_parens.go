package main

func IsValidParens(s string) bool {
	if len(s) <= 1 {
		return false
	}

	parens := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}

	var stack []string

	for _, c := range s {
		paren, exists := parens[string(c)]
		if exists {
			stack = append(stack, paren)
		} else {
			if len(stack) > 0 {
				paren := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if paren != string(c) {
					return false
				}
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
