package static

import (
	"fmt"
	"github.com/minio/minio-go"
	"github.com/sirupsen/logrus"
)

type Storage struct {
	log            *logrus.Entry
	accessKey      string
	secretKey      string
	bucketName     string
	bucketLocation string
	bucketSSL      bool
	client         *minio.Client
	endpoint       string
}

func New(log *logrus.Entry, cfg Config) (*Storage, error) {
	endpoint := fmt.Sprintf("%s.digitaloceanspaces.com", cfg.S3BucketLocation)
	client, err := minio.New(endpoint, cfg.S3AccessKey, cfg.S3SecretKey, cfg.S3BucketSSL)
	if err != nil {
		return nil, err
	}

	exists, err := client.BucketExists(cfg.S3BucketName)
	if err != nil {
		return nil, err
	}

	if !exists {
		if err = client.MakeBucket(cfg.S3BucketName, cfg.S3BucketLocation); err != nil {
			return nil, err
		}
	}

	return &Storage{
		log:            log,
		accessKey:      cfg.S3AccessKey,
		secretKey:      cfg.S3SecretKey,
		bucketName:     cfg.S3BucketName,
		bucketLocation: cfg.S3BucketLocation,
		bucketSSL:      cfg.S3BucketSSL,
		client:         client,
		endpoint:       endpoint,
	}, nil
}
