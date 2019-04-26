package rancher2

import (
	"github.com/hashicorp/terraform/helper/schema"
	managementClient "github.com/rancher/types/client/management/v3"
)

// Flatteners

func flattenAuthConfigOpenLdap(d *schema.ResourceData, in *managementClient.LdapConfig) error {
	d.SetId(AuthConfigOpenLdapName)

	err := d.Set("name", AuthConfigOpenLdapName)
	if err != nil {
		return err
	}
	err = d.Set("type", managementClient.OpenLdapConfigType)
	if err != nil {
		return err
	}

	err = flattenAuthConfigLdap(d, in)
	if err != nil {
		return err
	}

	return nil
}

// Expanders

func expandAuthConfigOpenLdap(in *schema.ResourceData) (*managementClient.LdapConfig, error) {
	obj, err := expandAuthConfigLdap(in)
	if err != nil {
		return nil, err
	}

	obj.Name = AuthConfigOpenLdapName
	obj.Type = managementClient.OpenLdapConfigType

	return obj, nil
}