package main

import (
	"encoding/xml" // import "github.com/zzzhr1990/go-server-common/repo"
	"fmt"

	"github.com/zzzhr1990/go-server-common/config"
)

// WopiDiscovery WP
type WopiDiscovery struct {
	XMLName xml.Name `xml:"wopi-discovery"`
	Text    string   `xml:",chardata"`
	NetZone []struct {
		Text string `xml:",chardata"`
		Name string `xml:"name,attr"`
		App  []struct {
			Text         string `xml:",chardata"`
			Name         string `xml:"name,attr"`
			FavIconURL   string `xml:"favIconUrl,attr"`
			CheckLicense string `xml:"checkLicense,attr"`
			Action       []struct {
				Text      string `xml:",chardata"`
				Name      string `xml:"name,attr"`
				Ext       string `xml:"ext,attr"`
				Default   string `xml:"default,attr"`
				Urlsrc    string `xml:"urlsrc,attr"`
				Requires  string `xml:"requires,attr"`
				Targetext string `xml:"targetext,attr"`
				Progid    string `xml:"progid,attr"`
				UseParent string `xml:"useParent,attr"`
				Newprogid string `xml:"newprogid,attr"`
				Newext    string `xml:"newext,attr"`
			} `xml:"action"`
		} `xml:"app"`
	} `xml:"net-zone"`
	ProofKey struct {
		Text        string `xml:",chardata"`
		Oldvalue    string `xml:"oldvalue,attr"`
		Oldmodulus  string `xml:"oldmodulus,attr"`
		Oldexponent string `xml:"oldexponent,attr"`
		Value       string `xml:"value,attr"`
		Modulus     string `xml:"modulus,attr"`
		Exponent    string `xml:"exponent,attr"`
	} `xml:"proof-key"`
}

func main() {
	fmt.Println("s")
	// fmt.Println(repo.Version())
	xx := &WopiDiscovery{}
	config.LoadXMLFromURL("", xx)
	fmt.Println(xx.XMLName)
}
