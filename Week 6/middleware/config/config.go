package config

func GetConfig() map[string]string {
	return map[string]string{
		"JWT_SECRET": "your_secret_key",
	}
}
