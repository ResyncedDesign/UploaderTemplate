package services

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	region = "apac" // It's worth noting that the region for Cloudflare's R2 *should* be automatic so apac is a valid option
)

type R2Service struct {
	S3Client *s3.Client
	Bucket   string
}

// Our functions are limited to only the ones we need for this project but you might need functions for listing files, getting files, etc.
// And I'm not commenting allat so I'll just ask chatgpt to generate some comments for me

// NewR2Service initializes a new instance of the R2Service struct, connecting to Cloudflare R2.
// It reads configuration from environment variables and validates their presence.
func NewR2Service() (*R2Service, error) {
	// Load necessary environment variables for Cloudflare R2 access.
	accessKey := os.Getenv("CLOUDFLARE_R2_ACCESS_KEY_ID")
	secretKey := os.Getenv("CLOUDFLARE_R2_SECRET_ACCESS_KEY")
	bucket := os.Getenv("CLOUDFLARE_R2_BUCKET")
	endpoint := os.Getenv("CLOUDFLARE_R2_ENDPOINT")

	// Check if all required environment variables are set.
	if accessKey == "" || secretKey == "" || bucket == "" || endpoint == "" {
		return nil, errors.New("missing one or more required environment variables for R2 service")
	}

	// Load the AWS SDK configuration, specifying the credentials and region.
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		// Provide static credentials using the access key and secret key.
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
		// Specify the region (assuming it's globally applicable for R2).
		config.WithRegion(region),
	)
	if err != nil {
		// Return an error if the SDK configuration fails to load.
		return nil, fmt.Errorf("failed to load AWS SDK config: %w", err)
	}

	// Initialize an S3 client with a custom endpoint pointing to Cloudflare R2.
	s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = &endpoint
	})

	// Return the initialized R2Service with the S3 client and bucket information.
	return &R2Service{
		S3Client: s3Client,
		Bucket:   bucket,
	}, nil
}

// UploadFile uploads a file to the specified bucket in Cloudflare R2.
// Parameters:
// - ctx: The context for managing the request lifecycle.
// - key: The destination key (path) for the file in the bucket.
// - file: The file content to upload, provided as an io.ReadSeeker.
// - size: The size of the file in bytes.
// - contentType: The MIME type of the file.
func (s *R2Service) UploadFile(ctx context.Context, key string, file io.ReadSeeker, size int64, contentType string) error {
	// Perform the file upload using the S3 client's PutObject method.
	_, err := s.S3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:        aws.String(s.Bucket),    // Specify the target bucket.
		Key:           aws.String(key),         // Specify the key (path) in the bucket.
		Body:          file,                    // Provide the file content.
		ContentLength: &size,                   // Set the content length.
		ContentType:   aws.String(contentType), // Set the file's content type.
	})
	if err != nil {
		// Return an error if the upload fails.
		return fmt.Errorf("failed to upload file %s: %w", key, err)
	}
	return nil
}

// DeleteFile deletes a file from the specified bucket in Cloudflare R2.
// Parameters:
// - ctx: The context for managing the request lifecycle.
// - key: The key (path) of the file to delete in the bucket.
func (s *R2Service) DeleteFile(ctx context.Context, key string) error {
	// Perform the file deletion using the S3 client's DeleteObject method.
	_, err := s.S3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.Bucket), // Specify the target bucket.
		Key:    aws.String(key),      // Specify the key (path) of the file to delete.
	})
	if err != nil {
		// Return an error if the deletion fails.
		return fmt.Errorf("failed to delete file %s: %w", key, err)
	}
	return nil
}
