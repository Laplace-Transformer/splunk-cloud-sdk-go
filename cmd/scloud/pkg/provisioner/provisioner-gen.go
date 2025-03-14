// Package provisioner -- generated by scloudgen
// !! DO NOT EDIT !!
//
package provisioner

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/Laplace-Transformer/splunk-cloud-sdk-go/cmd/scloud/auth"
	"github.com/Laplace-Transformer/splunk-cloud-sdk-go/cmd/scloud/flags"
	"github.com/Laplace-Transformer/splunk-cloud-sdk-go/cmd/scloud/jsonx"
	model "github.com/Laplace-Transformer/splunk-cloud-sdk-go/services/provisioner"
)

// CreateInvite Creates an invitation for a person to join the tenant using their email address.
func CreateInvite(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var commentDefault string
	comment := &commentDefault
	err = flags.ParseFlag(cmd.Flags(), "comment", &comment)
	if err != nil {
		return fmt.Errorf(`error parsing "comment": ` + err.Error())
	}
	var email string
	err = flags.ParseFlag(cmd.Flags(), "email", &email)
	if err != nil {
		return fmt.Errorf(`error parsing "email": ` + err.Error())
	}
	var groups []string
	err = flags.ParseFlag(cmd.Flags(), "groups", &groups)
	if err != nil {
		return fmt.Errorf(`error parsing "groups": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.InviteBody{

		Comment: comment,
		Email:   email,
		Groups:  groups,
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.ProvisionerService.CreateInvite(generated_request_body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// DeleteInvite Removes an invitation in the given tenant.
func DeleteInvite(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var inviteId string
	err = flags.ParseFlag(cmd.Flags(), "invite-id", &inviteId)
	if err != nil {
		return fmt.Errorf(`error parsing "invite-id": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	err = client.ProvisionerService.DeleteInvite(inviteId)
	if err != nil {
		return err
	}

	return nil
}

// GetInvite Returns an invitation in the given tenant.
func GetInvite(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var inviteId string
	err = flags.ParseFlag(cmd.Flags(), "invite-id", &inviteId)
	if err != nil {
		return fmt.Errorf(`error parsing "invite-id": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.ProvisionerService.GetInvite(inviteId)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// GetTenant Returns a specific tenant.
func GetTenant(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClientSystemTenant()

	if err != nil {
		return err
	}
	// Parse all flags

	var tenantName string
	err = flags.ParseFlag(cmd.Flags(), "tenant-name", &tenantName)
	if err != nil {
		return fmt.Errorf(`error parsing "tenant-name": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.ProvisionerService.GetTenant(tenantName)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListInvites Returns a list of invitations in a given tenant.
func ListInvites(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.ProvisionerService.ListInvites()
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListTenants Returns all tenants that the user can read.
func ListTenants(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClientSystemTenant()

	if err != nil {
		return err
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.ProvisionerService.ListTenants()
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// UpdateInvite Modifies an invitation in the given tenant.
func UpdateInvite(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var action model.UpdateInviteBodyAction
	err = flags.ParseFlag(cmd.Flags(), "action", &action)
	if err != nil {
		return fmt.Errorf(`error parsing "action": ` + err.Error())
	}
	var inviteId string
	err = flags.ParseFlag(cmd.Flags(), "invite-id", &inviteId)
	if err != nil {
		return fmt.Errorf(`error parsing "invite-id": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.UpdateInviteBody{

		Action: action,
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.ProvisionerService.UpdateInvite(inviteId, generated_request_body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}
