package main

import (
	"errors"
	"github.com/edsrzf/mmap-go"
	"log"
	"os"
	"sync/atomic"
	"time"
	"unsafe"
)

const MMFile = "mmping.file"

func main() {

	println("Two Processes: One Pings the other Pongs; using memory mapped files.")

	// creates a file only if it does not exist
	if _, err := os.Stat(MMFile); errors.Is(err, os.ErrNotExist) {
		if f, err := os.OpenFile(MMFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644); err != nil {
			log.Panic(err)
		} else {
			if _, err = f.Seek(256-1, 1); err != nil {
				log.Panic(err)
			}
			if _, err = f.Write([]byte{0}); err != nil {
				log.Panic(err)
			}
			if err = f.Close(); err != nil {
				log.Panic(err)
			}
		}
	}

	// opens the memory mapped file
	f, err := os.OpenFile(MMFile, os.O_RDWR, 0644)
	defer func() { _ = f.Close() }()
	if err != nil {
		log.Panic(err)
	}
	mapped, err := mmap.Map(f, mmap.RDWR, 0)
	if err != nil {
		log.Panic(err)
	}
	// converts the byte array into an array of uint64
	a32 := (*[32]uint64)(unsafe.Pointer(&mapped[0]))

	// ping ping
	t := time.Now().UnixNano()
	if len(os.Args) == 2 && os.Args[1] == "PING" {
		atomic.StoreUint64(&a32[0], 1) // resets counter
		for {
			ping := atomic.LoadUint64(&a32[0])
			pong := atomic.LoadUint64(&a32[1])
			if ping == pong {
				ping++
				atomic.CompareAndSwapUint64(&a32[0], ping-1, ping)
				if ping%10_000_000 == 0 {
					t = time.Now().UnixNano() - t
					pingTimeNs := float64(t) / 10_000_000.0
					log.Printf("PING %vM, %.2f/ns per ping, or %.1fM pings/sec\n",
						ping/1_000_000,
						pingTimeNs,
						1_000_000_000/pingTimeNs/1_000_000)
					t = time.Now().UnixNano()
				}
			}
		}

	} else if len(os.Args) == 2 && os.Args[1] == "PONG" {
		for {
			ping := atomic.LoadUint64(&a32[0])
			pong := atomic.LoadUint64(&a32[1])
			if pong != ping {
				atomic.CompareAndSwapUint64(&a32[1], pong, ping)
				if ping%10_000_000 == 0 {
					log.Println("PONG", ping)
				}
			}
		}

	} else {
		println("Either use PING or PONG as parameter")
	}

}
