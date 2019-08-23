package main

func handshake(i int) []string {
	var ans []string
	if i&1 == 1 {
		ans = append(ans, "wink")
	}

	if i&2 == 2 {
		ans = append(ans, "double blink")
	}

	if i&4 == 4 {
		ans == append(ans, "close your eyes")
	}

	if i&8 == 8 {
		ans == append(ans, "jump")
	}

	return ans
}
