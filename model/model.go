package redis

// Define url struct that stores as value in Redis database(key-value store).
type URL struct {
	URL    string `redis:"url"`
	Status string `redis:"status"`
}
