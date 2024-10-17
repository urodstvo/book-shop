package orders

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Masterminds/squirrel"
	"github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"
	"github.com/urodstvo/book-shop/libs/models"
)

type PDFOrder struct {
	OrderID            string
	OrderStatus        string
	OrderDate          string
	OrderPaymentNumber string
	CustomerName       string
	Items              []PDFOrderItem
	TotalPrice         float64
}

// OrderItem представляет структуру данных товара в заказе
type PDFOrderItem struct {
	Name     string
	Price    float64
	Quantity int
}

func (h *Orders) GetReport(w http.ResponseWriter, r *http.Request) {
	order_id := mux.Vars(r)["OrderId"]
	user := h.SessionManager.Get(r.Context(), "user").(models.User)
	pdforder := PDFOrder{
		OrderID:      order_id,
		CustomerName: user.Name,
	}

	getOrderQuery := squirrel.Select("o.status, p.card_number, o.created_at, o.price").From(models.Order{}.TableName() + " o").Join(models.Payment{}.TableName() + " p ON o.payment_id = p.id").Where(squirrel.Eq{"o.id": order_id}).Where(squirrel.Eq{"o.user_id": user.Id})
	err := getOrderQuery.RunWith(h.DB).QueryRow().Scan(&pdforder.OrderStatus, &pdforder.OrderPaymentNumber, &pdforder.OrderDate, &pdforder.TotalPrice)
	if err != nil {
		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	getOrderItemsQuery := squirrel.Select("b.name, b.price, ob.amount").From(models.OrderBook{}.TableName() + " ob").Join(models.Book{}.TableName() + " b ON ob.book_id = b.id").Where(squirrel.Eq{"ob.order_id": order_id})

	rows, err := getOrderItemsQuery.RunWith(h.DB).Query()
	if err != nil {
		h.Logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for rows.Next() {
		item := PDFOrderItem{}
		err = rows.Scan(&item.Name, &item.Price, &item.Quantity)
		if err != nil {
			h.Logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		pdforder.Items = append(pdforder.Items, item)
	}

	pdforder.OrderDate = pdforder.OrderDate[:10]
	pdforder.OrderPaymentNumber = fmt.Sprintf("xxxx-xxxx-xxxx-%s", pdforder.OrderPaymentNumber[len(pdforder.OrderPaymentNumber)-4:])

	pdf := createInvoice(pdforder)

	var buf bytes.Buffer

	err = pdf.Output(&buf)
	pdf.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	// w.Header().Set("Content-Disposition", "attachment; filename=отчет.pdf")
	buf.WriteTo(w)

}

func createInvoice(order PDFOrder) *gofpdf.Fpdf {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(20, 10, 10)
	pdf.AddUTF8Font("Roboto", "", "./static/roboto.ttf")
	pdf.SetTitle("Отчет о заказе", true)
	pdf.AddPage()

	pdf.SetFont("Roboto", "", 10)
	pdf.CellFormat(90, 5, `Поставщик: Книжный интернет-магазин`, "", 2, "L", false, 0, "")
	pdf.CellFormat(90, 5, `ИНН: xxxxxxxxx / КПП: xxxxxxxx`, "", 2, "L", false, 0, "")
	pdf.CellFormat(90, 5, `430000 г. Саранск, ул. Пролетарская, д. 1`, "", 0, "L", false, 0, "")

	pdf.CellFormat(90, 15, "", "", 1, "R", false, 0, "") //logo
	pdf.Ln(5)

	pdf.SetFillColor(200, 244, 255)
	pdf.CellFormat(90, 10, `Данные покупателя`, "", 0, "L", true, 0, "")
	pdf.CellFormat(90, 10, `Данные о заказе`, "", 1, "L", true, 0, "")
	pdf.CellFormat(90, 5, "Покупатель: "+order.CustomerName, "", 0, "L", false, 0, "")

	pdf.CellFormat(90, 5, "Заказ: №"+order.OrderID, "", 2, "L", false, 0, "")
	pdf.CellFormat(90, 5, "Дата заказа: "+order.OrderDate, "", 2, "L", false, 0, "")
	pdf.CellFormat(90, 5, "Статус заказа: "+order.OrderStatus, "", 2, "L", false, 0, "")
	pdf.CellFormat(90, 5, "Способ оплаты: "+order.OrderPaymentNumber, "", 1, "L", false, 0, "")

	pdf.Ln(10)
	pdf.CellFormat(180, 10, `Заказ №`+order.OrderID+" от "+order.OrderDate, "", 1, "C", true, 0, "")

	pdf.CellFormat(20, 5, "№", "1", 0, "L", false, 0, "")
	pdf.CellFormat(70, 5, "Наименование товара", "1", 0, "L", false, 0, "")
	pdf.CellFormat(30, 5, "Количество", "1", 0, "L", false, 0, "")
	pdf.CellFormat(30, 5, "Цена за ед., руб", "1", 0, "L", false, 0, "")
	pdf.CellFormat(30, 5, "Всего, руб.", "1", 1, "L", false, 0, "")

	for i, v := range order.Items {
		pdf.CellFormat(20, 5, strconv.Itoa(i+1), "1", 0, "L", false, 0, "")
		pdf.CellFormat(70, 5, v.Name, "1", 0, "L", false, 0, "")
		pdf.CellFormat(30, 5, strconv.Itoa(v.Quantity), "1", 0, "L", false, 0, "")
		pdf.CellFormat(30, 5, fmt.Sprintf("%.2f", v.Price), "1", 0, "L", false, 0, "")
		pdf.CellFormat(30, 5, fmt.Sprintf("%.2f", float64(v.Quantity)*v.Price), "1", 1, "L", false, 0, "")
	}
	pdf.CellFormat(150, 5, "Итого", "1", 0, "L", false, 0, "")
	pdf.CellFormat(30, 5, fmt.Sprintf("%.2f", order.TotalPrice), "1", 1, "L", false, 0, "")

	return pdf

}
