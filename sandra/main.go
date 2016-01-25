package main

import (
	"log"
	"sync"
	"time"

	"github.com/gocql/gocql"
)

var (
	db       *gocql.ClusterConfig
	testHash = "testhash"
)

type Worker interface {
	Run()
}

type Writer struct {
	Count int
	conn  *gocql.ClusterConfig
}

func NewWriter(conn *gocql.ClusterConfig, count int) *Writer {
	return &Writer{Count: count, conn: conn}
}

func (w *Writer) Run() {
	var start = time.Now()
	session, _ := w.conn.CreateSession()
	defer session.Close()

	for i := 0; i < w.Count; i++ {
		if err := session.Query(`INSERT INTO test (id, resptime) VALUES (?,?)`, i, i).Exec(); err != nil {
			log.Println(err)
		}
	}
	log.Println("Write", w.Count, "took", time.Now().Sub(start))

}

func main() {
	log.Println("Main")
	db := gocql.NewCluster("127.0.0.1")
	db.Keyspace = "test"
	db.ProtoVersion = 4
	db.Consistency = gocql.Quorum

	log.Println(db)

	now := time.Now()

	//numberOfWorker := 1
	//w := Writer{100000, db}
	//w.Run()

	numberOfWorker := 3
	var wg sync.WaitGroup
	wg.Add(numberOfWorker)
	for i := 0; i < numberOfWorker; i++ {
		go func(i int, conn *gocql.ClusterConfig, wg *sync.WaitGroup) {
			w := Writer{100000, conn}
			w.Run()
			log.Println("Writer", i, "finished")
			wg.Done()
		}(i, db, &wg)

	}

	wg.Wait()

	log.Println(numberOfWorker, "Writer", "took", time.Now().Sub(now))

}
