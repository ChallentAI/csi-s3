package s3_test

import (
	"os"
	"testing"

	"github.com/ctrox/csi-s3/pkg/s3"
)

func TestXxx(t *testing.T) {
	os.Setenv("AWS_DEFAULT_REGION", "us-west-2")
	os.Setenv("AWS_REGION", "us-west-2")
	os.Setenv("AWS_WEB_IDENTITY_TOKEN_FILE", "/Users/aidan/Development/csi-s3/token")
	os.Setenv("AWS_ROLE_ARN", "arn:aws:iam::693848528138:role/csi-s3-sandbox")
	c, err := s3.NewClient(
		&s3.Config{
			AccessKeyID:     "AKIAIOSFODNN7EXAMPLE",
			SecretAccessKey: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
			Region:          "us-west-2",
			Endpoint:        "https://s3.us-west-2.amazonaws.com",
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	_, err = c.BucketExists("aidan-delete-me")
	if err != nil {
		t.Fatal(err)
	}

}
