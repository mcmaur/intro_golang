package main

import (
    "net/http"
    "log"

    "github.com/gorilla/mux"
    "github.com/jung-kurt/gofpdf"

)

func YourHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello World!\n"))
}

func PdfDownload(w http.ResponseWriter, r *http.Request) {
  pdf := gofpdf.New("P", "mm", "A4", "")
  pdf.AddPage()
  pdf.SetFont("Arial", "B", 16)

  pdf.Image("../demo3_pdfcreation/gopher.png", 10, 5, 50, 0, false, "", 0, "")

  pdf.SetXY(56.2, 30)
  pdf.Write(0, "FACCIAMO PRATICA CON GOLANG") //PRAT-ICA

  pdf.Line(81, 34, 133, 34)

  pdf.SetTextColor(255, 0, 0)
  pdf.SetXY(41, 50)
  pdf.Write(0, "Tutti erano in fila, tranne i byte che erano in files...")//trann-e

  w.Header().Set("Content-Disposition", "attachment; filename=download.pdf")
  w.Header().Set("Content-Type","application/pdf")

  pdf.Output(w)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", YourHandler)
    r.HandleFunc("/pdf", PdfDownload)

    log.Fatal(http.ListenAndServe(":8080", r))
}
