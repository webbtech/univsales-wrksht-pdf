package pdf

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/jung-kurt/gofpdf"
	"github.com/pulpfree/univsales-wrksht-pdf/awsservices"
	"github.com/pulpfree/univsales-wrksht-pdf/config"
	"github.com/pulpfree/univsales-wrksht-pdf/model"
)

const (
	dateFormat        = "January 2, 2006"
	coAddressStreet   = "2514 Hwy 20 E"
	coAddressCity     = "Welland"
	coAddressPostal   = "L3B 5N5"
	coAddressProvince = "Ontario"
	coDomain          = "universalwindows.ca"
)

// PDF struct
type PDF struct {
	Request        *Request
	cfg            *config.Config
	outputFileName string
	pdf            *gofpdf.Fpdf
	q              *model.Quote
}

// Request struct
type Request struct {
	QuoteID string `json:"quoteID"`
}

// New function
func New(r *Request, q *model.Quote, cfg *config.Config) *PDF {
	return &PDF{
		Request: r,
		cfg:     cfg,
		q:       q,
	}
}

// OutputToDisk method
func (p *PDF) OutputToDisk() (err error) {
	outputPath := "../tmp/wrksht.pdf"
	err = p.pdf.OutputFileAndClose(outputPath)
	return err
}

// SaveToS3 method
func (p *PDF) SaveToS3() (location string, err error) {
	var buf bytes.Buffer
	if err := p.pdf.Output(&buf); err != nil {
		return "", err
	}
	location, err = awsservices.PutFile(p.outputFileName, &buf, p.cfg)
	return location, err
}

// ================================ Helper Methods

func (p *PDF) setOutputFileName() {

	var retStr []string

	retStr = []string{"worksheet/sht-", strconv.Itoa(p.q.Number), ".pdf"}

	p.outputFileName = strings.Join(retStr, "")
}
