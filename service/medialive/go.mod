module github.com/aws/aws-sdk-go-v2/service/medialive

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v1.17.7
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.31
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.25
	github.com/aws/smithy-go v1.13.5
	github.com/jmespath/go-jmespath v0.4.0
)

replace github.com/aws/aws-sdk-go-v2 => ../../

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => ../../internal/endpoints/v2/
