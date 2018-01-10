package v1

import (
	"github.com/gophercloud/gophercloud/acceptance/clients"
	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/gophercloud/gophercloud/openstack/networking/v1/subnets"
	"testing"
)

func TestSubnetList(t *testing.T) {
	client, err := clients.NewSubnetV1Client()
	if err != nil {
		t.Fatalf("Unable to create a subnet : %v", err)
	}
	allPages, err := subnets.List(client, subnets.ListOpts{})
	tools.PrintResource(t, allPages)

}

func TestSubnetsCRUD(t *testing.T) {
	client, err := clients.NewSubnetV1Client()
	if err != nil {
		t.Fatalf("Unable to create a subnet : %v", err)
	}

	// Create a subnet
	subnet, err := CreateSubnet(t, client)
	if err != nil {
		t.Fatalf("Unable to create subnet: %v", err)
	}

	// Delete a subnet
	defer DeleteSubnet(t, client,subnet.VPC_ID, subnet.ID)
	tools.PrintResource(t, subnet)

	// Update a subnet
	newName := tools.RandomString("ACPTTEST-", 8)
	updateOpts := &subnets.UpdateOpts{
		Name: newName,
	}
	_, err = subnets.Update(client, subnet.VPC_ID, subnet.ID, updateOpts).Extract()

	// Query a subnet
	newSubnet, err := subnets.Get(client, subnet.ID).Extract()
	if err != nil {
		t.Fatalf("Unable to retrieve subnet: %v", err)
	}

	tools.PrintResource(t, newSubnet)
}

