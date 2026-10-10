# ⚡ GoFuzz Ultra - Multi-Threaded Web Directory Fuzzer

Yuqori tezlikda ishlovchi, **Golang Core** yordamida yozilgan va **Advanced Bash Wrapper** bilan o'ralgan professional veb-katalog va yashirin yo'llarni qidiruvchi (fuzzer) vosita.

Dastur **Matrix Green & Cyberpunk** stilidagi interfeys, dynamic progress-bar, avtomatik Linux ulimit optimizatsiyasi hamda mashhur kiberxavfsizlik lug'atlari (SecLists/Dirb) bilan integratsiyaga ega.

---

## 🔥 Asosiy Imkoniyatlari

*   **Ultra Tezyurar Dvigatel (Golang Core):** Go-lang goroutinalari va kanallari orqali soniyasiga minglab so'rovlarni asinxron qayta ishlaydi.
*   **Intellektual Lug'at Menejeri:** Tizimda `SecLists` yoki `Dirb` lug'atlarini avtomatik qidiradi, agar ular arxivlangan (`.gz`) bo'lsa, ochadi va topilmasa, o'zi o'rnatadi.
*   **Avtomatik Resurs Optimizatsiyasi:** `GOMAXPROCS` va Linux `ulimit` (Open Files limit) parametrlarini maksimal tezlik uchun avtomatik 65,535 ga ko'taradi.
*   **Dinamik Progress Bar:** Real vaqt rejimida foizlar, bajarilgan so'rovlar soni va chiroyli yuklanish animatsiyasini ko'rsatib turadi.
*   **Aqlli Log Tizimi:** Topilgan yo'llarni (`HTTP 200, 301, 302, 403`) filtrlab, ularning to'liq fayl tizimidagi absolyut manzilini (Absolute Path) ko'rsatgan holda saqlaydi.

---

## 🖥️ Terminal Interfeysi (Interface Preview)

Dasturning terminaldagi ish faoliyati va chiroyli interaktiv menyusi quyidagicha ko'rinishga ega:

### 1. Ishga Tushirish va Lug'at Tekshiruvi
Dastur ishga tushganda tizim imkoniyatlarini tekshiradi va parametrlarni kiritishni so'raydi:

```text
 ██████╗  ██████╗ ███████╗██╗   ██╗███████╗███████╗
██╔════╝ ██╔═══██╗██╔════╝██║   ██║╚══███╔╝╚══███╔╝
██║  ███╗██║   ██║█████╗  ██║   ██║  ███╔╝   ███╔╝ 
██║   ██║██║   ██║██╔══╝  ██║   ██║ ███╔╝   ███╔╝  
╚██████╔╝╚██████╔╝██║     ╚██████╔╝███████╗███████╗
 ╚═════╝  ╚═════╝ ╚═╝      ╚═════╝ ╚══════╝╚══════╝
           [ AUTOMATED BASH AUTOMATION ]

┌───[ ENUMERATION ENGINE INFO ]
│ Vazifasi => Multi-threaded Web Directory Fuzzer
│ Arxitektura => Go (Golang Core) & Advanced Bash Wrapper
│ Muallif   => Siz (Kali Hacker)
└───[ PRO EDITION v2.2 ]

[*] Tizim lug'atlar bazasi tekshirilmoqda...
[+] Jangovar lug'at (SecLists) muvaffaqiyatli aniqlandi!

┌───[ HUJUM MAQSADINI SOZLASH ]
│ Nishon URL manzilini kiriting:
│ » https://target-site.com
│ Tarmoqlar soni (Threads) [Standart: 50]:
│ » 100
│ Natijalar fayli nomi [Standart: scan_results.txt]:
│ » target_report.txt
└───[ PARAMETRLAR SAQLANDI ]

[*] GoFuzz yadrosi yuklanmoqda... Barcha tizim limitlari ochildi.
------------------------------------------------------------------------
```

### 2. Reyd va Dinamik Progress-Bar (Cracking Process)
Go yadrosi yuklangach, terminal tozalanadi va har bir yadro holati hamda jonli progress aks etadi:

```text
┌───[ TARMOQ PARAMETRLARI ]
│ Nishon URL  => https://target-site.com
│ Lug'atda    => 220551 ta maxfiy yo'l
│ Tarmoqlar   => 100 concurrent workers
│ CPU Aktiv   => 7/8 Yadrolar
└───[ CRACKING IN PROGRESS... ]

[SUCCESS-200] ───> https://target-site.comadmin
[REDIRECT-301] ───> https://target-site.comimages
[FORBIDDEN-403] ───> https://target-site.com.htaccess

[◢] [████████░░░░░░░░░░░░] 42.15% (92962/220551)
```

### 3. Natijalarni Mukammal Saqlash
Jarayon yakunlangach, natija saqlangan faylning Linux tizimidagi aniq to'liq manzili chiqariladi:

```text
[+] REYD YAKUNLANDI! Barcha maxfiy yo'llar tekshirildi.

┌───[ REYD NATIJALARI SAQLANDI ]
│ Fayl Nomi   => target_report.txt
│ To'liq Yo'l  => /home/user/tools/fuzzer/target_report.txt
└───[ LOG FAYL TAYYOR ]
```

---

## 🛠️ O'rnatish va Talablar

Skript ishlashi uchun kompyuteringizda **Go (Golang)** va **Bash** muhiti bo'lishi kerak.

### 1. Go tilini o'rnatish (agar mavjud bo'lmasa):
```bash
sudo apt update
sudo apt install golang -y
```

### 2. Loyihani sozlash:
Loyihangiz papkasiga o'ting va fayllar joylashuvini tekshiring (`main.go` va `fuzzer.sh` bitta papkada bo'lishi kerak):
```bash
git clone https://github.com
cd LOYIHA-NOMI
chmod +x fuzzer.sh
```

---

## 🚀 Foydalanish

Skriptni ishga tushirish uchun quyidagi buyruqni bering:
```bash
./fuzzer.sh
```
Qolgan barcha jarayonlarni (URL kiritish, Threads soni va natija fayli) dasturning o'zi interaktiv tarzda so'raydi.

---

## ⚠️ OGOHLANTIRISH (DISCLAIMER)

Ushbu vosita faqat **resurslarni test qilish, kiberxavfsizlik auditlari va ta'lim** maqsadida yaratilgan. Ruxsat berilmagan begona serverlarda ushbu vosita orqali fuzzing amalga oshirish noqonuniy hisoblanadi. Foydalanish qoidalari buzilgan taqdirda muallif javobgarlikni o'z zimmasiga olmaydi.

---

## 🧑‍💻 Loyiha Ijodkori

*   **Sizning Ismingiz** - [@SizningTelegramingiz](https://t.me)
*   GitHub: [@SizningGitHubingiz](https://github.com)
