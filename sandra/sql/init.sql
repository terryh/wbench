CREATE KEYSPACE test WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };;

CREATE TABLE test
(
    id      int,
    resp   text,
    resptime     int,
    created timestamp,
    primary key (id, resptime)

) ;
//CREATE INDEX test_created ON test.test ( created );

