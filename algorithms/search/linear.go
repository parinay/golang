package main

func linear1(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}
func linear2(a []string, x string) bool {

	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
