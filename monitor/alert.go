package monitor

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

// interface contract check
var _ Alert = AlertLog()
var _ Alert = AlertSNS("")

// AlertLog will just log the message on stdout.
func AlertLog() Alert {

	return func(message interface{}) error {
		// test ignore
		if ignoreAlert(message) {
			return nil
		}

		// log actual message
		fmt.Println("===================== ALERT ===========================")
		fmt.Println(time.Now())
		fmt.Println(message)
		fmt.Println("=======================================================")
		return nil
	}
}

// AlertSNS sends the message to the specified
// AWS SNS notification topic.
func AlertSNS(snsTopic string) Alert {
	return func(message interface{}) (err error) {

		// test ignore
		if ignoreAlert(message) || len(snsTopic) == 0 {
			return nil
		}

		ses, err := session.NewSession(
			&aws.Config{
				Region: aws.String(endpoints.EuWest1RegionID),
			})
		if err != nil {
			fmt.Println(err)
			return err
		}

		svc := sns.New(ses)

		result, err := svc.Publish(&sns.PublishInput{
			Message:  aws.String(fmt.Sprint(message)),
			TopicArn: aws.String(snsTopic),
		})
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("Message envoyé :", *result.MessageId)

		return err
	}
}

// ignoreAlert tells us to ignore alert.
func ignoreAlert(message interface{}) bool {

	// ignore nil message
	if message == nil {
		// Do nothing
		return true
	}

	// ignore interface{} that prints as an empty string
	if s := fmt.Sprint(message); len(s) == 0 {
		fmt.Println("DEBUG : Empty string alert message ingnored ...")
		return true
	}
	return false
}

// TODO
// Other Alert should be able to send sms or emails.
