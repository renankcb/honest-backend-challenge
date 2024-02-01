package loanevaluation

func IsAllowedAreaCode(phoneNumber string) bool {
	allowedAreaCodes := []string{"0", "2", "5", "8"}
	firstDigit := string(phoneNumber[0])
	return contains(allowedAreaCodes, firstDigit)
}

func contains(arr []string, target string) bool {
	for _, item := range arr {
		if item == target {
			return true
		}
	}
	return false
}
