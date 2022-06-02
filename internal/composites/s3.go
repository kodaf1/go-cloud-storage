package composites

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Composite struct {
	instance *s3.S3
	bucket   string
}

func NewS3Composite(accessKey, secretKey, sessionToken, endpointURL, apiRegion, bucket string) *S3Composite {
	defaultResolver := endpoints.DefaultResolver()
	s3CustomResolverFunc := func(
		service,
		region string,
		optFns ...func(*endpoints.Options),
	) (endpoints.ResolvedEndpoint, error) {
		if service == "s3" {
			return endpoints.ResolvedEndpoint{
				URL:           endpointURL,
				SigningRegion: apiRegion,
			}, nil
		}

		return defaultResolver.EndpointFor(service, region, optFns...)
	}

	sess := session.Must(session.NewSession(&aws.Config{
		EndpointResolver: endpoints.ResolverFunc(s3CustomResolverFunc),
		Region:           aws.String(apiRegion),
		Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, sessionToken),
	}))
	return &S3Composite{s3.New(sess), bucket}
}
