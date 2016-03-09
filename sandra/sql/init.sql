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

CREATE TABLE data
(
    id      int,
    customer   text,
    service text,
    resptime     int,
    created timestamp,
    primary key (service, customer, created)

) ;
INSERT INTO data ( id, customer, service, resptime, created) VALUES ( 1, 'NXG', 'WWW', 10, '2016-01-01 10:10:10');
