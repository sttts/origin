// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2instanceconnect

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
)

const opSendSSHPublicKey = "SendSSHPublicKey"

// SendSSHPublicKeyRequest generates a "aws/request.Request" representing the
// client's request for the SendSSHPublicKey operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See SendSSHPublicKey for more information on using the SendSSHPublicKey
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the SendSSHPublicKeyRequest method.
//    req, resp := client.SendSSHPublicKeyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/ec2-instance-connect-2018-04-02/SendSSHPublicKey
func (c *EC2InstanceConnect) SendSSHPublicKeyRequest(input *SendSSHPublicKeyInput) (req *request.Request, output *SendSSHPublicKeyOutput) {
	op := &request.Operation{
		Name:       opSendSSHPublicKey,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SendSSHPublicKeyInput{}
	}

	output = &SendSSHPublicKeyOutput{}
	req = c.newRequest(op, input, output)
	return
}

// SendSSHPublicKey API operation for AWS EC2 Instance Connect.
//
// Pushes an SSH public key to a particular OS user on a given EC2 instance
// for 60 seconds.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS EC2 Instance Connect's
// API operation SendSSHPublicKey for usage and error information.
//
// Returned Error Types:
//   * AuthException
//   Indicates that either your AWS credentials are invalid or you do not have
//   access to the EC2 instance.
//
//   * InvalidArgsException
//   Indicates that you provided bad input. Ensure you have a valid instance ID,
//   the correct zone, and a valid SSH public key.
//
//   * ServiceException
//   Indicates that the service encountered an error. Follow the message's instructions
//   and try again.
//
//   * ThrottlingException
//   Indicates you have been making requests too frequently and have been throttled.
//   Wait for a while and try again. If higher call volume is warranted contact
//   AWS Support.
//
//   * EC2InstanceNotFoundException
//   Indicates that the instance requested was not found in the given zone. Check
//   that you have provided a valid instance ID and the correct zone.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/ec2-instance-connect-2018-04-02/SendSSHPublicKey
func (c *EC2InstanceConnect) SendSSHPublicKey(input *SendSSHPublicKeyInput) (*SendSSHPublicKeyOutput, error) {
	req, out := c.SendSSHPublicKeyRequest(input)
	return out, req.Send()
}

// SendSSHPublicKeyWithContext is the same as SendSSHPublicKey with the addition of
// the ability to pass a context and additional request options.
//
// See SendSSHPublicKey for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EC2InstanceConnect) SendSSHPublicKeyWithContext(ctx aws.Context, input *SendSSHPublicKeyInput, opts ...request.Option) (*SendSSHPublicKeyOutput, error) {
	req, out := c.SendSSHPublicKeyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// Indicates that either your AWS credentials are invalid or you do not have
// access to the EC2 instance.
type AuthException struct {
	_            struct{} `type:"structure"`
	respMetadata protocol.ResponseMetadata

	Message_ *string `locationName:"Message" type:"string"`
}

// String returns the string representation
func (s AuthException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s AuthException) GoString() string {
	return s.String()
}

func newErrorAuthException(v protocol.ResponseMetadata) error {
	return &AuthException{
		respMetadata: v,
	}
}

// Code returns the exception type name.
func (s AuthException) Code() string {
	return "AuthException"
}

// Message returns the exception's message.
func (s AuthException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s AuthException) OrigErr() error {
	return nil
}

func (s AuthException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s AuthException) StatusCode() int {
	return s.respMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s AuthException) RequestID() string {
	return s.respMetadata.RequestID
}

// Indicates that the instance requested was not found in the given zone. Check
// that you have provided a valid instance ID and the correct zone.
type EC2InstanceNotFoundException struct {
	_            struct{} `type:"structure"`
	respMetadata protocol.ResponseMetadata

	Message_ *string `locationName:"Message" type:"string"`
}

// String returns the string representation
func (s EC2InstanceNotFoundException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s EC2InstanceNotFoundException) GoString() string {
	return s.String()
}

func newErrorEC2InstanceNotFoundException(v protocol.ResponseMetadata) error {
	return &EC2InstanceNotFoundException{
		respMetadata: v,
	}
}

// Code returns the exception type name.
func (s EC2InstanceNotFoundException) Code() string {
	return "EC2InstanceNotFoundException"
}

// Message returns the exception's message.
func (s EC2InstanceNotFoundException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s EC2InstanceNotFoundException) OrigErr() error {
	return nil
}

func (s EC2InstanceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s EC2InstanceNotFoundException) StatusCode() int {
	return s.respMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s EC2InstanceNotFoundException) RequestID() string {
	return s.respMetadata.RequestID
}

// Indicates that you provided bad input. Ensure you have a valid instance ID,
// the correct zone, and a valid SSH public key.
type InvalidArgsException struct {
	_            struct{} `type:"structure"`
	respMetadata protocol.ResponseMetadata

	Message_ *string `locationName:"Message" type:"string"`
}

// String returns the string representation
func (s InvalidArgsException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s InvalidArgsException) GoString() string {
	return s.String()
}

func newErrorInvalidArgsException(v protocol.ResponseMetadata) error {
	return &InvalidArgsException{
		respMetadata: v,
	}
}

// Code returns the exception type name.
func (s InvalidArgsException) Code() string {
	return "InvalidArgsException"
}

// Message returns the exception's message.
func (s InvalidArgsException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s InvalidArgsException) OrigErr() error {
	return nil
}

func (s InvalidArgsException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s InvalidArgsException) StatusCode() int {
	return s.respMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s InvalidArgsException) RequestID() string {
	return s.respMetadata.RequestID
}

type SendSSHPublicKeyInput struct {
	_ struct{} `type:"structure"`

	// The availability zone the EC2 instance was launched in.
	//
	// AvailabilityZone is a required field
	AvailabilityZone *string `min:"6" type:"string" required:"true"`

	// The EC2 instance you wish to publish the SSH key to.
	//
	// InstanceId is a required field
	InstanceId *string `min:"10" type:"string" required:"true"`

	// The OS user on the EC2 instance whom the key may be used to authenticate
	// as.
	//
	// InstanceOSUser is a required field
	InstanceOSUser *string `min:"1" type:"string" required:"true"`

	// The public key to be published to the instance. To use it after publication
	// you must have the matching private key.
	//
	// SSHPublicKey is a required field
	SSHPublicKey *string `min:"256" type:"string" required:"true"`
}

// String returns the string representation
func (s SendSSHPublicKeyInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s SendSSHPublicKeyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SendSSHPublicKeyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "SendSSHPublicKeyInput"}
	if s.AvailabilityZone == nil {
		invalidParams.Add(request.NewErrParamRequired("AvailabilityZone"))
	}
	if s.AvailabilityZone != nil && len(*s.AvailabilityZone) < 6 {
		invalidParams.Add(request.NewErrParamMinLen("AvailabilityZone", 6))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceId != nil && len(*s.InstanceId) < 10 {
		invalidParams.Add(request.NewErrParamMinLen("InstanceId", 10))
	}
	if s.InstanceOSUser == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceOSUser"))
	}
	if s.InstanceOSUser != nil && len(*s.InstanceOSUser) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("InstanceOSUser", 1))
	}
	if s.SSHPublicKey == nil {
		invalidParams.Add(request.NewErrParamRequired("SSHPublicKey"))
	}
	if s.SSHPublicKey != nil && len(*s.SSHPublicKey) < 256 {
		invalidParams.Add(request.NewErrParamMinLen("SSHPublicKey", 256))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAvailabilityZone sets the AvailabilityZone field's value.
func (s *SendSSHPublicKeyInput) SetAvailabilityZone(v string) *SendSSHPublicKeyInput {
	s.AvailabilityZone = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *SendSSHPublicKeyInput) SetInstanceId(v string) *SendSSHPublicKeyInput {
	s.InstanceId = &v
	return s
}

// SetInstanceOSUser sets the InstanceOSUser field's value.
func (s *SendSSHPublicKeyInput) SetInstanceOSUser(v string) *SendSSHPublicKeyInput {
	s.InstanceOSUser = &v
	return s
}

// SetSSHPublicKey sets the SSHPublicKey field's value.
func (s *SendSSHPublicKeyInput) SetSSHPublicKey(v string) *SendSSHPublicKeyInput {
	s.SSHPublicKey = &v
	return s
}

type SendSSHPublicKeyOutput struct {
	_ struct{} `type:"structure"`

	// The request ID as logged by EC2 Connect. Please provide this when contacting
	// AWS Support.
	RequestId *string `type:"string"`

	// Indicates request success.
	Success *bool `type:"boolean"`
}

// String returns the string representation
func (s SendSSHPublicKeyOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s SendSSHPublicKeyOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *SendSSHPublicKeyOutput) SetRequestId(v string) *SendSSHPublicKeyOutput {
	s.RequestId = &v
	return s
}

// SetSuccess sets the Success field's value.
func (s *SendSSHPublicKeyOutput) SetSuccess(v bool) *SendSSHPublicKeyOutput {
	s.Success = &v
	return s
}

// Indicates that the service encountered an error. Follow the message's instructions
// and try again.
type ServiceException struct {
	_            struct{} `type:"structure"`
	respMetadata protocol.ResponseMetadata

	Message_ *string `locationName:"Message" type:"string"`
}

// String returns the string representation
func (s ServiceException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ServiceException) GoString() string {
	return s.String()
}

func newErrorServiceException(v protocol.ResponseMetadata) error {
	return &ServiceException{
		respMetadata: v,
	}
}

// Code returns the exception type name.
func (s ServiceException) Code() string {
	return "ServiceException"
}

// Message returns the exception's message.
func (s ServiceException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s ServiceException) OrigErr() error {
	return nil
}

func (s ServiceException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s ServiceException) StatusCode() int {
	return s.respMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s ServiceException) RequestID() string {
	return s.respMetadata.RequestID
}

// Indicates you have been making requests too frequently and have been throttled.
// Wait for a while and try again. If higher call volume is warranted contact
// AWS Support.
type ThrottlingException struct {
	_            struct{} `type:"structure"`
	respMetadata protocol.ResponseMetadata

	Message_ *string `locationName:"Message" type:"string"`
}

// String returns the string representation
func (s ThrottlingException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ThrottlingException) GoString() string {
	return s.String()
}

func newErrorThrottlingException(v protocol.ResponseMetadata) error {
	return &ThrottlingException{
		respMetadata: v,
	}
}

// Code returns the exception type name.
func (s ThrottlingException) Code() string {
	return "ThrottlingException"
}

// Message returns the exception's message.
func (s ThrottlingException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s ThrottlingException) OrigErr() error {
	return nil
}

func (s ThrottlingException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s ThrottlingException) StatusCode() int {
	return s.respMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s ThrottlingException) RequestID() string {
	return s.respMetadata.RequestID
}
