/*
 * Copyright © 2020 Splunk, Inc.
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
 *
 * Data Stream Processing REST API
 *
 * Use the Streams service to perform create, read, update, and delete (CRUD) operations on your data pipeline. The Streams service also has metrics and preview session endpoints and gives you full control over your data pipeline.
 *
 * API version: v3beta1.1 (recommended default)
 * Generated by: OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
 */

package streams

type ActivatePipelineRequest struct {
	// Set to true to activate the latest version of the pipeline. Set to false to use the previously activated version of the pipeline. Defaults to true.
	ActivateLatestVersion *bool `json:"activateLatestVersion,omitempty"`
	// Set to true to allow the pipeline to ignore any unused progress states. In some cases, when a data pipeline is changed, the progress state will be stored for functions that no longer exist, so this must be set to activate a pipeline in this state. Defaults to false.
	AllowNonRestoredState *bool `json:"allowNonRestoredState,omitempty"`
	// Set to true to start reading from the latest input rather than from where the pipeline's previous run left off, which can cause data loss. Defaults to false.
	SkipRestoreState *bool `json:"skipRestoreState,omitempty"`
}

type ArgumentModel struct {
	Name     *string `json:"name,omitempty"`
	Optional *bool   `json:"optional,omitempty"`
	Type     *string `json:"type,omitempty"`
}

type CollectJobRequest struct {
	// The ID of the connection that is assigned to this collect job.
	ConnectionId string `json:"connectionId"`
	// The ID of the connector this collect job uses.
	ConnectorId string `json:"connectorId"`
	// The description of the collect job.
	Description string `json:"description"`
	// The name of the collect job.
	Name string `json:"name"`
	// The key-value pairs of parameters for this collect job. Collect jobs may have some configurations that are required, which all collect jobs must provide values for. For configuration values of type BYTES, the provided values must be Base64 encoded.
	Parameters map[string]interface{} `json:"parameters"`
}

type CollectJobResponse struct {
	ActivePipelinesUsing []map[string]interface{} `json:"activePipelinesUsing,omitempty"`
	ConnectionId         *string                  `json:"connectionId,omitempty"`
	ConnectorId          *string                  `json:"connectorId,omitempty"`
	CreateDate           *int64                   `json:"createDate,omitempty"`
	CreateUserId         *string                  `json:"createUserId,omitempty"`
	Description          *string                  `json:"description,omitempty"`
	Id                   *string                  `json:"id,omitempty"`
	LastUpdateDate       *int64                   `json:"lastUpdateDate,omitempty"`
	LastUpdateUserId     *string                  `json:"lastUpdateUserId,omitempty"`
	Name                 *string                  `json:"name,omitempty"`
	Parameters           map[string]interface{}   `json:"parameters,omitempty"`
	Status               *string                  `json:"status,omitempty"`
	Version              *int64                   `json:"version,omitempty"`
}

type CollectJobStartStopResponse struct {
	Id     *string `json:"id,omitempty"`
	Status *string `json:"status,omitempty"`
}

type ConnectionPatchRequest struct {
	// The key-value pairs of configurations for this connection. Connectors may have some configurations that are required, which all connections must provide values for. For configuration values of type BYTES, the provided values must be Base64 encoded.
	Data map[string]interface{} `json:"data,omitempty"`
	// The description of the connection.
	Description *string `json:"description,omitempty"`
	// The name of the connection.
	Name *string `json:"name,omitempty"`
}

type ConnectionPutRequest struct {
	// The key-value pairs of configurations for this connection. Connectors may have some configurations that are required, which all connections must provide values for. For configuration values of type BYTES, the provided values must be Base64 encoded.
	Data map[string]interface{} `json:"data"`
	// The description of the connection.
	Description string `json:"description"`
	// The name of the connection.
	Name string `json:"name"`
}

type ConnectionRequest struct {
	// The ID of the parent connector.
	ConnectorId string `json:"connectorId"`
	// The key-value pairs of configurations for this connection. Connectors may have some configurations that are required, which all connections must provide values for. For configuration values of type BYTES, the provided values must be Base64 encoded.
	Data map[string]interface{} `json:"data"`
	// The description of the connection.
	Description string `json:"description"`
	// The name of the connection.
	Name string `json:"name"`
}

type ConnectionResponse struct {
	ActivePipelinesUsing []map[string]interface{}    `json:"activePipelinesUsing,omitempty"`
	ConnectorId          *string                     `json:"connectorId,omitempty"`
	ConnectorName        *string                     `json:"connectorName,omitempty"`
	CreateDate           *int64                      `json:"createDate,omitempty"`
	CreateUserId         *string                     `json:"createUserId,omitempty"`
	Id                   *string                     `json:"id,omitempty"`
	LastUpdateDate       *int64                      `json:"lastUpdateDate,omitempty"`
	LastUpdateUserId     *string                     `json:"lastUpdateUserId,omitempty"`
	Versions             []ConnectionVersionResponse `json:"versions,omitempty"`
}

type ConnectionSaveResponse struct {
	ConnectorId  *string                `json:"connectorId,omitempty"`
	CreateDate   *int64                 `json:"createDate,omitempty"`
	CreateUserId *string                `json:"createUserId,omitempty"`
	Data         map[string]interface{} `json:"data,omitempty"`
	Description  *string                `json:"description,omitempty"`
	Id           *string                `json:"id,omitempty"`
	Name         *string                `json:"name,omitempty"`
	Version      *int64                 `json:"version,omitempty"`
}

type ConnectionVersionResponse struct {
	CreateDate   *int64                 `json:"createDate,omitempty"`
	CreateUserId *string                `json:"createUserId,omitempty"`
	Data         map[string]interface{} `json:"data,omitempty"`
	Description  *string                `json:"description,omitempty"`
	Name         *string                `json:"name,omitempty"`
	Version      *int64                 `json:"version,omitempty"`
}

type ConnectorResponse struct {
	Config        map[string]interface{}   `json:"config,omitempty"`
	ConnectorType *string                  `json:"connectorType,omitempty"`
	Description   *string                  `json:"description,omitempty"`
	Functions     []map[string]interface{} `json:"functions,omitempty"`
	Hidden        *bool                    `json:"hidden,omitempty"`
	Id            *string                  `json:"id,omitempty"`
	Name          *string                  `json:"name,omitempty"`
	PanelUrl      *string                  `json:"panelUrl,omitempty"`
	Tag           *string                  `json:"tag,omitempty"`
}

type DeactivatePipelineRequest struct {
	// Set to true to skip saving the state of a deactivated pipeline. When the pipeline is later activated, it will start with the newest data and skip any data that arrived after this deactivation, which can cause data loss. Defaults to false.
	SkipSavepoint *bool `json:"skipSavepoint,omitempty"`
}

type DecompileRequest struct {
	Upl Pipeline `json:"upl"`
}

type DecompileResponse struct {
	Spl *string `json:"spl,omitempty"`
}

type ErrorResponse struct {
	Code     *string                `json:"code,omitempty"`
	Details  []MapOfstringAndstring `json:"details,omitempty"`
	Message  *string                `json:"message,omitempty"`
	MoreInfo *string                `json:"moreInfo,omitempty"`
}

type FilesMetaDataResponse struct {
	Files []UploadFile `json:"files,omitempty"`
}

type FunctionModel struct {
	Arguments  []ArgumentModel        `json:"arguments,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Op         *string                `json:"op,omitempty"`
	Outputs    []ArgumentModel        `json:"outputs,omitempty"`
	ResolvedId *string                `json:"resolvedId,omitempty"`
}

type GetInputSchemaRequest struct {
	// The function ID.
	NodeUuid string `json:"nodeUuid"`
	// The name of the input port.
	TargetPortName string   `json:"targetPortName"`
	UplJson        Pipeline `json:"uplJson"`
}

type GetOutputSchemaRequest struct {
	UplJson Pipeline `json:"uplJson"`
	// The function ID. If omitted, returns the output schema for all functions.
	NodeUuid *string `json:"nodeUuid,omitempty"`
	// The name of the output port. Deprecated.
	SourcePortName *string `json:"sourcePortName,omitempty"`
}

type LookupTableResponse struct {
	Data  []map[string]interface{} `json:"data,omitempty"`
	Types map[string]string        `json:"types,omitempty"`
}

type MapOfstringAndobject map[string]interface{}

type MapOfstringAndstring map[string]interface{}

type MetricsResponse struct {
	Nodes map[string]NodeMetrics `json:"nodes,omitempty"`
}

type NodeMetrics struct {
	Metrics map[string]interface{} `json:"metrics,omitempty"`
}

type PaginatedResponseOfConnectionResponse struct {
	Items []ConnectionResponse `json:"items,omitempty"`
	Total *int64               `json:"total,omitempty"`
}

type PaginatedResponseOfConnectorResponse struct {
	Items []ConnectorResponse `json:"items,omitempty"`
	Total *int64              `json:"total,omitempty"`
}

type PaginatedResponseOfPipelineJobStatus struct {
	Items []PipelineJobStatus `json:"items,omitempty"`
	Total *int64              `json:"total,omitempty"`
}

type PaginatedResponseOfPipelineResponse struct {
	Items []PipelineResponse `json:"items,omitempty"`
	Total *int64             `json:"total,omitempty"`
}

type PaginatedResponseOfPlugin struct {
	Items []Plugin `json:"items,omitempty"`
	Total *int64   `json:"total,omitempty"`
}

type PaginatedResponseOfTemplateResponse struct {
	Items []TemplateResponse `json:"items,omitempty"`
	Total *int64             `json:"total,omitempty"`
}

type Pipeline struct {
	// A list of links between the output of one pipeline function and the input of another pipeline function.
	Edges []PipelineEdge `json:"edges"`
	// The functions in your entire pipeline, including each function's operations, attributes, and properties.
	Nodes []PipelineNode `json:"nodes"`
	// Optional metadata specified by client. Not to be interpreted by the platform
	Attributes    map[string]interface{} `json:"attributes,omitempty"`
	MigrationInfo *PipelineMigrationInfo `json:"migrationInfo,omitempty"`
}

type PipelineEdge struct {
	SourceNode *string `json:"sourceNode,omitempty"`
	SourcePort *string `json:"sourcePort,omitempty"`
	TargetNode *string `json:"targetNode,omitempty"`
	TargetPort *string `json:"targetPort,omitempty"`
}

type PipelineJobStatus struct {
	JobId      *string `json:"jobId,omitempty"`
	JobStatus  *string `json:"jobStatus,omitempty"`
	PipelineId *string `json:"pipelineId,omitempty"`
}

type PipelineMigrationInfo struct {
	CanMigrate       *bool  `json:"canMigrate,omitempty"`
	CreateApiVersion *int32 `json:"createApiVersion,omitempty"`
	TargetApiVersion *int32 `json:"targetApiVersion,omitempty"`
}

type PipelineNode struct {
	// A unique identifier for the function.
	Id string `json:"id"`
	// The operation name.
	Op string `json:"op"`
	// Function arguments keyed by argument name.
	Arguments map[string]interface{} `json:"arguments,omitempty"`
	// For internal use only. This field is ignored.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// The identifier describing the operation name and required argument types. This field is ignored.
	ResolvedId *string `json:"resolvedId,omitempty"`
	// If present, creates a named statement when decompiling to SPL2. For example, \"$events = | from splunk_firehose();\".
	StatementName *string `json:"statementName,omitempty"`
}

type PipelinePatchRequest struct {
	// Set to true to bypass initial pipeline validation upon creation. The pipeline still needs to be validated before activation. Defaults to false.
	BypassValidation *bool `json:"bypassValidation,omitempty"`
	// The user that created the pipeline. Deprecated.
	CreateUserId *string   `json:"createUserId,omitempty"`
	Data         *Pipeline `json:"data,omitempty"`
	// The description of the pipeline. Defaults to null.
	Description *string `json:"description,omitempty"`
	// The name of the pipeline.
	Name *string `json:"name,omitempty"`
}

type PipelineReactivateResponse struct {
	CurrentlyActiveVersion     *int64                                                `json:"currentlyActiveVersion,omitempty"`
	PipelineId                 *string                                               `json:"pipelineId,omitempty"`
	PipelineReactivationStatus *PipelineReactivateResponsePipelineReactivationStatus `json:"pipelineReactivationStatus,omitempty"`
}

type PipelineReactivateResponsePipelineReactivationStatus string

// List of PipelineReactivateResponsePipelineReactivationStatus
const (
	PipelineReactivateResponsePipelineReactivationStatusActivated                        PipelineReactivateResponsePipelineReactivationStatus = "activated"
	PipelineReactivateResponsePipelineReactivationStatusCurrentVersionInvalid            PipelineReactivateResponsePipelineReactivationStatus = "currentVersionInvalid"
	PipelineReactivateResponsePipelineReactivationStatusFailedToDeactivateCurrentVersion PipelineReactivateResponsePipelineReactivationStatus = "failedToDeactivateCurrentVersion"
	PipelineReactivateResponsePipelineReactivationStatusNotActivated                     PipelineReactivateResponsePipelineReactivationStatus = "notActivated"
	PipelineReactivateResponsePipelineReactivationStatusRolledBack                       PipelineReactivateResponsePipelineReactivationStatus = "rolledBack"
	PipelineReactivateResponsePipelineReactivationStatusRolledBackError                  PipelineReactivateResponsePipelineReactivationStatus = "rolledBackError"
)

type PipelineRequest struct {
	Data Pipeline `json:"data"`
	// The name of the pipeline.
	Name string `json:"name"`
	// Set to true to bypass initial pipeline validation upon creation. The pipeline still needs to be validated before activation. Defaults to false.
	BypassValidation *bool `json:"bypassValidation,omitempty"`
	// The description of the pipeline. Defaults to null.
	Description *string `json:"description,omitempty"`
}

type PipelineResponse struct {
	ActivatedDate            *int64                  `json:"activatedDate,omitempty"`
	ActivatedUserId          *string                 `json:"activatedUserId,omitempty"`
	ActivatedVersion         *int64                  `json:"activatedVersion,omitempty"`
	CreateDate               *int64                  `json:"createDate,omitempty"`
	CreateUserId             *string                 `json:"createUserId,omitempty"`
	CurrentVersion           *int64                  `json:"currentVersion,omitempty"`
	Data                     *Pipeline               `json:"data,omitempty"`
	Description              *string                 `json:"description,omitempty"`
	Id                       *string                 `json:"id,omitempty"`
	LastUpdateDate           *int64                  `json:"lastUpdateDate,omitempty"`
	LastUpdateUserId         *string                 `json:"lastUpdateUserId,omitempty"`
	Name                     *string                 `json:"name,omitempty"`
	Status                   *PipelineResponseStatus `json:"status,omitempty"`
	StatusMessage            *string                 `json:"statusMessage,omitempty"`
	StreamingConfigurationId *int64                  `json:"streamingConfigurationId,omitempty"`
	TenantId                 *string                 `json:"tenantId,omitempty"`
	ValidationMessages       []string                `json:"validationMessages,omitempty"`
	Version                  *int64                  `json:"version,omitempty"`
}

type PipelineResponseStatus string

// List of PipelineResponseStatus
const (
	PipelineResponseStatusCreated    PipelineResponseStatus = "CREATED"
	PipelineResponseStatusActivated  PipelineResponseStatus = "ACTIVATED"
	PipelineResponseStatusFailed     PipelineResponseStatus = "FAILED"
	PipelineResponseStatusRestarting PipelineResponseStatus = "RESTARTING"
	PipelineResponseStatusFinished   PipelineResponseStatus = "FINISHED"
)

type Plugin struct {
	Description      *string `json:"description,omitempty"`
	IsDeleted        *bool   `json:"isDeleted,omitempty"`
	LastUpdateDate   *int64  `json:"lastUpdateDate,omitempty"`
	LastUpdateUserId *string `json:"lastUpdateUserId,omitempty"`
	Name             *string `json:"name,omitempty"`
	PluginId         *string `json:"pluginId,omitempty"`
	TenantId         *string `json:"tenantId,omitempty"`
}

type PluginPatchRequest struct {
	// Plugin description
	Description *string `json:"description,omitempty"`
	// Plugin name
	Name *string `json:"name,omitempty"`
}

type PluginRequest struct {
	// Plugin description
	Description string `json:"description"`
	// Plugin name
	Name string `json:"name"`
}

type PreviewData struct {
	CurrentNumberOfRecords *int32                 `json:"currentNumberOfRecords,omitempty"`
	Nodes                  map[string]PreviewNode `json:"nodes,omitempty"`
	PreviewId              *string                `json:"previewId,omitempty"`
	RecordsPerPipeline     *int32                 `json:"recordsPerPipeline,omitempty"`
	RuleMetrics            []RuleMetrics          `json:"ruleMetrics,omitempty"`
	TenantId               *string                `json:"tenantId,omitempty"`
}

type PreviewNode struct {
	NodeName *string                  `json:"nodeName,omitempty"`
	Records  []map[string]interface{} `json:"records,omitempty"`
}

type PreviewSessionStartRequest struct {
	Upl Pipeline `json:"upl"`
	// The maximum number of events per function. Defaults to 100.
	RecordsLimit *int32 `json:"recordsLimit,omitempty"`
	// The maximum number of events per pipeline. Defaults to 10000.
	RecordsPerPipeline *int32 `json:"recordsPerPipeline,omitempty"`
	// The maximum lifetime of a session, in milliseconds. Defaults to 300,000.
	SessionLifetimeMs *int64 `json:"sessionLifetimeMs,omitempty"`
}

type PreviewStartResponse struct {
	PreviewId *int64 `json:"previewId,omitempty"`
}

type PreviewState struct {
	ActivatedDate          *int64  `json:"activatedDate,omitempty"`
	CreatedDate            *int64  `json:"createdDate,omitempty"`
	CurrentNumberOfRecords *int32  `json:"currentNumberOfRecords,omitempty"`
	JobId                  *string `json:"jobId,omitempty"`
	PreviewId              *int64  `json:"previewId,omitempty"`
	RecordsPerPipeline     *int32  `json:"recordsPerPipeline,omitempty"`
}

type ReactivatePipelineRequest struct {
	// Set to true to allow the pipeline to ignore any unused progress states. In some cases, when a data pipeline is changed, the progress state will be stored for functions that no longer exist, so this must be set to reactivate a pipeline in this state. Defaults to false.
	AllowNonRestoredState *bool `json:"allowNonRestoredState,omitempty"`
	// Set to true to start reading from the latest input rather than from where the pipeline's previous run left off, which can cause data loss. Defaults to false.
	SkipRestoreState *bool `json:"skipRestoreState,omitempty"`
}

type RegistryModel struct {
	Categories []string        `json:"categories,omitempty"`
	Functions  []FunctionModel `json:"functions,omitempty"`
}

type Response struct {
	// Only set for /activate endpoint
	Activated *string `json:"activated,omitempty"`
	// Only set for /deactivate endpoint
	Deactivated *string `json:"deactivated,omitempty"`
}

// Represents the metrics for a particular rule.
type RuleMetrics struct {
	MatchCount      *int64  `json:"matchCount,omitempty"`
	MatchPercentage *int64  `json:"matchPercentage,omitempty"`
	Name            *string `json:"name,omitempty"`
}

type SplCompileRequest struct {
	// The SPL2 representation of a pipeline or function parameters.
	Spl string `json:"spl"`
	// A boolean flag to indicate whether the pipeline should be validated.
	Validate *bool `json:"validate,omitempty"`
}

type TemplatePatchRequest struct {
	Data *Pipeline `json:"data,omitempty"`
	// Template description
	Description *string `json:"description,omitempty"`
	// Template name
	Name *string `json:"name,omitempty"`
}

type TemplatePutRequest struct {
	Data Pipeline `json:"data"`
	// Template description
	Description string `json:"description"`
	// Template name
	Name string `json:"name"`
}

type TemplateRequest struct {
	Data Pipeline `json:"data"`
	// Template description
	Description string `json:"description"`
	// Template name
	Name string `json:"name"`
}

type TemplateResponse struct {
	CreateDate    *int64    `json:"createDate,omitempty"`
	CreateUserId  *string   `json:"createUserId,omitempty"`
	Data          *Pipeline `json:"data,omitempty"`
	Description   *string   `json:"description,omitempty"`
	Name          *string   `json:"name,omitempty"`
	OwnerTenantId *string   `json:"ownerTenantId,omitempty"`
	TemplateId    *string   `json:"templateId,omitempty"`
	Version       *int64    `json:"version,omitempty"`
}

type UplType struct {
	FieldName  *string   `json:"fieldName,omitempty"`
	Parameters []UplType `json:"parameters,omitempty"`
	Type       *string   `json:"type,omitempty"`
}

type UploadFile struct {
	CreatedDate   *int64  `json:"createdDate,omitempty"`
	CreatedUserId *string `json:"createdUserId,omitempty"`
	Filename      *string `json:"filename,omitempty"`
	Id            *string `json:"id,omitempty"`
	Sha256        *string `json:"sha256,omitempty"`
	TenantId      *string `json:"tenantId,omitempty"`
}

type ValidateRequest struct {
	Upl Pipeline `json:"upl"`
}

type ValidateResponse struct {
	NodeId             *string  `json:"nodeId,omitempty"`
	Spl                *string  `json:"spl,omitempty"`
	Success            *bool    `json:"success,omitempty"`
	ValidationMessages []string `json:"validationMessages,omitempty"`
}
