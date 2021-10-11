module github.com/aws/aws-sdk-go-v2/service/s3control

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v1.9.2
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.7.2
	github.com/aws/smithy-go v1.8.1-0.20211011213402-2f2385299b79
)

replace github.com/aws/aws-sdk-go-v2 => ../../

replace github.com/aws/aws-sdk-go-v2/service/internal/s3shared => ../../service/internal/s3shared/
