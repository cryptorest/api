package config

import (
	"os"
	"log"
	"strconv"
)

func envString(envName string, pDefault string) (string) {
	e := os.Getenv(envName)
	if e == "" {
		e = pDefault
	}

	return e
}

func envBool(envName string, pDefault bool) bool {
	e := os.Getenv(envName)
	if e == "" {
		return pDefault
	}

	b, err := strconv.ParseBool(e)
	if err != nil {
		log.Fatalf("Variable %s: %v", envName, err)
	}

	return b
}

func envInt(envName string, pDefault int) int {
	e := os.Getenv(envName)
	if e == "" {
		return pDefault
	}

	i, err := strconv.ParseInt(e, 10, 16)
	if err != nil {
		log.Fatalf("Variable %s: %v", envName, err)
	}

	return int(i)
}

func InitEnvironment(c *Configuration) {
	c.CertFile = envString("CRYPTOREST_CERT_FILE", Default.CertFile)
	c.KeyFile  = envString("CRYPTOREST_KEY_FILE",  Default.KeyFile)
	c.Host     = envString("CRYPTOREST_HOST",      Default.Host)
	c.Port     = envInt("CRYPTOREST_PORT",      Default.Port)
	c.Verbose  = envBool("CRYPTOREST_VERBOSE", Default.Verbose)
}
