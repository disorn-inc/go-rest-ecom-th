package config

import (
	"log"
	"math"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func LoadConfig(path string) IConfig {
	envMap, err := godotenv.Read(path)
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}

	return &config{
		app: &app{
			host: envMap["APP_HOST"],
			port: func() int {
				if envMap["APP_PORT"] == "" {
					return 8080
				}

				port, err := strconv.Atoi(envMap["APP_PORT"])
				if err != nil {
					log.Fatalf("Error loading .env file %v", err)
				}
				return port
			}(),
			name:    envMap["APP_NAME"],
			version: envMap["APP_VERSION"],
			reedTimeout: func() time.Duration {
				if envMap["APP_READ_TIMEOUT"] == "" {
					return 15
				}

				readTimeout, err := strconv.Atoi(envMap["APP_READ_TIMEOUT"])
				if err != nil {
					log.Fatalf("Error loading .env file %v", err)
				}
				return time.Duration(int64(readTimeout) * int64(math.Pow10(9)))
			}(),
			writeTimeout: func() time.Duration {
				writeTimeout, err := strconv.Atoi(envMap["APP_WRITE_TIMEOUT"])
				if err != nil {
					log.Fatalf("Error loading .env file %v", err)
				}
				return time.Duration(int64(writeTimeout) * int64(math.Pow10(9)))
			}(),
			bodyLimit: func() int {
				if envMap["APP_BODY_LIMIT"] == "" {
					return 10490000
				}

				bodyLimit, err := strconv.Atoi(envMap["APP_BODY_LIMIT"])
				if err != nil {
					log.Fatalf("Error loading body limit .env file %v", err)
				}
				return bodyLimit
			}(),
			fileLimit: func() int {
				if envMap["APP_FILE_LIMIT"] == "" {
					return 1048576
				}

				fileLimit, err := strconv.Atoi(envMap["APP_FILE_LIMIT"])
				if err != nil {
					log.Fatalf("Error loading file limit .env file %v", err)
				}
				return fileLimit
			}(),
			gcpBucket: envMap["APP_GCP_BUCKET"],
		},
		db: &db{
			host: envMap["DB_HOST"],
			port: func() int {
				if envMap["DB_PORT"] == "" {
					return 5432
				}

				port, err := strconv.Atoi(envMap["DB_PORT"])
				if err != nil {
					log.Fatalf("Error loading .env file %v", err)
				}
				return port
			}(),
			protocol: envMap["DB_PROTOCOL"],
			username: envMap["DB_USERNAME"],
			password: envMap["DB_PASSWORD"],
			database: envMap["DB_DATABASE"],
			sslMode:  envMap["DB_SSL_MODE"],
			maxConnections: func() int {
				if envMap["DB_MAX_CONNECTIONS"] == "" {
					return 10
				}

				maxConnections, err := strconv.Atoi(envMap["DB_MAX_CONNECTIONS"])
				if err != nil {
					log.Fatalf("Error loading max connection .env file %v", err)
				}
				return maxConnections
			}(),
		},
		jwt: &jwt{
			adminKey:  envMap["JWT_ADMIN_KEY"],
			secretKey: envMap["JWT_SECRET_KEY"],
			apiKey:    envMap["JWT_API_KEY"],
			accessExpiresAt: func() int {
				accessExpiresAt, err := strconv.Atoi(envMap["JWT_ACCESS_EXPIRES"])
				if err != nil {
					log.Fatalf("Error loading jwt expire .env file %v", err)
				}
				return accessExpiresAt
			}(),
			refreshExpiresAt: func() int {
				refreshExpiresAt, err := strconv.Atoi(envMap["JWT_REFRESH_EXPIRES"])
				if err != nil {
					log.Fatalf("Error loading jwt refresh .env file %v", err)
				}
				return refreshExpiresAt
			}(),
		},
	}
}

type IConfig interface {
	App() IAppConfig
	Db() IDbConfig
	Jwt() IJwtConfig
}

type config struct {
	app *app
	db  *db
	jwt *jwt
}

func (c *config) App() IAppConfig {
	return c.app
}

func (c *config) Db() IDbConfig {
	return c.db
}

func (c *config) Jwt() IJwtConfig {
	return c.jwt
}

type IAppConfig interface {
	Url() string // host:port
	Name() string
	Version() string
	ReadTimeout() time.Duration
	WriteTimeout() time.Duration
	BodyLimit() int
	FileLimit() int
	GcpBucket() string
}

type app struct {
	host         string
	port         int
	name         string
	version      string
	reedTimeout  time.Duration
	writeTimeout time.Duration
	bodyLimit    int
	fileLimit    int
	gcpBucket    string
}

func (a *app) Url() string {
	return a.host + ":" + strconv.Itoa(a.port)
}

func (a *app) Name() string {
	return a.name
}

func (a *app) Version() string {
	return a.version
}

func (a *app) ReadTimeout() time.Duration {
	return a.reedTimeout
}

func (a *app) WriteTimeout() time.Duration {
	return a.writeTimeout
}

func (a *app) BodyLimit() int {
	return a.bodyLimit
}

func (a *app) FileLimit() int {
	return a.fileLimit
}

func (a *app) GcpBucket() string {
	return a.gcpBucket
}

type IDbConfig interface {
	Url() string // protocol://username:password@host:port/database?sslmode=sslMode
	MaxConnections() int
}

func (d *db) Url() string {
	return d.protocol + "://" + d.username + ":" + d.password + "@" +
		d.host + ":" + strconv.Itoa(d.port) + "/" + d.database + "?sslmode=" + d.sslMode
}

func (d *db) MaxConnections() int {
	return d.maxConnections
}

type db struct {
	host           string
	port           int
	protocol       string
	username       string
	password       string
	database       string
	sslMode        string
	maxConnections int
}

type IJwtConfig interface {
	AdminKey() []byte
	SecretKey() []byte
	ApiKey() []byte
	AccessExpiresAt() int
	RefreshExpiresAt() int
	SetJwtAccessExpires(t int)
	SetJwtRefreshExpires(t int)
}

func (j *jwt) AdminKey() []byte {
	return []byte(j.adminKey)
}

func (j *jwt) SecretKey() []byte {
	return []byte(j.secretKey)
}

func (j *jwt) ApiKey() []byte {
	return []byte(j.apiKey)
}

func (j *jwt) AccessExpiresAt() int {
	return j.accessExpiresAt
}

func (j *jwt) RefreshExpiresAt() int {
	return j.refreshExpiresAt
}

func (j *jwt) SetJwtAccessExpires(t int) {
	j.accessExpiresAt = t
}

func (j *jwt) SetJwtRefreshExpires(t int) {
	j.refreshExpiresAt = t
}

type jwt struct {
	adminKey         string
	secretKey        string
	apiKey           string
	accessExpiresAt  int
	refreshExpiresAt int // sec
}
