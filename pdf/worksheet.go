package pdf

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/jung-kurt/gofpdf"
)

// Various constants
const (
	hdrSize       float64 = 8.5
	pfSize        float64 = 10
	topBr         float64 = 1.5
	midBr         float64 = 9
	hSep          float64 = 2
	newLnMaxLen   int     = 30
	HSTMultiplier float32 = 1.15
	moneyFormat           = "#,###.##"
)

// WorkSheet method
func (p *PDF) WorkSheet() (err error) {

	p.setOutputFileName()
	titleStr := "Worksheet " + strconv.Itoa(p.q.Number) + " PDF"

	p.pdf = gofpdf.New("P", "mm", "Letter", "")
	pdf := p.pdf
	pdf.SetTitle(titleStr, false)
	pdf.SetAuthor(p.cfg.DocAuthor, false)

	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFont("Arial", "I", 9)
		pdf.CellFormat(0, 10, fmt.Sprintf("Page %d of {nb}", pdf.PageNo()), "", 0, "C", false, 0, "")
	})
	pdf.AliasNbPages("")

	pdf.AddPage()
	p.quoteTitle()
	p.groupList()
	p.windowList()
	p.otherList()
	p.featureList()

	return err
}

func (p *PDF) quoteTitle() {

	pdf := p.pdf
	q := p.q
	fmt.Printf("q %+v\n", q.Customer)

	var (
		rsp     *http.Response
		tp      string
		imgInfo gofpdf.ImageOptions
	)

	rsp, err := http.Get(p.cfg.LogoURI)
	if err == nil {
		tp = pdf.ImageTypeFromMime(rsp.Header["Content-Type"][0])
		imgInfo = gofpdf.ImageOptions{ImageType: tp}
		pdf.RegisterImageReader(p.cfg.LogoURI, tp, rsp.Body)
	} else {
		pdf.SetError(err)
	}
	custName := fmt.Sprintf("%s %s", q.Customer.Name.First, q.Customer.Name.Last)
	address2 := fmt.Sprintf("%s, %s. %s", q.Customer.Address.City, q.Customer.Address.Province, q.Customer.Address.PostalCode)
	quoteNo := fmt.Sprintf("%d", q.Number)

	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(0, 6, "Worksheet", "", 2, "", false, 0, "")

	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(0, 5.5, custName, "", 2, "", false, 0, "")
	pdf.CellFormat(0, 5.5, q.Customer.Address.Street1, "", 2, "", false, 0, "")
	pdf.CellFormat(0, 5.5, address2, "", 2, "", false, 0, "")
	if v, ok := q.Customer.PhoneMap["mobile"]; ok {
		pdf.CellFormat(0, 5.5, fmt.Sprintf("Mobile %s", v), "", 2, "", false, 0, "")
	}
	if v, ok := q.Customer.PhoneMap["home"]; ok {
		pdf.CellFormat(0, 5.5, fmt.Sprintf("Home %s", v), "", 2, "", false, 0, "")
	}
	pdf.SetTextColor(0, 0, 200)
	pdf.SetFont("Arial", "U", 12)
	if q.Customer.Email != "" {
		pdf.CellFormat(0, 5.5, q.Customer.Email, "", 2, "", false, 0, fmt.Sprintf("mailto:%s", q.Customer.Email))
	}
	pdf.MoveTo(90, 10)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(0, 6, "Invoice No", "", 2, "", false, 0, "")
	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(0, 7, quoteNo, "", 2, "", false, 0, "")
	pdf.ImageOptions(p.cfg.LogoURI, 160, 10, 45, 0, false, imgInfo, 0, fmt.Sprintf("https://%s", coDomain))

	pdf.MoveTo(160, 30)
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(0, 5, coAddressStreet, "", 2, "", false, 0, "")
	pdf.CellFormat(0, 5, fmt.Sprintf("%s, %s %s", coAddressCity, coAddressProvince, coAddressPostal), "", 2, "", false, 0, "")
	pdf.SetTextColor(0, 0, 200)
	pdf.SetFont("Arial", "U", 10)
	pdf.CellFormat(0, 5, coDomain, "", 2, "", false, 0, fmt.Sprintf("https://%s", coDomain))

	pdf.MoveTo(10, 50)
	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(10, 5, "Notes:", "", 0, "", false, 0, "")
	pdf.SetFont("Arial", "I", 9)
	pdf.CellFormat(0, 5, q.Customer.Notes, "", 2, "", false, 0, "")

	pdf.Ln(4)
}

func (p *PDF) groupList() {

	pdf := p.pdf
	q := p.q

	if len(q.Items.Group) <= 0 {
		return
	}

	pdf.SetFillColor(100, 100, 100)
	pdf.SetDrawColor(100, 100, 100)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(0, 6, "Groups", "B", 0, "", false, 0, "")
	pdf.Ln(midBr)

	ctr := 1
	for _, g := range q.Items.Group {
		openWidth := fmt.Sprintf("%d %s x %d %s", g.Dims.Width.Inch, g.Dims.Width.Fraction, g.Dims.Height.Inch, g.Dims.Height.Fraction)

		installType := "none"
		if g.Specs["installType"] != nil {
			installType = g.Specs["installType"].(string)
		}

		pdf.SetFont("Arial", "B", pfSize)
		pdf.CellFormat(0, 6, fmt.Sprintf("%d%s", ctr, ")"), "", 1, "", false, 0, "")

		pdf.SetDrawColor(200, 200, 200)
		pdf.SetFont("Arial", "", hdrSize)
		pdf.MoveTo(pdf.GetX()+4, pdf.GetY())
		pdf.CellFormat(7, 5, "Qty", "", 0, "", false, 0, "")
		pdf.CellFormat(26, 5, "Rooms", "", 0, "", false, 0, "")
		pdf.CellFormat(35, 5, "Opening Width", "", 0, "", false, 0, "")
		pdf.CellFormat(45, 5, "Type", "", 1, "", false, 0, "")

		pdf.SetFont("Arial", "", pfSize)
		pdf.MoveTo(pdf.GetX()+4, pdf.GetY())
		pdf.CellFormat(7, 6, fmt.Sprintf("%d", g.Qty), "TB", 0, "", false, 0, "")
		pdf.CellFormat(26, 6, fmt.Sprintf("%s", strings.Join(g.Rooms, ", ")), "TB", 0, "", false, 0, "")
		pdf.CellFormat(35, 6, openWidth, "TB", 0, "", false, 0, "")
		pdf.CellFormat(45, 6, g.Specs["groupTypeDescription"].(string), "TB", 1, "", false, 0, "")

		pdf.MoveTo(pdf.GetX()+4, pdf.GetY()+2)
		pdf.SetFont("Arial", "", hdrSize)
		pdf.CellFormat(5, 6, "Windows", "", 1, "", false, 0, "")

		pdf.SetFont("Arial", "", pfSize)
		for _, item := range g.Items {
			winSize := fmt.Sprintf("%d %s x %d %s", item.Dims.Width.Inch, item.Dims.Width.Fraction, item.Dims.Height.Inch, item.Dims.Height.Fraction)

			pdf.MoveTo(pdf.GetX()+8, pdf.GetY())
			pdf.CellFormat(7, 6, fmt.Sprintf("%d", item.Qty), "B", 0, "", false, 0, "")
			pdf.CellFormat(30, 6, winSize, "B", 0, "", false, 0, "")
			pdf.CellFormat(60, 6, fmt.Sprintf("%s", item.Product["name"]), "B", 1, "", false, 0, "")
		}

		// start specs
		pdf.MoveTo(pdf.GetX()+4, pdf.GetY()+2)
		pdf.SetFont("Arial", "", hdrSize)
		pdf.CellFormat(5, 6, "Specifications", "", 1, "", false, 0, "")

		pdf.SetFont("Arial", "", pfSize)
		pdf.MoveTo(pdf.GetX()+8, pdf.GetY())
		pdf.CellFormat(25, 6, "Install Type", "", 0, "", false, 0, "")
		pdf.CellFormat(60, 6, installType, "", 1, "", false, 0, "")

		pdf.MoveTo(pdf.GetX()+8, pdf.GetY()+2)
		pdf.CellFormat(25, 6, "Trim", "", 0, "", false, 0, "")
		pdf.MultiCell(70, 4.5, setNewLines(g.Specs["trim"]), "", "LB", false)

		pdf.MoveTo(pdf.GetX()+8, pdf.GetY()+2)
		pdf.CellFormat(25, 6, "Options", "", 0, "", false, 0, "")
		pdf.MultiCell(95, 4.5, setNewLines(g.Specs["options"]), "", "LB", false)
		// end specs

		pdf.Ln(4)
		ctr = ctr + 1
	}
	pdf.Ln(4)
}

func (p *PDF) windowList() {

	pdf := p.pdf
	q := p.q
	if len(q.Items.Window) <= 0 {
		return
	}

	pdf.SetFillColor(100, 100, 100)
	pdf.SetDrawColor(100, 100, 100)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(0, 6, "Windows", "B", 0, "", false, 0, "")
	pdf.Ln(midBr)

	ctr := 1
	for _, g := range q.Items.Window {

		windowSize := fmt.Sprintf("%d %s x %d %s", g.Dims.Width.Inch, g.Dims.Width.Fraction, g.Dims.Height.Inch, g.Dims.Height.Fraction)
		installType := "none"
		if g.Specs["installType"] != nil {
			installType = g.Specs["installType"].(string)
		}
		trim := "none"
		if g.Specs["trim"] != nil {
			trim = setNewLines(g.Specs["trim"])
		}

		pdf.SetFont("Arial", "B", pfSize)
		pdf.CellFormat(0, 6, fmt.Sprintf("%d%s", ctr, ")"), "", 1, "", false, 0, "")

		pdf.SetDrawColor(200, 200, 200)
		pdf.SetFont("Arial", "", hdrSize)
		pdf.MoveTo(pdf.GetX()+4, pdf.GetY())
		pdf.CellFormat(7, 5, "Qty", "", 0, "", false, 0, "")
		pdf.CellFormat(26, 5, "Rooms", "", 0, "", false, 0, "")
		pdf.CellFormat(35, 5, "Size", "", 0, "", false, 0, "")
		pdf.CellFormat(45, 5, "Type", "", 1, "", false, 0, "")

		pdf.SetFont("Arial", "", pfSize)
		pdf.MoveTo(pdf.GetX()+4, pdf.GetY())
		pdf.CellFormat(7, 6, fmt.Sprintf("%d", g.Qty), "TB", 0, "", false, 0, "")
		pdf.CellFormat(26, 6, fmt.Sprintf("%s", strings.Join(g.Rooms, ", ")), "TB", 0, "", false, 0, "")
		pdf.CellFormat(35, 6, windowSize, "TB", 0, "", false, 0, "")
		pdf.CellFormat(60, 6, g.ProductName, "TB", 1, "", false, 0, "")

		pdf.MoveTo(pdf.GetX()+4, pdf.GetY()+2)
		pdf.SetFont("Arial", "", hdrSize)
		pdf.CellFormat(5, 6, "Specifications", "", 1, "", false, 0, "")

		pdf.MoveTo(pdf.GetX()+8, pdf.GetY())
		pdf.SetFont("Arial", "", pfSize)
		pdf.CellFormat(25, 6, "Install Type", "", 0, "", false, 0, "")
		pdf.CellFormat(95, 6, installType, "", 1, "", false, 0, "")

		pdf.MoveTo(pdf.GetX()+8, pdf.GetY()+2)
		pdf.CellFormat(25, 6, "Trim", "", 0, "", false, 0, "")
		pdf.MultiCell(70, 4.5, trim, "", "LB", false)

		pdf.MoveTo(pdf.GetX()+8, pdf.GetY()+2)
		pdf.CellFormat(25, 6, "Options", "", 0, "", false, 0, "")
		pdf.MultiCell(95, 4.5, setNewLines(g.Specs["options"]), "", "LB", false)

		pdf.Ln(4)
		ctr = ctr + 1
	}
	pdf.Ln(4)
}

func (p *PDF) otherList() {

	pdf := p.pdf
	q := p.q
	if len(q.Items.Other) <= 0 {
		return
	}

	pdf.SetFillColor(100, 100, 100)
	pdf.SetDrawColor(100, 100, 100)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(0, 6, "Misc Items", "B", 0, "", false, 0, "")
	pdf.Ln(midBr)

	ctr := 1
	for _, g := range q.Items.Other {
		pdf.SetFont("Arial", "B", pfSize)
		pdf.CellFormat(0, 6, fmt.Sprintf("%d%s", ctr, ")"), "", 1, "", false, 0, "")

		pdf.SetDrawColor(200, 200, 200)
		pdf.SetFont("Arial", "", hdrSize)
		pdf.MoveTo(pdf.GetX()+4, pdf.GetY())
		pdf.CellFormat(7, 5, "Qty", "", 0, "", false, 0, "")
		pdf.CellFormat(26, 5, "Rooms", "", 0, "", false, 0, "")
		pdf.CellFormat(35, 5, "Description", "", 1, "", false, 0, "")

		pdf.SetFont("Arial", "", pfSize)
		pdf.MoveTo(pdf.GetX()+4, pdf.GetY())
		pdf.CellFormat(7, 6, fmt.Sprintf("%d", g.Qty), "TB", 0, "", false, 0, "")
		pdf.CellFormat(26, 6, fmt.Sprintf("%s", strings.Join(g.Rooms, ", ")), "TB", 0, "", false, 0, "")
		pdf.CellFormat(60, 6, g.Description, "TB", 1, "", false, 0, "")

		pdf.MoveTo(pdf.GetX()+4, pdf.GetY()+2)
		pdf.SetFont("Arial", "", hdrSize)
		pdf.CellFormat(5, 6, "Specifications", "", 1, "", false, 0, "")

		pdf.MoveTo(pdf.GetX()+8, pdf.GetY())
		pdf.SetFont("Arial", "", pfSize)
		pdf.CellFormat(30, 6, "Options", "", 0, "", false, 0, "")
		// pdf.CellFormat(95, 6, replaceNLDash(g.Specs.Options), "B", 1, "", false, 0, "")
		pdf.MultiCell(70, 4.5, setNewLines(g.Specs.Options), "", "LB", false)

		pdf.Ln(4)
		ctr = ctr + 1
	}
	pdf.Ln(4)
}

func (p *PDF) featureList() {
	pdf := p.pdf
	q := p.q
	if q.Features == "" {
		return
	}

	pdf.SetFillColor(100, 100, 100)
	pdf.SetDrawColor(100, 100, 100)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(0, 6, "Job Features", "B", 0, "", false, 0, "")
	pdf.Ln(midBr)

	pdf.SetFont("Arial", "", pfSize)
	pdf.MultiCell(70, 4.5, setNewLines(q.Features), "", "L", false)
}

// ================================ Helper Methods

func setNewLines(name interface{}) string {

	str := name.(string)
	if len(str) <= newLnMaxLen {
		return str
	}

	pcs := strings.Split(str, ",")
	strSl := make([]string, len(pcs))
	for i := 0; i < len(pcs); i++ {
		strSl[i] = strings.TrimSpace(pcs[i])
	}
	retVal := strings.Join(strSl, "\n")

	return retVal
}

func formatMoney(num float64, prefix string) string {
	moneyFormat := "#,###.##"
	return fmt.Sprintf("%s%s", prefix, humanize.FormatFloat(moneyFormat, num))
}

func replaceNLDash(val interface{}) string {
	str := val.(string)
	return strings.Replace(str, "\n", ",  ", -1)
}
