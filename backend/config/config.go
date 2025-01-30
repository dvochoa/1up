package config

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type Config struct {
	Port           string
	AllowedOrigins []string
	ListenAddress  string
}

// TODO: Understand this method better.
// TODO: Look online, are other people doing something similar?
// get the non-loopback local IP of the host
func getLocalIP() string {
	log.Print("inside of getLocalIP")
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				log.Printf("local IP address = %s", ipnet.IP.String())
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func LoadConfig() *Config {
	log.Print("Inside of LoadConfig")
	// Default to development settings
	config := &Config{
		Port: "8080",
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	switch env {
	case "dev":
		config.ListenAddress = "0.0.0.0" // TODO: Is this safe to do, even in development?
		localIP := getLocalIP()

		config.AllowedOrigins = []string{
			"http://localhost:3000", // TODO: Aren't these two redundant? Fails when I remove this tho hmm
			"http://127.0.0.1:3000",
		}

		// Add the LAN IP if available
		if localIP != "" {
			lanOrigin := fmt.Sprintf("http://%s:3000", localIP)
			config.AllowedOrigins = append(config.AllowedOrigins, lanOrigin)
		}
	case "prod":
		config.ListenAddress = "0.0.0.0" // TODO: I thought this wasn't safe, especially in prod
		if origins := os.Getenv("ALLOWED_ORIGINS"); origins != "" {
			config.AllowedOrigins = strings.Split(origins, ",")
		}
	}

	return config
}
