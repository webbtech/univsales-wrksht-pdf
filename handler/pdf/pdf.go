package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pulpfree/univsales-wrksht-pdf/config"
	"github.com/pulpfree/univsales-wrksht-pdf/model"
	"github.com/pulpfree/univsales-wrksht-pdf/model/mongo"
	"github.com/pulpfree/univsales-wrksht-pdf/pdf"
	log "github.com/sirupsen/logrus"
	"github.com/thundra-io/thundra-lambda-agent-go/thundra"
)

// Response data format
type Response struct {
	Code      int         `json:"code"`      // HTTP status code
	Data      interface{} `json:"data"`      // Data payload
	Message   string      `json:"message"`   // Error or status message
	Status    string      `json:"status"`    // Status code (error|fail|success)
	Timestamp int64       `json:"timestamp"` // Machine-readable UTC timestamp in nanoseconds since EPOCH
}

var cfg *config.Config

func init() {
	cfg = &config.Config{}
	err := cfg.Load()
	if err != nil {
		log.Fatal(err)
	}
}

// HandleRequest function
func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	hdrs := make(map[string]string)
	hdrs["Content-Type"] = "application/json"
	t := time.Now()

	// If this is a ping test, intercept and return
	if req.HTTPMethod == "GET" {
		log.Info("Ping test in handleRequest")
		return gatewayResponse(Response{
			Code:      200,
			Data:      "pong",
			Status:    "success",
			Timestamp: t.Unix(),
		}, hdrs), nil
	}

	var r *pdf.Request
	json.Unmarshal([]byte(req.Body), &r)

	db, err := mongo.NewDB(cfg.GetMongoConnectURL(), cfg.DBName)
	if err != nil {
		return gatewayResponse(Response{
			Code:      500,
			Message:   fmt.Sprintf("error: %s", err.Error()),
			Status:    "error",
			Timestamp: t.Unix(),
		}, hdrs), nil
	}

	var q *model.Quote
	q, err = db.FetchQuote(r.QuoteID)
	if err != nil {
		return gatewayResponse(Response{
			Code:      500,
			Message:   fmt.Sprintf("error: %s", err.Error()),
			Status:    "error",
			Timestamp: t.Unix(),
		}, hdrs), nil
	}

	var p *pdf.PDF
	p = pdf.New(r, q, cfg)
	err = p.WorkSheet()

	if err != nil {
		return gatewayResponse(Response{
			Code:      500,
			Message:   fmt.Sprintf("error: %s", err.Error()),
			Status:    "error",
			Timestamp: t.Unix(),
		}, hdrs), nil
	}

	location, err := p.SaveToS3()
	if err != nil {
		return gatewayResponse(Response{
			Code:      500,
			Message:   fmt.Sprintf("error: %s", err.Error()),
			Status:    "error",
			Timestamp: t.Unix(),
		}, hdrs), nil
	}
	log.Infof("Successfully created PDF with location: %s", location)

	return gatewayResponse(Response{
		Code:      201,
		Data:      location,
		Status:    "success",
		Timestamp: t.Unix(),
	}, hdrs), nil
}

func main() {
	lambda.Start(thundra.Wrap(HandleRequest))
}

func gatewayResponse(resp Response, hdrs map[string]string) events.APIGatewayProxyResponse {
	body, _ := json.Marshal(&resp)
	if resp.Status == "error" {
		log.Errorf("Error: status: %s, code: %d, message: %s", resp.Status, resp.Code, resp.Message)
	}
	return events.APIGatewayProxyResponse{Body: string(body), Headers: hdrs, StatusCode: resp.Code}
}
