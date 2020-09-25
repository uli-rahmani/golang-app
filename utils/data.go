package utils

func gettokenOrgs() map[string]string {
	orgsTokenStringMap := make(map[string]string)
	orgsTokenStringMap["save-comment"] = "9ab295deccf9bd67d9ee4b70d63e0a27e907fdeb"
	orgsTokenStringMap["get-comment"] = "1d11c5dfea82e591ae045eed4380372a2edb8e64"
	orgsTokenStringMap["delete-comment"] = "d64c0d7da8e1639ea7a64e7bdfe398c286fde3b1"

	orgsTokenStringMap["get-member"] = "eb922e4b26a175a9b40a276760a8f7119f20ede5"

	return orgsTokenStringMap
}
