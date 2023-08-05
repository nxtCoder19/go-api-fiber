package configs

import (
	"os"
)

func EnvMongoUri() string{

	os.Setenv("MONGODB_URI", "mongodb+srv://golang-api:golang-api@cluster0.uek6tzj.mongodb.net/?retryWrites=true&w=majority")
	MONGODB_URI := os.Getenv("MONGODB_URI")

	return MONGODB_URI


    // getEnv := func(key string) {
    //     val, ok := os.LookupEnv(key)
    //     if !ok {
    //         fmt.Printf("%s not set\n", key)
    //     } else {
    //         fmt.Printf("%s=%s\n", key, val)
    //     }
    // }
	// getEnv("MONGODB_URI")
}