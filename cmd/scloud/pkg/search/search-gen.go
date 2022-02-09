// Package search -- generated by scloudgen
// !! DO NOT EDIT !!
//
package search

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/auth"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/flags"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/jsonx"
	model "github.com/splunk/splunk-cloud-sdk-go/services/search"
)

// CreateDataset Creates a dataset.
func CreateDataset(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var fields []model.FieldPost
	err = flags.ParseFlag(cmd.Flags(), "fields", &fields)
	if err != nil {
		return fmt.Errorf(`error parsing "fields": ` + err.Error())
	}
	var idDefault string
	id := &idDefault
	err = flags.ParseFlag(cmd.Flags(), "id", &id)
	if err != nil {
		return fmt.Errorf(`error parsing "id": ` + err.Error())
	}
	var moduleDefault string
	module := &moduleDefault
	err = flags.ParseFlag(cmd.Flags(), "module", &module)
	if err != nil {
		return fmt.Errorf(`error parsing "module": ` + err.Error())
	}
	var name string
	err = flags.ParseFlag(cmd.Flags(), "name", &name)
	if err != nil {
		return fmt.Errorf(`error parsing "name": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.DatasetPost{

		Fields: fields,
		Id:     id,
		Module: module,
		Name:   name,
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.SearchService.CreateDataset(generated_request_body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// CreateFederatedConnection Creates a new federated connection with information about how to connect to a remote index.
func CreateFederatedConnection(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var hostnameip string
	err = flags.ParseFlag(cmd.Flags(), "hostnameip", &hostnameip)
	if err != nil {
		return fmt.Errorf(`error parsing "hostnameip": ` + err.Error())
	}
	var name string
	err = flags.ParseFlag(cmd.Flags(), "name", &name)
	if err != nil {
		return fmt.Errorf(`error parsing "name": ` + err.Error())
	}
	var port float32
	err = flags.ParseFlag(cmd.Flags(), "port", &port)
	if err != nil {
		return fmt.Errorf(`error parsing "port": ` + err.Error())
	}
	var serviceaccountpassword string
	err = flags.ParseFlag(cmd.Flags(), "serviceaccountpassword", &serviceaccountpassword)
	if err != nil {
		return fmt.Errorf(`error parsing "serviceaccountpassword": ` + err.Error())
	}
	var serviceaccountuser string
	err = flags.ParseFlag(cmd.Flags(), "serviceaccountuser", &serviceaccountuser)
	if err != nil {
		return fmt.Errorf(`error parsing "serviceaccountuser": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.FederatedConnectionInput{

		Hostnameip:             hostnameip,
		Name:                   name,
		Port:                   port,
		Serviceaccountpassword: serviceaccountpassword,
		Serviceaccountuser:     serviceaccountuser,
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.SearchService.CreateFederatedConnection(generated_request_body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// CreateJob Creates a search job.
func CreateJob(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var allowSideEffectsDefault bool
	allowSideEffects := &allowSideEffectsDefault
	err = flags.ParseFlag(cmd.Flags(), "allow-side-effects", &allowSideEffects)
	if err != nil {
		return fmt.Errorf(`error parsing "allow-side-effects": ` + err.Error())
	}
	var collectEventSummaryDefault bool
	collectEventSummary := &collectEventSummaryDefault
	err = flags.ParseFlag(cmd.Flags(), "collect-event-summary", &collectEventSummary)
	if err != nil {
		return fmt.Errorf(`error parsing "collect-event-summary": ` + err.Error())
	}
	var collectFieldSummaryDefault bool
	collectFieldSummary := &collectFieldSummaryDefault
	err = flags.ParseFlag(cmd.Flags(), "collect-field-summary", &collectFieldSummary)
	if err != nil {
		return fmt.Errorf(`error parsing "collect-field-summary": ` + err.Error())
	}
	var collectTimeBucketsDefault bool
	collectTimeBuckets := &collectTimeBucketsDefault
	err = flags.ParseFlag(cmd.Flags(), "collect-time-buckets", &collectTimeBuckets)
	if err != nil {
		return fmt.Errorf(`error parsing "collect-time-buckets": ` + err.Error())
	}
	var earliestDefault string
	earliest := &earliestDefault
	err = flags.ParseFlag(cmd.Flags(), "earliest", &earliest)
	if err != nil {
		return fmt.Errorf(`error parsing "earliest": ` + err.Error())
	}
	var enablePreviewDefault bool
	enablePreview := &enablePreviewDefault
	err = flags.ParseFlag(cmd.Flags(), "enable-preview", &enablePreview)
	if err != nil {
		return fmt.Errorf(`error parsing "enable-preview": ` + err.Error())
	}
	var extractAllFieldsDefault bool
	extractAllFields := &extractAllFieldsDefault
	err = flags.ParseFlag(cmd.Flags(), "extract-all-fields", &extractAllFields)
	if err != nil {
		return fmt.Errorf(`error parsing "extract-all-fields": ` + err.Error())
	}
	var extractFieldsDefault string
	extractFields := &extractFieldsDefault
	err = flags.ParseFlag(cmd.Flags(), "extract-fields", &extractFields)
	if err != nil {
		return fmt.Errorf(`error parsing "extract-fields": ` + err.Error())
	}
	var latestDefault string
	latest := &latestDefault
	err = flags.ParseFlag(cmd.Flags(), "latest", &latest)
	if err != nil {
		return fmt.Errorf(`error parsing "latest": ` + err.Error())
	}
	var maxTimeDefault int32
	maxTime := &maxTimeDefault
	err = flags.ParseFlag(cmd.Flags(), "max-time", &maxTime)
	if err != nil {
		return fmt.Errorf(`error parsing "max-time": ` + err.Error())
	}
	var messages []model.Message
	err = flags.ParseFlag(cmd.Flags(), "messages", &messages)
	if err != nil {
		return fmt.Errorf(`error parsing "messages": ` + err.Error())
	}
	var moduleDefault string
	module := &moduleDefault
	err = flags.ParseFlag(cmd.Flags(), "module", &module)
	if err != nil {
		return fmt.Errorf(`error parsing "module": ` + err.Error())
	}
	var query string
	err = flags.ParseFlag(cmd.Flags(), "query", &query)
	if err != nil {
		return fmt.Errorf(`error parsing "query": ` + err.Error())
	}
	var relativeTimeAnchorDefault string
	relativeTimeAnchor := &relativeTimeAnchorDefault
	err = flags.ParseFlag(cmd.Flags(), "relative-time-anchor", &relativeTimeAnchor)
	if err != nil {
		return fmt.Errorf(`error parsing "relative-time-anchor": ` + err.Error())
	}
	var requiredFreshnessDefault int32
	requiredFreshness := &requiredFreshnessDefault
	err = flags.ParseFlag(cmd.Flags(), "required-freshness", &requiredFreshness)
	if err != nil {
		return fmt.Errorf(`error parsing "required-freshness": ` + err.Error())
	}
	var statusDefault model.SearchStatus
	status := &statusDefault
	err = flags.ParseFlag(cmd.Flags(), "status", &status)
	if err != nil {
		return fmt.Errorf(`error parsing "status": ` + err.Error())
	}
	var timezone interface{}
	err = flags.ParseFlag(cmd.Flags(), "timezone", &timezone)
	if err != nil {
		return fmt.Errorf(`error parsing "timezone": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.SearchJob{

		AllowSideEffects:    allowSideEffects,
		CollectEventSummary: collectEventSummary,
		CollectFieldSummary: collectFieldSummary,
		CollectTimeBuckets:  collectTimeBuckets,
		EnablePreview:       enablePreview,
		ExtractAllFields:    extractAllFields,
		ExtractFields:       extractFields,
		MaxTime:             maxTime,
		Messages:            messages,
		Module:              module,
		Query:               query,
		QueryParameters: &model.QueryParameters{
			Earliest:           earliest,
			Latest:             latest,
			RelativeTimeAnchor: relativeTimeAnchor,
			Timezone:           timezone,
		},
		RequiredFreshness: requiredFreshness,
		Status:            status,
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.SearchService.CreateJob(generated_request_body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// DeleteDatasetById Deletes a dataset with a specified dataset ID (datasetid).
func DeleteDatasetById(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var datasetid string
	err = flags.ParseFlag(cmd.Flags(), "datasetid", &datasetid)
	if err != nil {
		return fmt.Errorf(`error parsing "datasetid": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	err = client.SearchService.DeleteDatasetById(datasetid)
	if err != nil {
		return err
	}

	return nil
}

// DeleteFederatedConnection Deletes a federated connection with the specified name (connectionName)
func DeleteFederatedConnection(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var connectionName string
	err = flags.ParseFlag(cmd.Flags(), "connection-name", &connectionName)
	if err != nil {
		return fmt.Errorf(`error parsing "connection-name": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	err = client.SearchService.DeleteFederatedConnection(connectionName)
	if err != nil {
		return err
	}

	return nil
}

// DeleteJob Creates a search job that deletes events from an index. The events are deleted from the index in the specified module, based on the search criteria as specified by the predicate.
func DeleteJob(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var earliestDefault string
	earliest := &earliestDefault
	err = flags.ParseFlag(cmd.Flags(), "earliest", &earliest)
	if err != nil {
		return fmt.Errorf(`error parsing "earliest": ` + err.Error())
	}
	var extractFieldsDefault string
	extractFields := &extractFieldsDefault
	err = flags.ParseFlag(cmd.Flags(), "extract-fields", &extractFields)
	if err != nil {
		return fmt.Errorf(`error parsing "extract-fields": ` + err.Error())
	}
	var index string
	err = flags.ParseFlag(cmd.Flags(), "index", &index)
	if err != nil {
		return fmt.Errorf(`error parsing "index": ` + err.Error())
	}
	var latestDefault string
	latest := &latestDefault
	err = flags.ParseFlag(cmd.Flags(), "latest", &latest)
	if err != nil {
		return fmt.Errorf(`error parsing "latest": ` + err.Error())
	}
	var maxTimeDefault int32
	maxTime := &maxTimeDefault
	err = flags.ParseFlag(cmd.Flags(), "max-time", &maxTime)
	if err != nil {
		return fmt.Errorf(`error parsing "max-time": ` + err.Error())
	}
	var messages []model.Message
	err = flags.ParseFlag(cmd.Flags(), "messages", &messages)
	if err != nil {
		return fmt.Errorf(`error parsing "messages": ` + err.Error())
	}
	var module string
	err = flags.ParseFlag(cmd.Flags(), "module", &module)
	if err != nil {
		return fmt.Errorf(`error parsing "module": ` + err.Error())
	}
	var predicate string
	err = flags.ParseFlag(cmd.Flags(), "predicate", &predicate)
	if err != nil {
		return fmt.Errorf(`error parsing "predicate": ` + err.Error())
	}
	var relativeTimeAnchorDefault string
	relativeTimeAnchor := &relativeTimeAnchorDefault
	err = flags.ParseFlag(cmd.Flags(), "relative-time-anchor", &relativeTimeAnchor)
	if err != nil {
		return fmt.Errorf(`error parsing "relative-time-anchor": ` + err.Error())
	}
	var statusDefault model.SearchStatus
	status := &statusDefault
	err = flags.ParseFlag(cmd.Flags(), "status", &status)
	if err != nil {
		return fmt.Errorf(`error parsing "status": ` + err.Error())
	}
	var timezone interface{}
	err = flags.ParseFlag(cmd.Flags(), "timezone", &timezone)
	if err != nil {
		return fmt.Errorf(`error parsing "timezone": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.DeleteSearchJob{

		ExtractFields: extractFields,
		Index:         index,
		MaxTime:       maxTime,
		Messages:      messages,
		Module:        module,
		Predicate:     predicate,
		QueryParameters: &model.QueryParameters{
			Earliest:           earliest,
			Latest:             latest,
			RelativeTimeAnchor: relativeTimeAnchor,
			Timezone:           timezone,
		},
		Status: status,
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.SearchService.DeleteJob(generated_request_body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ExportResults Exports the search results for the job with the specified search ID (SID). Export the results as a CSV file or JSON file.
func ExportResults(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var countDefault int32
	count := &countDefault
	err = flags.ParseFlag(cmd.Flags(), "count", &count)
	if err != nil {
		return fmt.Errorf(`error parsing "count": ` + err.Error())
	}
	var filename string
	err = flags.ParseFlag(cmd.Flags(), "filename", &filename)
	if err != nil {
		return fmt.Errorf(`error parsing "filename": ` + err.Error())
	}
	var outputMode model.ExportResultsoutputMode
	err = flags.ParseFlag(cmd.Flags(), "output-mode", &outputMode)
	if err != nil {
		return fmt.Errorf(`error parsing "output-mode": ` + err.Error())
	}
	var sid string
	err = flags.ParseFlag(cmd.Flags(), "sid", &sid)
	if err != nil {
		return fmt.Errorf(`error parsing "sid": ` + err.Error())
	}
	// Form query params
	generated_query := model.ExportResultsQueryParams{}
	generated_query.Count = count
	generated_query.Filename = filename
	generated_query.OutputMode = outputMode

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.SearchService.ExportResults(sid, &generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// GetAllFederatedConnections Returns all federated connections.
func GetAllFederatedConnections(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.SearchService.GetAllFederatedConnections()
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// GetDatasetById Returns a dataset with a specified dataset ID (datasetid).
func GetDatasetById(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var datasetid string
	err = flags.ParseFlag(cmd.Flags(), "datasetid", &datasetid)
	if err != nil {
		return fmt.Errorf(`error parsing "datasetid": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.SearchService.GetDatasetById(datasetid)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// GetFederatedConnectionByName Returns the federated connection with the specified name (connectionName).
func GetFederatedConnectionByName(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var connectionName string
	err = flags.ParseFlag(cmd.Flags(), "connection-name", &connectionName)
	if err != nil {
		return fmt.Errorf(`error parsing "connection-name": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.SearchService.GetFederatedConnectionByName(connectionName)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// GetJob Returns the search job with the specified search ID (SID).
func GetJob(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var sid string
	err = flags.ParseFlag(cmd.Flags(), "sid", &sid)
	if err != nil {
		return fmt.Errorf(`error parsing "sid": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.SearchService.GetJob(sid)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListDatasets Returns a list of all datasets.
func ListDatasets(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.SearchService.ListDatasets()
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListEventsSummary Returns an events summary for search ID (SID) search.
func ListEventsSummary(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var countDefault int32
	count := &countDefault
	err = flags.ParseFlag(cmd.Flags(), "count", &count)
	if err != nil {
		return fmt.Errorf(`error parsing "count": ` + err.Error())
	}
	var earliest string
	err = flags.ParseFlag(cmd.Flags(), "earliest", &earliest)
	if err != nil {
		return fmt.Errorf(`error parsing "earliest": ` + err.Error())
	}
	var field string
	err = flags.ParseFlag(cmd.Flags(), "field", &field)
	if err != nil {
		return fmt.Errorf(`error parsing "field": ` + err.Error())
	}
	var latest string
	err = flags.ParseFlag(cmd.Flags(), "latest", &latest)
	if err != nil {
		return fmt.Errorf(`error parsing "latest": ` + err.Error())
	}
	var offsetDefault int32
	offset := &offsetDefault
	err = flags.ParseFlag(cmd.Flags(), "offset", &offset)
	if err != nil {
		return fmt.Errorf(`error parsing "offset": ` + err.Error())
	}
	var sid string
	err = flags.ParseFlag(cmd.Flags(), "sid", &sid)
	if err != nil {
		return fmt.Errorf(`error parsing "sid": ` + err.Error())
	}
	// Form query params
	generated_query := model.ListEventsSummaryQueryParams{}
	generated_query.Count = count
	generated_query.Earliest = earliest
	generated_query.Field = field
	generated_query.Latest = latest
	generated_query.Offset = offset

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.SearchService.ListEventsSummary(sid, &generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListFieldsSummary Returns a fields stats summary of the events to-date, for search ID (SID) search.
func ListFieldsSummary(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var earliest string
	err = flags.ParseFlag(cmd.Flags(), "earliest", &earliest)
	if err != nil {
		return fmt.Errorf(`error parsing "earliest": ` + err.Error())
	}
	var latest string
	err = flags.ParseFlag(cmd.Flags(), "latest", &latest)
	if err != nil {
		return fmt.Errorf(`error parsing "latest": ` + err.Error())
	}
	var sid string
	err = flags.ParseFlag(cmd.Flags(), "sid", &sid)
	if err != nil {
		return fmt.Errorf(`error parsing "sid": ` + err.Error())
	}
	// Form query params
	generated_query := model.ListFieldsSummaryQueryParams{}
	generated_query.Earliest = earliest
	generated_query.Latest = latest

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.SearchService.ListFieldsSummary(sid, &generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListJobs Returns the matching list of search jobs.
func ListJobs(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var countDefault int32
	count := &countDefault
	err = flags.ParseFlag(cmd.Flags(), "count", &count)
	if err != nil {
		return fmt.Errorf(`error parsing "count": ` + err.Error())
	}
	var filter string
	err = flags.ParseFlag(cmd.Flags(), "filter", &filter)
	if err != nil {
		return fmt.Errorf(`error parsing "filter": ` + err.Error())
	}
	var status model.SearchStatus
	err = flags.ParseFlag(cmd.Flags(), "status", &status)
	if err != nil {
		return fmt.Errorf(`error parsing "status": ` + err.Error())
	}
	// Form query params
	generated_query := model.ListJobsQueryParams{}
	generated_query.Count = count
	generated_query.Filter = filter
	generated_query.Status = status

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.SearchService.ListJobs(&generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListPreviewResults Returns the preview search results for the job with the specified search ID (SID). Can be used when a job is running to return interim results.
func ListPreviewResults(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var countDefault int32
	count := &countDefault
	err = flags.ParseFlag(cmd.Flags(), "count", &count)
	if err != nil {
		return fmt.Errorf(`error parsing "count": ` + err.Error())
	}
	var offsetDefault int32
	offset := &offsetDefault
	err = flags.ParseFlag(cmd.Flags(), "offset", &offset)
	if err != nil {
		return fmt.Errorf(`error parsing "offset": ` + err.Error())
	}
	var sid string
	err = flags.ParseFlag(cmd.Flags(), "sid", &sid)
	if err != nil {
		return fmt.Errorf(`error parsing "sid": ` + err.Error())
	}
	// Form query params
	generated_query := model.ListPreviewResultsQueryParams{}
	generated_query.Count = count
	generated_query.Offset = offset

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.SearchService.ListPreviewResults(sid, &generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListResults Returns the search results for the job with the specified search ID (SID).
func ListResults(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var countDefault int32
	count := &countDefault
	err = flags.ParseFlag(cmd.Flags(), "count", &count)
	if err != nil {
		return fmt.Errorf(`error parsing "count": ` + err.Error())
	}
	var field string
	err = flags.ParseFlag(cmd.Flags(), "field", &field)
	if err != nil {
		return fmt.Errorf(`error parsing "field": ` + err.Error())
	}
	var offsetDefault int32
	offset := &offsetDefault
	err = flags.ParseFlag(cmd.Flags(), "offset", &offset)
	if err != nil {
		return fmt.Errorf(`error parsing "offset": ` + err.Error())
	}
	var sid string
	err = flags.ParseFlag(cmd.Flags(), "sid", &sid)
	if err != nil {
		return fmt.Errorf(`error parsing "sid": ` + err.Error())
	}
	// Form query params
	generated_query := model.ListResultsQueryParams{}
	generated_query.Count = count
	generated_query.Field = field
	generated_query.Offset = offset

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.SearchService.ListResults(sid, &generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListTimeBuckets Returns the event distribution over time of the untransformed events read to-date, for search ID(SID) search.
func ListTimeBuckets(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var sid string
	err = flags.ParseFlag(cmd.Flags(), "sid", &sid)
	if err != nil {
		return fmt.Errorf(`error parsing "sid": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.SearchService.ListTimeBuckets(sid)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// PutFederatedConnectionByName Creates or updates a federated connection with a specified name (connectionName).
func PutFederatedConnectionByName(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var connectionName string
	err = flags.ParseFlag(cmd.Flags(), "connection-name", &connectionName)
	if err != nil {
		return fmt.Errorf(`error parsing "connection-name": ` + err.Error())
	}
	var hostnameip string
	err = flags.ParseFlag(cmd.Flags(), "hostnameip", &hostnameip)
	if err != nil {
		return fmt.Errorf(`error parsing "hostnameip": ` + err.Error())
	}
	var name string
	err = flags.ParseFlag(cmd.Flags(), "name", &name)
	if err != nil {
		return fmt.Errorf(`error parsing "name": ` + err.Error())
	}
	var port float32
	err = flags.ParseFlag(cmd.Flags(), "port", &port)
	if err != nil {
		return fmt.Errorf(`error parsing "port": ` + err.Error())
	}
	var serviceaccountpassword string
	err = flags.ParseFlag(cmd.Flags(), "serviceaccountpassword", &serviceaccountpassword)
	if err != nil {
		return fmt.Errorf(`error parsing "serviceaccountpassword": ` + err.Error())
	}
	var serviceaccountuser string
	err = flags.ParseFlag(cmd.Flags(), "serviceaccountuser", &serviceaccountuser)
	if err != nil {
		return fmt.Errorf(`error parsing "serviceaccountuser": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.FederatedConnectionInput{

		Hostnameip:             hostnameip,
		Name:                   name,
		Port:                   port,
		Serviceaccountpassword: serviceaccountpassword,
		Serviceaccountuser:     serviceaccountuser,
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.SearchService.PutFederatedConnectionByName(connectionName, generated_request_body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// RefreshFederatedConnection Refresh a federated connection to fetch new remote indexes and add/delete corresponding federated datasets.
func RefreshFederatedConnection(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var connectionName string
	err = flags.ParseFlag(cmd.Flags(), "connection-name", &connectionName)
	if err != nil {
		return fmt.Errorf(`error parsing "connection-name": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	err = client.SearchService.RefreshFederatedConnection(connectionName)
	if err != nil {
		return err
	}

	return nil
}

// TestFederatedConnection Test connection with remote EC instance using federated connection parameters.
func TestFederatedConnection(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var connectionName string
	err = flags.ParseFlag(cmd.Flags(), "connection-name", &connectionName)
	if err != nil {
		return fmt.Errorf(`error parsing "connection-name": ` + err.Error())
	}

	// Silence Usage
	cmd.SilenceUsage = true

	err = client.SearchService.TestFederatedConnection(connectionName)
	if err != nil {
		return err
	}

	return nil
}

// UpdateDatasetByIdFederatedDataset Modifies a dataset with a specified dataset ID (datasetid).
func UpdateDatasetByIdFederatedDataset(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var datasetid string
	err = flags.ParseFlag(cmd.Flags(), "datasetid", &datasetid)
	if err != nil {
		return fmt.Errorf(`error parsing "datasetid": ` + err.Error())
	}
	var federatedConnectionDefault string
	federatedConnection := &federatedConnectionDefault
	err = flags.ParseFlag(cmd.Flags(), "federated-connection", &federatedConnection)
	if err != nil {
		return fmt.Errorf(`error parsing "federated-connection": ` + err.Error())
	}
	var federatedDatasetDefault string
	federatedDataset := &federatedDatasetDefault
	err = flags.ParseFlag(cmd.Flags(), "federated-dataset", &federatedDataset)
	if err != nil {
		return fmt.Errorf(`error parsing "federated-dataset": ` + err.Error())
	}
	var federatedDatasetKindDefault string
	federatedDatasetKind := &federatedDatasetKindDefault
	err = flags.ParseFlag(cmd.Flags(), "federated-dataset-kind", &federatedDatasetKind)
	if err != nil {
		return fmt.Errorf(`error parsing "federated-dataset-kind": ` + err.Error())
	}
	var kindDefault model.FederatedDatasetKind
	kind := &kindDefault
	err = flags.ParseFlag(cmd.Flags(), "kind", &kind)
	if err != nil {
		return fmt.Errorf(`error parsing "kind": ` + err.Error())
	}
	var moduleDefault string
	module := &moduleDefault
	err = flags.ParseFlag(cmd.Flags(), "module", &module)
	if err != nil {
		return fmt.Errorf(`error parsing "module": ` + err.Error())
	}
	var nameDefault string
	name := &nameDefault
	err = flags.ParseFlag(cmd.Flags(), "name", &name)
	if err != nil {
		return fmt.Errorf(`error parsing "name": ` + err.Error())
	}
	var ownerDefault string
	owner := &ownerDefault
	err = flags.ParseFlag(cmd.Flags(), "owner", &owner)
	if err != nil {
		return fmt.Errorf(`error parsing "owner": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.FederatedDatasetPatch{

		FederatedConnection:  federatedConnection,
		FederatedDataset:     federatedDataset,
		FederatedDatasetKind: federatedDatasetKind,
		Kind:                 kind,
		Module:               module,
		Name:                 name,
		Owner:                owner,
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.SearchService.UpdateDatasetById(datasetid, model.MakeDatasetPatchFromFederatedDatasetPatch(generated_request_body))
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// UpdateDatasetByIdIndexDataset Modifies a dataset with a specified dataset ID (datasetid).
func UpdateDatasetByIdIndexDataset(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var datasetid string
	err = flags.ParseFlag(cmd.Flags(), "datasetid", &datasetid)
	if err != nil {
		return fmt.Errorf(`error parsing "datasetid": ` + err.Error())
	}
	var disabledDefault bool
	disabled := &disabledDefault
	err = flags.ParseFlag(cmd.Flags(), "disabled", &disabled)
	if err != nil {
		return fmt.Errorf(`error parsing "disabled": ` + err.Error())
	}
	var frozenTimePeriodInSecsDefault int32
	frozenTimePeriodInSecs := &frozenTimePeriodInSecsDefault
	err = flags.ParseFlag(cmd.Flags(), "frozen-time-period-in-secs", &frozenTimePeriodInSecs)
	if err != nil {
		return fmt.Errorf(`error parsing "frozen-time-period-in-secs": ` + err.Error())
	}
	var kindDefault model.IndexDatasetKind
	kind := &kindDefault
	err = flags.ParseFlag(cmd.Flags(), "kind", &kind)
	if err != nil {
		return fmt.Errorf(`error parsing "kind": ` + err.Error())
	}
	var moduleDefault string
	module := &moduleDefault
	err = flags.ParseFlag(cmd.Flags(), "module", &module)
	if err != nil {
		return fmt.Errorf(`error parsing "module": ` + err.Error())
	}
	var nameDefault string
	name := &nameDefault
	err = flags.ParseFlag(cmd.Flags(), "name", &name)
	if err != nil {
		return fmt.Errorf(`error parsing "name": ` + err.Error())
	}
	var ownerDefault string
	owner := &ownerDefault
	err = flags.ParseFlag(cmd.Flags(), "owner", &owner)
	if err != nil {
		return fmt.Errorf(`error parsing "owner": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.IndexDatasetPatch{

		Disabled:               disabled,
		FrozenTimePeriodInSecs: frozenTimePeriodInSecs,
		Kind:                   kind,
		Module:                 module,
		Name:                   name,
		Owner:                  owner,
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.SearchService.UpdateDatasetById(datasetid, model.MakeDatasetPatchFromIndexDatasetPatch(generated_request_body))
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// UpdateDatasetByIdKvCollectionDataset Modifies a dataset with a specified dataset ID (datasetid).
func UpdateDatasetByIdKvCollectionDataset(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var datasetid string
	err = flags.ParseFlag(cmd.Flags(), "datasetid", &datasetid)
	if err != nil {
		return fmt.Errorf(`error parsing "datasetid": ` + err.Error())
	}
	var kindDefault model.KvCollectionDatasetKind
	kind := &kindDefault
	err = flags.ParseFlag(cmd.Flags(), "kind", &kind)
	if err != nil {
		return fmt.Errorf(`error parsing "kind": ` + err.Error())
	}
	var moduleDefault string
	module := &moduleDefault
	err = flags.ParseFlag(cmd.Flags(), "module", &module)
	if err != nil {
		return fmt.Errorf(`error parsing "module": ` + err.Error())
	}
	var nameDefault string
	name := &nameDefault
	err = flags.ParseFlag(cmd.Flags(), "name", &name)
	if err != nil {
		return fmt.Errorf(`error parsing "name": ` + err.Error())
	}
	var ownerDefault string
	owner := &ownerDefault
	err = flags.ParseFlag(cmd.Flags(), "owner", &owner)
	if err != nil {
		return fmt.Errorf(`error parsing "owner": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.KvCollectionDatasetPatch{

		Kind:   kind,
		Module: module,
		Name:   name,
		Owner:  owner,
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.SearchService.UpdateDatasetById(datasetid, model.MakeDatasetPatchFromKvCollectionDatasetPatch(generated_request_body))
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// UpdateDatasetByIdLookupDataset Modifies a dataset with a specified dataset ID (datasetid).
func UpdateDatasetByIdLookupDataset(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var caseSensitiveMatchDefault bool
	caseSensitiveMatch := &caseSensitiveMatchDefault
	err = flags.ParseFlag(cmd.Flags(), "case-sensitive-match", &caseSensitiveMatch)
	if err != nil {
		return fmt.Errorf(`error parsing "case-sensitive-match": ` + err.Error())
	}
	var datasetid string
	err = flags.ParseFlag(cmd.Flags(), "datasetid", &datasetid)
	if err != nil {
		return fmt.Errorf(`error parsing "datasetid": ` + err.Error())
	}
	var externalKindDefault model.LookupDatasetExternalKind
	externalKind := &externalKindDefault
	err = flags.ParseFlag(cmd.Flags(), "external-kind", &externalKind)
	if err != nil {
		return fmt.Errorf(`error parsing "external-kind": ` + err.Error())
	}
	var externalNameDefault string
	externalName := &externalNameDefault
	err = flags.ParseFlag(cmd.Flags(), "external-name", &externalName)
	if err != nil {
		return fmt.Errorf(`error parsing "external-name": ` + err.Error())
	}
	var filterDefault string
	filter := &filterDefault
	err = flags.ParseFlag(cmd.Flags(), "filter", &filter)
	if err != nil {
		return fmt.Errorf(`error parsing "filter": ` + err.Error())
	}
	var kindDefault model.LookupDatasetKind
	kind := &kindDefault
	err = flags.ParseFlag(cmd.Flags(), "kind", &kind)
	if err != nil {
		return fmt.Errorf(`error parsing "kind": ` + err.Error())
	}
	var moduleDefault string
	module := &moduleDefault
	err = flags.ParseFlag(cmd.Flags(), "module", &module)
	if err != nil {
		return fmt.Errorf(`error parsing "module": ` + err.Error())
	}
	var nameDefault string
	name := &nameDefault
	err = flags.ParseFlag(cmd.Flags(), "name", &name)
	if err != nil {
		return fmt.Errorf(`error parsing "name": ` + err.Error())
	}
	var ownerDefault string
	owner := &ownerDefault
	err = flags.ParseFlag(cmd.Flags(), "owner", &owner)
	if err != nil {
		return fmt.Errorf(`error parsing "owner": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.LookupDatasetPatch{

		CaseSensitiveMatch: caseSensitiveMatch,
		ExternalKind:       externalKind,
		ExternalName:       externalName,
		Filter:             filter,
		Kind:               kind,
		Module:             module,
		Name:               name,
		Owner:              owner,
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.SearchService.UpdateDatasetById(datasetid, model.MakeDatasetPatchFromLookupDatasetPatch(generated_request_body))
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// UpdateDatasetByIdMetricDataset Modifies a dataset with a specified dataset ID (datasetid).
func UpdateDatasetByIdMetricDataset(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var datasetid string
	err = flags.ParseFlag(cmd.Flags(), "datasetid", &datasetid)
	if err != nil {
		return fmt.Errorf(`error parsing "datasetid": ` + err.Error())
	}
	var disabledDefault bool
	disabled := &disabledDefault
	err = flags.ParseFlag(cmd.Flags(), "disabled", &disabled)
	if err != nil {
		return fmt.Errorf(`error parsing "disabled": ` + err.Error())
	}
	var frozenTimePeriodInSecsDefault int32
	frozenTimePeriodInSecs := &frozenTimePeriodInSecsDefault
	err = flags.ParseFlag(cmd.Flags(), "frozen-time-period-in-secs", &frozenTimePeriodInSecs)
	if err != nil {
		return fmt.Errorf(`error parsing "frozen-time-period-in-secs": ` + err.Error())
	}
	var kindDefault model.MetricDatasetKind
	kind := &kindDefault
	err = flags.ParseFlag(cmd.Flags(), "kind", &kind)
	if err != nil {
		return fmt.Errorf(`error parsing "kind": ` + err.Error())
	}
	var moduleDefault string
	module := &moduleDefault
	err = flags.ParseFlag(cmd.Flags(), "module", &module)
	if err != nil {
		return fmt.Errorf(`error parsing "module": ` + err.Error())
	}
	var nameDefault string
	name := &nameDefault
	err = flags.ParseFlag(cmd.Flags(), "name", &name)
	if err != nil {
		return fmt.Errorf(`error parsing "name": ` + err.Error())
	}
	var ownerDefault string
	owner := &ownerDefault
	err = flags.ParseFlag(cmd.Flags(), "owner", &owner)
	if err != nil {
		return fmt.Errorf(`error parsing "owner": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.MetricDatasetPatch{

		Disabled:               disabled,
		FrozenTimePeriodInSecs: frozenTimePeriodInSecs,
		Kind:                   kind,
		Module:                 module,
		Name:                   name,
		Owner:                  owner,
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.SearchService.UpdateDatasetById(datasetid, model.MakeDatasetPatchFromMetricDatasetPatch(generated_request_body))
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// UpdateJob Updates the search job with the specified search ID (SID) with an action.
func UpdateJob(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var sid string
	err = flags.ParseFlag(cmd.Flags(), "sid", &sid)
	if err != nil {
		return fmt.Errorf(`error parsing "sid": ` + err.Error())
	}
	var status model.UpdateJobStatus
	err = flags.ParseFlag(cmd.Flags(), "status", &status)
	if err != nil {
		return fmt.Errorf(`error parsing "status": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.UpdateJob{

		Status: status,
	}

	// Silence Usage
	cmd.SilenceUsage = true

	resp, err := client.SearchService.UpdateJob(sid, generated_request_body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}
