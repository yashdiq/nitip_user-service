# Go Microservices

Clean microservices structures for scalability project crafted with Go + Cassandra + Protobuf + gRPC + Wire.

## Database Configuration
```
  docker exec -it cassandra cqlsh
    CREATE KEYSPACE <keyspace> WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};
    CREATE TABLE up_users (id text PRIMARY KEY, username text, email text, password text, first_name text, last_name text,  phone text, avatar text, blocked boolean, fcm_token text)
```

## Running proto
```
  protoc --go_out=. --go-grpc_out=. pkg/pb/*.proto
```