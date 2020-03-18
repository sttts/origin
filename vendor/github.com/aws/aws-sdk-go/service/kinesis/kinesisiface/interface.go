// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package kinesisiface provides an interface to enable mocking the Amazon Kinesis service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package kinesisiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

// KinesisAPI provides an interface to enable mocking the
// kinesis.Kinesis service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Kinesis.
//    func myFunc(svc kinesisiface.KinesisAPI) bool {
//        // Make svc.AddTagsToStream request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := kinesis.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockKinesisClient struct {
//        kinesisiface.KinesisAPI
//    }
//    func (m *mockKinesisClient) AddTagsToStream(input *kinesis.AddTagsToStreamInput) (*kinesis.AddTagsToStreamOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockKinesisClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type KinesisAPI interface {
	AddTagsToStream(*kinesis.AddTagsToStreamInput) (*kinesis.AddTagsToStreamOutput, error)
	AddTagsToStreamWithContext(aws.Context, *kinesis.AddTagsToStreamInput, ...request.Option) (*kinesis.AddTagsToStreamOutput, error)
	AddTagsToStreamRequest(*kinesis.AddTagsToStreamInput) (*request.Request, *kinesis.AddTagsToStreamOutput)

	CreateStream(*kinesis.CreateStreamInput) (*kinesis.CreateStreamOutput, error)
	CreateStreamWithContext(aws.Context, *kinesis.CreateStreamInput, ...request.Option) (*kinesis.CreateStreamOutput, error)
	CreateStreamRequest(*kinesis.CreateStreamInput) (*request.Request, *kinesis.CreateStreamOutput)

	DecreaseStreamRetentionPeriod(*kinesis.DecreaseStreamRetentionPeriodInput) (*kinesis.DecreaseStreamRetentionPeriodOutput, error)
	DecreaseStreamRetentionPeriodWithContext(aws.Context, *kinesis.DecreaseStreamRetentionPeriodInput, ...request.Option) (*kinesis.DecreaseStreamRetentionPeriodOutput, error)
	DecreaseStreamRetentionPeriodRequest(*kinesis.DecreaseStreamRetentionPeriodInput) (*request.Request, *kinesis.DecreaseStreamRetentionPeriodOutput)

	DeleteStream(*kinesis.DeleteStreamInput) (*kinesis.DeleteStreamOutput, error)
	DeleteStreamWithContext(aws.Context, *kinesis.DeleteStreamInput, ...request.Option) (*kinesis.DeleteStreamOutput, error)
	DeleteStreamRequest(*kinesis.DeleteStreamInput) (*request.Request, *kinesis.DeleteStreamOutput)

	DeregisterStreamConsumer(*kinesis.DeregisterStreamConsumerInput) (*kinesis.DeregisterStreamConsumerOutput, error)
	DeregisterStreamConsumerWithContext(aws.Context, *kinesis.DeregisterStreamConsumerInput, ...request.Option) (*kinesis.DeregisterStreamConsumerOutput, error)
	DeregisterStreamConsumerRequest(*kinesis.DeregisterStreamConsumerInput) (*request.Request, *kinesis.DeregisterStreamConsumerOutput)

	DescribeLimits(*kinesis.DescribeLimitsInput) (*kinesis.DescribeLimitsOutput, error)
	DescribeLimitsWithContext(aws.Context, *kinesis.DescribeLimitsInput, ...request.Option) (*kinesis.DescribeLimitsOutput, error)
	DescribeLimitsRequest(*kinesis.DescribeLimitsInput) (*request.Request, *kinesis.DescribeLimitsOutput)

	DescribeStream(*kinesis.DescribeStreamInput) (*kinesis.DescribeStreamOutput, error)
	DescribeStreamWithContext(aws.Context, *kinesis.DescribeStreamInput, ...request.Option) (*kinesis.DescribeStreamOutput, error)
	DescribeStreamRequest(*kinesis.DescribeStreamInput) (*request.Request, *kinesis.DescribeStreamOutput)

	DescribeStreamPages(*kinesis.DescribeStreamInput, func(*kinesis.DescribeStreamOutput, bool) bool) error
	DescribeStreamPagesWithContext(aws.Context, *kinesis.DescribeStreamInput, func(*kinesis.DescribeStreamOutput, bool) bool, ...request.Option) error

	DescribeStreamConsumer(*kinesis.DescribeStreamConsumerInput) (*kinesis.DescribeStreamConsumerOutput, error)
	DescribeStreamConsumerWithContext(aws.Context, *kinesis.DescribeStreamConsumerInput, ...request.Option) (*kinesis.DescribeStreamConsumerOutput, error)
	DescribeStreamConsumerRequest(*kinesis.DescribeStreamConsumerInput) (*request.Request, *kinesis.DescribeStreamConsumerOutput)

	DescribeStreamSummary(*kinesis.DescribeStreamSummaryInput) (*kinesis.DescribeStreamSummaryOutput, error)
	DescribeStreamSummaryWithContext(aws.Context, *kinesis.DescribeStreamSummaryInput, ...request.Option) (*kinesis.DescribeStreamSummaryOutput, error)
	DescribeStreamSummaryRequest(*kinesis.DescribeStreamSummaryInput) (*request.Request, *kinesis.DescribeStreamSummaryOutput)

	DisableEnhancedMonitoring(*kinesis.DisableEnhancedMonitoringInput) (*kinesis.EnhancedMonitoringOutput, error)
	DisableEnhancedMonitoringWithContext(aws.Context, *kinesis.DisableEnhancedMonitoringInput, ...request.Option) (*kinesis.EnhancedMonitoringOutput, error)
	DisableEnhancedMonitoringRequest(*kinesis.DisableEnhancedMonitoringInput) (*request.Request, *kinesis.EnhancedMonitoringOutput)

	EnableEnhancedMonitoring(*kinesis.EnableEnhancedMonitoringInput) (*kinesis.EnhancedMonitoringOutput, error)
	EnableEnhancedMonitoringWithContext(aws.Context, *kinesis.EnableEnhancedMonitoringInput, ...request.Option) (*kinesis.EnhancedMonitoringOutput, error)
	EnableEnhancedMonitoringRequest(*kinesis.EnableEnhancedMonitoringInput) (*request.Request, *kinesis.EnhancedMonitoringOutput)

	GetRecords(*kinesis.GetRecordsInput) (*kinesis.GetRecordsOutput, error)
	GetRecordsWithContext(aws.Context, *kinesis.GetRecordsInput, ...request.Option) (*kinesis.GetRecordsOutput, error)
	GetRecordsRequest(*kinesis.GetRecordsInput) (*request.Request, *kinesis.GetRecordsOutput)

	GetShardIterator(*kinesis.GetShardIteratorInput) (*kinesis.GetShardIteratorOutput, error)
	GetShardIteratorWithContext(aws.Context, *kinesis.GetShardIteratorInput, ...request.Option) (*kinesis.GetShardIteratorOutput, error)
	GetShardIteratorRequest(*kinesis.GetShardIteratorInput) (*request.Request, *kinesis.GetShardIteratorOutput)

	IncreaseStreamRetentionPeriod(*kinesis.IncreaseStreamRetentionPeriodInput) (*kinesis.IncreaseStreamRetentionPeriodOutput, error)
	IncreaseStreamRetentionPeriodWithContext(aws.Context, *kinesis.IncreaseStreamRetentionPeriodInput, ...request.Option) (*kinesis.IncreaseStreamRetentionPeriodOutput, error)
	IncreaseStreamRetentionPeriodRequest(*kinesis.IncreaseStreamRetentionPeriodInput) (*request.Request, *kinesis.IncreaseStreamRetentionPeriodOutput)

	ListShards(*kinesis.ListShardsInput) (*kinesis.ListShardsOutput, error)
	ListShardsWithContext(aws.Context, *kinesis.ListShardsInput, ...request.Option) (*kinesis.ListShardsOutput, error)
	ListShardsRequest(*kinesis.ListShardsInput) (*request.Request, *kinesis.ListShardsOutput)

	ListStreamConsumers(*kinesis.ListStreamConsumersInput) (*kinesis.ListStreamConsumersOutput, error)
	ListStreamConsumersWithContext(aws.Context, *kinesis.ListStreamConsumersInput, ...request.Option) (*kinesis.ListStreamConsumersOutput, error)
	ListStreamConsumersRequest(*kinesis.ListStreamConsumersInput) (*request.Request, *kinesis.ListStreamConsumersOutput)

	ListStreamConsumersPages(*kinesis.ListStreamConsumersInput, func(*kinesis.ListStreamConsumersOutput, bool) bool) error
	ListStreamConsumersPagesWithContext(aws.Context, *kinesis.ListStreamConsumersInput, func(*kinesis.ListStreamConsumersOutput, bool) bool, ...request.Option) error

	ListStreams(*kinesis.ListStreamsInput) (*kinesis.ListStreamsOutput, error)
	ListStreamsWithContext(aws.Context, *kinesis.ListStreamsInput, ...request.Option) (*kinesis.ListStreamsOutput, error)
	ListStreamsRequest(*kinesis.ListStreamsInput) (*request.Request, *kinesis.ListStreamsOutput)

	ListStreamsPages(*kinesis.ListStreamsInput, func(*kinesis.ListStreamsOutput, bool) bool) error
	ListStreamsPagesWithContext(aws.Context, *kinesis.ListStreamsInput, func(*kinesis.ListStreamsOutput, bool) bool, ...request.Option) error

	ListTagsForStream(*kinesis.ListTagsForStreamInput) (*kinesis.ListTagsForStreamOutput, error)
	ListTagsForStreamWithContext(aws.Context, *kinesis.ListTagsForStreamInput, ...request.Option) (*kinesis.ListTagsForStreamOutput, error)
	ListTagsForStreamRequest(*kinesis.ListTagsForStreamInput) (*request.Request, *kinesis.ListTagsForStreamOutput)

	MergeShards(*kinesis.MergeShardsInput) (*kinesis.MergeShardsOutput, error)
	MergeShardsWithContext(aws.Context, *kinesis.MergeShardsInput, ...request.Option) (*kinesis.MergeShardsOutput, error)
	MergeShardsRequest(*kinesis.MergeShardsInput) (*request.Request, *kinesis.MergeShardsOutput)

	PutRecord(*kinesis.PutRecordInput) (*kinesis.PutRecordOutput, error)
	PutRecordWithContext(aws.Context, *kinesis.PutRecordInput, ...request.Option) (*kinesis.PutRecordOutput, error)
	PutRecordRequest(*kinesis.PutRecordInput) (*request.Request, *kinesis.PutRecordOutput)

	PutRecords(*kinesis.PutRecordsInput) (*kinesis.PutRecordsOutput, error)
	PutRecordsWithContext(aws.Context, *kinesis.PutRecordsInput, ...request.Option) (*kinesis.PutRecordsOutput, error)
	PutRecordsRequest(*kinesis.PutRecordsInput) (*request.Request, *kinesis.PutRecordsOutput)

	RegisterStreamConsumer(*kinesis.RegisterStreamConsumerInput) (*kinesis.RegisterStreamConsumerOutput, error)
	RegisterStreamConsumerWithContext(aws.Context, *kinesis.RegisterStreamConsumerInput, ...request.Option) (*kinesis.RegisterStreamConsumerOutput, error)
	RegisterStreamConsumerRequest(*kinesis.RegisterStreamConsumerInput) (*request.Request, *kinesis.RegisterStreamConsumerOutput)

	RemoveTagsFromStream(*kinesis.RemoveTagsFromStreamInput) (*kinesis.RemoveTagsFromStreamOutput, error)
	RemoveTagsFromStreamWithContext(aws.Context, *kinesis.RemoveTagsFromStreamInput, ...request.Option) (*kinesis.RemoveTagsFromStreamOutput, error)
	RemoveTagsFromStreamRequest(*kinesis.RemoveTagsFromStreamInput) (*request.Request, *kinesis.RemoveTagsFromStreamOutput)

	SplitShard(*kinesis.SplitShardInput) (*kinesis.SplitShardOutput, error)
	SplitShardWithContext(aws.Context, *kinesis.SplitShardInput, ...request.Option) (*kinesis.SplitShardOutput, error)
	SplitShardRequest(*kinesis.SplitShardInput) (*request.Request, *kinesis.SplitShardOutput)

	StartStreamEncryption(*kinesis.StartStreamEncryptionInput) (*kinesis.StartStreamEncryptionOutput, error)
	StartStreamEncryptionWithContext(aws.Context, *kinesis.StartStreamEncryptionInput, ...request.Option) (*kinesis.StartStreamEncryptionOutput, error)
	StartStreamEncryptionRequest(*kinesis.StartStreamEncryptionInput) (*request.Request, *kinesis.StartStreamEncryptionOutput)

	StopStreamEncryption(*kinesis.StopStreamEncryptionInput) (*kinesis.StopStreamEncryptionOutput, error)
	StopStreamEncryptionWithContext(aws.Context, *kinesis.StopStreamEncryptionInput, ...request.Option) (*kinesis.StopStreamEncryptionOutput, error)
	StopStreamEncryptionRequest(*kinesis.StopStreamEncryptionInput) (*request.Request, *kinesis.StopStreamEncryptionOutput)

	SubscribeToShard(*kinesis.SubscribeToShardInput) (*kinesis.SubscribeToShardOutput, error)
	SubscribeToShardWithContext(aws.Context, *kinesis.SubscribeToShardInput, ...request.Option) (*kinesis.SubscribeToShardOutput, error)
	SubscribeToShardRequest(*kinesis.SubscribeToShardInput) (*request.Request, *kinesis.SubscribeToShardOutput)

	UpdateShardCount(*kinesis.UpdateShardCountInput) (*kinesis.UpdateShardCountOutput, error)
	UpdateShardCountWithContext(aws.Context, *kinesis.UpdateShardCountInput, ...request.Option) (*kinesis.UpdateShardCountOutput, error)
	UpdateShardCountRequest(*kinesis.UpdateShardCountInput) (*request.Request, *kinesis.UpdateShardCountOutput)

	WaitUntilStreamExists(*kinesis.DescribeStreamInput) error
	WaitUntilStreamExistsWithContext(aws.Context, *kinesis.DescribeStreamInput, ...request.WaiterOption) error

	WaitUntilStreamNotExists(*kinesis.DescribeStreamInput) error
	WaitUntilStreamNotExistsWithContext(aws.Context, *kinesis.DescribeStreamInput, ...request.WaiterOption) error
}

var _ KinesisAPI = (*kinesis.Kinesis)(nil)
