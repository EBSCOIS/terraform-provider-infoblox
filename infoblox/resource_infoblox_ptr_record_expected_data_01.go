package infoblox

import ibclient "github.com/ebscois/platform.infrastructure.infoblox-go-client/v2"

var testCasePtrRecordExpectedData01 = map[string]*ibclient.RecordPTR{
	"infoblox_ptr_record.rec1": {
		View:     "default",
		Ipv4Addr: "10.0.0.1",
		Ipv6Addr: "",
		Name:     "1.0.0.10.in-addr.arpa",
		PtrdName: "ptr-target1.example1.org",
		Zone:     "0.0.10.in-addr.arpa",
		UseTtl:   true,
		Ttl:      5,
		Comment:  "non-empty comment",
		Ea:       ibclient.EA{"Location": "Test location"},
	},
	"infoblox_ptr_record.rec2": {
		View:     "default",
		Ipv4Addr: "10.0.0.32",
		Ipv6Addr: "",
		Name:     "32.0.0.10.in-addr.arpa",
		PtrdName: "ptr-target2.example1.org",
		Zone:     "0.0.10.in-addr.arpa",
		UseTtl:   false,
		Ttl:      0,
		Comment:  "",
		Ea:       ibclient.EA{},
	},
	"infoblox_ptr_record.rec3": {
		View:     "default",
		Ipv4Addr: "",
		Ipv6Addr: "",
		Name:     "ptr-rec3-2.example1.org",
		PtrdName: "ptr-target3.example1.org",
		Zone:     "example1.org",
		UseTtl:   false,
		Ttl:      0,
		Comment:  "",
		Ea:       ibclient.EA{},
	},
	"infoblox_ptr_record.rec4": {
		View:     "default",
		Ipv4Addr: "172.16.0.201",
		Ipv6Addr: "",
		Name:     "201.0.16.172.in-addr.arpa",
		PtrdName: "ptr-target4.example1.org",
		Zone:     "0.16.172.in-addr.arpa",
		UseTtl:   false,
		Ttl:      0,
		Comment:  "",
		Ea:       ibclient.EA{},
	},
	"infoblox_ptr_record.rec5": {
		View:     "default",
		Ipv4Addr: "",
		Ipv6Addr: "2002:1f93::cb",
		Name:     "b.c.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.3.9.f.1.2.0.0.2.ip6.arpa",
		PtrdName: "ptr-target5.example1.org",
		Zone:     "0.0.0.0.0.0.0.0.3.9.f.1.2.0.0.2.ip6.arpa",
		UseTtl:   true,
		Ttl:      302,
		Comment:  "workstation #5-2",
		Ea:       ibclient.EA{"Location": "the new office"},
	},
	"infoblox_ptr_record.rec6": {
		View:     "default",
		Ipv4Addr: "",
		Ipv6Addr: "2002:1f93::cc",
		Name:     "c.c.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.3.9.f.1.2.0.0.2.ip6.arpa",
		PtrdName: "ptr-target6.example1.org",
		Zone:     "0.0.0.0.0.0.0.0.3.9.f.1.2.0.0.2.ip6.arpa",
		UseTtl:   false,
		Ttl:      0,
		Comment:  "",
		Ea:       ibclient.EA{},
	},
	"infoblox_ptr_record.rec7": {
		View:     "default",
		Ipv4Addr: "",
		Ipv6Addr: "2002:1f93::c9",
		Name:     "9.c.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.3.9.f.1.2.0.0.2.ip6.arpa",
		PtrdName: "ptr-target7.example1.org",
		Zone:     "0.0.0.0.0.0.0.0.3.9.f.1.2.0.0.2.ip6.arpa",
		UseTtl:   false,
		Ttl:      0,
		Comment:  "",
		Ea:       ibclient.EA{},
	},
	"infoblox_ptr_record.rec8": {
		View:     "default",
		Ipv4Addr: "",
		Ipv6Addr: "",
		Name:     "ptr-rec8.example1.org",
		PtrdName: "ptr-target8.example1.org",
		Zone:     "example1.org",
		UseTtl:   true,
		Ttl:      301,
		Comment:  "workstation #8",
		Ea:       ibclient.EA{"Location": "the main office"},
	},
	"infoblox_ptr_record.rec9": {
		View:     "nondefault_dnsview1",
		Ipv4Addr: "",
		Ipv6Addr: "2002:1f93::ca",
		Name:     "a.c.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.3.9.f.1.2.0.0.2.ip6.arpa",
		PtrdName: "ptr-target9.example2.org",
		Zone:     "0.0.0.0.0.0.0.0.3.9.f.1.2.0.0.2.ip6.arpa",
		UseTtl:   true,
		Ttl:      300,
		Comment:  "workstation #9",
		Ea:       ibclient.EA{"Location": "the main office"},
	},
	"infoblox_ptr_record.rec10": {
		View:     "nondefault_dnsview1",
		Ipv4Addr: "",
		Ipv6Addr: "2002:1f93::b",
		Name:     "b.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.3.9.f.1.2.0.0.2.ip6.arpa",
		PtrdName: "ptr-target10.example2.org",
		Zone:     "0.0.0.0.0.0.0.0.3.9.f.1.2.0.0.2.ip6.arpa",
		UseTtl:   true,
		Ttl:      30,
		Comment:  "workstation #10",
		Ea:       ibclient.EA{"Location": "the main office"},
	},
	"infoblox_ptr_record.rec11": {
		View:     "nondefault_dnsview1",
		Ipv4Addr: "",
		Ipv6Addr: "",
		Name:     "ptr-rec11.example2.org",
		PtrdName: "ptr-target11.example2.org",
		Zone:     "example2.org",
		UseTtl:   true,
		Ttl:      301,
		Comment:  "workstation #11",
		Ea:       ibclient.EA{"Location": "the main office"},
	},
	"infoblox_ptr_record.rec12": {
		View:     "nondefault_dnsview1",
		Ipv4Addr: "10.0.0.32",
		Ipv6Addr: "",
		Name:     "32.0.0.10.in-addr.arpa",
		PtrdName: "ptr-target12.example2.org",
		Zone:     "10.in-addr.arpa",
		UseTtl:   true,
		Ttl:      30,
		Comment:  "workstation #12",
		Ea:       ibclient.EA{"Location": "the main office"},
	},
	"infoblox_ptr_record.rec13": {
		View:     "nondefault_dnsview2",
		Ipv4Addr: "",
		Ipv6Addr: "2002:1f93::13",
		Name:     "3.1.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.3.9.f.1.2.0.0.2.ip6.arpa",
		PtrdName: "ptr-target13.example4.org",
		Zone:     "0.0.0.0.0.0.0.0.3.9.f.1.2.0.0.2.ip6.arpa",
		UseTtl:   true,
		Ttl:      30,
		Comment:  "workstation #13",
		Ea:       ibclient.EA{"Location": "the main office"},
	},
	"infoblox_ptr_record.rec14": {
		View:     "default",
		Ipv4Addr: "10.0.0.44",
		Ipv6Addr: "",
		Name:     "44.0.0.10.in-addr.arpa",
		PtrdName: "ptr-target14.example1.org",
		Zone:     "0.0.10.in-addr.arpa",
		UseTtl:   false,
		Ttl:      0,
		Comment:  "",
		Ea:       ibclient.EA{},
	},
}
