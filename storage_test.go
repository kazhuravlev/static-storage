package static_test

import (
	"github.com/kazhuravlev/static-storage"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

var (
	cfg = static.Config{
		S3AccessKey:      os.Getenv("TEST_S3_ACCESS_KEY"),
		S3SecretKey:      os.Getenv("TEST_S3_SECRET_KEY"),
		S3BucketName:     os.Getenv("TEST_S3_BUCKET_NAME"),
		S3BucketLocation: os.Getenv("TEST_S3_BUCKET_LOCATION"),
		S3BucketSSL:      os.Getenv("TEST_S3_BUCKET_SSL") != "no",
	}
)

func TestStorage_MakeURL(t *testing.T) {
	table := []struct {
		name     string
		objectID string
		want     string
	}{
		{
			name:     "Test 1",
			objectID: "id-1",
			want:     "https://example-bucket.ams3.digitaloceanspaces.com/id-1",
		},
	}

	s, err := static.New(logrus.New().WithField("test", "test"), cfg)
	require.NoError(t, err)

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			got := s.MakeURL(tt.objectID)
			assert.Equal(t, got, tt.want)
		})
	}
}
