package main

// TODO: Create a testing file to test this function!
func toTest(str string, num int) int{
	if str == "" {
		return 0
	}
	switch str {
	case "distributed":
		return 1380
	case "golang":
		if num < 10 {
			return 5
		} else {
			return 1
		}
	default:
		return 10
	}
}