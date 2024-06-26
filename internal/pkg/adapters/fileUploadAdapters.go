package adapters

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gofiber/fiber/v2"
	// "github.com/rs/zerolog/log"
)

// func init() {
// 	var err error
// 	cfg, err = config.Load()
// 	if err != nil {
// 		log.Println("error loading config for file upload adapter ::: ", zap.Error(err))
// 	}
// }

type IFileUploader interface {
	UploadFile(fileName string, c *fiber.Ctx) (*string, error)
	UploadFileLocally(file *multipart.FileHeader, c *fiber.Ctx, newFileName string) (*string, error)
	UploadFileToProduction(newFileName string, fileData multipart.File) (*string, error)
}

type FileUploader struct {
}

func NewFileUploadAdapter() IFileUploader {
	return &FileUploader{}
}

// UploadFile handles file upload and returns the file URL
func (f *FileUploader) UploadFile(fileName string, c *fiber.Ctx) (*string, error) {
	/*
		This is used to upload a file to an S3 bucket or locally to our personal computer.
	*/
	file, err := c.FormFile(fileName)
	if err != nil {
		return nil, err
	}

	// Open the file
	fileData, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer fileData.Close()

	// Generate a unique filename by appending a timestamp to the original filename
	timestamp := time.Now().Unix()
	fileExt := filepath.Ext(file.Filename)
	newFileName := fmt.Sprintf("%s_%d%s", strings.TrimSuffix(file.Filename, fileExt), timestamp, fileExt)

	if strings.ToLower(cfg.LearniumStorageType) != "local" {
		return f.UploadFileToProduction(newFileName, fileData)
	} else {
		return f.UploadFileLocally(file, c, newFileName)
	}
}

func (f *FileUploader) UploadFileToProduction(newFileName string, fileData multipart.File) (*string, error) {
	/*
		This is used to upload the file to an s3 bucket
	*/

	// cfg, err := config.Load()
	// if err != nil {
	// 	return nil, err
	// }
	// Use S3 bucket
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(cfg.S3BucketRegion),
		Credentials: credentials.NewStaticCredentials(cfg.S3AccessKeyId, cfg.S3AccessKeySecret, ""),
	}))

	uploader := s3manager.NewUploader(sess)

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(cfg.S3BucketName),
		Key:    aws.String(newFileName), // Use the new filename
		Body:   fileData,
	})
	if err != nil {
		return nil, err
	}
	// Return the S3 file URL
	return &result.Location, nil

}

func (f *FileUploader) UploadFileLocally(file *multipart.FileHeader, c *fiber.Ctx, newFileName string) (*string, error) {
	/*
		This is used to upload a file to the local filesystem
	*/
	// cfg, err := config.Load()
	// if err != nil {
	// 	return nil, err
	// }

	dirPath := "uploads"

	// Use os.Stat to check if the directory exists
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		// Define the desired permissions (read and write for all users)
		// 0777 allows reading and write permissions for an owner, group, and others
		// Use os.ModeDir to indicate that it's a directory
		permissions := os.FileMode(0777) | os.ModeDir

		// Update the directory with the specified permissions
		if err := os.MkdirAll(dirPath, permissions); err != nil {
			return nil, errors.New(fmt.Sprintf("Error creating directory: %v", err))
		}
	}

	err = c.SaveFile(file, "./"+cfg.LearniumUploadDirectory+"/"+newFileName) // Use the new filename
	if err != nil {
		return nil, err
	}

	// Return the local file URL
	fileUrl := c.BaseURL() + "/" + cfg.LearniumUploadDirectory + "/" + newFileName
	return &fileUrl, nil
}
