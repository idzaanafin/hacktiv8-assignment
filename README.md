# Hacktiv8 Assignment - Product Management API

REST API untuk manajemen produk, sumber (sources), dan transaksi menggunakan Go Gin framework.

## Instalasi & Menjalankan

1. Clone repository ini
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Jalankan aplikasi:
   ```bash
   go run main.go
   ```
4. Server akan berjalan di `http://localhost:8080`

## Response Format

Semua endpoint menggunakan format response yang konsisten:

```json
{
  "message": "string",
  "data": "object/array/null",
  "error": "string/null"
}
```

## Endpoints API

### Products

#### 1. Get All Products
- **URL**: `GET /products`
- **Deskripsi**: Mendapatkan semua produk
- **cURL**:
```bash
curl -X GET http://localhost:8080/products
```
- **Response**:
```json
{
  "message": "List of products",
  "data": [
    {
      "id": "uuid-string",
      "name": "Nama Produk",
      "description": "Deskripsi produk",
      "price": 100000,
      "stock": 50,
      "source_id": "uuid-string"
    }
  ],
  "error": null
}
```

#### 2. Get Product by ID
- **URL**: `GET /products/:id`
- **Deskripsi**: Mendapatkan produk berdasarkan ID
- **Parameter**: `id` (string) - ID produk
- **cURL**:
```bash
curl -X GET http://localhost:8080/products/your-product-id
```
- **Response Sukses**:
```json
{
  "message": "Product found",
  "data": {
    "id": "uuid-string",
    "name": "Nama Produk",
    "description": "Deskripsi produk",
    "price": 100000,
    "stock": 50,
    "source_id": "uuid-string"
  },
  "error": null
}
```
- **Response Error (404)**:
```json
{
  "message": "Product not found",
  "data": null,
  "error": "invalid_id"
}
```

#### 3. Create Product
- **URL**: `POST /products`
- **Deskripsi**: Membuat produk baru
- **Request Body**:
```json
{
  "name": "Nama Produk",
  "description": "Deskripsi produk",
  "price": 100000,
  "stock": 50,
  "source_id": "uuid-string"
}
```
- **cURL**:
```bash
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Laptop Gaming",
    "description": "Laptop untuk gaming dengan spesifikasi tinggi",
    "price": 15000000,
    "stock": 10,
    "source_id": "your-source-id"
  }'
```
- **Response Sukses (201)**:
```json
{
  "message": "Product created",
  "data": {
    "id": "generated-uuid",
    "name": "Nama Produk",
    "description": "Deskripsi produk",
    "price": 100000,
    "stock": 50,
    "source_id": "uuid-string"
  },
  "error": null
}
```
- **Response Error (400)**:
```json
{
  "message": "Invalid price or stock",
  "data": null,
  "error": "Validation failed"
}
```

#### 4. Update Product
- **URL**: `PUT /products/:id`
- **Deskripsi**: Update produk berdasarkan ID
- **Parameter**: `id` (string) - ID produk
- **Request Body**:
```json
{
  "name": "Nama Produk Updated",
  "description": "Deskripsi produk updated",
  "price": 150000,
  "stock": 30,
  "source_id": "uuid-string"
}
```
- **cURL**:
```bash
curl -X PUT http://localhost:8080/products/your-product-id \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Laptop Gaming Updated",
    "description": "Laptop gaming dengan spesifikasi yang diupdate",
    "price": 18000000,
    "stock": 8,
    "source_id": "your-source-id"
  }'
```
- **Response Sukses**:
```json
{
  "message": "Product updated",
  "data": {
    "id": "uuid-string",
    "name": "Nama Produk Updated",
    "description": "Deskripsi produk updated",
    "price": 150000,
    "stock": 30,
    "source_id": "uuid-string"
  },
  "error": null
}
```

#### 5. Delete Product
- **URL**: `DELETE /products/:id`
- **Deskripsi**: Hapus produk berdasarkan ID
- **Parameter**: `id` (string) - ID produk
- **cURL**:
```bash
curl -X DELETE http://localhost:8080/products/your-product-id
```
- **Response Sukses**:
```json
{
  "message": "Product deleted",
  "data": null,
  "error": null
}
```

### Sources

#### 1. Get All Sources
- **URL**: `GET /sources`
- **Deskripsi**: Mendapatkan semua sumber
- **cURL**:
```bash
curl -X GET http://localhost:8080/sources
```
- **Response**:
```json
{
  "message": "List of sources",
  "data": [
    {
      "id": "uuid-string",
      "name": "Nama Sumber"
    }
  ],
  "error": null
}
```

#### 2. Get Source by ID
- **URL**: `GET /sources/:id`
- **Deskripsi**: Mendapatkan sumber berdasarkan ID
- **Parameter**: `id` (string) - ID sumber
- **cURL**:
```bash
curl -X GET http://localhost:8080/sources/your-source-id
```
- **Response Sukses**:
```json
{
  "message": "Source found",
  "data": {
    "id": "uuid-string",
    "name": "Nama Sumber"
  },
  "error": null
}
```

#### 3. Create Source
- **URL**: `POST /sources`
- **Deskripsi**: Membuat sumber baru
- **Request Body**:
```json
{
  "name": "Nama Sumber Baru"
}
```
- **cURL**:
```bash
curl -X POST http://localhost:8080/sources \
  -H "Content-Type: application/json" \
  -d '{"name": "Supplier A"}'
```
- **Response Sukses (201)**:
```json
{
  "message": "Source created",
  "data": {
    "id": "generated-uuid",
    "name": "Nama Sumber Baru"
  },
  "error": null
}
```

#### 4. Update Source
- **URL**: `PUT /sources/:id`
- **Deskripsi**: Update sumber berdasarkan ID
- **Parameter**: `id` (string) - ID sumber
- **Request Body**:
```json
{
  "name": "Nama Sumber Updated"
}
```
- **cURL**:
```bash
curl -X PUT http://localhost:8080/sources/your-source-id \
  -H "Content-Type: application/json" \
  -d '{"name": "Supplier A Updated"}'
```
- **Response Sukses**:
```json
{
  "message": "Source updated",
  "data": {
    "id": "uuid-string",
    "name": "Nama Sumber Updated"
  },
  "error": null
}
```

#### 5. Delete Source
- **URL**: `DELETE /sources/:id`
- **Deskripsi**: Hapus sumber berdasarkan ID
- **Parameter**: `id` (string) - ID sumber
- **cURL**:
```bash
curl -X DELETE http://localhost:8080/sources/your-source-id
```
- **Response Sukses**:
```json
{
  "message": "Source deleted",
  "data": null,
  "error": null
}
```

### Transactions

#### 1. Get All Transactions
- **URL**: `GET /transactions`
- **Deskripsi**: Mendapatkan semua transaksi
- **cURL**:
```bash
curl -X GET http://localhost:8080/transactions
```
- **Response**:
```json
{
  "message": "List of transactions",
  "data": [
    {
      "id": "uuid-string",
      "product_id": "uuid-string",
      "quantity": 2,
      "total": 200000
    }
  ],
  "error": null
}
```

#### 2. Get Transaction by ID
- **URL**: `GET /transactions/:id`
- **Deskripsi**: Mendapatkan transaksi berdasarkan ID
- **Parameter**: `id` (string) - ID transaksi
- **cURL**:
```bash
curl -X GET http://localhost:8080/transactions/your-transaction-id
```
- **Response Sukses**:
```json
{
  "message": "Transaction found",
  "data": {
    "id": "uuid-string",
    "product_id": "uuid-string",
    "quantity": 2,
    "total": 200000
  },
  "error": null
}
```

#### 3. Create Transaction
- **URL**: `POST /transactions`
- **Deskripsi**: Membuat transaksi baru (mengurangi stok produk)
- **Request Body**:
```json
{
  "product_id": "uuid-string",
  "quantity": 2
}
```
- **cURL**:
```bash
curl -X POST http://localhost:8080/transactions \
  -H "Content-Type: application/json" \
  -d '{
    "product_id": "your-product-id",
    "quantity": 1
  }'
```
- **Response Sukses (201)**:
```json
{
  "message": "Transaction success",
  "data": {
    "id": "generated-uuid",
    "product_id": "uuid-string",
    "quantity": 2,
    "total": 200000
  },
  "error": null
}
```
- **Response Error - Stok Tidak Cukup (400)**:
```json
{
  "message": "Invalid quantity",
  "data": null,
  "error": "Stock not enough or invalid quantity"
}
```
- **Response Error - Produk Tidak Ditemukan (404)**:
```json
{
  "message": "Product not found",
  "data": null,
  "error": "invalid_product_id"
}
```

## Contoh Workflow Lengkap dengan cURL

### 1. Membuat Source terlebih dahulu
```bash
curl -X POST http://localhost:8080/sources \
  -H "Content-Type: application/json" \
  -d '{"name": "Supplier A"}'
```

### 2. Membuat Product dengan source_id dari step 1
```bash
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Laptop Gaming",
    "description": "Laptop untuk gaming dengan spesifikasi tinggi",
    "price": 15000000,
    "stock": 10,
    "source_id": "source-id-dari-step-1"
  }'
```

### 3. Membuat Transaksi dengan product_id dari step 2
```bash
curl -X POST http://localhost:8080/transactions \
  -H "Content-Type: application/json" \
  -d '{
    "product_id": "product-id-dari-step-2",
    "quantity": 1
  }'
```

### 4. Melihat semua data
```bash
# Lihat semua products
curl -X GET http://localhost:8080/products

# Lihat semua sources  
curl -X GET http://localhost:8080/sources

# Lihat semua transactions
curl -X GET http://localhost:8080/transactions
```

## Validasi

### Product
- `price` harus lebih dari 0
- `stock` harus 0 atau lebih

### Transaction
- `quantity` harus lebih dari 0
- `quantity` tidak boleh melebihi stok yang tersedia
- `product_id` harus valid (produk harus ada)

## Error Codes

- **200**: OK - Request berhasil
- **201**: Created - Resource berhasil dibuat
- **400**: Bad Request - Input tidak valid atau validasi gagal
- **404**: Not Found - Resource tidak ditemukan

## Data Models

### Product
```go
type Product struct {
    ID          string  `json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
    Stock       int     `json:"stock"`
    SourceID    string  `json:"source_id"`
}
```

### Source
```go
type Source struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}
```

### Transaction
```go
type Transaction struct {
    ID        string  `json:"id"`
    ProductID string  `json:"product_id"`
    Quantity  int     `json:"quantity"`
    Total     float64 `json:"total"`
}
```

## Catatan

- ID untuk semua resource di-generate otomatis menggunakan UUID
- Total transaksi dihitung otomatis: `quantity * price`
- Stok produk akan berkurang otomatis saat transaksi dibuat
- Data disimpan dalam memory (akan hilang saat aplikasi di-restart)
