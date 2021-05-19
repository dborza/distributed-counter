package main

import (
    "log"
    "net/http"
    "context"

    "github.com/jackc/pgx/v4/pgxpool"
)

func main() {


    // Set connection pool configuration, with maximum connection pool size.
    config, err := pgxpool.ParseConfig("postgres://root@localhost:26257/counter_db.counter?sslmode=disable&pool_max_conns=40")
    if err != nil {
        log.Fatal("error configuring the database: ", err)
    }

    // Create a connection pool to the "bank" database.
    conn, err := pgxpool.ConnectConfig(context.Background(), config)
    if err != nil {
        log.Fatal("error connecting to the database: ", err)
    }
    defer conn.Close()

    http.HandleFunc("/inc-cockroachdb", func(w http.ResponseWriter, r *http.Request){
        if _, err := conn.Exec(context.Background(),
            "UPDATE counter_db.counter SET value=value+1 where id=1;"); err != nil {
            log.Fatal(err)
        }
    })

    log.Fatal(http.ListenAndServe(":8081", nil))

}
