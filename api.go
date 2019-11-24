package static

import (
	"context"
	"fmt"
	"github.com/minio/minio-go"
	"io"
)

func (s *Storage) MakeURL(objectID string) string {
	return fmt.Sprintf(fmt.Sprintf("https://%s.%s/%s", s.bucketName, s.endpoint, objectID))
}

func (s *Storage) PutPublic(refID, objectID, contentType string, data io.Reader) (string, error) {
	opts := minio.PutObjectOptions{
		UserMetadata: map[string]string{
			"refID":     refID,
			"x-amz-acl": "public-read",
		},
		ContentType: contentType,
	}

	_, err := s.client.PutObjectWithContext(context.Background(), s.bucketName, objectID, data, -1, opts)
	if err != nil {
		return "", err
	}

	return s.MakeURL(objectID), nil
}

func (s *Storage) Exists(objectID string) (bool, error) {
	_, err := s.client.StatObject(s.bucketName, objectID, minio.StatObjectOptions{})
	if err != nil {
		if e, ok := err.(minio.ErrorResponse); ok {
			if e.Code == "NoSuchKey" {
				return false, nil
			}
		}
		return false, err
	}

	return true, nil
}
