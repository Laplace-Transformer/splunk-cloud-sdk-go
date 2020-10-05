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

package ml

import (
	"net/http"
)

// ServicerGenerated represents the interface for implementing all endpoints for this service
type ServicerGenerated interface {
	/*
		CreateWorkflow - Creates a workflow configuration.
		Parameters:
			workflow: Workflow configuration to be created.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateWorkflow(workflow Workflow, resp ...*http.Response) (*Workflow, error)
	/*
		CreateWorkflowBuild - Creates a workflow build.
		Parameters:
			id: The workflow ID.
			workflowBuild: Input data used to build the workflow.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateWorkflowBuild(id string, workflowBuild WorkflowBuild, resp ...*http.Response) (*WorkflowBuild, error)
	/*
		CreateWorkflowDeployment - Creates a workflow deployment.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			workflowDeployment: Input data used to build the workflow deployment.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateWorkflowDeployment(id string, buildId string, workflowDeployment WorkflowDeployment, resp ...*http.Response) (*WorkflowDeployment, error)
	/*
		CreateWorkflowInference - Creates a workflow inference request.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			deploymentId: The workflow deployment ID.
			workflowInference: Input data to the inference request.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateWorkflowInference(id string, buildId string, deploymentId string, workflowInference WorkflowInference, resp ...*http.Response) (*WorkflowInference, error)
	/*
		CreateWorkflowRun - Creates a workflow run.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			workflowRun: Input data used to build the workflow.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateWorkflowRun(id string, buildId string, workflowRun WorkflowRun, resp ...*http.Response) (*WorkflowRun, error)
	/*
		CreateWorkflowStreamDeployment - Creates a workflow streaming deployment.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			workflowStreamDeployment: Configuration for the workflow streaming deployment.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateWorkflowStreamDeployment(id string, buildId string, workflowStreamDeployment WorkflowStreamDeployment, resp ...*http.Response) (*WorkflowStreamDeployment, error)
	/*
		DeleteWorkflow - Removes a workflow configuration.
		Parameters:
			id: The workflow ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteWorkflow(id string, resp ...*http.Response) error
	/*
		DeleteWorkflowBuild - Removes a workflow build.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteWorkflowBuild(id string, buildId string, resp ...*http.Response) error
	/*
		DeleteWorkflowDeployment - Removes a workflow deployment.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			deploymentId: The workflow deployment ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteWorkflowDeployment(id string, buildId string, deploymentId string, resp ...*http.Response) error
	/*
		DeleteWorkflowRun - Removes a workflow run.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			runId: The workflow run ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteWorkflowRun(id string, buildId string, runId string, resp ...*http.Response) error
	/*
		DeleteWorkflowStreamDeployment - Removes a workflow streaming deployment.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			streamDeploymentId: The workflow streaming deployment ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteWorkflowStreamDeployment(id string, buildId string, streamDeploymentId string, resp ...*http.Response) error
	/*
		GetWorkflow - Returns a workflow configuration.
		Parameters:
			id: The workflow ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetWorkflow(id string, resp ...*http.Response) (*Workflow, error)
	/*
		GetWorkflowBuild - Returns the status of a workflow build.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetWorkflowBuild(id string, buildId string, resp ...*http.Response) (*WorkflowBuild, error)
	/*
		GetWorkflowBuildError - Returns a list of workflow errors.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetWorkflowBuildError(id string, buildId string, resp ...*http.Response) (*WorkflowBuildError, error)
	/*
		GetWorkflowBuildLog - Returns the logs from a workflow build.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetWorkflowBuildLog(id string, buildId string, resp ...*http.Response) (*WorkflowBuildLog, error)
	/*
		GetWorkflowDeployment - Returns the status of a workflow deployment.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			deploymentId: The workflow deployment ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetWorkflowDeployment(id string, buildId string, deploymentId string, resp ...*http.Response) (*WorkflowDeployment, error)
	/*
		GetWorkflowDeploymentError - Returns a list of workflow deployment errors.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			deploymentId: The workflow deployment ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetWorkflowDeploymentError(id string, buildId string, deploymentId string, resp ...*http.Response) (*WorkflowDeploymentError, error)
	/*
		GetWorkflowDeploymentLog - Returns the logs from a workflow deployment.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			deploymentId: The workflow deployment ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetWorkflowDeploymentLog(id string, buildId string, deploymentId string, resp ...*http.Response) (*WorkflowDeploymentLog, error)
	/*
		GetWorkflowRun - Returns the status of a workflow run.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			runId: The workflow run ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetWorkflowRun(id string, buildId string, runId string, resp ...*http.Response) (*WorkflowRun, error)
	/*
		GetWorkflowRunError - Returns the errors for a workflow run.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			runId: The workflow run ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetWorkflowRunError(id string, buildId string, runId string, resp ...*http.Response) (*WorkflowRunError, error)
	/*
		GetWorkflowRunLog - Returns the logs for a workflow run.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			runId: The workflow run ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetWorkflowRunLog(id string, buildId string, runId string, resp ...*http.Response) (*WorkflowRunLog, error)
	/*
		GetWorkflowStreamDeployment - Returns the status of a workflow streaming deployment.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			streamDeploymentId: The workflow streaming deployment ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetWorkflowStreamDeployment(id string, buildId string, streamDeploymentId string, resp ...*http.Response) (*WorkflowStreamDeployment, error)
	/*
		ListWorkflowBuilds - Returns a list of workflow builds.
		Parameters:
			id: The workflow ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListWorkflowBuilds(id string, resp ...*http.Response) ([]WorkflowBuild, error)
	/*
		ListWorkflowDeployments - Returns a list of workflow deployments.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListWorkflowDeployments(id string, buildId string, resp ...*http.Response) ([]WorkflowDeployment, error)
	/*
		ListWorkflowRuns - Returns a list of workflow runs.
		Parameters:
			id: The workflow ID.
			buildId: The workflow build ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListWorkflowRuns(id string, buildId string, resp ...*http.Response) ([]WorkflowRun, error)
	/*
		ListWorkflows - Returns a list of workflow configurations.
		Parameters:
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListWorkflows(resp ...*http.Response) ([]WorkflowsGetResponse, error)
}
