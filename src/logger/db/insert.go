package db

import (
	"context"
	"fmt"
	"time"
)

func InsertRecord(logMessage string, userID string){
	Conn := Connect()
	defer Conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := Conn.ExecContext(ctx,
		"INSERT INTO log(description,userID) " +
		"VALUES('" + logMessage  +"','" +  userID + "')")
	if err != nil {
		fmt.Printf("Error %s when inserting row\n\n", err)
		return
	}
	Conn.Close()

}



