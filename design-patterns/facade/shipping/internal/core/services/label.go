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
		PageSize: *gopdf.PageSizeA4,
		Unit: gopdf.UnitPT,
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

	pdFile.AddHeader(func()  {
		const (
			xHeaderLeft float64 = 20
			// xHeaderRight float64 = 710
			yTitleOneLineOne, yTitleOneLineTwo float64 = 12, 24 
		)
		err = pdFile.SetFont("graphik-bold", "", 10)
		if err != nil {
			log.Fatal(err)
			return
		}
		pdFile.SetXY(xHeaderLeft, yTitleOneLineOne)
		err = pdFile.Text(fmt.Sprintf("Shipping Identity: %s", name))
		if err != nil {
			log.Fatal(err)
			return
		}
		pdFile.SetXY(xHeaderLeft, yTitleOneLineTwo)
		err = pdFile.Text(fmt.Sprintf("Shipping Address: %s", name))
		if err != nil {
			log.Fatal(err)
			return
		}
	})
	// pdFile.SetLineWidth(2)
	pdFile.AddPage()
	err = pdFile.WritePdf(fmt.Sprintf("%s/internal/pdf/%s.pdf", workingDir, name))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Shipping Label created successfully.")
	return nil
}
