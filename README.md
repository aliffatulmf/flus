# Flus (⚠️eksperimental)

[English](README.en.md)

**Flus** adalah sebuah program eksperimental yang berfungsi untuk mengatur atau memindahkan file berdasarkan ekstensi ke direktori yang sesuai.

## Insalasi

Untuk menggunakan program ini, Anda dapat mengikuti langkah-langkah berikut:

1. Download repository ini.
2. Navigasi terminal.
3. Jalankan perintah `build.bat` atau `build.sh`.

## Penjelasan

### target

Argumen `-target` adalah opsi dalam program Flus yang digunakan untuk menentukan direktori target yang ingin Anda atur. Ketika Anda menjalankan program dengan opsi ini, Flus akan memindai berkas-berkas dalam direktori target dan memindahkannya ke direktori yang sesuai berdasarkan jenis file mereka.

Sebagai contoh, jika Anda memiliki sebuah direktori target dengan berkas-berkas berbagai jenis seperti `.jpg`, `.pdf`, dan `.zip`, Flus akan memindahkan berkas `.jpg` ke direktori bernama `Images`, berkas `.pdf` ke direktori bernama `Documents`, dan berkas `.zip` ke direktori bernama `Archives`.

Anda dapat menentukan direktori target dengan menggunakan opsi `-target` diikuti dengan jalur direktori. Misalnya, jika direktori target Anda berada di `/home/user/documents`, Anda dapat menjalankan program dengan perintah berikut:

```powershell
flus -target /home/user/documents
```

### unsafe

Argumen `-unsafe` adalah opsi yang memungkinkan pengguna untuk mengaktifkan mode `unsafe` atau melewati langkah verifikasi setelah proses penyalinan file berhasil dilakukan. Penggunaan opsi ini bertujuan untuk meningkatkan kecepatan proses penyalinan, namun perlu diingat bahwa metode ini juga membawa potensi risiko kerusakan data karena ketiadaan verifikasi setelah penyalinan selesai.

Dalam program ini, metode verifikasi yang biasa digunakan adalah dengan membandingkan hasil _Hash_ dari file asli dan file hasil salinan. Namun, saat mode `unsafe` diaktifkan, proses verifikasi ini dilewati untuk mempercepat eksekusi.

Jika Anda ingin mengaktifkan mode `unsafe`, Anda dapat menjalankan program dengan perintah berikut:

```powershell
flus -target /home/user/documents -unsafe
```

Namun, disarankan untuk berhati-hati saat menggunakan opsi ini. Pastikan Anda hanya mengaktifkan mode `unsafe` ketika Anda yakin bahwa sumber file dan proses penyalinannya dapat dipercaya sepenuhnya, dan risiko kehilangan atau kerusakan data akibat kelalaian dalam verifikasi dapat diterima.

### move

Secara default program ini akan menyalin file ke direktori yang sesuai berdasarkan jenis file. Namun, jika Anda ingin menghapus file asli setelah proses penyalinan selesai, Anda dapat menggunakan opsi `-move`.

⚠️ PERHATIAN!  Penggunaan opsi `-move` dengan `-unsafe` akan menghapus file asli tanpa melakukan verifikasi. Pastikan Anda telah memahami risiko yang mungkin terjadi sebelum menggunakan opsi ini.

Jika Anda ingin mengaktifkan opsi `-move`, Anda dapat menjalankan program dengan perintah berikut:

```powershell
flus -target /home/user/documents -move
```

Disarankan untuk tidak menggunakan opsi `-move` jika Anda tidak yakin dengan proses penyalinan yang akan dilakukan. Hapus file yang telah dicopy secara manual dan pastikan sebelum menghapus, data yang disalin telah terverifikasi dan tidak rusak.
