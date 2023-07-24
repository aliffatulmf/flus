# Flus (⚠️eksperimental)
[English](README.en.md)

**Flus** adalah sebuah program eksperimental yang berfungsi untuk mengatur atau memindahkan file berdasarkan ekstensi ke direktori yang sesuai.

## Insalasi

Untuk menggunakan program ini, Anda dapat mengikuti langkah-langkah berikut:

1. Download repository ini.
2. Navigasi terminal.
3. Jalankan perintah `build.bat` atau `build.sh`.

## Penjelasan -target

Argumen `-target` adalah opsi dalam program Flus yang digunakan untuk menentukan direktori target yang ingin Anda atur. Ketika Anda menjalankan program dengan opsi ini, Flus akan memindai berkas-berkas dalam direktori target dan memindahkannya ke direktori yang sesuai berdasarkan jenis file mereka.

Sebagai contoh, jika Anda memiliki sebuah direktori target dengan berkas-berkas berbagai jenis seperti `.jpg`, `.pdf`, dan `.zip`, Flus akan memindahkan berkas `.jpg` ke direktori bernama `Images`, berkas `.pdf` ke direktori bernama `Documents`, dan berkas `.zip` ke direktori bernama `Archives`.

Anda dapat menentukan direktori target dengan menggunakan opsi `-target` diikuti dengan jalur direktori. Misalnya, jika direktori target Anda berada di `/home/user/documents`, Anda dapat menjalankan program dengan perintah berikut:

```powershell
flus -target /home/user/documents
```

