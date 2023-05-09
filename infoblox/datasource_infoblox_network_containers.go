package infoblox

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	ibclient "github.com/ebscois/platform.infrastructure.infoblox-go-client/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Networkoutput struct {
	Result []struct {
		Ref         string                       `json:"_ref"`
		Extattrs    map[string]map[string]string `json:"extattrs"`
		Network     string                       `json:"network"`
		NetworkView string                       `json:"network_view"`
	}
}

func dataSourceIpv4NetworkContainers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIpv4NetworkContainersRead,
		Schema: map[string]*schema.Schema{
			"ext_attrs": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Extensible attributes for the network container.",
			},
			"network_view": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"ea_search": {
				Type:     schema.TypeMap,
				Optional: true,
			},
			"network": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Network.",
			},
		},
	}
}

func dataSourceIpv4NetworkContainersRead(d *schema.ResourceData, m interface{}) error {
	var networkoutput Networkoutput
	var returnCIDR []string
	extAttrJSON := d.Get("ext_attrs").(string)
	newview := d.Get("network_view").(string)
	eaSearch := d.Get("ea_search").(map[string]interface{})
	extAttrs := make(map[string]interface{})
	if extAttrJSON != "" {
		if err := json.Unmarshal([]byte(extAttrJSON), &extAttrs); err != nil {
			return fmt.Errorf("cannot process 'ext_attrs' field: %s", err.Error())
		}
	}

	log.Printf("[INFO] Ext Attrs: %v", extAttrs)
	connector := m.(ibclient.IBConnector)
	objMgr := ibclient.NewObjectManager(connector, "Terraform", "")
	obj, err := objMgr.GetNetworkByEA(newview, eaSearch)
	if err != nil {
		log.Println(err)
	}
	jsonout, err := json.Marshal(obj)
	err = json.Unmarshal(jsonout, &networkoutput)
	if err != nil {
		log.Println(err)
	}
	for _, res := range networkoutput.Result {
		splitref := strings.Split(strings.Split(res.Ref, ":")[1], "/")
		cidr := splitref[0] + "/" + splitref[1]
		returnCIDR = append(returnCIDR, cidr)
	}
	log.Printf("[INFO] cidr: %v\n", returnCIDR)
	// log.Printf("Network: %v\n", networkoutput.Result[0].Network)
	schemaCIDR := strings.Join(returnCIDR, ", ")
	log.Printf("[INFO] cidr: %v\n", schemaCIDR)
	if err = d.Set("network", schemaCIDR); err != nil {
		return err
	}
	d.SetId(string(jsonout))
	return nil
}

func searchByEA(ea string) error {

	return nil
}
