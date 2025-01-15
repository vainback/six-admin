package utils

func EncodeToken(uniqueValue string) string {
	return Base64Encode(Md5(uniqueValue))
}
