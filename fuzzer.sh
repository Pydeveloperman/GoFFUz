#!/bin/bash

# Professional ANSI Matrix & Cyberpunk ranglar palitrasi
Reset="\033[0m"
Bold="\033[1m"
MatrixGreen="\033[38;5;46m"
NeonCyan="\033[38;5;51m"
CyberPink="\033[38;5;198m"
GoldYellow="\033[38;5;220m"
BloodRed="\033[38;5;196m"
DarkGray="\033[38;5;240m"
White="\033[38;5;255m"

# 1. CYBERPUNK ASCII ART BANNER & INTERFEYS
clear
echo -e "${MatrixGreen}${Bold}"
echo " ██████╗  ██████╗ ███████╗██╗   ██╗███████╗███████╗"
echo "██╔════╝ ██╔═══██╗██╔════╝██║   ██║╚══███╔╝╚══███╔╝"
echo "██║  ███╗██║   ██║█████╗  ██║   ██║  ███╔╝   ███╔╝ "
echo "██║   ██║██║   ██║██╔══╝  ██║   ██║ ███╔╝   ███╔╝  "
echo "╚██████╔╝╚██████╔╝██║     ╚██████╔╝███████╗███████╗"
echo " ╚═════╝  ╚═════╝ ╚═╝      ╚═════╝ ╚══════╝╚══════╝"
echo -e "           ${CyberPink}[ AUTOMATED BASH AUTOMATION ]${Reset}\n"

echo -e "${DarkGray}┌───[ ENUMERATION ENGINE INFO ]${Reset}"
echo -e "${DarkGray}│${Reset} ${NeonCyan}Vazifasi ${DarkGray}=>${Reset} Multi-threaded Web Directory Fuzzer"
echo -e "${DarkGray}│${Reset} ${NeonCyan}Arxitektura ${DarkGray}=>${Reset} Go (Golang Core) & Advanced Bash Wrapper"
echo -e "${DarkGray}│${Reset} ${NeonCyan}Muallif   ${DarkGray}=>${Reset} ${MatrixGreen}Siz (Kali Hacker)${Reset}"
echo -e "${DarkGray}└───${DarkGray}[ PRO EDITION v2.2 ]${Reset}\n"

# 2. SECLISTS LUG'ATINI INTEGRATSIYA QILISH VA TEKSHIRUV
WORDLIST_PATH="/usr/share/seclists/Discovery/Web-Content/directory-list-2.3-small.txt"
DIRB_FALLBACK="/usr/share/wordlists/dirb/common.txt"

echo -e "${DarkGray}[*]${Reset} Tizim lug'atlar bazasi tekshirilmoqda..."

# Agar seclists tizimda siqilgan (.gz) holatda bo'lsa, avtomatik ochish
if [ -f "/usr/share/seclists/Discovery/Web-Content/directory-list-2.3-small.txt.gz" ]; then
    echo -e "${GoldYellow}[!] Lug'at fayli arxivda aniqlandi. Ochilmoqda...${Reset}"
    sudo gunzip "/usr/share/seclists/Discovery/Web-Content/directory-list-2.3-small.txt.gz"
fi

# Agar seclists topilmasa, dirb lug'atiga o'tish
if [ ! -f "$WORDLIST_PATH" ]; then
    echo -e "${GoldYellow}[!] SecLists ko'rsatilgan manzilda topilmadi. Muqobil lug'at tekshirilmoqda...${Reset}"
    if [ -f "$DIRB_FALLBACK" ]; then
        WORDLIST_PATH="$DIRB_FALLBACK"
        echo -e "${MatrixGreen}[+] Muqobil lug'at muvaffaqiyatli ulandi: DIRB common.txt${Reset}"
    else
        echo -e "${CyberPink}[*] Paketlar tizimda yo'q. Avtomatik o'rnatish yuklanmoqda...${Reset}"
        sudo apt update && sudo apt install seclists dirb -y
    fi
else
    echo -e "${MatrixGreen}[+] Jangovar lug'at (SecLists) muvaffaqiyatli aniqlandi!${Reset}"
fi
echo ""

# 3. TAJRIBA PARAMETRLARINI INPUT QILISH
echo -e "${DarkGray}┌───[ HUJUM MAQSADINI SOZLASH ]${Reset}"

# Nishon URL kiritish
echo -ne "${DarkGray}│${Reset} ${NeonCyan}Nishon URL manzilini kiriting:${Reset}\n${DarkGray}│${Reset} ${CyberPink}»${Reset} "
read TARGET_URL
if [ -z "$TARGET_URL" ]; then
    echo -e "\n${BloodRed}[XATO] URL manzili bo'sh bo'lishi mumkin emas!${Reset}"
    exit 1
fi

# Tarmoqlar sonini kiritish
echo -ne "${DarkGray}│${Reset} ${NeonCyan}Tarmoqlar soni (Threads) [Standart: 50]:${Reset}\n${DarkGray}│${Reset} ${CyberPink}»${Reset} "
read THREADS
THREADS=${THREADS:-50}

# Output faylini kiritish
echo -ne "${DarkGray}│${Reset} ${NeonCyan}Natijalar fayli nomi [Standart: scan_results.txt]:${Reset}\n${DarkGray}│${Reset} ${CyberPink}»${Reset} "
read OUTPUT_FILE
OUTPUT_FILE=${OUTPUT_FILE:-scan_results.txt}

echo -e "${DarkGray}└───[ PARAMETRLAR SAQLANDI ]${Reset}\n"

echo -e "${GoldYellow}[*] GoFuzz yadrosi yuklanmoqda... Barcha tizim limitlari ochildi.${Reset}"
echo -e "${DarkGray}------------------------------------------------------------------------${Reset}"

# Tizim limitini dastur uchun ham ochib qo'yamiz
ulimit -n 100000 2>/dev/null

# 4. GO CORE NI ISHGA TUSHIRISH
if [ -f "main.go" ]; then
    go run main.go -u "$TARGET_URL" -w "$WORDLIST_PATH" -t "$THREADS" -o "$OUTPUT_FILE"
else
    echo -e "${BloodRed}[XATO] Tizim xatoligi: 'main.go' joriy papkada topilmadi!${Reset}"
    exit 1
fi
