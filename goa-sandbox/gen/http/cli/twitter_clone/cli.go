// Code generated by goa v3.7.10, DO NOT EDIT.
//
// twitter-clone HTTP client CLI support package
//
// Command:
// $ goa gen goa-sandbox/design

package cli

import (
	"flag"
	"fmt"
	twitterclonec "goa-sandbox/gen/http/twitter_clone/client"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `twitter-clone public-timeline
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` twitter-clone public-timeline` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		twitterCloneFlags = flag.NewFlagSet("twitter-clone", flag.ContinueOnError)

		twitterClonePublicTimelineFlags = flag.NewFlagSet("public-timeline", flag.ExitOnError)
	)
	twitterCloneFlags.Usage = twitterCloneUsage
	twitterClonePublicTimelineFlags.Usage = twitterClonePublicTimelineUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "twitter-clone":
			svcf = twitterCloneFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "twitter-clone":
			switch epn {
			case "public-timeline":
				epf = twitterClonePublicTimelineFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "twitter-clone":
			c := twitterclonec.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "public-timeline":
				endpoint = c.PublicTimeline()
				data = nil
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// twitter-cloneUsage displays the usage of the twitter-clone command and its
// subcommands.
func twitterCloneUsage() {
	fmt.Fprintf(os.Stderr, `Twitter clone
Usage:
    %[1]s [globalflags] twitter-clone COMMAND [flags]

COMMAND:
    public-timeline: PublicTimeline implements publicTimeline.

Additional help:
    %[1]s twitter-clone COMMAND --help
`, os.Args[0])
}
func twitterClonePublicTimelineUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] twitter-clone public-timeline

PublicTimeline implements publicTimeline.

Example:
    %[1]s twitter-clone public-timeline
`, os.Args[0])
}
