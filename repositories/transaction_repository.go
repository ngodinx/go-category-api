package repositories

import (
	"database/sql"
	"fmt"
	"time"
	"kasir-api/models"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}
	
func (repo *TransactionRepository) CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	totalAmount := 0
	details := make([]models.TransactionDetail, 0)

	for _, item := range items {
		var productPrice, stock int
		var productName string

		err := tx.QueryRow("SELECT name, price, stock FROM products WHERE id = $1", item.ProductID).Scan(&productName, &productPrice, &stock)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product id %d not found", item.ProductID)
		}
		if err != nil {
			return nil, err
		}
  
		subtotal := productPrice * item.Quantity
		totalAmount += subtotal

		_, err = tx.Exec("UPDATE products SET stock = stock - $1 WHERE id = $2", item.Quantity, item.ProductID)
		if err != nil {
			return nil, err
		}

		details = append(details, models.TransactionDetail{
			ProductID:   item.ProductID,
			ProductName: productName,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	var transactionID int
	err = tx.QueryRow("INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id", totalAmount).Scan(&transactionID)
	if err != nil {
		return nil, err
	}

	for i := range details {
		details[i].TransactionID = transactionID
		_, err = tx.Exec("INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES ($1, $2, $3, $4)",
			transactionID, details[i].ProductID, details[i].Quantity, details[i].Subtotal)
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &models.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		Details:     details,
	}, nil
}

func (repo *TransactionRepository) GetReport(start, end time.Time) (*models.ReportResponse, error) {
	var totalRevenue int
	var totalTransaksi int

	err := repo.db.QueryRow(`
		SELECT COALESCE(SUM(total_amount), 0) AS total_revenue,
		       COUNT(*) AS total_transaksi
		FROM transactions
		WHERE created_at >= $1 AND created_at < $2
	`, start, end).Scan(&totalRevenue, &totalTransaksi)
	if err != nil {
		return nil, err
	}

	report := &models.ReportResponse{
		TotalRevenue:   totalRevenue,
		TotalTransaksi: totalTransaksi,
		ProdukTerlaris: models.ProdukTerlaris{Nama: "", QtyTerjual: 0},
	}

	// produk terlaris berdasarkan total qty terjual
	var nama string
	var qty int
	err = repo.db.QueryRow(`
		SELECT p.name, SUM(td.quantity) AS qty_terjual
		FROM transaction_details td
		JOIN transactions t ON t.id = td.transaction_id
		JOIN products p ON p.id = td.product_id
		WHERE t.created_at >= $1 AND t.created_at < $2
		GROUP BY p.id, p.name
		ORDER BY qty_terjual DESC
		LIMIT 1
	`, start, end).Scan(&nama, &qty)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == nil {
		report.ProdukTerlaris = models.ProdukTerlaris{Nama: nama, QtyTerjual: qty}
	}

	return report, nil
}