package action

import (
	"github.com/Laplace-Transformer/splunk-cloud-sdk-go/cmd/scloud/auth"
	model "github.com/Laplace-Transformer/splunk-cloud-sdk-go/services/action"
)

// TriggerActionOverride
func TriggerActionOverride(actionName string, body model.TriggerEvent) (*model.TriggerResponse, error) {

	client, err := auth.GetClient()
	if err != nil {
		return nil, err
	}

	return client.ActionService.TriggerActionWithStatus(actionName, body)
}
