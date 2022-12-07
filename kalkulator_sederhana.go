/*
Created By Anwar Sanusi
*/

package main

import (
   "fmt";
)

var (
   menu int
   celcius,farenheit,reamur int
   kelvin,cel,far float64
)

func main() {
   fmt.Println("=================WELCOME================")
   fmt.Println("=========Menu Utama Konversi Suhu=======")
   Suhu()
}

//========================SUHU==================//
func Suhu(){
   fmt.Println("Konversi Suhu")
   fmt.Println("1. Celcius")
   fmt.Println("2. Farenheit")
   fmt.Println("3. Reamur")
   fmt.Println("4. Kelvin")
   fmt.Println("5. Menu Utama")
   fmt.Print("Pilih salah satu suhu yang ingin dikonversi pada menu diatas (ketik nomor menu): ")
   fmt.Scanf("%d", &menu)

switch menu{
   case 1 :
      fmt.Println("1. Celcius - Farenheit")
      fmt.Println("2. Celcius - Reamur")
      fmt.Println("3. Celcius - Kelvin")
      fmt.Println("4. Menu Utama")
      fmt.Print("Pilih salah satu menu diatas (ketik nomor menu): ")
      fmt.Scanf("%d", &celcius)

      switch celcius{
         case 1 :
            CF()
         case 2 :
            CR()
         case 3 :
            CK()
         default :
            Suhu()
      }
   case 2 :
      fmt.Println("1. Farenheit - Celcius")
      fmt.Println("2. Farenheit - Reamur")
      fmt.Println("3. Menu Utama")
      fmt.Print("Pilih salah satu menu diatas (ketik nomor menu): ")
      fmt.Scanf("%d", &farenheit)

      switch farenheit{
         case 1 :
            FC()
         case 2 :
            FR()
         default :
            Suhu()
      }
   case 3 :
      fmt.Println("1. Reamur - Celcius")
      fmt.Println("2. Reamur - Farenheit")
      fmt.Println("3. Menu Utama")
      fmt.Print("Pilih salah satu menu diatas (ketik nomor menu): ")
      fmt.Scanf("%d", &reamur)

      switch reamur{
         case 1 :
            RC()
         case 2 :
            RF()
         default :
            Suhu()
         }
   case 4 :
      fmt.Println("konversi Kelvin - Celcius")
      fmt.Print("Masukkan nilai suhu Kelvin = ")
      fmt.Scanf("%f", &kelvin)
      KC := kelvin-273
      fmt.Print("Nilai Suhu dalam Celcius = ", KC)
      Suhu()
   default :
      main()
}
}

//-----------------------------------------------------------------------------------------------//
//CELCIUS

func CF(){
   var cek int
   fmt.Println("konversi Celcius - Farenheit")
   fmt.Print("Masukkan nilai suhu Celcius = ")
   fmt.Scanf("%f", &cel)
   hasil := ((9*cel)+32)/5
   fmt.Println("Nilai Suhu dalam Fareinheit =  ", hasil)

   fmt.Println("Apa anda ingin keluar ? (jika ya tekan 1, jika tidak tekan 0) ")
   fmt.Scanf("%d", &cek)

   switch cek{
      case 1 :
         exit()
      default :
         Suhu()
   }
}

func CR(){
   var cek int
   fmt.Println("konversi Celcius - Reamur")
   fmt.Print("Masukkan nilai suhu Celcius = ")
   fmt.Scanf("%f", &cel)
   hasil := cel*4/5
   fmt.Println("Nilai Suhu dalam Reamur = ", hasil)

   fmt.Println("Apa anda ingin keluar ? (jika ya tekan 1, jika tidak tekan 0) ")
   fmt.Scanf("%d", &cek)

   switch cek{
      case 1 :
      exit()
   default :
      Suhu()
   }
}

func CK(){
   var cek int
   fmt.Println("konversi Celcius - Kelvin")
   fmt.Print("Masukkan nilai suhu Celcius = ")
   fmt.Scanf("%f", &cel)
   hasil := cel+273
   fmt.Println("Nilai Suhu dalam Kelvin = ", hasil)

   fmt.Println("Apa anda ingin keluar ? (jika ya tekan 1, jika tidak tekan 0) ")
   fmt.Scanf("%d", &cek)

   switch cek{
      case 1 :
         exit()
      default :
         Suhu()
   }
}

//--------------------------------------------------------------------------------------------//
// FARENHEIT

func FC(){
   var cek int
   fmt.Println("konversi Farenheit - Celcius")
   fmt.Print("Masukkan nilai suhu Farenheit = ")
   fmt.Scanf("%f", &far)
   hasil := 5/9*(far-32)
   fmt.Println("Nilai Suhu dalam Celcius = ", hasil)

   fmt.Println("Apa anda ingin keluar ? (jika ya tekan 1, jika tidak tekan 0) ")
   fmt.Scanf("%d", &cek)

   switch cek{
      case 1 :
        exit()
      default :
         Suhu()
   }
}

func FR(){
   var cek int
   fmt.Println("konversi Farenheit - Reamur")
   fmt.Print("Masukkan nilai suhu Farenheit = ")
   fmt.Scanf("%f", &far)
   hasil := 4*(far-32)/9
   fmt.Println("Nilai Suhu dalam Reamur = ", hasil)

   fmt.Println("Apa anda ingin keluar ? (jika ya tekan 1, jika tidak tekan 0) ")
   fmt.Scanf("%d", &cek)

   switch cek{
      case 1 :
         exit()
      default :
         Suhu()
   }
}

//-------------------------------------------------------------------------------------------//
// REAMUR

func RC(){
   var r1 float64
   var cek int

   fmt.Println("konversi Reamur - Celcius")
   fmt.Print("Masukkan nilai suhu Reamur = ")
   fmt.Scanf("%f", &r1)
   hasil := (r1*5)/4
   fmt.Println("Nilai Suhu dalam Celcius = ", hasil)

   fmt.Println("Apa anda ingin keluar ? (jika ya tekan 1, jika tidak tekan 0) ")
   fmt.Scanf("%d", &cek)

   switch cek{
      case 1 :
         exit()
      default :
         Suhu()
   }
}

func RF(){
   var r2 float64
   var cek int

   fmt.Println("konversi Reamur - Farenheit")
   fmt.Print("Masukkan nilai suhu Reamur = ")
   fmt.Scanf("%f", &r2)
   hasil := (r2*9)/4+32
   fmt.Println("Nilai Suhu dalam Farenheit = ", hasil)

   fmt.Println("Apa anda ingin keluar ? (jika ya tekan 1, jika tidak tekan 0) ")
   fmt.Scanf("%d", &cek)

   switch cek{
      case 1 :
         exit()
      default :
         Suhu()
   }
}

//=========================EXIT=====================//
func exit(){
   fmt.Println("====Selesai===")
}