package helpers

import (
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	exc "github.com/kangman53/project-sprint-halo-suster/exceptions"
	"github.com/spf13/viper"
)

func FileUpload(file *multipart.FileHeader, fileName string) (string, error) {
	// Get Buffer from file
	buffer, err := file.Open()
	if err != nil {
		return "", exc.InternalServerException("Processing File Failed")
	}
	defer buffer.Close()

	endpointUrl := viper.GetString("AWS_S3_BUCKET_URL")
	if endpointUrl == "" {
		endpointUrl = "s3.ap-southeast-1.amazonaws.com"
	}

	// CHECK AWS FILENAME
	// fmt.Println(fileName)
	// Configure to use MinIO Server
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(viper.GetString("AWS_ACCES_KEY_ID"), viper.GetString("AWS_SECRET_ACCESS_KEY"), ""),
		Endpoint:         aws.String(endpointUrl),
		Region:           aws.String(viper.GetString("AWS_REGION")),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}

	newSession, err := session.NewSession(s3Config)
	if err != nil {
		return "", exc.InternalServerException(err.Error())
	}

	s3Client := s3.New(newSession)
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Body:   buffer,
		Bucket: aws.String(viper.GetString("AWS_S3_BUCKET_NAME")),
		ACL:    aws.String("public-read"),
		Key:    aws.String(fileName),
	})
	if err != nil {
		return "", exc.InternalServerException(err.Error())
	}

	// CHECK AWS RESULT
	// output, err := s3Client.ListObjectsV2(&s3.ListObjectsV2Input{
	// 	Bucket: aws.String(viper.GetString("AWS_S3_BUCKET_NAME")),
	// })

	// for _, object := range output.Contents {
	// 	if fileName == aws.StringValue(object.Key) {
	// 		fmt.Printf("key=%s size=%d\n", aws.StringValue(object.Key), object.Size)
	// 	}
	// }
	// if err != nil {
	// 	return exc.InternalServerException(err.Error())
	// }
	url := fmt.Sprintf("https://%s.%s/%s", viper.GetString("AWS_S3_BUCKET_NAME"), endpointUrl, fileName)
	return url, nil
}
