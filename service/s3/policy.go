package s3

import "encoding/json"

// AWS POLICY GENERATOR
// https://awspolicygen.s3.amazonaws.com/policygen.html

type AwsPolicy struct {
	Id        string               `json:"Id"`
	Version   string               `json:"Version"`
	Statement []AwsPolicyStatement `json:"Statement"`
}

type AwsPolicyStatement struct {
	Sid       string            `json:"Sid"`
	Effect    string            `json:"Effect"`
	Principal map[string]string `json:"Principal"`
	Action    []string          `json:"Action"`
	Resource  []string          `json:"Resource"`
}

const (
	POLICY_VERSION       = "2012-10-17"
	POLICY_ID_THUMBNAILS = "2-12-2019"
	POLICY_ID_ASSETS     = "2-12-2019"
)

func getNewBucketPolicy(bucketName string) AwsPolicy {
	switch bucketName {
	case BUCKET_NAME_THUMBNAILS:
		return AwsPolicy{
			Id:      POLICY_ID_THUMBNAILS,
			Version: POLICY_VERSION,
			Statement: []AwsPolicyStatement{
				AwsPolicyStatement{
					Sid:       "PublicReadAccess",
					Effect:    "Allow",
					Principal: map[string]string{"AWS": "*"},
					Action:    []string{"s3:GetObject"},
					Resource:  []string{"arn:aws:s3:::thumbnails/*"},
				},
			},
		}

	case BUCKET_NAME_ASSETS:
		return AwsPolicy{
			Id:      POLICY_ID_ASSETS,
			Version: POLICY_VERSION,
			Statement: []AwsPolicyStatement{
				AwsPolicyStatement{
					Sid:       "PublicReadAccess",
					Effect:    "Allow",
					Principal: map[string]string{"AWS": "*"},
					Action:    []string{"s3:GetObject"},
					Resource:  []string{"arn:aws:s3:::assets/*"},
				},
			},
		}

	default:
		return AwsPolicy{}
	}
}

func marshalPolicy(policy AwsPolicy) string {
	x, _ := json.Marshal(policy)
	return string(x)
}

func unmarshalPolicy(s string) (policy AwsPolicy) {
	json.Unmarshal([]byte(s), &policy)
	return
}
