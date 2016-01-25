package main

import (
	"log"
	"sync"
	"time"

	"github.com/ssdb/gossdb/ssdb"
)

var (
	db       *ssdb.Client
	testHash = "testhash"
)

type Worker interface {
	Run()
}

type Writer struct {
	Count int
	conn  *ssdb.Client
}

func NewWriter(conn *ssdb.Client, count int) *Writer {
	return &Writer{Count: count, conn: conn}
}

func (w *Writer) Run() {
	var start = time.Now()
	for i := 0; i < w.Count; i++ {
		_, err := w.conn.Do("hincr", testHash, i, 1)
		if err != nil {
			log.Println(err)
		}
	}
	log.Println("Write", w.Count, "took", time.Now().Sub(start))

}

func main() {
	log.Println("Main")
	conn, err := ssdb.Connect("127.0.0.1", 8888)
	log.Println(conn, err)
	//w := Writer{10000, conn}
	//log.Println(w)
	//w.Run()
	now := time.Now()
	numberOfWorker := 3
	var wg sync.WaitGroup
	wg.Add(numberOfWorker)
	for i := 0; i < numberOfWorker; i++ {
		go func(i int, conn *ssdb.Client, wg *sync.WaitGroup) {
			w := Writer{100000, conn}
			w.Run()
			log.Println("Writer", i, "finished")
			wg.Done()
		}(i, conn, &wg)

	}

	wg.Wait()

	log.Println(numberOfWorker, "Writer", "took", time.Now().Sub(now))

}
