package utils

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/spf13/viper"

	"github.com/gobuffalo/packr"
	//go get -u github.com/aws/aws-sdk-go
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const (
	// Replace sender@example.com with your "From" address.
	// This address must be verified with Amazon SES.
	Sender = "lifetrusty@lifetrusty.com"

	// Replace recipient@example.com with a "To" address. If your account
	// is still in the sandbox, this address must be verified.
	Recipient = "enquiry@lifetrusty.com"

	// Specify a configuration set. To use a configuration
	// set, comment the next line and line 92.
	//ConfigurationSet = "ConfigSet"

	// The subject line for the email.
	Subject = "LifeTrusty"

	// The HTML body for the email.




	// The character encoding for the email.
	CharSet = "UTF-8"
)

func SendEmail(em,fm,ln,ph,cnt string) {
	//Create a new session in the us-west-2 region.
	//Replace us-west-2 with the AWS Region you're using for Amazon SES.
	//
	//



	str :=  "<span>Email :  </span><strong>"+em +" </strong> <br/>" +
		"<span>First Name :   </span><strong>"+ fm +" </strong> <br/>" +
		"<span>Last Name  :   </span><strong>"+ ln +" </strong> <br/>" +
		"<span>Last Name  :   </span><strong>"+ ph +" </strong> <br/>" +
		"<span>Complains  :  </span><strong>"+ cnt +" </strong> <br/>"


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
				aws.String(Recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(str),
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

type Person struct {
	Name   string
	Age    int

}

func Send(email string) {


	InitializeViper()
	region  := viper.GetString("email_region")

	accessKey  := viper.GetString("accessKey")
	accessSecret  := viper.GetString("accessSecret")
	//sess, err := session.NewSession(&aws.Config{
	//	Region:aws.String(region)},
	//)

	//// Create an SES session.
	box := packr.NewBox("./")

	s, err := box.FindString("temp.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(s)





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
				aws.String(email),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(s),
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

	fmt.Println("Email Sent to address: " + email)


	fmt.Println(result)
}

