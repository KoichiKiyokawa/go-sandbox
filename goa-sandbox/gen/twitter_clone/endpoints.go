// Code generated by goa v3.7.10, DO NOT EDIT.
//
// twitter-clone endpoints
//
// Command:
// $ goa gen goa-sandbox/design

package twitterclone

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "twitter-clone" service endpoints.
type Endpoints struct {
	PublicTimeline goa.Endpoint
}

// NewEndpoints wraps the methods of the "twitter-clone" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		PublicTimeline: NewPublicTimelineEndpoint(s),
	}
}

// Use applies the given middleware to all the "twitter-clone" service
// endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.PublicTimeline = m(e.PublicTimeline)
}

// NewPublicTimelineEndpoint returns an endpoint function that calls the method
// "publicTimeline" of service "twitter-clone".
func NewPublicTimelineEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, s.PublicTimeline(ctx)
	}
}
