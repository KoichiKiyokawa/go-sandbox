package design // The convention consists of naming the design
// package "design"
import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("svc", func() { // API defines the microservice endpoint and
	Title("The virtual wine cellar")    // other global properties. There should be one
	Description("A simple goa service") // and exactly one API definition appearing in
})

var _ = Service("twitter-clone", func() { // Resources group related API endpoints
	Description("Twitter clone")
	Method("publicTimeline", func() {
		HTTP(func() {
			GET("/timeline/public")
			Response(StatusOK)
		})
	})
})
