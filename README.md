# go-igdb #
-

**WIP**. Golang Wrapper for [IGDB.com](https://www.igdb.com) v2 API.

Read more: https://market.mashape.com/igdbcom/internet-game-database

###Usage

```go
import "github.com/aldidana/go-igdb/igdb"
```

Construct a new IGDB client, then use the various services on the client to access IGDB API. For example:

```go
client := igdb.NewClient(os.Getenv("igdbKey"))

companies, _ := client.GetCompanies("name", 1, 0)

genres, _ := client.GetGenres("*", 10)
```