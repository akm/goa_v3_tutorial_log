// Code generated by goa v3.0.3, DO NOT EDIT.
//
// calc HTTP client CLI support package
//
// Command:
// $ goa gen calcsvc/design

package cli

import (
	accountc "calcsvc/gen/http/account/client"
	calcc "calcsvc/gen/http/calc/client"
	"flag"
	"fmt"
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
	return `account signup
calc add
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` account signup --body '{
      "id_token": "Quam minima."
   }'` + "\n" +
		os.Args[0] + ` calc add --a 3523480046783923250 --b 5487677923305002774 --token "Ut itaque sit corrupti velit."` + "\n" +
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
		accountFlags = flag.NewFlagSet("account", flag.ContinueOnError)

		accountSignupFlags    = flag.NewFlagSet("signup", flag.ExitOnError)
		accountSignupBodyFlag = accountSignupFlags.String("body", "REQUIRED", "")

		calcFlags = flag.NewFlagSet("calc", flag.ContinueOnError)

		calcAddFlags     = flag.NewFlagSet("add", flag.ExitOnError)
		calcAddAFlag     = calcAddFlags.String("a", "REQUIRED", "Left operand")
		calcAddBFlag     = calcAddFlags.String("b", "REQUIRED", "Right operand")
		calcAddTokenFlag = calcAddFlags.String("token", "REQUIRED", "")
	)
	accountFlags.Usage = accountUsage
	accountSignupFlags.Usage = accountSignupUsage

	calcFlags.Usage = calcUsage
	calcAddFlags.Usage = calcAddUsage

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
		case "account":
			svcf = accountFlags
		case "calc":
			svcf = calcFlags
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
		case "account":
			switch epn {
			case "signup":
				epf = accountSignupFlags

			}

		case "calc":
			switch epn {
			case "add":
				epf = calcAddFlags

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
		case "account":
			c := accountc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "signup":
				endpoint = c.Signup()
				data, err = accountc.BuildSignupPayload(*accountSignupBodyFlag)
			}
		case "calc":
			c := calcc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "add":
				endpoint = c.Add()
				data, err = calcc.BuildAddPayload(*calcAddAFlag, *calcAddBFlag, *calcAddTokenFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// accountUsage displays the usage of the account command and its subcommands.
func accountUsage() {
	fmt.Fprintf(os.Stderr, `Create and delete account
Usage:
    %s [globalflags] account COMMAND [flags]

COMMAND:
    signup: Sign up  account with ID token from Google

Additional help:
    %s account COMMAND --help
`, os.Args[0], os.Args[0])
}
func accountSignupUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] account signup -body JSON

Sign up  account with ID token from Google
    -body JSON: 

Example:
    `+os.Args[0]+` account signup --body '{
      "id_token": "Quam minima."
   }'
`, os.Args[0])
}

// calcUsage displays the usage of the calc command and its subcommands.
func calcUsage() {
	fmt.Fprintf(os.Stderr, `The calc service performs operations on numbers
Usage:
    %s [globalflags] calc COMMAND [flags]

COMMAND:
    add: Add implements add.

Additional help:
    %s calc COMMAND --help
`, os.Args[0], os.Args[0])
}
func calcAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] calc add -a INT -b INT -token STRING

Add implements add.
    -a INT: Left operand
    -b INT: Right operand
    -token STRING: 

Example:
    `+os.Args[0]+` calc add --a 3523480046783923250 --b 5487677923305002774 --token "Ut itaque sit corrupti velit."
`, os.Args[0])
}
