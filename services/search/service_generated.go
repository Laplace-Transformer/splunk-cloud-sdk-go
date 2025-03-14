/*
 * Copyright © 2022 Splunk, Inc.
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
 * API version: v2 (recommended default)
 * Generated by: OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
 */

package search

import (
	"io"
	"net/http"

	"fmt"

	"github.com/splunk/go-dependencies/services"
	"github.com/splunk/go-dependencies/util"
)

const serviceCluster = "api"

type Service services.BaseService

// NewService creates a new search service client from the given Config
func NewService(iClient services.IClient) *Service {
	return &Service{Client: iClient}
}

/*
CreateDataset - search service endpoint
Creates a dataset.
Parameters:

	datasetPost
	resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) CreateDataset(datasetPost DatasetPost, resp ...*http.Response) (*Dataset, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/search/v2/datasets`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: datasetPost})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb Dataset
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
CreateFederatedConnection - search service endpoint
Creates a new federated connection with information about how to connect to a remote index.
Parameters:

	federatedConnectionInput
	resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) CreateFederatedConnection(federatedConnectionInput FederatedConnectionInput, resp ...*http.Response) (*FederatedConnection, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/search/v2/connections`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: federatedConnectionInput})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb FederatedConnection
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
CreateJob - search service endpoint
Creates a search job.
Parameters:

	searchJob
	resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) CreateJob(searchJob SearchJob, resp ...*http.Response) (*SearchJob, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/search/v2/jobs`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: searchJob})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb SearchJob
	b, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error reading response body: ", err)
	}
	fmt.Println("response body: \n", string(b))
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
DeleteDatasetById - search service endpoint
Deletes a dataset with a specified dataset ID (datasetid).
Parameters:

	datasetid: The dataset ID.
	resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) DeleteDatasetById(datasetid string, resp ...*http.Response) error {
	pp := struct {
		Datasetid string
	}{
		Datasetid: datasetid,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/search/v2/datasets/{{.Datasetid}}`, pp)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	return err
}

/*
DeleteFederatedConnection - search service endpoint
Deletes a federated connection with the specified name (connectionName)
Parameters:

	connectionName: The name of the federated connection.
	resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) DeleteFederatedConnection(connectionName string, resp ...*http.Response) error {
	pp := struct {
		ConnectionName string
	}{
		ConnectionName: connectionName,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/search/v2/connections/{{.ConnectionName}}`, pp)
	if err != nil {
		return err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	return err
}

/*
DeleteJob - search service endpoint
Creates a search job that deletes events from an index. The events are deleted from the index in the specified module, based on the search criteria as specified by the predicate.
Parameters:

	deleteSearchJob
	resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) DeleteJob(deleteSearchJob DeleteSearchJob, resp ...*http.Response) (*DeleteSearchJob, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/search/v2/jobs/delete`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: deleteSearchJob})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb DeleteSearchJob
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
ExportResults - search service endpoint
Exports the search results for the job with the specified search ID (SID). Export the results as a CSV file or JSON file.
Parameters:

	sid: The search ID.
	query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
	resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ExportResults(sid string, query *ExportResultsQueryParams, resp ...*http.Response) (*map[string]interface{}, error) {
	values := util.ParseURLParams(query)
	pp := struct {
		Sid string
	}{
		Sid: sid,
	}
	u, err := s.Client.BuildURLFromPathParams(values, serviceCluster, `/search/v2/jobs/{{.Sid}}/export`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb map[string]interface{}
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
GetAllFederatedConnections - search service endpoint
Returns all federated connections.
Parameters:

	resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetAllFederatedConnections(resp ...*http.Response) (*ListFederatedConnections, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/search/v2/connections`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb ListFederatedConnections
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
GetDatasetById - search service endpoint
Returns a dataset with a specified dataset ID (datasetid).
Parameters:

	datasetid: The dataset ID.
	resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetDatasetById(datasetid string, resp ...*http.Response) (*Dataset, error) {
	pp := struct {
		Datasetid string
	}{
		Datasetid: datasetid,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/search/v2/datasets/{{.Datasetid}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb Dataset
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
GetFederatedConnectionByName - search service endpoint
Returns the federated connection with the specified name (connectionName).
Parameters:

	connectionName: The name of the federated connection.
	resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetFederatedConnectionByName(connectionName string, resp ...*http.Response) (*FederatedConnection, error) {
	pp := struct {
		ConnectionName string
	}{
		ConnectionName: connectionName,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/search/v2/connections/{{.ConnectionName}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb FederatedConnection
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
GetJob - search service endpoint
Returns the search job with the specified search ID (SID).
Parameters:

	sid: The search ID.
	resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetJob(sid string, resp ...*http.Response) (*SearchJob, error) {
	pp := struct {
		Sid string
	}{
		Sid: sid,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/search/v2/jobs/{{.Sid}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb SearchJob
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
ListDatasets - search service endpoint
Returns a list of all datasets.
Parameters:

	resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListDatasets(resp ...*http.Response) (*ListDatasets, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/search/v2/datasets`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb ListDatasets
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
ListEventsSummary - search service endpoint
Returns an events summary for search ID (SID) search.
Parameters:

	sid: The search ID.
	query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
	resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListEventsSummary(sid string, query *ListEventsSummaryQueryParams, resp ...*http.Response) (*ListSearchResultsResponse, error) {
	values := util.ParseURLParams(query)
	pp := struct {
		Sid string
	}{
		Sid: sid,
	}
	u, err := s.Client.BuildURLFromPathParams(values, serviceCluster, `/search/v2/jobs/{{.Sid}}/timeline-metadata/auto/events-summary`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb ListSearchResultsResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
ListFieldsSummary - search service endpoint
Returns a fields stats summary of the events to-date, for search ID (SID) search.
Parameters:

	sid: The search ID.
	query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
	resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListFieldsSummary(sid string, query *ListFieldsSummaryQueryParams, resp ...*http.Response) (*FieldsSummary, error) {
	values := util.ParseURLParams(query)
	pp := struct {
		Sid string
	}{
		Sid: sid,
	}
	u, err := s.Client.BuildURLFromPathParams(values, serviceCluster, `/search/v2/jobs/{{.Sid}}/timeline-metadata/auto/fields-summary`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb FieldsSummary
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
ListJobs - search service endpoint
Returns the matching list of search jobs.
Parameters:

	query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
	resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListJobs(query *ListJobsQueryParams, resp ...*http.Response) ([]SearchJob, error) {
	values := util.ParseURLParams(query)
	u, err := s.Client.BuildURLFromPathParams(values, serviceCluster, `/search/v2/jobs`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb []SearchJob
	err = util.ParseResponse(&rb, response)
	return rb, err
}

/*
ListPreviewResults - search service endpoint
Returns the preview search results for the job with the specified search ID (SID). Can be used when a job is running to return interim results.
Parameters:

	sid: The search ID.
	query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
	resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListPreviewResults(sid string, query *ListPreviewResultsQueryParams, resp ...*http.Response) (*ListPreviewResultsResponse, error) {
	values := util.ParseURLParams(query)
	pp := struct {
		Sid string
	}{
		Sid: sid,
	}
	u, err := s.Client.BuildURLFromPathParams(values, serviceCluster, `/search/v2/jobs/{{.Sid}}/results-preview`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb ListPreviewResultsResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
ListResults - search service endpoint
Returns the search results for the job with the specified search ID (SID).
Parameters:

	sid: The search ID.
	query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
	resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListResults(sid string, query *ListResultsQueryParams, resp ...*http.Response) (*ListSearchResultsResponse, error) {
	values := util.ParseURLParams(query)
	pp := struct {
		Sid string
	}{
		Sid: sid,
	}
	u, err := s.Client.BuildURLFromPathParams(values, serviceCluster, `/search/v2/jobs/{{.Sid}}/results`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb ListSearchResultsResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
ListTimeBuckets - search service endpoint
Returns the event distribution over time of the untransformed events read to-date, for search ID(SID) search.
Parameters:

	sid: The search ID.
	resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListTimeBuckets(sid string, resp ...*http.Response) (*TimeBucketsSummary, error) {
	pp := struct {
		Sid string
	}{
		Sid: sid,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/search/v2/jobs/{{.Sid}}/timeline-metadata/auto/time-buckets`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb TimeBucketsSummary
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
PutFederatedConnectionByName - search service endpoint
Creates or updates a federated connection with a specified name (connectionName).
Parameters:

	connectionName: The name of the federated connection.
	federatedConnectionInput
	resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) PutFederatedConnectionByName(connectionName string, federatedConnectionInput FederatedConnectionInput, resp ...*http.Response) (*FederatedConnection, error) {
	pp := struct {
		ConnectionName string
	}{
		ConnectionName: connectionName,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/search/v2/connections/{{.ConnectionName}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Put(services.RequestParams{URL: u, Body: federatedConnectionInput})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb FederatedConnection
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
RefreshFederatedConnection - search service endpoint
Refresh a federated connection to fetch new remote indexes and add/delete corresponding federated datasets.
Parameters:

	connectionName: The name of the federated connection.
	resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) RefreshFederatedConnection(connectionName string, resp ...*http.Response) error {
	pp := struct {
		ConnectionName string
	}{
		ConnectionName: connectionName,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/search/v2/connections/{{.ConnectionName}}/refresh`, pp)
	if err != nil {
		return err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	return err
}

/*
TestFederatedConnection - search service endpoint
Test connection with remote EC instance using federated connection parameters.
Parameters:

	connectionName: The name of the federated connection.
	resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) TestFederatedConnection(connectionName string, resp ...*http.Response) error {
	pp := struct {
		ConnectionName string
	}{
		ConnectionName: connectionName,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/search/v2/connections/{{.ConnectionName}}/test`, pp)
	if err != nil {
		return err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	return err
}

/*
UpdateDatasetById - search service endpoint
Modifies a dataset with a specified dataset ID (datasetid).
Parameters:

	datasetid: The dataset ID.
	datasetPatch
	resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) UpdateDatasetById(datasetid string, datasetPatch DatasetPatch, resp ...*http.Response) (*Dataset, error) {
	pp := struct {
		Datasetid string
	}{
		Datasetid: datasetid,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/search/v2/datasets/{{.Datasetid}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Patch(services.RequestParams{URL: u, Body: datasetPatch})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb Dataset
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
UpdateJob - search service endpoint
Updates the search job with the specified search ID (SID) with an action.
Parameters:

	sid: The search ID.
	updateJob
	resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) UpdateJob(sid string, updateJob UpdateJob, resp ...*http.Response) (*SearchJob, error) {
	pp := struct {
		Sid string
	}{
		Sid: sid,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/search/v2/jobs/{{.Sid}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Patch(services.RequestParams{URL: u, Body: updateJob})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb SearchJob
	err = util.ParseResponse(&rb, response)
	return &rb, err
}
