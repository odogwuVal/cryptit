// decrypt package consists of all the decryption algorithms
package decrypt

// decrypts by reducing the ascii code by 3 from each character
func Nimbus(str string) string {
	decryptStr := ""
	for _, c := range str {
		asciiCode := int32(c)
		character := string(asciiCode - 3)
		decryptStr += character
	}
	return decryptStr
}
