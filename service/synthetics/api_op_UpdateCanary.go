// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package synthetics

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateCanaryInput struct {
	_ struct{} `type:"structure"`

	// A structure that includes the entry point from which the canary should start
	// running your script. If the script is stored in an S3 bucket, the bucket
	// name, key, and version are also included.
	Code *CanaryCodeInput `type:"structure"`

	// The ARN of the IAM role to be used to run the canary. This role must already
	// exist, and must include lambda.amazonaws.com as a principal in the trust
	// policy. The role must also have the following permissions:
	//
	//    * s3:PutObject
	//
	//    * s3:GetBucketLocation
	//
	//    * s3:ListAllMyBuckets
	//
	//    * cloudwatch:PutMetricData
	//
	//    * logs:CreateLogGroup
	//
	//    * logs:CreateLogStream
	//
	//    * logs:CreateLogStream
	ExecutionRoleArn *string `type:"string"`

	// The number of days to retain data about failed runs of this canary.
	FailureRetentionPeriodInDays *int64 `min:"1" type:"integer"`

	// The name of the canary that you want to update. To find the names of your
	// canaries, use DescribeCanaries.
	//
	// You cannot change the name of a canary that has already been created.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" min:"1" type:"string" required:"true"`

	// A structure that contains the timeout value that is used for each individual
	// run of the canary.
	RunConfig *CanaryRunConfigInput `type:"structure"`

	// Specifies the runtime version to use for the canary. Currently, the only
	// valid value is syn-1.0. For more information about runtime versions, see
	// Canary Runtime Versions (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Library.html).
	RuntimeVersion *string `min:"1" type:"string"`

	// A structure that contains information about how often the canary is to run,
	// and when these runs are to stop.
	Schedule *CanaryScheduleInput `type:"structure"`

	// The number of days to retain data about successful runs of this canary.
	SuccessRetentionPeriodInDays *int64 `min:"1" type:"integer"`

	// If this canary is to test an endpoint in a VPC, this structure contains information
	// about the subnet and security groups of the VPC endpoint. For more information,
	// see Running a Canary in a VPC (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_VPC.html).
	VpcConfig *VpcConfigInput `type:"structure"`
}

// String returns the string representation
func (s UpdateCanaryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateCanaryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateCanaryInput"}
	if s.FailureRetentionPeriodInDays != nil && *s.FailureRetentionPeriodInDays < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("FailureRetentionPeriodInDays", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.RuntimeVersion != nil && len(*s.RuntimeVersion) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuntimeVersion", 1))
	}
	if s.SuccessRetentionPeriodInDays != nil && *s.SuccessRetentionPeriodInDays < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("SuccessRetentionPeriodInDays", 1))
	}
	if s.Code != nil {
		if err := s.Code.Validate(); err != nil {
			invalidParams.AddNested("Code", err.(aws.ErrInvalidParams))
		}
	}
	if s.RunConfig != nil {
		if err := s.RunConfig.Validate(); err != nil {
			invalidParams.AddNested("RunConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.Schedule != nil {
		if err := s.Schedule.Validate(); err != nil {
			invalidParams.AddNested("Schedule", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateCanaryInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Code != nil {
		v := s.Code

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Code", v, metadata)
	}
	if s.ExecutionRoleArn != nil {
		v := *s.ExecutionRoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ExecutionRoleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FailureRetentionPeriodInDays != nil {
		v := *s.FailureRetentionPeriodInDays

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FailureRetentionPeriodInDays", protocol.Int64Value(v), metadata)
	}
	if s.RunConfig != nil {
		v := s.RunConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "RunConfig", v, metadata)
	}
	if s.RuntimeVersion != nil {
		v := *s.RuntimeVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RuntimeVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Schedule != nil {
		v := s.Schedule

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Schedule", v, metadata)
	}
	if s.SuccessRetentionPeriodInDays != nil {
		v := *s.SuccessRetentionPeriodInDays

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SuccessRetentionPeriodInDays", protocol.Int64Value(v), metadata)
	}
	if s.VpcConfig != nil {
		v := s.VpcConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "VpcConfig", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateCanaryOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateCanaryOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateCanaryOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateCanary = "UpdateCanary"

// UpdateCanaryRequest returns a request value for making API operation for
// Synthetics.
//
// Use this operation to change the settings of a canary that has already been
// created.
//
// You can't use this operation to update the tags of an existing canary. To
// change the tags of an existing canary, use TagResource (https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_TagResource.html).
//
//    // Example sending a request using UpdateCanaryRequest.
//    req := client.UpdateCanaryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/synthetics-2017-10-11/UpdateCanary
func (c *Client) UpdateCanaryRequest(input *UpdateCanaryInput) UpdateCanaryRequest {
	op := &aws.Operation{
		Name:       opUpdateCanary,
		HTTPMethod: "PATCH",
		HTTPPath:   "/canary/{name}",
	}

	if input == nil {
		input = &UpdateCanaryInput{}
	}

	req := c.newRequest(op, input, &UpdateCanaryOutput{})
	return UpdateCanaryRequest{Request: req, Input: input, Copy: c.UpdateCanaryRequest}
}

// UpdateCanaryRequest is the request type for the
// UpdateCanary API operation.
type UpdateCanaryRequest struct {
	*aws.Request
	Input *UpdateCanaryInput
	Copy  func(*UpdateCanaryInput) UpdateCanaryRequest
}

// Send marshals and sends the UpdateCanary API request.
func (r UpdateCanaryRequest) Send(ctx context.Context) (*UpdateCanaryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateCanaryResponse{
		UpdateCanaryOutput: r.Request.Data.(*UpdateCanaryOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateCanaryResponse is the response type for the
// UpdateCanary API operation.
type UpdateCanaryResponse struct {
	*UpdateCanaryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateCanary request.
func (r *UpdateCanaryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}