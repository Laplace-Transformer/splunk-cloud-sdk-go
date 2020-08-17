// Package kvstore -- generated by scloudgen
// !! DO NOT EDIT !!
//
package kvstore

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/auth"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/flags"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/jsonx"
	model "github.com/splunk/splunk-cloud-sdk-go/services/kvstore"
)

// CreateIndex Creates an index on a collection.
func CreateIndex(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var collection string
	err = flags.ParseFlag(cmd.Flags(), "collection", &collection)
	if err != nil {
		return fmt.Errorf(`error parsing "collection": ` + err.Error())
	}
	var fields []model.IndexFieldDefinition
	err = flags.ParseFlag(cmd.Flags(), "fields", &fields)
	if err != nil {
		return fmt.Errorf(`error parsing "fields": ` + err.Error())
	}
	var name string
	err = flags.ParseFlag(cmd.Flags(), "name", &name)
	if err != nil {
		return fmt.Errorf(`error parsing "name": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.IndexDefinition{

		Fields: fields,
		Name:   name,
	}

	resp, err := client.KVStoreService.CreateIndex(collection, generated_request_body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// DeleteIndex Removes an index from a collection.
func DeleteIndex(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var collection string
	err = flags.ParseFlag(cmd.Flags(), "collection", &collection)
	if err != nil {
		return fmt.Errorf(`error parsing "collection": ` + err.Error())
	}
	var index string
	err = flags.ParseFlag(cmd.Flags(), "index", &index)
	if err != nil {
		return fmt.Errorf(`error parsing "index": ` + err.Error())
	}

	err = client.KVStoreService.DeleteIndex(collection, index)
	if err != nil {
		return err
	}

	return nil
}

// DeleteRecordByKey Deletes a record with a given key.
func DeleteRecordByKey(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var collection string
	err = flags.ParseFlag(cmd.Flags(), "collection", &collection)
	if err != nil {
		return fmt.Errorf(`error parsing "collection": ` + err.Error())
	}
	var key string
	err = flags.ParseFlag(cmd.Flags(), "key", &key)
	if err != nil {
		return fmt.Errorf(`error parsing "key": ` + err.Error())
	}

	err = client.KVStoreService.DeleteRecordByKey(collection, key)
	if err != nil {
		return err
	}

	return nil
}

// DeleteRecords Removes records in a collection that match the query.
func DeleteRecords(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var collection string
	err = flags.ParseFlag(cmd.Flags(), "collection", &collection)
	if err != nil {
		return fmt.Errorf(`error parsing "collection": ` + err.Error())
	}
	var query string
	err = flags.ParseFlag(cmd.Flags(), "query", &query)
	if err != nil {
		return fmt.Errorf(`error parsing "query": ` + err.Error())
	}
	// Form query params
	generated_query := model.DeleteRecordsQueryParams{}
	generated_query.Query = query

	err = client.KVStoreService.DeleteRecords(collection, &generated_query)
	if err != nil {
		return err
	}

	return nil
}

// GetRecordByKey Returns a record with a given key.
func GetRecordByKey(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var collection string
	err = flags.ParseFlag(cmd.Flags(), "collection", &collection)
	if err != nil {
		return fmt.Errorf(`error parsing "collection": ` + err.Error())
	}
	var key string
	err = flags.ParseFlag(cmd.Flags(), "key", &key)
	if err != nil {
		return fmt.Errorf(`error parsing "key": ` + err.Error())
	}

	resp, err := client.KVStoreService.GetRecordByKey(collection, key)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// InsertRecord Inserts a record into a collection.
func InsertRecord(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var body map[string]interface{}
	err = flags.ParseFlag(cmd.Flags(), "body", &body)
	if err != nil {
		return fmt.Errorf(`error parsing "body": ` + err.Error())
	}
	var collection string
	err = flags.ParseFlag(cmd.Flags(), "collection", &collection)
	if err != nil {
		return fmt.Errorf(`error parsing "collection": ` + err.Error())
	}

	resp, err := client.KVStoreService.InsertRecord(collection, body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// InsertRecords Inserts multiple records in a single request.
func InsertRecords(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var allow_updatesDefault bool
	allow_updates := &allow_updatesDefault
	err = flags.ParseFlag(cmd.Flags(), "allow-updates", &allow_updates)
	if err != nil {
		return fmt.Errorf(`error parsing "allow-updates": ` + err.Error())
	}
	var body []map[string]interface{}
	err = flags.ParseFlag(cmd.Flags(), "body", &body)
	if err != nil {
		return fmt.Errorf(`error parsing "body": ` + err.Error())
	}
	var collection string
	err = flags.ParseFlag(cmd.Flags(), "collection", &collection)
	if err != nil {
		return fmt.Errorf(`error parsing "collection": ` + err.Error())
	}
	// Form query params
	generated_query := model.InsertRecordsQueryParams{}
	generated_query.Allow_updates = allow_updates

	resp, err := client.KVStoreService.InsertRecords(collection, body, &generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListIndexes Returns a list of all indexes on a collection.
func ListIndexes(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var collection string
	err = flags.ParseFlag(cmd.Flags(), "collection", &collection)
	if err != nil {
		return fmt.Errorf(`error parsing "collection": ` + err.Error())
	}

	resp, err := client.KVStoreService.ListIndexes(collection)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListRecords Use key-value query parameters to filter fields. Fields are implicitly ANDed and values for the same field are implicitly ORed.
func ListRecords(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var collection string
	err = flags.ParseFlag(cmd.Flags(), "collection", &collection)
	if err != nil {
		return fmt.Errorf(`error parsing "collection": ` + err.Error())
	}
	var countDefault int32
	count := &countDefault
	err = flags.ParseFlag(cmd.Flags(), "count", &count)
	if err != nil {
		return fmt.Errorf(`error parsing "count": ` + err.Error())
	}
	var fields []string
	err = flags.ParseFlag(cmd.Flags(), "fields", &fields)
	if err != nil {
		return fmt.Errorf(`error parsing "fields": ` + err.Error())
	}
	var filters map[string]interface{}
	err = flags.ParseFlag(cmd.Flags(), "filters", &filters)
	if err != nil {
		return fmt.Errorf(`error parsing "filters": ` + err.Error())
	}
	var offsetDefault int32
	offset := &offsetDefault
	err = flags.ParseFlag(cmd.Flags(), "offset", &offset)
	if err != nil {
		return fmt.Errorf(`error parsing "offset": ` + err.Error())
	}
	var orderby []string
	err = flags.ParseFlag(cmd.Flags(), "orderby", &orderby)
	if err != nil {
		return fmt.Errorf(`error parsing "orderby": ` + err.Error())
	}
	// Form query params
	generated_query := model.ListRecordsQueryParams{}
	generated_query.Count = count
	generated_query.Fields = fields
	generated_query.Filters = filters
	generated_query.Offset = offset
	generated_query.Orderby = orderby

	resp, err := client.KVStoreService.ListRecords(collection, &generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// Ping Returns the health status from the database.
func Ping(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}

	resp, err := client.KVStoreService.Ping()
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// PutRecord Updates the record with a given key, either by inserting or replacing the record.
func PutRecord(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var body map[string]interface{}
	err = flags.ParseFlag(cmd.Flags(), "body", &body)
	if err != nil {
		return fmt.Errorf(`error parsing "body": ` + err.Error())
	}
	var collection string
	err = flags.ParseFlag(cmd.Flags(), "collection", &collection)
	if err != nil {
		return fmt.Errorf(`error parsing "collection": ` + err.Error())
	}
	var key string
	err = flags.ParseFlag(cmd.Flags(), "key", &key)
	if err != nil {
		return fmt.Errorf(`error parsing "key": ` + err.Error())
	}

	resp, err := client.KVStoreService.PutRecord(collection, key, body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// QueryRecords Returns a list of query records in a collection.
func QueryRecords(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var collection string
	err = flags.ParseFlag(cmd.Flags(), "collection", &collection)
	if err != nil {
		return fmt.Errorf(`error parsing "collection": ` + err.Error())
	}
	var countDefault int32
	count := &countDefault
	err = flags.ParseFlag(cmd.Flags(), "count", &count)
	if err != nil {
		return fmt.Errorf(`error parsing "count": ` + err.Error())
	}
	var fields []string
	err = flags.ParseFlag(cmd.Flags(), "fields", &fields)
	if err != nil {
		return fmt.Errorf(`error parsing "fields": ` + err.Error())
	}
	var offsetDefault int32
	offset := &offsetDefault
	err = flags.ParseFlag(cmd.Flags(), "offset", &offset)
	if err != nil {
		return fmt.Errorf(`error parsing "offset": ` + err.Error())
	}
	var orderby []string
	err = flags.ParseFlag(cmd.Flags(), "orderby", &orderby)
	if err != nil {
		return fmt.Errorf(`error parsing "orderby": ` + err.Error())
	}
	var query string
	err = flags.ParseFlag(cmd.Flags(), "query", &query)
	if err != nil {
		return fmt.Errorf(`error parsing "query": ` + err.Error())
	}
	// Form query params
	generated_query := model.QueryRecordsQueryParams{}
	generated_query.Count = count
	generated_query.Fields = fields
	generated_query.Offset = offset
	generated_query.Orderby = orderby
	generated_query.Query = query

	resp, err := client.KVStoreService.QueryRecords(collection, &generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}
