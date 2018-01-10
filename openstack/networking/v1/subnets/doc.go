/*
Package subnets enables management and retrieval of Subnets from the Open Telekom Cloud
Subnet service.

Example to List Subnets

	listOpts := subnets.ListOpts{VPC_ID:"d4f2c817-d5df-4a66-994a-6571312b470e"}
	allSubnets, err := subnets.List(client, listOpts)
	if err != nil {
		panic(err)
	}

	for _, subnet := range allSubnets {
		fmt.Printf("%+v\n", subnet)
	}

Example to Create a Subnet

	createOpts := subnets.CreateOpts{
		Name:"test_subnets",
		CIDR:"192.168.0.0/16",
		GatewayIP:"192.168.0.1",
		AvailabilityZone:"eu-de-02",
		VPC_ID:"3b9740a0-b44d-48f0-84ee-42eb166e54f7"

	}

	subnet, err := subnets.Create(client, createOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Update a Subnet

	VPC_ID := "78d27109-e714-4591-9724-8162703a88d1",
	ID := "ff832470-95ce-4694-a57d-b4447ca0680a"

	updateOpts := vpcs.UpdateOpts{
		Name:         "update_subnet"
	}

	subnet, err := subnets.Update(client, VPC_ID, ID, updateOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Delete a Subnet

	VPC_ID := "78d27109-e714-4591-9724-8162703a88d1",
	ID := "ff832470-95ce-4694-a57d-b4447ca0680a"

	err := subnets.Delete(client, VPC_ID, ID).ExtractErr()
	if err != nil {
		panic(err)
	}
*/
package subnets
