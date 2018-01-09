package subnets

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"

)

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the floating IP attributes you want to see returned. SortKey allows you to
// sort by a particular network attribute. SortDir sets the direction, and is
// either `asc' or `desc'. Marker and Limit are used for pagination.

type ListOpts struct {
	// ID is the unique identifier for the vpc.
	VPC_ID string `json:"vpc_id"`

	//Specifies the number of records returned on each page.
	//The value ranges from 0 to intmax.
	Limit        int    `q:"limit"`

	//Specifies the resource ID of pagination query.
	//If the parameter is left blank, only resources on the first page are queried.
	Marker       string `q:"marker"`
}

// List returns collection of
// subnets. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
//
// Default policy settings return only those subnets that are owned by the
// tenant who submits the request, unless an admin user submits the request.
func List(c *gophercloud.ServiceClient, opts ListOpts) pagination.Pager {
	q, err := gophercloud.BuildQueryString(&opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}
	u := rootURL(c) + q.String()
	return pagination.NewPager(c, u, func(r pagination.PageResult) pagination.Page {
		return SubnetPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToSubnetCreateMap() (map[string]interface{}, error)
}

// CreateOpts contains all the values needed to create a new subnets. There are
// no required values.
type CreateOpts struct {
	Name string `json:"name,omitempty"`
	CIDR string `json:"cidr,omitempty"`
	GatewayIP string `json:"gateway_ip,omitempty"`
	EnableDHCP string `json:"dhcp_enable,omitempty"`
	PRIMARY_DNS string `json:"primary_dns,omitempty"`
	SECONDARY_DNS string `json:"secondary_dns,omitempty"`
	AvailabilityZone string `json:"availability_zone,omitempty"`
	VPC_ID string `json:"vpc_id,omitempty"`

}

// ToSubnetCreateMap builds a create request body from CreateOpts.
func (opts CreateOpts) ToSubnetCreateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "subnet")
}

// Create accepts a CreateOpts struct and uses the values to create a new
// logical subnets. When it is created, the subnets does not have an internal
// interface - it is not associated to any subnet.
//
func Create(c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToSubnetCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	reqOpt := &gophercloud.RequestOpts{OkCodes: []int{200}}
	_, r.Err = c.Post(rootURL(c), b, &r.Body, reqOpt)
	return
}

// Get retrieves a particular subnets based on its unique ID.
func Get(c *gophercloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(resourceURL(c, id), &r.Body, nil)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	//ToSubnetUpdateMap() (map[string]interface{}, error)
	ToSubnetUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts contains the values used when updating a subnets.
type UpdateOpts struct {
	Name string `json:"name,omitempty"`
	EnableDHCP bool `json:"dhcp_enable,omitempty"`
	PRIMARY_DNS string `json:"primary_dns,omitempty"`
	SECONDARY_DNS string `json:"secondary_dns,omitempty"`

}


// ToSubnetUpdateMap builds an update body based on UpdateOpts.
func (opts UpdateOpts) ToSubnetUpdateMap() (map[string]interface{}, error) {
	return gophercloud.BuildRequestBody(opts, "subnet")
}

// Update allows subnets to be updated. You can update the name, administrative
// state, and the external gateway.
func Update(c *gophercloud.ServiceClient, vpcid string, id string,opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToSubnetUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(updateURL(c, vpcid, id), b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Delete will permanently delete a particular subnets based on its unique ID.
func Delete(c *gophercloud.ServiceClient, vpcid string, id string) (r DeleteResult) {
	_, r.Err = c.Delete(updateURL(c, vpcid, id), nil)
	return
}

