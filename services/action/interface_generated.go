/*
 * Copyright © 2019 Splunk, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"): you may
 * not use this file except in compliance with the License. You may obtain
 * a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 */

// Code generated by gen_interface.go. DO NOT EDIT.

package action

import (
	"net/http"
)

// ServicerGenerated represents the interface for implementing all endpoints for this service
type ServicerGenerated interface {
	/*
		CreateAction - Creates an action template.
		Parameters:
			action: The action template to create.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateAction(action Action, resp ...*http.Response) (*Action, error)
	/*
		DeleteAction - Removes an action template.
		Parameters:
			actionName: The name of the action as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteAction(actionName string, resp ...*http.Response) error
	/*
		GetAction - Returns a specific action template.
		Parameters:
			actionName: The name of the action as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetAction(actionName string, resp ...*http.Response) (*Action, error)
	/*
		GetActionStatus - Returns the status of an action that was invoked. The status is available for 4 days after the last status change.
		Parameters:
			actionName: The name of the action as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.
			statusId: The ID of the action status.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetActionStatus(actionName string, statusId string, resp ...*http.Response) (*ActionResult, error)
	/*
		GetActionStatusDetails - Returns the status details of the invoked email action. The status is available for 4 days after the last status change.
		Parameters:
			actionName: The name of the action as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.
			statusId: The ID of the action status.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetActionStatusDetails(actionName string, statusId string, resp ...*http.Response) ([]ActionResultEmailDetail, error)
	/*
		GetPublicWebhookKeys - Returns an array of one or two webhook keys. The first key is active. The second key, if present, is expired.
		Parameters:
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetPublicWebhookKeys(resp ...*http.Response) ([]PublicWebhookKey, error)
	/*
		ListActions - Returns the list of action templates.
		Parameters:
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListActions(resp ...*http.Response) ([]Action, error)
	/*
		TriggerAction - Invokes an action.
		Parameters:
			actionName: The name of the action as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.
			triggerEvent: The action payload, which should include values for any templated fields.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	TriggerAction(actionName string, triggerEvent TriggerEvent, resp ...*http.Response) error
	/*
		UpdateAction - Modifies an action template.
		Parameters:
			actionName: The name of the action as one or more identifier strings separated by periods. Each identifier string consists of lowercase letters, digits, and underscores, and cannot start with a digit.
			actionMutable: Updates to the action template.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UpdateAction(actionName string, actionMutable ActionMutable, resp ...*http.Response) (*Action, error)
}
