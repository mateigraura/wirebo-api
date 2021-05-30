package utils

const (
	DbHost    = "DB_HOST"
	DbUser    = "DB_USER"
	DbPsw     = "DB_PSW"
	DbName    = "DB_NAME"
	Port      = "API_PORT"
	MaxConn   = "MAX_CONN"
	MinConn   = "MIN_CONN"
	RedisHost = "REDIS_HOST"
	JWTSecret = "JWT_SECRET"
	JWTIssuer = "JWT_ISSUER"
	JWTExpiry = "JWT_EXPIRY_MINUTES"
)

const (
	Red   = "\u001B[31m"
	Green = "\033[32m"
)

const (
	prodEnv = ".env"
	devEnv  = ".env.dev"
	testEnv = ".env.testing"
)
