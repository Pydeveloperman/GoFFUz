package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath" // Fayl manzilini aniqlash uchun qo'shildi
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

const (
	Reset       = "\033[0m"
	Bold        = "\033[1m"
	MatrixGreen = "\033[38;5;46m"
	NeonCyan    = "\033[38;5;51m"
	CyberPink   = "\033[38;5;198m"
	GoldYellow  = "\033[38;5;220m"
	BloodRed    = "\033[38;5;196m"
	DarkGray    = "\033[38;5;240m"
)

func init() {
	numCPU := runtime.NumCPU()
	if numCPU > 1 {
		runtime.GOMAXPROCS(numCPU - 1)
	} else {
		runtime.GOMAXPROCS(1)
	}

	var rLimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err == nil {
		rLimit.Cur = 65535
		rLimit.Max = 65535
		syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	}
}

func main() {
	targetURL := flag.String("u", "", "Nishon URL")
	wordlistPath := flag.String("w", "", "Wordlist")
	threads := flag.Int("t", 50, "Tarmoqlar soni")
	outputPath := flag.String("o", "", "Natija fayli")

	flag.Parse()

	if *targetURL == "" || *wordlistPath == "" {
		fmt.Printf("%s[!] XATOLIK: Noto'g'ri argumentlar kiritildi.%s\n", BloodRed, Reset)
		return
	}

	if !strings.HasSuffix(*targetURL, "/") {
		*targetURL += "/"
	}

	fileCount, err := os.Open(*wordlistPath)
	if err != nil {
		fmt.Printf("%s[!] XATOLIK: Wordlist ochilmadi: %v%s\n", BloodRed, err, Reset)
		return
	}
	var totalWords int64
	countScanner := bufio.NewScanner(fileCount)
	for countScanner.Scan() {
		if strings.TrimSpace(countScanner.Text()) != "" {
			totalWords++
		}
	}
	fileCount.Close()

	fmt.Print("\033[H\033[2J") // Terminalni tozalash
	fmt.Printf("%s%s", MatrixGreen, Bold)
	fmt.Println(" ██████╗  ██████╗ ███████╗██╗   ██╗███████╗███████╗")
	fmt.Println("██╔════╝ ██╔═══██╗██╔════╝██║   ██║╚══███╔╝╚══███╔╝")
	fmt.Println("██║  ███╗██║   ██║█████╗  ██║   ██║  ███╔╝   ███╔╝ ")
	fmt.Println("██║   ██║██║   ██║██╔══╝  ██║   ██║ ███╔╝   ███╔╝  ")
	fmt.Println("╚██████╔╝╚██████╔╝██║     ╚██████╔╝███████╗███████╗")
	fmt.Println(" ╚═════╝  ╚═════╝ ╚═╝      ╚═════╝ ╚══════╝╚══════╝")
	fmt.Printf("           %s[ V2.2 - ULTRA WEBFUZZER ]%s\n\n", CyberPink, Reset)

	fmt.Printf("%s┌───[ TARMOQ PARAMETRLARI ]%s\n", DarkGray, Reset)
	fmt.Printf("%s│ %sNishon URL  %s=> %s%s\n", DarkGray, NeonCyan, DarkGray, Reset, *targetURL)
	fmt.Printf("%s│ %sLug'atda    %s=> %s%d ta maxfiy yo'l\n", DarkGray, NeonCyan, DarkGray, Reset, totalWords)
	fmt.Printf("%s│ %sTarmoqlar   %s=> %s%d concurrent workers\n", DarkGray, NeonCyan, DarkGray, Reset, *threads)
	fmt.Printf("%s│ %sCPU Aktiv   %s=> %s%d/%d Yadrolar\n", DarkGray, NeonCyan, DarkGray, Reset, runtime.GOMAXPROCS(0), runtime.NumCPU())
	fmt.Printf("%s└───%s[ CRACKING IN PROGRESS... ]%s\n\n", DarkGray, CyberPink, Reset)

	file, err := os.Open(*wordlistPath)
	if err != nil {
		return
	}
	defer file.Close()

	var outFile *os.File
	if *outputPath != "" {
		outFile, err = os.OpenFile(*outputPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return
		}
		defer outFile.Close()
	}

	jobs := make(chan string, *threads*2)
	var wg sync.WaitGroup
	var processedWords int64

	go func() {
		chars := []string{"◢", "◣", "◤", "◥"}
		i := 0
		for atomic.LoadInt64(&processedWords) < totalWords {
			current := atomic.LoadInt64(&processedWords)
			percent := float64(current) / float64(totalWords) * 100

			barLength := 20
			filledLength := int(float64(barLength) * (float64(current) / float64(totalWords)))
			bar := strings.Repeat("█", filledLength) + strings.Repeat(" ", barLength-filledLength)

			fmt.Printf("\r%s[%s] %s[%s] %s%.2f%% %s(%d/%d)%s", 
				CyberPink, chars[i], MatrixGreen, bar, NeonCyan, percent, DarkGray, current, totalWords, Reset)
			
			i = (i + 1) % len(chars)
			time.Sleep(80 * time.Millisecond)
		}
	}()

	for w := 1; w <= *threads; w++ {
		wg.Add(1)
		go worker(*targetURL, jobs, &wg, outFile, &processedWords)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word == "" || strings.HasPrefix(word, "#") {
			continue
		}
		jobs <- word
	}

	close(jobs)
	wg.Wait()

	fmt.Printf("\r\033[K") // Progress barni tozalash
	
	// FAYL MANZILINI VA NOMINI INTERFEYSDA KO'RSATISH QISMI
	fmt.Printf("%s[+] REYD YAKUNLANDI! Barcha maxfiy yo'llar tekshirildi.%s\n\n", MatrixGreen, Reset)
	
	if *outputPath != "" {
		// Kiritilgan nom bo'yicha to'liq absolut manzilni hisoblash
		absPath, err := filepath.Abs(*outputPath)
		if err == nil {
			fmt.Printf("%s┌───[ REYD NATIJALARI SAQLANDI ]%s\n", DarkGray, Reset)
			fmt.Printf("%s│ %sFayl Nomi   %s=> %s%s\n", DarkGray, GoldYellow, DarkGray, Reset, filepath.Base(absPath))
			fmt.Printf("%s│ %sTo'liq Yo'l  %s=> %s%s\n", DarkGray, GoldYellow, DarkGray, Reset, absPath)
			fmt.Printf("%s└───%s[ LOG FAYL TAYYOR ]%s\n", DarkGray, MatrixGreen, Reset)
		}
	} else {
		fmt.Printf("%s[!] Eslatma: Natijalarni faylga saqlash uchun '-o' flagidan foydalanilmagan.%s\n", GoldYellow, Reset)
	}
}

func worker(baseURL string, jobs <-chan string, wg *sync.WaitGroup, outFile *os.File, processed *int64) {
	defer wg.Done()

	tr := &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		MaxIdleConns:        10000,
		MaxIdleConnsPerHost: 10000,
		MaxConnsPerHost:     10000,
		IdleConnTimeout:     15 * time.Second,
	}

	client := &http.Client{
		Timeout:   4 * time.Second,
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	var mu sync.Mutex

	for word := range jobs {
		fullURL := baseURL + word

		req, err := http.NewRequest("HEAD", fullURL, nil)
		if err != nil {
			atomic.AddInt64(processed, 1)
			continue
		}

		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
		req.Header.Set("Connection", "keep-alive")

		resp, err := client.Do(req)
		if err != nil {
			atomic.AddInt64(processed, 1)
			continue
		}

		if resp.StatusCode == 200 || resp.StatusCode == 301 || resp.StatusCode == 302 || resp.StatusCode == 403 {
			var statusText string
			switch resp.StatusCode {
			case 200:
				statusText = fmt.Sprintf("%s[SUCCESS-200]%s", MatrixGreen, Reset)
			case 403:
				statusText = fmt.Sprintf("%s[FORBIDDEN-403]%s", BloodRed, Reset)
			default:
				statusText = fmt.Sprintf("%s[REDIRECT-%d]%s", GoldYellow, resp.StatusCode, Reset)
			}

			fmt.Printf("\r\033[K%s ───> %s%s%s\n", statusText, Bold, fullURL, Reset)

			if outFile != nil {
				mu.Lock()
				outFile.WriteString(fmt.Sprintf("[%d] -> %s\n", resp.StatusCode, fullURL))
				mu.Unlock()
			}
		}

		resp.Body.Close()
		atomic.AddInt64(processed, 1)
	}
}
