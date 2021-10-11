module github.com/aws/aws-sdk-go-v2/service/rds

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v1.9.2
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.3.2
	github.com/aws/smithy-go v1.8.1-0.20211011213402-2f2385299b79
	github.com/jmespath/go-jmespath v0.4.0
)

replace github.com/aws/aws-sdk-go-v2 => ../../

replace github.com/aws/aws-sdk-go-v2/service/internal/presigned-url => ../../service/internal/presigned-url/
