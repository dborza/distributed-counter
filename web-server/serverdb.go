package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
    "sync/atomic"
    "context"

    //"github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgx"
    "github.com/jackc/pgx/v4"
)

func main() {

    config, err := pgx.ParseConfig("postgres://root@localhost:26257/counter_db.counter?sslmode=disable")
    if err != nil {
        log.Fatal("error configuring the database: ", err)
    }

    //config.TLSConfig.ServerName = "localhost"

    conn, err := pgx.ConnectConfig(context.Background(), config)
    if err != nil {
        log.Fatal("error connecting to the database: ", err)
    }
    defer conn.Close(context.Background())

    if _, err := conn.Exec(context.Background(),
        "UPDATE counter_db.counter SET value=value+1 where id=1;"); err != nil {
        log.Fatal(err)
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi")
    })

    var ops uint64

    http.HandleFunc("/inc-cockroachdb", func(w http.ResponseWriter, r *http.Request){
        if _, err := conn.Exec(context.Background(),
            "UPDATE counter_db.counter SET value=value+1 where id=1;"); err != nil {
            log.Fatal(err)
        }
        fmt.Fprintf(w, "inc-cockroachdb")
    })

    http.HandleFunc("/inc", func(w http.ResponseWriter, r *http.Request){
        atomic.AddUint64(&ops, 1)
        fmt.Fprintf(w, "Count: %d", ops)
    })

    http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Count: %d", ops)
    })

    log.Fatal(http.ListenAndServe(":8081", nil))

}
