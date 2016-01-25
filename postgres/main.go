package main

import (
	"log"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	db       *sqlx.DB
	testHash = "testhash"
)

type Worker interface {
	Run()
}

type Writer struct {
	Count int
	conn  *sqlx.DB
}

func NewWriter(conn *sqlx.DB, count int) *Writer {
	return &Writer{Count: count, conn: conn}
}

func (w *Writer) Run() {
	var start = time.Now()
	for i := 0; i < w.Count; i++ {
		_, err := db.Exec(`INSERT INTO test (resptime) VALUES ( $1)`, i)
		if err != nil {
			log.Println(err)
		}
	}
	log.Println("Write", w.Count, "took", time.Now().Sub(start))

}

func main() {
	log.Println("Main")
	db, _ = sqlx.Open("postgres", "dbname=terry host=localhost user=terry sslmode=disable")
	log.Println(db)

	now := time.Now()

	//numberOfWorker := 1
	//w := Writer{100000, db}
	//w.Run()

	numberOfWorker := 3
	var wg sync.WaitGroup
	wg.Add(numberOfWorker)
	for i := 0; i < numberOfWorker; i++ {
		go func(i int, conn *sqlx.DB, wg *sync.WaitGroup) {
			w := Writer{100000, conn}
			w.Run()
			log.Println("Writer", i, "finished")
			wg.Done()
		}(i, db, &wg)

	}

	wg.Wait()

	log.Println(numberOfWorker, "Writer", "took", time.Now().Sub(now))

}
