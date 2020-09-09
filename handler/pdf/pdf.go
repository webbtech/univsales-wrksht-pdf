package main

import (
	"encoding/json"
	"time"

	pres "github.com/pulpfree/lambda-go-proxy-response"
	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/epsagon/epsagon-go/epsagon"
	"github.com/pulpfree/univsales-wrksht-pdf/config"
	"github.com/pulpfree/univsales-wrksht-pdf/model"
	"github.com/pulpfree/univsales-wrksht-pdf/model/mongo"
	"github.com/pulpfree/univsales-wrksht-pdf/pdf"
)

const (
	epsagonAppName = "univsales"
	epsagonToken   = "73993039-d583-43ad-84eb-1a443e257274"
)

var (
	cfg *config.Config
)

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
	hdrs["Access-Control-Allow-Origin"] = "*"
	hdrs["Access-Control-Allow-Methods"] = "GET,OPTIONS,POST,PUT"
	hdrs["Access-Control-Allow-Headers"] = "Authorization,Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token"

	if req.HTTPMethod == "OPTIONS" {
		return events.APIGatewayProxyResponse{Body: string("null"), Headers: hdrs, StatusCode: 200}, nil
	}

	t := time.Now()

	// If this is a ping test, intercept and return
	if req.HTTPMethod == "GET" {
		log.Info("Ping test in handleRequest")
		return pres.ProxyRes(pres.Response{
			Code:      200,
			Data:      "pong",
			Status:    "success",
			Timestamp: t.Unix(),
		}, hdrs, nil), nil
	}

	var r *pdf.Request
	json.Unmarshal([]byte(req.Body), &r)

	db, err := mongo.NewDB(cfg.GetMongoConnectURL(), cfg.DBName)
	if err != nil {
		return pres.ProxyRes(pres.Response{
			Timestamp: t.Unix(),
		}, hdrs, err), nil
	}

	var q *model.Quote
	q, err = db.FetchQuote(r.QuoteID)
	if err != nil {
		return pres.ProxyRes(pres.Response{
			Timestamp: t.Unix(),
		}, hdrs, err), nil
	}

	var p *pdf.PDF
	p = pdf.New(r, q, cfg)
	err = p.WorkSheet()
	if err != nil {
		return pres.ProxyRes(pres.Response{
			Timestamp: t.Unix(),
		}, hdrs, err), nil
	}

	location, err := p.SaveToS3()
	if err != nil {
		return pres.ProxyRes(pres.Response{
			Timestamp: t.Unix(),
		}, hdrs, err), nil
	}
	log.Infof("Successfully created PDF with location: %s", location)

	return pres.ProxyRes(pres.Response{
		Code:      201,
		Data:      location,
		Status:    "success",
		Timestamp: t.Unix(),
	}, hdrs, nil), nil
}

func main() {
	log.Println("enter main")
	lambda.Start(epsagon.WrapLambdaHandler(
		epsagon.NewTracerConfig(epsagonAppName, epsagonToken),
		HandleRequest))
}
