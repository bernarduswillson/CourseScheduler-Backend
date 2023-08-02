# Seleksi 3 Labpro - Single Service

## 13521024 - Ahmad Nadil

## How to Run
### Run using Go
#### 1. Make sure you have Go Lang installed
#### 2. Copy the .env.example file and rename it to .env and change the following environment variables to your own
```
DB_USERNAME=
DB_PASSWORD=
DB_HOST=
DB_TABLE=
DB_PORT=
```
#### 3. Run the project
```
go run main.go
```
#### Alternatively, you can use compile daemon to automatically compile the project when there is a change
```
CompileDaemon -command="./singleservice"
```

#### 4. The project will be run on port 8080
```
localhost:8080
```

### Run using Docker
#### 1. Make sure you have docker installed
#### 2. Run the docker-compose
```
docker-compose up
```
#### 3. The project will be served on port 8080
```
localhost:8080
```

### Admin Account
```
Username : admin
Password : admin
```

## Design Pattern yang Digunakan
### 1. Singleton Pattern
Singleton pattern digunakan untuk membuat koneksi ke database. Koneksi ke database hanya perlu dibuat sekali saja dan disimpan dalam variabel static sehingga dapat digunakan oleh class lain tanpa perlu membuat koneksi baru.

### 2. Model-View-Controller (MVC) Pattern
MVC pattern digunakan untuk memisahkan antara logic, view, dan model. Logic diletakkan di dalam folder controller, digunakan untuk operasi CRUD. Untuk view karena aplikasi ini adalah aplikasi backend, hanya berupa routing yang menampilkan data dari controller, hal ini terdapat pada main.go. Sedangkan model digunakan untuk membuat model dari data yang ada di database, hal ini terdapat pada folder model. Ketiga bagian ini dipisahkan agar aplikasi dapat dikembangkan dengan mudah.

### 3. Repository Pattern
Repository pattern digunakan untuk memisahkan antara logic dan data. Logic diletakkan di dalam folder controller, digunakan untuk operasi CRUD. Sedangkan data diletakkan di dalam folder database, digunakan untuk membuat model dari data yang ada di database.

## Technology Stack yang Digunakan
- Go version 1.20.3
- PostgreSQL version 15.2
- Docker version 20.10.24

## Endpoint yang dibuat
- GET /self : Get user data
- GET /barang : Get all barang
- GET /barang/{id} : Get barang by id
- GET /perusahaan : Get all perusahaan
- GET /perusahaan/{id} : Get perusahaan by id
- POST /barang : Create barang
- POST /login : Login user
- POST /perusahaan : Create perusahaan
- UPDATE /barang/{id} : Update barang by id
- UPDATE /perusahaan/{id} : Update perusahaan by id
- DELETE /barang/{id} : Delete barang by id
- DELETE /perusahaan/{id} : Delete perusahaan by id

## Bonus
### B02 - Deployment
Aplikasi backend dan database ini sudah di-deploy menggunakan railway

https://single-service-production.up.railway.app/

### B04 - Polling
Penjelasan ada di repostitory [monolith](https://github.com/IceTeaXXD/Seleksi-3-Labpro-Monolith-Ahmad-Nadil#b04---polling)

### B05 - Lighthouse
Penjelasan ada di repostitory [monolith](https://github.com/IceTeaXXD/Seleksi-3-Labpro-Monolith-Ahmad-Nadil#b05---lighthouse)

### B06 - Responsive Layout
Penjelasan ada di repostitory [monolith](https://github.com/IceTeaXXD/Seleksi-3-Labpro-Monolith-Ahmad-Nadil#b06---responsive-layout)

### B11 - Search Feature
Penjelasan ada di repostitory [monolith](https://github.com/IceTeaXXD/Seleksi-3-Labpro-Monolith-Ahmad-Nadil#b11---search-feature)
