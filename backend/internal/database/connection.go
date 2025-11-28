package connection

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"

	"os"
)

func Connect() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("success!")
	defer conn.Close(context.Background())
}
