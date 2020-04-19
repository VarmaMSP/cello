package filestorage

import "github.com/varmamsp/cello/model"

const (
	BUCKET_THUMBNAILS        = "thumbnails"
	BUCKET_ASSETS            = "assets"
	BUCKET_PHENOPOD_CHARTS   = "phenopod-charts"
	BUCKET_PHENOPOD_DISCOVER = "phenopod-discover"
)

var BUCKET_POLICY_THUMBNAILS = model.NewAwsPolicy([]model.AwsPolicyStatement{
	{
		Sid:       "PublicReadAccess",
		Effect:    "Allow",
		Principal: map[string]string{"AWS": "*"},
		Action:    []string{"s3:GetObject"},
		Resource:  []string{"arn:aws:s3:::thumbnails/*"},
	},
})

var BUCKET_POLICY_ASSETS = model.NewAwsPolicy([]model.AwsPolicyStatement{
	{
		Sid:       "PublicReadAccess",
		Effect:    "Allow",
		Principal: map[string]string{"AWS": "*"},
		Action:    []string{"s3:GetObject"},
		Resource:  []string{"arn:aws:s3:::assets/*"},
	},
})
