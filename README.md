# JatisMarch23

Test from one company that i interest to join

## Description

Question Start with single choice answer, coding test, algorithm, and Query
I put all answer in this repo.

## Getting Started

Specific to coding task, you can run this repo with ```go run main.go``` in your go environment. You re free to run which "Soal" to be executed by editing in main.go on main func.

Notes:
I worked with go v1.18

## ANSWERS (i will switch to bahasa)

### Pilihan Ganda & Uraian Singkat
1. b. Karena constructor tidak mengembalikan nilai apapun, Anda menempatkan keyword “void” sebelum nama constructor.
2. c. Akan meng-override fungsi base class.
3. b. Meng-assign nilai 36.96 pada values[2]
4. a. Meng-assign nilai 3 pada p
5. c. Kondisi test selalu bernilai salah
6. Polymorphism adalah perbedaan fungi method dalam class walau memiliki Nama yang sama. Polimorphism ada 2 jenis, statis(overloading) dan dynamis(overriding). Overloading adalah dalam satu class terdapat beberapa method dengan nama sama namun parameter dan return value nya berbeda. Overriding adalah penggunaan kembali method parent class Dari child class yang dimodifikasi, sehingga minimpa penggunaan method base class its sender.

Encapsulation dalam OOP adalah konsep membungkus informasi/atribut dalam objek, mengatur tiap atributnya sehingga memiliki hak akses yang berbeda sesuai kebutuhan penggunaanya.

Inheritance adalah konsep pewarisan method/atribut dari parent class keypads derived class sehingga derived class dapat menggunakan fungi yang sama tau bahkan memiliki hak menciptakan fungsi sendiri atau memodifikasi fungsai turunan.

### Algoritma
Bisa di execute sesuai penjelasan di getting started

Soal nomer 6 bisa di lihat di gambar Soal6Jatis.jpg yang ada di repo

### DATABASE
1. Join Table
```
    SELECT X.A, X.B, X.C, X.E FROM X
    JOIN Y ON X.C = Y.D
```

2. Query
```
    SELECT
	EXTRACT(year FROM tanggal) AS Tahun,
        count(case when choice='A' then 0 else NULL end) as 'A',
        count(case when choice='B' then 0 else NULL end) as 'B',
        count(case when choice='C' then 0 else NULL end) as 'C'
    FROM Choices
    Group By EXTRACT(year FROM tanggal)
```
