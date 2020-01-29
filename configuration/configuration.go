package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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
	fmt.Println("Attempting to load configuration from aws S3")
	cb, err = readFromS3(fn)

	if err != nil {
		// S3 download failed, so let's try local files ?
		fmt.Println("Attempt to download configuration from AWS-S3 failed !")
		fmt.Println("Attempting to load configuration from local file : ")

		for _, f := range where {
			fmt.Println("\t", f)
			cb, err = ioutil.ReadFile(f)
			if err == nil {
				fmt.Println("Found usable configuration file : ", f)
				break
			}
		}

		if err != nil {
			fmt.Println("============================================================================================")
			fmt.Println("Make sure you create a configuration named mystock.json in a reasonnable location (see above).")
			fmt.Println("You may copy/rename the file mystock_example.json and modify its content.")
			fmt.Println("============================================================================================")
			panic("unable to find the configuration file ! ")
		}

	} else {
		fmt.Println("Successfully loaded configuration from aws s3 object : ", fn)
	}

	// Parse config from cb []bytes
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

// readFromS3 attempts getting a []byte from s3 object.
func readFromS3(fileName string) ([]byte, error) {

	svc := s3.New(session.New(&aws.Config{
		Region: aws.String(endpoints.EuWest1RegionID),
	}))
	input := &s3.GetObjectInput{
		Bucket: aws.String("xavier268.gandillot.com"),
		Key:    aws.String(fileName),
	}

	result, err := svc.GetObject(input)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer result.Body.Close()
	cb, err := ioutil.ReadAll(result.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return cb, nil
}
