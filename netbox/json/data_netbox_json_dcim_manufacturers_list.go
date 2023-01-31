// Code generated by util/generateJsonDatasources; DO NOT EDIT.
package json

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	netboxclient "github.com/smutel/go-netbox/v3/netbox/client"
	"github.com/smutel/go-netbox/v3/netbox/client/dcim"
	"github.com/smutel/terraform-provider-netbox/v6/netbox/internal/util"
)

// This file was generated by the util/generateJsonDatasources.
// Editing this file might prove futile when you re-run the util/generateJsonDatasources command

func DataNetboxJSONDcimManufacturersList() *schema.Resource {
	return &schema.Resource{
		Description: "Get json output from the dcim_manufacturers_list Netbox endpoint.",
		ReadContext: dataNetboxJSONDcimManufacturersListRead,

		Schema: map[string]*schema.Schema{
			"filter": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Filter the records returned by the query.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Name of the field to use for filtering.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Value of the field to use for filtering.",
						},
					},
				},
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "The max number of returned results. If 0 is specified, all records will be returned.",
			},
			"json": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "JSON output of the list of objects for this Netbox endpoint.",
			},
		},
	}
}

func dataNetboxJSONDcimManufacturersListRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*netboxclient.NetBoxAPI)

	params := dcim.NewDcimManufacturersListParams()
	limit := int64(d.Get("limit").(int))
	params.Limit = &limit

	if filter, ok := d.GetOk("filter"); ok {
		var filterParams = filter.(*schema.Set)
		for _, f := range filterParams.List() {
			k := f.(map[string]interface{})["name"]
			v := f.(map[string]interface{})["value"]
			kString := k.(string)
			vString := v.(string)
			field := reflect.ValueOf(params).Elem().FieldByName(util.FieldNameToStructName(kString))
			if field != (reflect.Value{}) {
				field.Set(reflect.ValueOf(&vString))
			} else {
				return diag.Errorf("Field %s does not exist in schema.", kString)
			}
		}
	}

	list, err := client.Dcim.DcimManufacturersList(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	tmp := list.Payload.Results
	resultLength := int64(len(tmp))
	desiredLength := *list.Payload.Count
	if limit > 0 && limit < desiredLength {
		desiredLength = limit
	}
	if limit == 0 || resultLength < desiredLength {
		limit = resultLength
	}
	offset := limit
	params.Offset = &offset
	for int64(len(tmp)) < desiredLength {
		offset = int64(len(tmp))
		if limit > desiredLength-offset {
			limit = desiredLength - offset
		}
		list, err = client.Dcim.DcimManufacturersList(params, nil)
		if err != nil {
			return diag.FromErr(err)
		}
		tmp = append(tmp, list.Payload.Results...)
	}

	j, _ := json.Marshal(tmp)

	if err = d.Set("json", string(j)); err != nil {
		return diag.FromErr(err)
	}
	d.SetId("NetboxJSONDcimManufacturersList")

	return nil
}
