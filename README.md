# Simple Crud Employee

## Todo
Create Simple CRUD Employee 

### Data:
1. EmployeeID, string
2. FullName, string
3. Address, string

### Output:
1. Output Daftar Karyawan
    | Employee ID | FullName | Address   |
    |-------------|----------|-----------|
    | 1001        | Budi     | Jakarta   |
    | 1002        | Adi      | Jakarta   |
    | 1003        | Muhammad | Tangerang |

### Specification
1. Buatlah program dengan menggunakan GoLang 
    1. Menambahkan data 
    2. Menampilkan data 
    3. Menghapus data 
    4. Mengedit Data
2. Program harus memiliki
    1. Controller, View
    2. Comment untuk descripsi singkat aplikasi
    3. Menggunakan Database Mysql
    4. Validation untuk valid data (null handling, date handling, duplicate data handling)
    5. Error handling ( exception )
3. Buatlah guide/dokumentasi untuk menjalankan program tersebut

## Layout
Project Layout ini menggunakan `Standar Project Layout` dibangun oleh [komunitas](https://github.com/golang-standards/project-layout)

Berikut detail dari project layout:
```
simple-crud-employee/
|-- cmd/
|   |-- app/            # Main entry point 
|-- internal/
|   |-- entity/         # Business entities
|   |-- infrastructure/ # Frameworks and drivers
|   |  |-- db/          # Database implementation
|   |  |-- server/      # HTTP server implementation
|   |-- interface/      # Adapters for the external world
|   |  |-- http/        # HTTP handlers (controllers)
|   |  |-- repository/  # Database repositories
|   |-- usecase/        # Application business rules
|-- web
    |-- template        # Server side template (view)
```

## How to run
1. Set environment sesuai [configs](configs/), atau buat sebuah `.env` file. Contoh:
    ```bash
    HTTP_PORT="8080"
    MYSQL_HOST="localhost"
    MYSQL_PORT="3306"
    MYSQL_DB_NAME="employee_db"
    MYSQL_USERNAME="root"
    MYSQL_PASSWORD="password"
    MYSQL_LOG_MODE="1"
    MYSQL_PARSE_TIME="true"
    MYSQL_CHARSET="utf8mb4"
    MYSQL_LOC="Local"
    MYSQL_MAX_LIFETIME_CONNECTION="10"
    MYSQL_MAX_OPEN_CONNECTION="50"
    MYSQL_MAX_IDLE_CONNECTION="10"
    ```
2. Di project root folder, jalankan di terminal:
    ```bash
    go run cmd/app/main.go 
    ```