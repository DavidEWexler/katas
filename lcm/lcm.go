package lcm

func LeastCommonMultiple(n1, n2 int) int {
	min := n1
	max := n2
	if min > max {
		min = n2
		max = n1
	}

	for test := max; test < max*min; test += max {
		if test%min == 0 {
			return test
		}
	}

	return max * min
}
