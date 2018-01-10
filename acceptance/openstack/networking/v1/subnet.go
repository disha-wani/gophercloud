package v1

import (
	"testing"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/gophercloud/gophercloud/openstack/networking/v1/subnets"
)

func CreateSubnet(t *testing.T, client *gophercloud.ServiceClient) (*subnets.Subnet, error) {

	subnetName := tools.RandomString("ACPTTEST-", 8)

	createOpts := subnets.CreateOpts{
		Name: subnetName,
		CIDR: "192.168.0.0/16",
		GatewayIP: "192.168.0.1",
		AvailabilityZone: "eu-de-02",
		VPC_ID: "5468ee72-54d4-413f-bd36-bea084c5b2e9",
	}

	t.Logf("Attempting to create subnet: %s", subnetName)

	subnet, err := subnets.Create(client, createOpts).Extract()
	if err != nil {
		return subnet, err
	}
	t.Logf("Created subnet: %s", subnet)

	return subnet, nil
}

func DeleteSubnet(t *testing.T, client *gophercloud.ServiceClient, vpcID string, id string) {
	t.Logf("Attempting to delete subnet: %s", id)

	err := subnets.Delete(client, vpcID, id).ExtractErr()
	if err != nil {
		t.Fatalf("Error deleting subnet: %v", err)
	}

	t.Logf("Deleted subnet: %s", id)
}