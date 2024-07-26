Program ini dibuat dengan bahasa pemrograman Golang

1. Buat database dengan nama `esbtest` di MySQL
2. Install makefile untuk menjalankan program ini dengan cara `brew install make` atau `sudo apt-get install make` atau `sudo yum install make` atau di Windows bisa download di https://sourceforge.net/projects/gnuwin32/files/make/3.81/make-3.81.exe/download
3. Install XAMPP atau program sejenisnya untuk menjalankan MySQL
4. Jalankan program ini dengan cara `make migrate.up` untuk membuat tabel di database dan `make run` untuk menjalankan program ini
5. Akses API melalui `http://localhost:3000/` untuk melihat hasil dari program ini