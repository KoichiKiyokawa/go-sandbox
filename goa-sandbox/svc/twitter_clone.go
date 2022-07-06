package svc

import (
	"context"
	twitterclone "goa-sandbox/gen/twitter_clone"
	"log"
)

// twitter-clone service example implementation.
// The example methods log the requests and return zero values.
type twitterClonesrvc struct {
	logger *log.Logger
}

// NewTwitterClone returns the twitter-clone service implementation.
func NewTwitterClone(logger *log.Logger) twitterclone.Service {
	return &twitterClonesrvc{logger}
}

// PublicTimeline implements publicTimeline.
func (s *twitterClonesrvc) PublicTimeline(ctx context.Context) (err error) {
	s.logger.Print("twitterClone.publicTimeline")
	return
}
