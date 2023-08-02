
# Flus

[EN-en](README.en.md)

Alat baris perintah **Flus** menawarkan solusi yang efisien untuk mengelola operasi direktori, seperti menyalin dan memindahkan berkas antara lokasi yang berbeda. Alat ini menyediakan berbagai opsi yang dapat disesuaikan untuk memenuhi beragam kebutuhan pada berbagai sistem operasi.

## Penggunaan

Untuk memanfaatkan kemampuan alat **Flus**, kamu perlu memberikan argumen baris perintah spesifik sebagai berikut:

```powershell
PS C:\> flus.exe [opsi]
```

## Opsi yang Tersedia

- `-target`: Tentukan direktori target untuk proses. Berkas akan disalin atau dipindahkan ke direktori ini.

- `-move`: Aktifkan mode pemindahan untuk memindahkan berkas dari direktori sumber ke direktori target, alih-alih menyalin.

- `-buffer`: Tentukan ukuran buffer untuk menyalin berkas. Ukuran buffer sangat mempengaruhi efisiensi penyalinan. Ukuran buffer default diatur ke 64 KB.

## Contoh

Berikut adalah contoh yang menggambarkan penggunaan alat **Flus**:

```powershell
PS C:\> flus.exe -move -buffer 128000 -target C:\path\ke\direktori\target
```

Dalam contoh ini, alat ini akan memindahkan berkas dari direktori sumber ke direktori target yang ditentukan menggunakan ukuran buffer 128 KB.

## Catatan Penting

- Gantikan `/path/ke/direktori/target` dengan jalur aktual dari direktori target.

- Berhati-hatilah saat menggunakan mode _move_, karena berkas akan dihapus dari direktori sumber setelah berhasil dipindahkan ke direktori target.

- Penyesuaian ukuran buffer dapat secara signifikan mempengaruhi kecepatan penyalinan atau pemindahan berkas, serta penggunaan memori. Coba variasikan ukuran buffer untuk menemukan konfigurasi optimal bagi sistem milikmu.

## Menjalankan Alat

Untuk menjalankan alat **Flus** pada sistem operasi milikmu, ikuti langkah berikut:

1. Unduh berkas eksekusi (`flus`) dari repositori resmi atau sumber yang tersedia.

2. Buka terminal atau _command prompt_.

3. Navigasikan ke direktori yang berisi berkas eksekusi `flus` dengan menggunakan perintah `cd`.

4. Jalankan alat dengan opsi yang diinginkan:

   ```powershell
   PS C:\> flus.exe [opsi]
   ```

## Catatan

Meskipun dokumentasi ini memberikan gambaran umum tentang fungsionalitas **Flus**, kompatibilitas dan performanya pada sistem operasi lain selain Windows belum sepenuhnya diuji.
