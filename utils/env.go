package utils

import (
	"log"
	"os"
	"strconv"
)

func MustGetEnv(env string) string {
	v := os.Getenv(env)
	if v == "" {
		log.Panicf("%s missing", env)
	}
	return v
}

func MustGetBool(env string) bool {
	v := os.Getenv(env)
	if v == "" {
		log.Panicf("%s env missing", env)
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		log.Panicf("Error '%q' while parsing env variable %s", err, env)
	}
	return b
}
