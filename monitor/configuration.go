package monitor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// configuration as read from configuration file.
type configuration struct {
	Lines        []Line
	APISecretKey string
	SNSTopic     string
}

// loadConfiguration from the mystock.json file.
func loadConfiguration() configuration {

	var err error

	// where defines the possible location to look for
	// the configuration file, in that order.
	// Only the first file found will be processed.
	fn := "mystock.json"

	where := []string{
		"secret_test.json", // for testing
		"secret.json",      // for testing
		fn,
		path.Join(".", fn),
		path.Join("..", fn),
		path.Join("..", "..", fn),
		path.Join("..", "..", "..", fn),
		path.Join("..", "conf", fn),
		path.Join("..", "..", "conf", fn)}

	user, okuser := os.LookupEnv("USER")
	if okuser {
		where = append(where,
			path.Join(user, fn),
			path.Join(user, "conf", fn),
			path.Join(user, ".mystock", fn),
			path.Join(user, "mystock", fn),
		)
	}

	var cb []byte
	fmt.Println("Attempting to load configuration file from : ")

	for _, f := range where {
		fmt.Println("\t", f)
		cb, err = ioutil.ReadFile(f)
		if err == nil {
			fmt.Println("Found usable configuration file : ", f)
			break
		}
	}

	if err != nil {
		panic("unable to find the configuration file ! ")
	}

	conf := new(configuration)
	err = json.Unmarshal(cb, conf)
	if err != nil {
		fmt.Println(err, conf)
		panic("Could not parse configuration")
	}
	fmt.Println("Configuration successfully parsed")
	return *conf

}
