package main

import "fmt"
import "github.com/jung-kurt/gofpdf"

func main() { //210*297mm
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	pdf.Image("gopher.png", 10, 5, 50, 0, false, "", 0, "")

	pdf.SetXY(56.2, 30)
	pdf.Write(0, "FACCIAMO PRATICA CON GOLANG") //PRAT-ICA

	pdf.Line(81, 34, 133, 34)

	pdf.SetTextColor(255, 0, 0)
	pdf.SetXY(41, 50)
	pdf.Write(0, "Tutti erano in fila, tranne i byte che erano in files...")//trann-e

  // pdf.Line(105, 0, 105, 297)

	err := pdf.OutputFileAndClose("generated.pdf")
	if err != nil {
		fmt.Printf("Error: ", err)
	}
}
