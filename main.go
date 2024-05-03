package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

type User struct {
    UserID   int
    Username string
    Email    string
    Password string
}

type Friendship struct {
    ID       int
    UserID   int
    FriendID int
    Status   string
}

func main() {
    db, err := sql.Open("postgres", "user=postgres password=vakhaboff dbname=shaxboz sslmode=disable")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()
    sendFriendRequest(db, 1, 2)
    acceptFriendRequest(db, 1, 2)
    blockUser(db, 1, 2) 
}

func sendFriendRequest(db *sql.DB, userID, friendID int) {
    _, err := db.Exec("INSERT INTO friendships (user_id, friend_id, status) VALUES ($1, $2, 'pending')", userID, friendID)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("Do'stlik so'rovi %d dan %d ga yuborildi\n", userID, friendID)
}

func acceptFriendRequest(db *sql.DB, userID, friendID int) {
    _, err := db.Exec("UPDATE friendships SET status = 'accepted' WHERE user_id = $1 AND friend_id = $2", friendID, userID)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("%d tomonidan %d ga do'stlik so'rovi qabul qilindi\n", friendID, userID)
}

func blockUser(db *sql.DB, userID, friendID int) {
    _, err := db.Exec("UPDATE friendships SET status = 'blocked' WHERE user_id = $1 AND friend_id = $2", userID, friendID)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("%d tomonidan %d bloklandi\n", userID, friendID)
}
