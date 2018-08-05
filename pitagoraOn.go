package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// type MyEvent struct {
// 	Name serialNumber   `json:"serialNumber"`
// 	Name clickType      `json:"clickType"`
// 	Name batteryVoltage `json:batteryVoltage`
// }

type MyResponse struct {
	Message string `json:"Answer:"`
}

//event MyEvent
func pitagoraOn() (MyResponse, error) {
	sess, sessionError := session.NewSession()
	if sessionError != nil {
		panic(sessionError)
	}

	service := ssm.New(sess)

	param := ssm.GetParameterInput{
		Name:           aws.String("pitagora-switch-ifttt-key"),
		WithDecryption: aws.Bool(true),
	}

	keyOutput, keyError := service.GetParameter(&param)
	if keyError != nil {
		panic(keyError)
	}

	url := fmt.Sprintf(
		"https://maker.ifttt.com/trigger/Pi-ta-go-ra-Switch/with/key/%s",
		*keyOutput.Parameter.Value)

	fmt.Println(url)

	return MyResponse{Message: fmt.Sprintf("Hello %s", "shinsuke")}, nil
}

func main() {

	lambda.Start(pitagoraOn)
}
