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

package search

import (
	"net/http"
)

// ServicerGenerated represents the interface for implementing all endpoints for this service
type ServicerGenerated interface {
	/*
		CreateJob - search service endpoint
		Creates a search job.
		Parameters:
			searchJob
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateJob(searchJob SearchJob, resp ...*http.Response) (*SearchJob, error)
	/*
		DeleteJob - search service endpoint
		Creates a search job that deletes events from an index. The events are deleted from the index in the specified module, based on the search criteria as specified by the predicate.
		Parameters:
			deleteSearchJob
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteJob(deleteSearchJob DeleteSearchJob, resp ...*http.Response) (*DeleteSearchJob, error)
	/*
		GetJob - search service endpoint
		Return the search job with the specified search ID (SID).
		Parameters:
			sid: The search ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetJob(sid string, resp ...*http.Response) (*SearchJob, error)
	/*
		ListEventsSummary - search service endpoint
		Return events summary, for search ID (SID) search.
		Parameters:
			sid: The search ID.
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListEventsSummary(sid string, query *ListEventsSummaryQueryParams, resp ...*http.Response) (*ListSearchResultsResponse, error)
	/*
		ListFieldsSummary - search service endpoint
		Return fields stats summary of the events to-date, for search ID (SID) search.
		Parameters:
			sid: The search ID.
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListFieldsSummary(sid string, query *ListFieldsSummaryQueryParams, resp ...*http.Response) (*FieldsSummary, error)
	/*
		ListJobs - search service endpoint
		Return the matching list of search jobs.
		Parameters:
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListJobs(query *ListJobsQueryParams, resp ...*http.Response) ([]SearchJob, error)
	/*
		ListPreviewResults - search service endpoint
		Return the preview search results for the job with the specified search ID (SID). Can be used when a job is running to return interim results.
		Parameters:
			sid: The search ID.
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListPreviewResults(sid string, query *ListPreviewResultsQueryParams, resp ...*http.Response) (*ListPreviewResultsResponse, error)
	/*
		ListResults - search service endpoint
		Return the search results for the job with the specified search ID (SID).
		Parameters:
			sid: The search ID.
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListResults(sid string, query *ListResultsQueryParams, resp ...*http.Response) (*ListSearchResultsResponse, error)
	/*
		ListTimeBuckets - search service endpoint
		Return event distribution over time of the untransformed events read to-date, for search ID(SID) search.
		Parameters:
			sid: The search ID.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListTimeBuckets(sid string, resp ...*http.Response) (*TimeBucketsSummary, error)
	/*
		UpdateJob - search service endpoint
		Update the search job with the specified search ID (SID) with an action.
		Parameters:
			sid: The search ID.
			updateJob
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UpdateJob(sid string, updateJob UpdateJob, resp ...*http.Response) (*SearchJob, error)
}
