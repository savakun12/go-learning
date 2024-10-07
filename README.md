# Proyek Go Sederhana

Proyek ini adalah contoh aplikasi sederhana yang ditulis dalam bahasa pemrograman Go. Aplikasi ini mengelola buku dengan fitur untuk membuat buku baru dan menampilkan ringkasan.

## Struktur Proyek

```
my_go_app/
│
├── main.go
└── README.md
```

## Deskripsi

Aplikasi ini mencakup contoh kode untuk:

1. **Struct dan Method**
2. **Validasi Input**
3. **Error Handling**
4. **Slice dan Iterasi**

## Cara Menjalankan

### Persyaratan

Pastikan kamu sudah menginstal [Go](https://golang.org/dl/).

### Langkah-langkah

1. **Clone atau Unduh Proyek**
   ```bash
   git clone <repository-url>
   cd my_go_app
   ```

2. **Jalankan Kode**
   - Untuk menjalankan kode Go, buka terminal dan gunakan perintah:
     ```bash
     go run main.go
     ```

## Contoh Kode

### Struct dan Method

```go
type Book struct {
    Title  string
    Author string
    Pages  int
}

func (b Book) Summary() string {
    return fmt.Sprintf("%s by %s, %d pages", b.Title, b.Author, b.Pages)
}
```

### Validasi Input

```go
func NewBook(title, author string, pages int) (*Book, error) {
    if title == "" {
        return nil, errors.New("title cannot be empty")
    }
    if author == "" {
        return nil, errors.New("author cannot be empty")
    }
    if pages <= 0 {
        return nil, errors.New("pages must be a positive number")
    }
    return &Book{Title: title, Author: author, Pages: pages}, nil
}
```

### Slice dan Iterasi

```go
books := []Book{*book1, {"Go in Action", "William Kennedy", 265}}

for _, book := range books {
    fmt.Println(book.Summary())
}
```

## Belajar Dasar Go

- Untuk mempelajari dasar-dasar Go, kamu bisa mengunjungi [Tour of Go](https://tour.golang.org/) dan [Dokumentasi Resmi Go](https://golang.org/doc/).
- Pelajari tentang:
  - Tipe Data dan Variabel
  - Kontrol Alur (if, switch, loop)
  - Fungsi dan Method
  - Struct dan Interface
  - Error Handling

## Lisensi

Proyek ini dirilis di bawah lisensi MIT. Silakan lihat file LICENSE untuk informasi lebih lanjut.
