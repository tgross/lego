// Code generated by 'make generate-dns'; DO NOT EDIT.

package dns

import (
	"fmt"

	"github.com/go-acme/lego/v4/challenge"
	"github.com/go-acme/lego/v4/challenge/dns01"
	"github.com/go-acme/lego/v4/providers/dns/acmedns"
	"github.com/go-acme/lego/v4/providers/dns/bunny"
	"github.com/go-acme/lego/v4/providers/dns/exec"
	"github.com/go-acme/lego/v4/providers/dns/route53"
)

// NewDNSChallengeProviderByName Factory for DNS providers.
func NewDNSChallengeProviderByName(name string) (challenge.Provider, error) {
	switch name {
	case "manual":
		return dns01.NewDNSProviderManual()
	case "acme-dns", "acmedns":
		return acmedns.NewDNSProvider()
	case "bunny":
		return bunny.NewDNSProvider()
	case "exec":
		return exec.NewDNSProvider()
	case "route53":
		return route53.NewDNSProvider()
	default:
		return nil, fmt.Errorf("unrecognized DNS provider: %s", name)
	}
}
