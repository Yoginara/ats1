package config

var allowedOrigins = []string{
	"https://Yoginara.github.io/",
}

// GetAllowedOrigins returns the list of allowed origins
func GetAllowedOrigins() []string {
	return allowedOrigins
}
