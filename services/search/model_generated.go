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
 * Splunk Search service
 *
 * Use the Search service in Splunk Cloud Services to dispatch, review, and manage searches and search jobs. You can finalize or cancel jobs, retrieve search results, and request search-related configurations from the Metadata Catalog service in Splunk Cloud Services.
 *
 * API version: v2beta1.1 (recommended default)
 * Generated by: OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
 */

package search

// A fully-constructed delete search job, including read-only fields.
type DeleteSearchJob struct {
	// The index to delete events from.
	Index string `json:"index"`
	// The module to run the delete search job in. The default module is used if module field is empty.
	Module string `json:"module"`
	// The predicate expression that identifies the events to delete from the index. This expression must return true or false. To delete all events from the index, specify \"true\" instead of an expression.
	Predicate string `json:"predicate"`
	// Specifies that the delete search job will contain side effects, with possible security risks.
	AllowSideEffects *bool `json:"allowSideEffects,omitempty"`
	// This field does not apply to delete search jobs and is defaulted to false.
	CollectEventSummary *bool `json:"collectEventSummary,omitempty"`
	// This field does not apply to delete search jobs and is defaulted to false.
	CollectFieldSummary *bool `json:"collectFieldSummary,omitempty"`
	// This field does not apply to delete search jobs and is defaulted to false.
	CollectTimeBuckets *bool `json:"collectTimeBuckets,omitempty"`
	// The time, in GMT, that the search job is finished. Empty if the search job has not completed.
	CompletionTime *string `json:"completionTime,omitempty"`
	// The time, in GMT, that the search job is dispatched.
	DispatchTime *string `json:"dispatchTime,omitempty"`
	// This field does not apply to delete search jobs and is defaulted to false.
	EnablePreview *bool `json:"enablePreview,omitempty"`
	// Specifies whether the Search service should extract all of the available fields in the data, including fields not mentioned in the SPL for the search job. Set to 'false' for better search performance. The 'extractAllFields' parameter is deprecated as of version v3alpha1. Although this parameter continues to function, it might be removed in a future version. Use the 'extractFields' parameter instead.
	ExtractAllFields *bool `json:"extractAllFields,omitempty"`
	// Specifies how the Search service should extract fields. Valid values include 'all', 'none', or 'indexed'. 'all' will extract all fields, 'indexed' will extract only indexed fields, and 'none' will extract only the default fields. This parameter overwrites the value of the 'extractAllFields' parameter. Set to 'none' for better search performance.
	ExtractFields *string `json:"extractFields,omitempty"`
	// The amount of time, in seconds, to run the delete search job before finalizing the search. The maximum value is 3600 seconds (1 hour).
	MaxTime  *int32    `json:"maxTime,omitempty"`
	Messages []Message `json:"messages,omitempty"`
	// The name of the created search job.
	Name *string `json:"name,omitempty"`
	// An estimate of the percent of time remaining before the delete search job completes.
	PercentComplete *int32 `json:"percentComplete,omitempty"`
	// This field does not apply to delete search jobs and is defaulted to false.
	PreviewAvailable *string `json:"previewAvailable,omitempty"`
	// The SPL search string that is generated based on index, module and predicate that are specified.
	Query *string `json:"query,omitempty"`
	// Represents parameters on the search job such as 'earliest' and 'latest'.
	QueryParameters *QueryParameters `json:"queryParameters,omitempty"`
	// This field does not apply to delete search jobs and is set to 0.
	RequiredFreshness *int32 `json:"requiredFreshness,omitempty"`
	// The earliest time speciifed as an absolute value in GMT. The time is computed based on the values you specify for the 'timezone' and 'earliest' queryParameters.
	ResolvedEarliest *string `json:"resolvedEarliest,omitempty"`
	// The latest time specified as an absolute value in GMT. The time is computed based on the values you specify for the 'timezone' and 'earliest' queryParameters.
	ResolvedLatest *string `json:"resolvedLatest,omitempty"`
	// The number of results produced so far by the delete search job that are going to be deleted.
	ResultsAvailable *int32 `json:"resultsAvailable,omitempty"`
	// This field does not apply to delete search jobs and is defaulted to 0.
	ResultsPreviewAvailable *int32 `json:"resultsPreviewAvailable,omitempty"`
	// The ID assigned to the delete search job.
	Sid    *string       `json:"sid,omitempty"`
	Status *SearchStatus `json:"status,omitempty"`
}

// Fields statistics summary model of the events to-date, for search ID (sid).
type FieldsSummary struct {
	// The amount of time, in seconds, that a time bucket spans from the earliest to the latest time.
	Duration *float64 `json:"duration,omitempty"`
	// If specified, the earliest timestamp, in UTC format, of the events to process.
	EarliestTime *string `json:"earliestTime,omitempty"`
	// The total number of events for all fields returned in the time range (earliestTime and latestTime) specified.
	EventCount *int32 `json:"eventCount,omitempty"`
	// A map of the fields in the time range specified.
	Fields map[string]SingleFieldSummary `json:"fields,omitempty"`
	// If specified, the latest timestamp, in UTC format, of the events to process.
	LatestTime *string `json:"latestTime,omitempty"`
}

// The structure of the response body for the preview search results that is returned for the job with the specified search ID (SID). When search is running, it might return incomplete or truncated search results. The isPreviewStable property indicates whether the returned preview results stucture is stable or not. Truncated preview results occur because the number of requested results exceeds the page limit. Follow the 'nextLink' URL to retrieve the next page of results.
type ListPreviewResultsResponse struct {
	IsPreviewStable bool                               `json:"isPreviewStable"`
	Results         []map[string]interface{}           `json:"results"`
	Fields          []ListPreviewResultsResponseFields `json:"fields,omitempty"`
	Messages        []Message                          `json:"messages,omitempty"`
	NextLink        *string                            `json:"nextLink,omitempty"`
	Wait            *string                            `json:"wait,omitempty"`
}

type ListPreviewResultsResponseFields struct {
	Name           string  `json:"name"`
	DataSource     *string `json:"dataSource,omitempty"`
	GroupbyRank    *string `json:"groupbyRank,omitempty"`
	SplitField     *string `json:"splitField,omitempty"`
	SplitValue     *string `json:"splitValue,omitempty"`
	SplitbySpecial *string `json:"splitbySpecial,omitempty"`
	TypeSpecial    *string `json:"typeSpecial,omitempty"`
}

// The structure of the  search results or events metadata that is returned for the job with the specified search ID (SID). When search is running, it might return incomplete or truncated search results. Incomplete search results occur when a search has not completed. Wait until search completes for full result set. Truncated search results occur because the number of requested results exceeds the page limit. Use the 'nextLink' URL to retrieve the next page of results.
type ListSearchResultsResponse struct {
	Results  []map[string]interface{}           `json:"results"`
	Fields   []ListPreviewResultsResponseFields `json:"fields,omitempty"`
	Messages []Message                          `json:"messages,omitempty"`
	NextLink *string                            `json:"nextLink,omitempty"`
	Wait     *string                            `json:"wait,omitempty"`
}

// The message field in search results or search jobs. The types of messages are INFO, DEBUG, FATAL, and ERROR.
type Message struct {
	Text *string      `json:"text,omitempty"`
	Type *MessageType `json:"type,omitempty"`
}

type MessageType string

// List of MessageType
const (
	MessageTypeInfo  MessageType = "INFO"
	MessageTypeDebug MessageType = "DEBUG"
	MessageTypeFatal MessageType = "FATAL"
	MessageTypeError MessageType = "ERROR"
)

// Represents parameters on the search job such as 'earliest' and 'latest'.
type QueryParameters struct {
	// The earliest time, in absolute or relative format, to retrieve events. When specifying an absolute time specify either UNIX time, or UTC in seconds using the ISO-8601 (%FT%T.%Q) format. For example 2019-01-25T13:15:30Z. GMT is the default timezone. You must specify GMT when you specify UTC. Any offset specified is ignored.
	Earliest *string `json:"earliest,omitempty"`
	// The latest time, in absolute or relative format, to retrieve events. When specifying an absolute time specify either UNIX time, or UTC in seconds using the ISO-8601 (%FT%T.%Q) format. For example 2019-01-25T13:15:30Z. GMT is the default timezone. You must specify GMT when you specify UTC. Any offset specified is ignored.
	Latest *string `json:"latest,omitempty"`
	// Relative values for the 'earliest' and 'latest' parameters snap to the unit that you specify. For example, if 'earliest' is set to -d@d, the unit is day. If the 'relativeTimeAnchor' is is set to '2020-10-05T13:15:30Z' then 'resolvedEarliest' is snapped to '2020-10-05T00:00:00Z', which is the day. Hours, minutes, and seconds are dropped. If no 'relativeTimeAnchor' is specified, the default value is set to the time the search job was created.
	RelativeTimeAnchor *string `json:"relativeTimeAnchor,omitempty"`
	// The timezone that relative time specifiers are based off of. Timezone only applies to relative time literals for 'earliest' and 'latest'. If UNIX time or UTC format is used for 'earliest' and 'latest', this field is ignored. For the list of supported timezone formats, see https://docs.splunk.com/Documentation/Splunk/latest/Data/Applytimezoneoffsetstotimestamps#zoneinfo_.28TZ.29_database type: string default: \"GMT\"
	Timezone interface{} `json:"timezone,omitempty"`
}

// A fully-constructed search job, including read-only fields.
type SearchJob struct {
	// The SPL search string.
	Query string `json:"query"`
	// Specifies whether a search that contains commands with side effects (with possible security risks) is allowed to run.
	AllowSideEffects *bool `json:"allowSideEffects,omitempty"`
	// Specified whether a search is allowed to collect events summary during the run time.
	CollectEventSummary *bool `json:"collectEventSummary,omitempty"`
	// Specified whether a search is allowed to collect Fields summary during the run time.
	CollectFieldSummary *bool `json:"collectFieldSummary,omitempty"`
	// Specified whether a search is allowed to collect Timeline Buckets summary during the run time.
	CollectTimeBuckets *bool `json:"collectTimeBuckets,omitempty"`
	// The time, in GMT, that the search job is finished. Empty if the search job has not completed.
	CompletionTime *string `json:"completionTime,omitempty"`
	// The time, in GMT, that the search job is dispatched.
	DispatchTime *string `json:"dispatchTime,omitempty"`
	// Specified whether a search is allowed to collect preview results during the run time.
	EnablePreview *bool `json:"enablePreview,omitempty"`
	// Specifies whether the Search service should extract all of the available fields in the data, including fields not mentioned in the SPL for the search job. Set to 'false' for better search performance. The 'extractAllFields' parameter is deprecated as of version v3alpha1. Although this parameter continues to function, it might be removed in a future version. Use the 'extractFields' parameter instead.
	ExtractAllFields *bool `json:"extractAllFields,omitempty"`
	// Specifies how the Search service should extract fields. Valid values include 'all', 'none', or 'indexed'. 'all' will extract all fields, 'indexed' will extract only indexed fields, and 'none' will extract only the default fields. This parameter overwrites the value of the 'extractAllFields' parameter. Set to 'none' for better search performance.
	ExtractFields *string `json:"extractFields,omitempty"`
	// The number of seconds to run the search before finalizing the search. The maximum value is 3600 seconds (1 hour).
	MaxTime  *int32    `json:"maxTime,omitempty"`
	Messages []Message `json:"messages,omitempty"`
	// The module to run the search in. The default module is used if a module is not specified.
	Module *string `json:"module,omitempty"`
	// The name of the created search job.
	Name *string `json:"name,omitempty"`
	// An estimate of the percent of time remaining before the job completes.
	PercentComplete *int32 `json:"percentComplete,omitempty"`
	// Specifies if preview results for the search job are available. The valid status values are 'unknown', 'true', and 'false'.
	PreviewAvailable *string `json:"previewAvailable,omitempty"`
	// Represents parameters on the search job such as 'earliest' and 'latest'.
	QueryParameters *QueryParameters `json:"queryParameters,omitempty"`
	// Specifies a maximum time interval, in seconds, between identical existing searches. The 'requiredFreshness' parameter is used to determine if an existing search with the same query and the same time boundaries can be reused, instead of running the same search again. Freshness is applied to the resolvedEarliest and resolvedLatest parameters. If an existing search has the same exact criteria as this search and the resolvedEarliest and resolvedLatest values are within the freshness interval, the existing search metadata is returned instead of initiating a new search job. By default, the requiredFreshness parameter is set to 0 which means that the platform does not attempt to use an existing search.
	RequiredFreshness *int32 `json:"requiredFreshness,omitempty"`
	// The earliest time speciifed as an absolute value in GMT. The time is computed based on the values you specify for the 'timezone' and 'earliest' queryParameters.
	ResolvedEarliest *string `json:"resolvedEarliest,omitempty"`
	// The latest time specified as an absolute value in GMT. The time is computed based on the values you specify for the 'timezone' and 'earliest' queryParameters.
	ResolvedLatest *string `json:"resolvedLatest,omitempty"`
	// The number of results produced so far for the search job.
	ResultsAvailable *int32 `json:"resultsAvailable,omitempty"`
	// The number of the preview search results for the job with the specified search ID (SID)
	ResultsPreviewAvailable *int32 `json:"resultsPreviewAvailable,omitempty"`
	// The ID assigned to the search job.
	Sid    *string       `json:"sid,omitempty"`
	Status *SearchStatus `json:"status,omitempty"`
}

// SearchStatus : The current status of the search job. The valid status values are 'running', 'done', 'canceled', and 'failed'.
type SearchStatus string

// List of SearchStatus
const (
	SearchStatusRunning  SearchStatus = "running"
	SearchStatusDone     SearchStatus = "done"
	SearchStatusCanceled SearchStatus = "canceled"
	SearchStatusFailed   SearchStatus = "failed"
)

// Summary of each field.
type SingleFieldSummary struct {
	// The total number of events that contain the field.
	Count *int32 `json:"count,omitempty"`
	// The total number of unique values in the field.
	DistinctCount *int32 `json:"distinctCount,omitempty"`
	// Specifies if the distinctCount is accurate. The isExact property is FALSE when the distinctCount exceeds the maximum count and an exact count is not available.
	IsExact *bool `json:"isExact,omitempty"`
	// The maximum numeric values in the field.
	Max *string `json:"max,omitempty"`
	// The mean (average) for the numeric values in the field.
	Mean *float64 `json:"mean,omitempty"`
	// The minimum numeric values in the field.
	Min *string `json:"min,omitempty"`
	// An array of the values in the field.
	Modes []SingleValueMode `json:"modes,omitempty"`
	// The count of the numeric values in the field.
	NumericCount *int32 `json:"numericCount,omitempty"`
	// Specifies if the field was added or changed by the search.
	Relevant *bool `json:"relevant,omitempty"`
	// The standard deviation for the numeric values in the field.
	Stddev *float64 `json:"stddev,omitempty"`
}

// Events summary in single time bucket.
type SingleTimeBucket struct {
	// Count of available events. Not all events in a bucket are retrievable. Typically this count is capped at 10000.
	AvailableCount *int32   `json:"availableCount,omitempty"`
	Duration       *float64 `json:"duration,omitempty"`
	// The timestamp of the earliest event in the current bucket, in UNIX format. This is the same time as 'earliestTimeStrfTime' in UNIX format.
	EarliestTime *float64 `json:"earliestTime,omitempty"`
	// The timestamp of the earliest event in the current bucket, in UTC format with seconds. For example 2019-01-25T13:15:30Z, which follows the ISO-8601 (%FT%T.%Q) format.
	EarliestTimeStrfTime *string `json:"earliestTimeStrfTime,omitempty"`
	// Specifies if all of the events in the current bucket have been finalized.
	IsFinalized *bool `json:"isFinalized,omitempty"`
	// The total count of the events in the current bucket.
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// Single value summary of the field.
type SingleValueMode struct {
	// The number of occurences that the value appears in a field.
	Count *int32 `json:"count,omitempty"`
	// Specifies if the count is accurate. The isExact property is FALSE when the count exceeds the maximum count and an exact count is not available.
	IsExact *bool `json:"isExact,omitempty"`
	// The value in the field.
	Value *string `json:"value,omitempty"`
}

// A timeline metadata model of the event distribution. The model shows the untransformed events that are read to date for a specific for search ID (sid).
type TimeBucketsSummary struct {
	// Specifies if the events are returned in time order.
	IsTimeCursored *bool              `json:"IsTimeCursored,omitempty"`
	Buckets        []SingleTimeBucket `json:"buckets,omitempty"`
	// Identifies where the cursor is in processing the events. The cursorTime is a timestamp specified in UNIX time.
	CursorTime *float32 `json:"cursorTime,omitempty"`
	// The number of events processed at the cursorTime.
	EventCount *int32 `json:"eventCount,omitempty"`
}

// Update a search job with a status.
type UpdateJob struct {
	// The status to PATCH to an existing search job. The only status values you can PATCH are 'canceled' and 'finalized'. You can PATCH the 'canceled' status only to a search job that is running.
	Status UpdateJobStatus `json:"status"`
}

// UpdateJobStatus : The status to PATCH to an existing search job. The only status values you can PATCH are 'canceled' and 'finalized'. You can PATCH the 'canceled' status only to a search job that is running.
type UpdateJobStatus string

// List of UpdateJobStatus
const (
	UpdateJobStatusCanceled  UpdateJobStatus = "canceled"
	UpdateJobStatusFinalized UpdateJobStatus = "finalized"
)
