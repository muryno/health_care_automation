package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/dgrijalva/jwt-go"
	"github.com/rebuy-de/aws-nuke/pkg/util"
	"regexp"

	"github.com/globalsign/mgo/bson"
	"github.com/spf13/viper"
	"net/http/httputil"
	"strings"

	"lifetrusty-brain/model"
	"math/rand"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"time"
)

var UserId = 0


func Message(status bool, message string) map[string]interface{} {

	return map[string]interface{}{"status.yml": status, "message": message}

}

func Responds(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	_ = json.NewEncoder(w).Encode(data)
}





func InitializeViper()  {
	// Set the file name of the configurations file
	viper.SetConfigName("configs")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading configs file, %s", err)
	}
}


func GenerateAuthToken(id uint) string {
	tk := &model.Token{UserId: id,}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(viper.GetString("token_password")))
	return "Bearer" + tokenString
}

func GenerateRandomPassword()string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	specials := "~=+%^*/()[]{}/!@#$?|"
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		digits + specials
	length := 8
	buf := make([]byte, length)
	buf[0] = digits[rand.Intn(len(digits))]
	buf[1] = specials[rand.Intn(len(specials))]
	for i := 2; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	str := string(buf) // E.g. "3i[g0|)z"
	return  str
}

func GetFileByte(fil multipart.File, fileHeader *multipart.FileHeader) []byte {
	size := fileHeader.Size
	buffer := make([]byte, size)
	fil.Read(buffer)
	return buffer
}

func GetTemp(fileHeader *multipart.FileHeader) string {
	tempFileName := "post/" + bson.NewObjectId().Hex() +  filepath.Ext(fileHeader.Filename)
	var myKey =tempFileName
	return myKey
}



func UploadFileToS3(buffer []byte, key string) (string, error) {



	InitializeViper()

	myBucket := viper.GetString("bucketName")
	accessKey := viper.GetString("accessKey")
	accessSecret  := viper.GetString("accessSecret")
	region  := viper.GetString("region")



	var awsConfig *aws.Config

	awsConfig = &aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, accessSecret, ""),
	}


	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession(awsConfig))

	// Create an uploader with the session and default options
	//uploader := s3manager.NewUploader(sess)

	// Create an uploader with the session and custom options
	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
		u.PartSize = 5 * 1024 * 1024 // The minimum/default allowed part size is 5MB
		u.Concurrency = 2           // default is 5
	})




	// Upload the file to S3.
	upParams := &s3manager.UploadInput{
		Bucket: aws.String(myBucket),
		Key:    aws.String(key),
		Body:    bytes.NewReader(buffer),
		ContentType:        aws.String(http.DetectContentType(buffer)),
		ACL:                  aws.String("public-read"),
		Metadata: map[string]*string{ "Key": aws.String("MetadataValue"),},
	}

	result, err := uploader.Upload(upParams, func(u *s3manager.Uploader) {
		u.PartSize = 10 * 1024 * 1024 // 10MB part size
		u.LeavePartsOnError = true    // Dont delete the parts if the upload fails.
	})

	//in case it fails to upload
	if err != nil {
		fmt.Printf("failed to upload file, %v", err)

	}
	fmt.Printf("file uploaded to, %s\n", result.Location)
	res:= result.Location
	return res, err
}


func DeleteFileS3(myKey string)   (map[string]interface{}, bool)  {


	myBucket := viper.GetString("bucketName")
	accessKey := viper.GetString("accessKey")
	accessSecret  := viper.GetString("accessSecret")
	region  := viper.GetString("region")

	res2 := strings.TrimLeft(myKey, "http://api-lf.eu-west-2.elasticbeanstalk.com/")
	fmt.Print(res2)

	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:      aws.String(region),
			Credentials: credentials.NewStaticCredentials(accessKey, accessSecret, ""),
			DisableRestProtocolURICleaning:aws.Bool(true),
		},
		SharedConfigState: session.SharedConfigEnable,
		Profile:           "default",

	})
	if err != nil {
		panic(err)
	}

	sess.Handlers.Send.PushFront(func(r *request.Request) {
		fmt.Printf("sending AWS request:\n%s\n", DumpRequest(r.HTTPRequest))
	})

	svc := s3.New(sess)
	params := &s3.DeleteObjectInput{
		Bucket: aws.String(myBucket),
		Key:   aws.String(res2),
	}
	fmt.Println(params)
	_, err = svc.DeleteObject(params)
	if err != nil {
		fmt.Printf("Error deleting image .. please try again! , %v", err)
		return Message(false, "Error fetching category. Please retry"),false
	}

	fmt.Printf("deleted successfully")

	return Message(false, "Requirement passed"),true
}




func DumpRequest(r *http.Request) string {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		panic(err)
	}

	dump = bytes.TrimSpace(dump)
	dump = HideSecureHeaders(dump)
	dump = util.IndentBytes(dump, []byte("    > "))
	return string(dump)
}


func HideSecureHeaders(dump []byte) []byte {
	return RESecretHeader.ReplaceAll(dump, []byte("$1: <hidden>"))
}

var (
	RESecretHeader = regexp.MustCompile(`(?m:^([^:]*(Auth|Security)[^:]*):.*$)`)
)