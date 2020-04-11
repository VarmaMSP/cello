package model

const (
	POLICY_VERSION       = "2012-10-17"
	POLICY_ID_THUMBNAILS = "2-12-2019"
	POLICY_ID_ASSETS     = "2-12-2019"
)

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

func NewAwsPolicy(statements []AwsPolicyStatement) AwsPolicy {
	return AwsPolicy{
		Id:        POLICY_ID_THUMBNAILS,
		Version:   POLICY_VERSION,
		Statement: statements,
	}
}
