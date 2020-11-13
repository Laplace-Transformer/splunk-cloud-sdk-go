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

// ListEventsSummaryQueryParams represents valid query parameters for the ListEventsSummary operation
// For convenience ListEventsSummaryQueryParams can be formed in a single statement, for example:
//     `v := ListEventsSummaryQueryParams{}.SetCount(...).SetEarliest(...).SetField(...).SetLatest(...).SetOffset(...)`
type ListEventsSummaryQueryParams struct {
	// Count : The maximum number of entries to return. Set to 0 to return all available entries.
	Count *int32 `key:"count"`
	// Earliest : The earliest time filter, in absolute time. When specifying an absolute time specify either UNIX time, or UTC in seconds using the ISO-8601 (%FT%T.%Q) format. For example 2019-01-25T13:15:30Z. GMT is the default timezone. You must specify GMT when you specify UTC. Any offset specified is ignored.
	Earliest string `key:"earliest"`
	// Field : A field to return for the result set. You can specify multiple fields of comma-separated values if multiple fields are required.
	Field string `key:"field"`
	// Latest : The latest time filter in absolute time. When specifying an absolute time specify either UNIX time, or UTC in seconds using the ISO-8601 (%FT%T.%Q) format. For example 2019-01-25T13:15:30Z. GMT is the default timezone. You must specify GMT when you specify UTC. Any offset specified is ignored.
	Latest string `key:"latest"`
	// Offset : Index of first item to return.
	Offset *int32 `key:"offset"`
}

func (q ListEventsSummaryQueryParams) SetCount(v int32) ListEventsSummaryQueryParams {
	q.Count = &v
	return q
}

func (q ListEventsSummaryQueryParams) SetEarliest(v string) ListEventsSummaryQueryParams {
	q.Earliest = v
	return q
}

func (q ListEventsSummaryQueryParams) SetField(v string) ListEventsSummaryQueryParams {
	q.Field = v
	return q
}

func (q ListEventsSummaryQueryParams) SetLatest(v string) ListEventsSummaryQueryParams {
	q.Latest = v
	return q
}

func (q ListEventsSummaryQueryParams) SetOffset(v int32) ListEventsSummaryQueryParams {
	q.Offset = &v
	return q
}

// ListFieldsSummaryQueryParams represents valid query parameters for the ListFieldsSummary operation
// For convenience ListFieldsSummaryQueryParams can be formed in a single statement, for example:
//     `v := ListFieldsSummaryQueryParams{}.SetEarliest(...).SetLatest(...)`
type ListFieldsSummaryQueryParams struct {
	// Earliest : The earliest time filter, in absolute time. When specifying an absolute time specify either UNIX time, or UTC in seconds using the ISO-8601 (%FT%T.%Q) format. For example 2019-01-25T13:15:30Z. GMT is the default timezone. You must specify GMT when you specify UTC. Any offset specified is ignored.
	Earliest string `key:"earliest"`
	// Latest : The latest time filter in absolute time. When specifying an absolute time specify either UNIX time, or UTC in seconds using the ISO-8601 (%FT%T.%Q) format. For example 2019-01-25T13:15:30Z. GMT is the default timezone. You must specify GMT when you specify UTC. Any offset specified is ignored.
	Latest string `key:"latest"`
}

func (q ListFieldsSummaryQueryParams) SetEarliest(v string) ListFieldsSummaryQueryParams {
	q.Earliest = v
	return q
}

func (q ListFieldsSummaryQueryParams) SetLatest(v string) ListFieldsSummaryQueryParams {
	q.Latest = v
	return q
}

// ListJobsQueryParams represents valid query parameters for the ListJobs operation
// For convenience ListJobsQueryParams can be formed in a single statement, for example:
//     `v := ListJobsQueryParams{}.SetCount(...).SetFilter(...).SetStatus(...)`
type ListJobsQueryParams struct {
	// Count : The maximum number of jobs that you want to return the status entries for.
	Count *int32 `key:"count"`
	// Filter : Filter the list of jobs by sid. Valid format is  &#x60;sid IN ({comma separated list of SIDs in quotes})&#x60;. A maximum of 50 SIDs can be specified in one query.
	Filter string `key:"filter"`
	// Status : Filter the list of jobs by status. Valid status values are &#39;running&#39;, &#39;done&#39;, &#39;canceled&#39;, or &#39;failed&#39;.
	Status SearchStatus `key:"status"`
}

func (q ListJobsQueryParams) SetCount(v int32) ListJobsQueryParams {
	q.Count = &v
	return q
}

func (q ListJobsQueryParams) SetFilter(v string) ListJobsQueryParams {
	q.Filter = v
	return q
}

func (q ListJobsQueryParams) SetStatus(v SearchStatus) ListJobsQueryParams {
	q.Status = v
	return q
}

// ListPreviewResultsQueryParams represents valid query parameters for the ListPreviewResults operation
// For convenience ListPreviewResultsQueryParams can be formed in a single statement, for example:
//     `v := ListPreviewResultsQueryParams{}.SetCount(...).SetOffset(...)`
type ListPreviewResultsQueryParams struct {
	// Count : The maximum number of entries to return. Set to 0 to return all available entries.
	Count *int32 `key:"count"`
	// Offset : Index of first item to return.
	Offset *int32 `key:"offset"`
}

func (q ListPreviewResultsQueryParams) SetCount(v int32) ListPreviewResultsQueryParams {
	q.Count = &v
	return q
}

func (q ListPreviewResultsQueryParams) SetOffset(v int32) ListPreviewResultsQueryParams {
	q.Offset = &v
	return q
}

// ListResultsQueryParams represents valid query parameters for the ListResults operation
// For convenience ListResultsQueryParams can be formed in a single statement, for example:
//     `v := ListResultsQueryParams{}.SetCount(...).SetField(...).SetOffset(...)`
type ListResultsQueryParams struct {
	// Count : The maximum number of entries to return. Set to 0 to return all available entries.
	Count *int32 `key:"count"`
	// Field : A field to return for the result set. You can specify multiple fields of comma-separated values if multiple fields are required.
	Field string `key:"field"`
	// Offset : Index of first item to return.
	Offset *int32 `key:"offset"`
}

func (q ListResultsQueryParams) SetCount(v int32) ListResultsQueryParams {
	q.Count = &v
	return q
}

func (q ListResultsQueryParams) SetField(v string) ListResultsQueryParams {
	q.Field = v
	return q
}

func (q ListResultsQueryParams) SetOffset(v int32) ListResultsQueryParams {
	q.Offset = &v
	return q
}
