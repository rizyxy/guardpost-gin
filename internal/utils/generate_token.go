package utils

func GenerateToken() (string, string, error) {
	// Generate Access Token
	accessToken, err := GenerateJWT()

	if err != nil {
		return "", "", err
	}

	// Generate Refresh Token
	refreshToken, err := GenerateJWT()

	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
