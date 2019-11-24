package static

type Config struct {
	S3AccessKey      string `toml:"s3_access_key"`
	S3SecretKey      string `toml:"s3_secret_key"`
	S3BucketName     string `toml:"s3_bucket_name"`
	S3BucketLocation string `toml:"s3_bucket_location"`
	S3BucketSSL      bool   `toml:"s3_bucket_ssl"`
}
