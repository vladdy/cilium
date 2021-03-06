// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CancelCapacityReservationInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Capacity Reservation to be cancelled.
	//
	// CapacityReservationId is a required field
	CapacityReservationId *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`
}

// String returns the string representation
func (s CancelCapacityReservationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelCapacityReservationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelCapacityReservationInput"}

	if s.CapacityReservationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CapacityReservationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CancelCapacityReservationOutput struct {
	_ struct{} `type:"structure"`

	// Returns true if the request succeeds; otherwise, it returns an error.
	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s CancelCapacityReservationOutput) String() string {
	return awsutil.Prettify(s)
}

const opCancelCapacityReservation = "CancelCapacityReservation"

// CancelCapacityReservationRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Cancels the specified Capacity Reservation, releases the reserved capacity,
// and changes the Capacity Reservation's state to cancelled.
//
// Instances running in the reserved capacity continue running until you stop
// them. Stopped instances that target the Capacity Reservation can no longer
// launch. Modify these instances to either target a different Capacity Reservation,
// launch On-Demand Instance capacity, or run in any open Capacity Reservation
// that has matching attributes and sufficient capacity.
//
//    // Example sending a request using CancelCapacityReservationRequest.
//    req := client.CancelCapacityReservationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CancelCapacityReservation
func (c *Client) CancelCapacityReservationRequest(input *CancelCapacityReservationInput) CancelCapacityReservationRequest {
	op := &aws.Operation{
		Name:       opCancelCapacityReservation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CancelCapacityReservationInput{}
	}

	req := c.newRequest(op, input, &CancelCapacityReservationOutput{})

	return CancelCapacityReservationRequest{Request: req, Input: input, Copy: c.CancelCapacityReservationRequest}
}

// CancelCapacityReservationRequest is the request type for the
// CancelCapacityReservation API operation.
type CancelCapacityReservationRequest struct {
	*aws.Request
	Input *CancelCapacityReservationInput
	Copy  func(*CancelCapacityReservationInput) CancelCapacityReservationRequest
}

// Send marshals and sends the CancelCapacityReservation API request.
func (r CancelCapacityReservationRequest) Send(ctx context.Context) (*CancelCapacityReservationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelCapacityReservationResponse{
		CancelCapacityReservationOutput: r.Request.Data.(*CancelCapacityReservationOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelCapacityReservationResponse is the response type for the
// CancelCapacityReservation API operation.
type CancelCapacityReservationResponse struct {
	*CancelCapacityReservationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelCapacityReservation request.
func (r *CancelCapacityReservationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
