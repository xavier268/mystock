package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// Conf as read from configuration file.
type Conf struct {
	Lines        []Line
	APISecretKey string
	SNSTopic     string
}

// Load from the mystock.json configuration file.
func Load() Conf {

	var err error

	// where defines the possible location to look for
	// the configuration file, in that order.
	// Only the first file found will be processed.
	fn := "mystock.json"

	where := []string{
		"secret_test.json",
		path.Join("configuration", "secret_test.json"),                   // for testing
		path.Join("..", "configuration", "secret_test.json"),             // for testing
		path.Join("..", "..", "configuration", "secret_test.json"),       // for testing
		path.Join("..", "..", "..", "configuration", "secret_test.json"), // for testing

		fn,
		path.Join(".", fn),
		path.Join("..", fn),
		path.Join("..", "..", fn),
		path.Join("configuration", fn),
		path.Join("..", "configuration", fn),
		path.Join("..", "..", "configuration", fn),
		path.Join("..", "..", "..", "configuration", fn)}

	user, okuser := os.LookupEnv("USER")
	if okuser {
		where = append(where,
			path.Join(user, fn),
			path.Join(user, ".mystock", fn),
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

	conf := new(Conf)
	err = json.Unmarshal(cb, conf)
	if err != nil {
		fmt.Println(err, conf)
		panic("Could not parse configuration")
	}
	fmt.Println("Configuration successfully parsed")
	return *conf

}

// Line defines a line of the shares portfolio
// as they are described in the mystock.json file
type Line struct {
	// Informative, human readable name
	Name string
	// Unique ticker identifying the  stock shares
	Ticker string
	// Date of purchase - use special type to allow easier unmarshalling
	Date MyTime
	// Historical/average purchase price
	Price float64
	// number of shares
	Volume float64
}

// Tickers returns the list of Tickers
// mentionned in the configuration.
// This is mostly used to cache them initially.
func (c *Conf) Tickers() []string {
	tl := make([]string, 0)
	m := make(map[string]bool)

	for _, l := range c.Lines {
		if _, ok := m[l.Ticker]; !ok {
			tl = append(tl, l.Ticker)
			m[l.Ticker] = true
		}
	}
	return tl
}
