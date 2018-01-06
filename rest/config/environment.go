package config

import (
	"os"
	"log"
	"strconv"
)

var envVars = [5]string {
	"CRYPTOREST_CERT_FILE",
	"CRYPTOREST_KEY_FILE",
	"CRYPTOREST_HOST",
	"CRYPTOREST_PORT",
	"CRYPTOREST_VERBOSE",
}

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
	c.CertFile = envString(envVars[0], Default.CertFile)
	c.KeyFile  = envString(envVars[1], Default.KeyFile)
	c.Host     = envString(envVars[2], Default.Host)
	c.Port     = envInt(envVars[3],    Default.Port)
	c.Verbose  = envBool(envVars[4],   Default.Verbose)
}
