package utils

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/spf13/viper"
	//go get -u github.com/aws/aws-sdk-go
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const (
	// Replace sender@example.com with your "From" address.
	// This address must be verified with Amazon SES.
	Sende = "murainoy@lifetrusty.com"

	// Replace recipient@example.com with a "To" address. If your account
	// is still in the sandbox, this address must be verified.
	Recipien = "enquiry@lifetrusty.com"

	// Specify a configuration set. To use a configuration
	// set, comment the next line and line 92.
	//ConfigurationSet = "ConfigSet"

	// The subject line for the email.
	Subjec = "lifetrusty"

	// The HTML body for the email.




	// The character encoding for the email.
	CharSe = "UTF-8"
)

func SendClientEmail(em,fm,ln,ph,cnt string) {
	// Create a new session in the us-west-2 region.
	// Replace us-west-2 with the AWS Region you're using for Amazon SES.


	HtmlBody := "<!DOCTYPE html>" +
			"<html lang='en'>" +
			"<head>" +
			"<meta charset='UTF-8'>" +
			"<title>LifeTrusty</title>" +
			"</head>" +
			"<body>" +
			"<style type='text/css' media='all'>" +
			"	span {" +
			"	font-weight: bold;" +
			"	padding-bottom: 4px;" +
			"	}" +
			"div{" +

			"	padding-bottom: 15px;" +
			"}" +
			"	</style>" +
			"	<header style='font-weight: bolder; margin: 30px; font-size: x-large; color: darkgreen'>lifetrusty Client Enquiry</header>" +
			"	<div> <span> Email :   </span> "+em +" </div>" +
			"	<div> <span> First Name :   </span> "+fm +"</div>" +
			"	<div> <span> Last Name :   </span> "+ln +"</div>" +
			"	<div> <span> Phone number :   </span> "+ph +"</div>" +
			"	<div> <span> Complains :   </span> "+cnt +"</div>" +
			"	</body>" +
			"	</html>"

	InitializeViper()
	region  := viper.GetString("email_region")

	accessKey  := viper.GetString("accessKey")
	accessSecret  := viper.GetString("accessSecret")
	//sess, err := session.NewSession(&aws.Config{
	//	Region:aws.String(region)},
	//)

	// Create an SES session.


	awsSession := session.New(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, accessSecret ,""),
	})

	svc := ses.New(awsSession)


	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{
			},
			ToAddresses: []*string{
				aws.String(Recipien),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSe),
					Data:    aws.String(HtmlBody),
				},

			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(Subject),
			},
		},
		Source: aws.String(Sender),
		// Uncomment to use a configuration set
		//ConfigurationSetName: aws.String(ConfigurationSet),
	}

	// Attempt to send the email.
	result, err := svc.SendEmail(input)

	// Display error messages if they occur.
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}

		return
	}

	fmt.Println("Email Sent to address: " + Recipient)


	fmt.Println(result)
}
