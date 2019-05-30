package pdf

import (
	"os"
	"regexp"
	"testing"

	"github.com/pulpfree/univsales-wrksht-pdf/config"
	"github.com/pulpfree/univsales-wrksht-pdf/model"
	"github.com/pulpfree/univsales-wrksht-pdf/model/mongo"
	"github.com/stretchr/testify/suite"
)

// IntegSuite struct
type IntegSuite struct {
	suite.Suite
	cfg *config.Config
	db  model.DBHandler
	q   *model.Quote
	p   *PDF
}

const (
	defaultsFP = "../config/defaults.yml"
	quoteID    = "5cd16f18699e0300c7b10d30"
	// quoteID = "5c9d21f0f1c8a86cac0adcbc"
	// quoteID = "5c880187b5342edbda202712"
)

// SetupTest method
func (suite *IntegSuite) SetupTest() {

	req := &Request{
		QuoteID: quoteID,
	}
	// setup config
	os.Setenv("Stage", "test")
	suite.cfg = &config.Config{DefaultsFilePath: defaultsFP}
	err := suite.cfg.Load()
	suite.NoError(err)

	suite.db, err = mongo.NewDB(suite.cfg.GetMongoConnectURL(), suite.cfg.DBName)
	suite.NoError(err)

	suite.q, err = suite.db.FetchQuote(quoteID)
	suite.NoError(err)

	suite.p = New(req, suite.q, suite.cfg)
	suite.NoError(err)
}

func (suite *IntegSuite) TestTypes() {
	suite.True(suite.q.Number > 0)
	suite.IsType(&model.Quote{}, suite.q)
	suite.IsType(&PDF{}, suite.p)
}

func (suite *IntegSuite) TestSetFileNameOutput() {

	// start with quote
	req := &Request{
		QuoteID: quoteID,
	}
	suite.p = New(req, suite.q, suite.cfg)
	suite.p.setOutputFileName()
	r, _ := regexp.Compile("^worksheet\\/sht-([0-9]+)\\.pdf?")
	suite.True(r.MatchString(suite.p.outputFileName))
}

func (suite *IntegSuite) TestOutputToDisk() {

	req := &Request{
		QuoteID: quoteID,
	}
	suite.cfg.SetStageEnv("test")
	suite.p = New(req, suite.q, suite.cfg)

	err := suite.p.WorkSheet()
	suite.NoError(err)

	err = suite.p.OutputToDisk()
	suite.NoError(err)
}

// TestIntegrationSuite function
func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, new(IntegSuite))
}
