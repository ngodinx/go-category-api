package main

// @title Kasir API
// @version 1.0
// @description Dokumentasi API Kasir
// @BasePath /

// GetProduk godoc
// @Summary List produk
// @Tags Produk
// @Produce json
// @Param name query string false "Filter nama produk (partial match)" example(Indomie)
// @Success 200 {array} models.Product
// @Router /api/produk [get]
func GetProdukSwaggerDoc() {}

// CreateProduk godoc
// @Summary Buat produk
// @Tags Produk
// @Accept json
// @Produce json
// @Param payload body models.Product true "payload"
// @Success 201 {object} models.Product
// @Router /api/produk [post]
func CreateProdukSwaggerDoc() {}

// GetProdukByID godoc
// @Summary Detail produk by ID
// @Tags Produk
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Product
// @Router /api/produk/{id} [get]
func GetProdukByIDSwaggerDoc() {}

// UpdateProduk godoc
// @Summary Update produk by ID
// @Tags Produk
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param payload body models.Product true "payload"
// @Success 200 {object} models.Product
// @Router /api/produk/{id} [put]
func UpdateProdukSwaggerDoc() {}

// DeleteProduk godoc
// @Summary Hapus produk by ID
// @Tags Produk
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]string
// @Router /api/produk/{id} [delete]
func DeleteProdukSwaggerDoc() {}

// GetKategori godoc
// @Summary List kategori
// @Tags Kategori
// @Produce json
// @Success 200 {array} models.Category
// @Router /api/kategori [get]
func GetKategoriSwaggerDoc() {}

// CreateKategori godoc
// @Summary Buat kategori
// @Tags Kategori
// @Accept json
// @Produce json
// @Param payload body models.Category true "payload"
// @Success 201 {object} models.Category
// @Router /api/kategori [post]
func CreateKategoriSwaggerDoc() {}

// GetKategoriByID godoc
// @Summary Detail kategori by ID
// @Tags Kategori
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} models.Category
// @Router /api/kategori/{id} [get]
func GetKategoriByIDSwaggerDoc() {}

// UpdateKategori godoc
// @Summary Update kategori by ID
// @Tags Kategori
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param payload body models.Category true "payload"
// @Success 200 {object} models.Category
// @Router /api/kategori/{id} [put]
func UpdateKategoriSwaggerDoc() {}

// DeleteKategori godoc
// @Summary Hapus kategori by ID
// @Tags Kategori
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} map[string]string
// @Router /api/kategori/{id} [delete]
func DeleteKategoriSwaggerDoc() {}

// ReportHariIni godoc
// @Summary      Report transaksi hari ini
// @Description  Mengambil total revenue, total transaksi, dan produk terlaris untuk hari ini (Asia/Jakarta)
// @Tags         Report
// @Produce      json
// @Success      200 {object} models.ReportResponse
// @Failure      500 {object} map[string]string
// @Router       /api/report/hari-ini [get]
func ReportHariIni() {}

// ReportRange godoc
// @Summary      Report transaksi berdasarkan rentang tanggal
// @Description  start_date dan end_date format YYYY-MM-DD
// @Tags         Report
// @Produce      json
// @Param        start_date query string false "Tanggal mulai (YYYY-MM-DD)"
// @Param        end_date   query string false "Tanggal akhir (YYYY-MM-DD)"
// @Success      200 {object} models.ReportResponse
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /api/report [get]
func ReportRange() {}
