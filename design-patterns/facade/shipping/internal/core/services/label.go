package services

import (
	"fmt"
	"log"
	"os"
	"github.com/signintech/gopdf"
)

type LabelService struct {}

func (label *LabelService) CreateShippingLabel(name string) error {
	fmt.Println("Creating shipping label...")
	var pdFile gopdf.GoPdf
	pdFile.Start(gopdf.Config{
		PageSize: gopdf.Rect{W: 50, H: 75},
		Unit: gopdf.Unit_CM,
	})

	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	err = pdFile.AddTTFFont("graphik-bold", fmt.Sprintf("%s/internal/core/fonts/image-master-font-gofont-ttfs/Go-Mono.ttf", workingDir))
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = pdFile.AddTTFFont("regular", fmt.Sprintf("%s/internal/core/fonts/image-master-font-gofont-ttfs/Go-Regular.ttf", workingDir))
	if err != nil {
		log.Fatal(err)
		return err
	}
	const (
		xHeaderLeft, xFooterLeft float64 = 25, 25
		yTitleOneLineOne, yFooterLine float64 = 3, 70
	)
	pdFile.AddHeader(func()  {
		err = pdFile.SetFont("graphik-bold", "", 45)
		if err != nil {
			log.Fatal(err)
			return
		}
		// CloudShipping Headings
		pdFile.SetXY(xHeaderLeft, yTitleOneLineOne)
		err = pdFile.Text("CloudShippings Inc.")
		if err != nil {
			log.Fatal(err)
			return
		}
		pdFile.SetLineWidth(0.1)
		pdFile.Line(2, 3.5, 48, 3.5)

		// From Configuration
		err = pdFile.SetFont("regular", "", 40)
		if err != nil {
			log.Fatal(err)
			return
		}
		pdFile.SetXY(xHeaderLeft - 23, 3.5 + 4.0)
		err = pdFile.Text("From:")
		if err != nil {
			log.Fatal(err)
			return
		}
		pdFile.SetXY(xHeaderLeft - 21, yTitleOneLineOne + 6.0)
		err = pdFile.Text(fmt.Sprintf("No: %s", "41"))
		if err != nil {
			log.Fatal(err)
			return
		}
		pdFile.SetXY(xHeaderLeft - 21, yTitleOneLineOne + 8.0)
		err = pdFile.Text(fmt.Sprintf("Street-Name: %s", "Jibowu Estate Road, U-turn Busstop,"))
		if err != nil {
			log.Fatal(err)
			return
		}
		pdFile.SetXY(xHeaderLeft - 21, yTitleOneLineOne + 10.0)
		err = pdFile.Text(fmt.Sprintf("Province: %s", "Abule-Egba,"))
		if err != nil {
			log.Fatal(err)
			return
		}
		pdFile.SetXY(xHeaderLeft - 21, yTitleOneLineOne + 12.0)
		err = pdFile.Text(fmt.Sprintf("State: %s", "Lagos,"))
		if err != nil {
			log.Fatal(err)
			return
		}
		pdFile.SetXY(xHeaderLeft - 21, yTitleOneLineOne + 14.0)
		err = pdFile.Text(fmt.Sprintf("Country: %s", "Nigeria."))
		if err != nil {
			log.Fatal(err)
			return
		}
		pdFile.SetLineWidth(0.1)
		pdFile.Line(2, 19.5, 48, 19.5)

		// To Configuration
		pdFile.SetXY(xHeaderLeft - 23, 19.5 + 3.0)
		err = pdFile.Text("To:")
		if err != nil {
			log.Fatal(err)
			return
		}
		pdFile.SetXY(xHeaderLeft - 21, 19.5 + 5.0)
		err = pdFile.Text(fmt.Sprintf("No: %s", "42"))
		if err != nil {
			log.Fatal(err)
			return
		}
		pdFile.SetXY(xHeaderLeft - 21, 19.5 + 7.0)
		err = pdFile.Text(fmt.Sprintf("Street-Name: %s", "Jibowu Estate Road, U-turn Busstop,"))
		if err != nil {
			log.Fatal(err)
			return
		}

		pdFile.SetXY(xHeaderLeft - 21, 19.5 + 9.0)
		err = pdFile.Text(fmt.Sprintf("Province: %s", "Abule-Egba,"))
		if err != nil {
			log.Fatal(err)
			return
		}

		pdFile.SetXY(xHeaderLeft - 21, 19.5 + 11.0)
		err = pdFile.Text(fmt.Sprintf("State: %s", "Lagos,"))
		if err != nil {
			log.Fatal(err)
			return
		}

		pdFile.SetXY(xHeaderLeft - 21, 19.5 + 13.0)
		err = pdFile.Text(fmt.Sprintf("Country: %s", "Nigeria."))
		if err != nil {
			log.Fatal(err)
			return
		}

		pdFile.SetLineWidth(0.1)
		pdFile.Line(2, 34.0, 48, 34.0)

		// Product Configuration
		pdFile.SetXY(xHeaderLeft - 23, 37)
		err = pdFile.Text("Product Details:")
		if err != nil {
			log.Fatal(err)
			return
		}
		pdFile.SetXY(xHeaderLeft - 21, 39)
		err = pdFile.Text(fmt.Sprintf("Product Identity: %s", name))
		if err != nil {
			log.Fatal(err)
			return
		}
		pdFile.SetXY(xHeaderLeft - 21, 41)
		err = pdFile.Text(fmt.Sprintf("Product Tracking Identity: %s.", name))
		if err != nil {
			log.Fatal(err)
			return
		}
		pdFile.SetXY(xHeaderLeft - 21, 43)
		err = pdFile.Text(fmt.Sprintf("Product Weight: %s %s", "10", "Kg."))
		if err != nil {
			log.Fatal(err)
			return
		}
		pdFile.SetXY(xHeaderLeft - 21, 45)
		err = pdFile.Text(fmt.Sprintf("Product Carrier: %s", "UPS"))
		if err != nil {
			log.Fatal(err)
			return
		}
	})

	pdFile.AddFooter(func() {
		err = pdFile.SetFont("graphik-bold", "", 20)
		if err != nil {
			log.Fatal(err)
			return
		}
		pdFile.SetLineWidth(0.1)
		pdFile.Line(2, 73, 48, 73)
		pdFile.SetXY(2, 74.0)
		err = pdFile.Text("Powered by: CloudShippings Inc.")
		if err != nil {
			log.Fatal(err)
			return
		}
	})
	pdFile.AddPage()
	// pdFile.AddFooter()
	err = pdFile.WritePdf(fmt.Sprintf("%s/internal/pdf/%s.pdf", workingDir, name))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Shipping Label created successfully.")
	return nil
}
