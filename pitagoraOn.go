package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"

	"net/http"
)

type MyEvent struct {
	SerialNumber   string `json:"serialNumber"`
	ClickType      string `json:"clickType"`
	BatteryVoltage string `json:"batteryVoltage"`
}

type MyResponse struct {
	Message string `json:"Answer:"`
}

func pitagoraOn(event MyEvent) (MyResponse, error) {
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

	values := fmt.Sprintf(
		"{Value1: serialNumber is %s, Value2: clickType is %s, Value3: batteryVoltage is %s}",
		event.SerialNumber,
		event.ClickType,
		event.BatteryVoltage)

	fmt.Println(url)
	fmt.Println(values)

	http.Post(url, "text/json", strings.NewReader(values))

	return MyResponse{Message: fmt.Sprintf("Pitagora Switch kicked!!!\n%s", values)}, nil
}

func main() {
	lambda.Start(pitagoraOn)
}
