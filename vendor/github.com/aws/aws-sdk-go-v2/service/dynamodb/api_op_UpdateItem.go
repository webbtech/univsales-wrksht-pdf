// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of an UpdateItem operation.
type UpdateItemInput struct {
	_ struct{} `type:"structure"`

	// This is a legacy parameter. Use UpdateExpression instead. For more information,
	// see AttributeUpdates (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.AttributeUpdates.html)
	// in the Amazon DynamoDB Developer Guide.
	AttributeUpdates map[string]AttributeValueUpdate `type:"map"`

	// A condition that must be satisfied in order for a conditional update to succeed.
	//
	// An expression can contain any of the following:
	//
	//    * Functions: attribute_exists | attribute_not_exists | attribute_type
	//    | contains | begins_with | size These function names are case-sensitive.
	//
	//    * Comparison operators: = | <> | < | > | <= | >= | BETWEEN | IN
	//
	//    * Logical operators: AND | OR | NOT
	//
	// For more information about condition expressions, see Specifying Conditions
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.SpecifyingConditions.html)
	// in the Amazon DynamoDB Developer Guide.
	ConditionExpression *string `type:"string"`

	// This is a legacy parameter. Use ConditionExpression instead. For more information,
	// see ConditionalOperator (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.ConditionalOperator.html)
	// in the Amazon DynamoDB Developer Guide.
	ConditionalOperator ConditionalOperator `type:"string" enum:"true"`

	// This is a legacy parameter. Use ConditionExpression instead. For more information,
	// see Expected (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.Expected.html)
	// in the Amazon DynamoDB Developer Guide.
	Expected map[string]ExpectedAttributeValue `type:"map"`

	// One or more substitution tokens for attribute names in an expression. The
	// following are some use cases for using ExpressionAttributeNames:
	//
	//    * To access an attribute whose name conflicts with a DynamoDB reserved
	//    word.
	//
	//    * To create a placeholder for repeating occurrences of an attribute name
	//    in an expression.
	//
	//    * To prevent special characters in an attribute name from being misinterpreted
	//    in an expression.
	//
	// Use the # character in an expression to dereference an attribute name. For
	// example, consider the following attribute name:
	//
	//    * Percentile
	//
	// The name of this attribute conflicts with a reserved word, so it cannot be
	// used directly in an expression. (For the complete list of reserved words,
	// see Reserved Words (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ReservedWords.html)
	// in the Amazon DynamoDB Developer Guide.) To work around this, you could specify
	// the following for ExpressionAttributeNames:
	//
	//    * {"#P":"Percentile"}
	//
	// You could then use this substitution in an expression, as in this example:
	//
	//    * #P = :val
	//
	// Tokens that begin with the : character are expression attribute values, which
	// are placeholders for the actual value at runtime.
	//
	// For more information about expression attribute names, see Specifying Item
	// Attributes (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html)
	// in the Amazon DynamoDB Developer Guide.
	ExpressionAttributeNames map[string]string `type:"map"`

	// One or more values that can be substituted in an expression.
	//
	// Use the : (colon) character in an expression to dereference an attribute
	// value. For example, suppose that you wanted to check whether the value of
	// the ProductStatus attribute was one of the following:
	//
	// Available | Backordered | Discontinued
	//
	// You would first need to specify ExpressionAttributeValues as follows:
	//
	// { ":avail":{"S":"Available"}, ":back":{"S":"Backordered"}, ":disc":{"S":"Discontinued"}
	// }
	//
	// You could then use these values in an expression, such as this:
	//
	// ProductStatus IN (:avail, :back, :disc)
	//
	// For more information on expression attribute values, see Condition Expressions
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.SpecifyingConditions.html)
	// in the Amazon DynamoDB Developer Guide.
	ExpressionAttributeValues map[string]AttributeValue `type:"map"`

	// The primary key of the item to be updated. Each element consists of an attribute
	// name and a value for that attribute.
	//
	// For the primary key, you must provide all of the attributes. For example,
	// with a simple primary key, you only need to provide a value for the partition
	// key. For a composite primary key, you must provide values for both the partition
	// key and the sort key.
	//
	// Key is a required field
	Key map[string]AttributeValue `type:"map" required:"true"`

	// Determines the level of detail about provisioned throughput consumption that
	// is returned in the response:
	//
	//    * INDEXES - The response includes the aggregate ConsumedCapacity for the
	//    operation, together with ConsumedCapacity for each table and secondary
	//    index that was accessed. Note that some operations, such as GetItem and
	//    BatchGetItem, do not access any indexes at all. In these cases, specifying
	//    INDEXES will only return ConsumedCapacity information for table(s).
	//
	//    * TOTAL - The response includes only the aggregate ConsumedCapacity for
	//    the operation.
	//
	//    * NONE - No ConsumedCapacity details are included in the response.
	ReturnConsumedCapacity ReturnConsumedCapacity `type:"string" enum:"true"`

	// Determines whether item collection metrics are returned. If set to SIZE,
	// the response includes statistics about item collections, if any, that were
	// modified during the operation are returned in the response. If set to NONE
	// (the default), no statistics are returned.
	ReturnItemCollectionMetrics ReturnItemCollectionMetrics `type:"string" enum:"true"`

	// Use ReturnValues if you want to get the item attributes as they appear before
	// or after they are updated. For UpdateItem, the valid values are:
	//
	//    * NONE - If ReturnValues is not specified, or if its value is NONE, then
	//    nothing is returned. (This setting is the default for ReturnValues.)
	//
	//    * ALL_OLD - Returns all of the attributes of the item, as they appeared
	//    before the UpdateItem operation.
	//
	//    * UPDATED_OLD - Returns only the updated attributes, as they appeared
	//    before the UpdateItem operation.
	//
	//    * ALL_NEW - Returns all of the attributes of the item, as they appear
	//    after the UpdateItem operation.
	//
	//    * UPDATED_NEW - Returns only the updated attributes, as they appear after
	//    the UpdateItem operation.
	//
	// There is no additional cost associated with requesting a return value aside
	// from the small network and processing overhead of receiving a larger response.
	// No read capacity units are consumed.
	//
	// The values returned are strongly consistent.
	ReturnValues ReturnValue `type:"string" enum:"true"`

	// The name of the table containing the item to update.
	//
	// TableName is a required field
	TableName *string `min:"3" type:"string" required:"true"`

	// An expression that defines one or more attributes to be updated, the action
	// to be performed on them, and new values for them.
	//
	// The following action values are available for UpdateExpression.
	//
	//    * SET - Adds one or more attributes and values to an item. If any of these
	//    attributes already exist, they are replaced by the new values. You can
	//    also use SET to add or subtract from an attribute that is of type Number.
	//    For example: SET myNum = myNum + :val SET supports the following functions:
	//    if_not_exists (path, operand) - if the item does not contain an attribute
	//    at the specified path, then if_not_exists evaluates to operand; otherwise,
	//    it evaluates to path. You can use this function to avoid overwriting an
	//    attribute that may already be present in the item. list_append (operand,
	//    operand) - evaluates to a list with a new element added to it. You can
	//    append the new element to the start or the end of the list by reversing
	//    the order of the operands. These function names are case-sensitive.
	//
	//    * REMOVE - Removes one or more attributes from an item.
	//
	//    * ADD - Adds the specified value to the item, if the attribute does not
	//    already exist. If the attribute does exist, then the behavior of ADD depends
	//    on the data type of the attribute: If the existing attribute is a number,
	//    and if Value is also a number, then Value is mathematically added to the
	//    existing attribute. If Value is a negative number, then it is subtracted
	//    from the existing attribute. If you use ADD to increment or decrement
	//    a number value for an item that doesn't exist before the update, DynamoDB
	//    uses 0 as the initial value. Similarly, if you use ADD for an existing
	//    item to increment or decrement an attribute value that doesn't exist before
	//    the update, DynamoDB uses 0 as the initial value. For example, suppose
	//    that the item you want to update doesn't have an attribute named itemcount,
	//    but you decide to ADD the number 3 to this attribute anyway. DynamoDB
	//    will create the itemcount attribute, set its initial value to 0, and finally
	//    add 3 to it. The result will be a new itemcount attribute in the item,
	//    with a value of 3. If the existing data type is a set and if Value is
	//    also a set, then Value is added to the existing set. For example, if the
	//    attribute value is the set [1,2], and the ADD action specified [3], then
	//    the final attribute value is [1,2,3]. An error occurs if an ADD action
	//    is specified for a set attribute and the attribute type specified does
	//    not match the existing set type. Both sets must have the same primitive
	//    data type. For example, if the existing data type is a set of strings,
	//    the Value must also be a set of strings. The ADD action only supports
	//    Number and set data types. In addition, ADD can only be used on top-level
	//    attributes, not nested attributes.
	//
	//    * DELETE - Deletes an element from a set. If a set of values is specified,
	//    then those values are subtracted from the old set. For example, if the
	//    attribute value was the set [a,b,c] and the DELETE action specifies [a,c],
	//    then the final attribute value is [b]. Specifying an empty set is an error.
	//    The DELETE action only supports set data types. In addition, DELETE can
	//    only be used on top-level attributes, not nested attributes.
	//
	// You can have many actions in a single expression, such as the following:
	// SET a=:value1, b=:value2 DELETE :value3, :value4, :value5
	//
	// For more information on update expressions, see Modifying Items and Attributes
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.Modifying.html)
	// in the Amazon DynamoDB Developer Guide.
	UpdateExpression *string `type:"string"`
}

// String returns the string representation
func (s UpdateItemInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateItemInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateItemInput"}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}

	if s.TableName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TableName"))
	}
	if s.TableName != nil && len(*s.TableName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("TableName", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of an UpdateItem operation.
type UpdateItemOutput struct {
	_ struct{} `type:"structure"`

	// A map of attribute values as they appear before or after the UpdateItem operation,
	// as determined by the ReturnValues parameter.
	//
	// The Attributes map is only present if ReturnValues was specified as something
	// other than NONE in the request. Each element represents one attribute.
	Attributes map[string]AttributeValue `type:"map"`

	// The capacity units consumed by the UpdateItem operation. The data returned
	// includes the total provisioned throughput consumed, along with statistics
	// for the table and any indexes involved in the operation. ConsumedCapacity
	// is only returned if the ReturnConsumedCapacity parameter was specified. For
	// more information, see Provisioned Throughput (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughputIntro.html)
	// in the Amazon DynamoDB Developer Guide.
	ConsumedCapacity *ConsumedCapacity `type:"structure"`

	// Information about item collections, if any, that were affected by the UpdateItem
	// operation. ItemCollectionMetrics is only returned if the ReturnItemCollectionMetrics
	// parameter was specified. If the table does not have any local secondary indexes,
	// this information is not returned in the response.
	//
	// Each ItemCollectionMetrics element consists of:
	//
	//    * ItemCollectionKey - The partition key value of the item collection.
	//    This is the same as the partition key value of the item itself.
	//
	//    * SizeEstimateRangeGB - An estimate of item collection size, in gigabytes.
	//    This value is a two-element array containing a lower bound and an upper
	//    bound for the estimate. The estimate includes the size of all the items
	//    in the table, plus the size of all attributes projected into all of the
	//    local secondary indexes on that table. Use this estimate to measure whether
	//    a local secondary index is approaching its size limit. The estimate is
	//    subject to change over time; therefore, do not rely on the precision or
	//    accuracy of the estimate.
	ItemCollectionMetrics *ItemCollectionMetrics `type:"structure"`
}

// String returns the string representation
func (s UpdateItemOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateItem = "UpdateItem"

// UpdateItemRequest returns a request value for making API operation for
// Amazon DynamoDB.
//
// Edits an existing item's attributes, or adds a new item to the table if it
// does not already exist. You can put, delete, or add attribute values. You
// can also perform a conditional update on an existing item (insert a new attribute
// name-value pair if it doesn't exist, or replace an existing name-value pair
// if it has certain expected attribute values).
//
// You can also return the item's attribute values in the same UpdateItem operation
// using the ReturnValues parameter.
//
//    // Example sending a request using UpdateItemRequest.
//    req := client.UpdateItemRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/UpdateItem
func (c *Client) UpdateItemRequest(input *UpdateItemInput) UpdateItemRequest {
	op := &aws.Operation{
		Name:       opUpdateItem,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateItemInput{}
	}

	req := c.newRequest(op, input, &UpdateItemOutput{})

	if req.Config.EnableEndpointDiscovery {
		de := discovererDescribeEndpoints{
			Client:        c,
			Required:      false,
			EndpointCache: c.endpointCache,
			Params: map[string]*string{
				"op": &req.Operation.Name,
			},
		}

		for k, v := range de.Params {
			if v == nil {
				delete(de.Params, k)
			}
		}

		req.Handlers.Build.PushFrontNamed(aws.NamedHandler{
			Name: "crr.endpointdiscovery",
			Fn:   de.Handler,
		})
	}

	return UpdateItemRequest{Request: req, Input: input, Copy: c.UpdateItemRequest}
}

// UpdateItemRequest is the request type for the
// UpdateItem API operation.
type UpdateItemRequest struct {
	*aws.Request
	Input *UpdateItemInput
	Copy  func(*UpdateItemInput) UpdateItemRequest
}

// Send marshals and sends the UpdateItem API request.
func (r UpdateItemRequest) Send(ctx context.Context) (*UpdateItemResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateItemResponse{
		UpdateItemOutput: r.Request.Data.(*UpdateItemOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateItemResponse is the response type for the
// UpdateItem API operation.
type UpdateItemResponse struct {
	*UpdateItemOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateItem request.
func (r *UpdateItemResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}