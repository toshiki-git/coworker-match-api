package common

// contextKey is a custom type for context keys to avoid collisions.
type contextKey string

// UserIDKey is the key used for storing the user ID in the context.
const UserIdKey contextKey = "userId"
