package utils

// Generic function to handle errors
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
