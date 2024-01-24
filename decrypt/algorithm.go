package decrypt

func Nimbus(str string) string {
	decryptStr := ""
	for _, c := range str {
		asciiCode := int32(c)
		character := string(asciiCode - 3)
		decryptStr += character
	}
	return decryptStr
}
