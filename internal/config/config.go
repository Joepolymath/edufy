package config

import (
	"log"

	// "github.com/joho/godotenv"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// EmailConfig /* this is used for the email configuration*/
type EmailConfig struct {
	EmailHostUser        string `envconfig:"EMAIL_USER" default:"info@learniumai.com"`
	EmailHostPassword    string `envconfig:"EMAIL_HOST_PASSWORD" default:""`
	EmailHost            string `envconfig:"EMAIL_HOST" default:"smtp.gmail.com"`
	EmailPort            string `envconfig:"EMAIL_PORT" default:"587"`
	EmailUseTLS          string `envconfig:"EMAIL_USE_TLS" default:"True"`
	EmailCustomerSupport string `envconfig:"EMAIL_CUSTOMER_SUPPORT" default:"dev.codertjay@gmail.com"`
	EmailInfoEmail       string `envconfig:"EMAIL_INFO_EMAIL" default:"dev.codertjay@gmail.com"`
}

// LearniumCustomConfig /* these are just custom settings and keys secret key  used in the project*/
type LearniumCustomConfig struct {
	LearniumServerPort         string `envconfig:"LEARNIUM_SERVER_PORT" default:"8008"`
	LearniumSecretKey          string `envconfig:"LEARNIUM_SECRET_KEY" default:""`
	LearniumSKHeader           string `envconfig:"LEARNIUM_SK_HEADER" default:""`
	LearniumEncryptionKey      string `envconfig:"LEARNIUM_ENCRYPTION_KEY" default:""`
	LearniumRandomStringLength int    `envconfig:"LEARNIUM_RANDOM_STRING_LENGTH" default:"5"`
	LearniumUploadDirectory    string `envconfig:"LEARNIUM_UPLOAD_DIRECTORY" default:"uploads"`
	LearniumStorageType        string `envconfig:"LEARNIUM_STORAGE_TYPE" default:"LOCAL"`
	LearniumSuperUserMail      string `envconfig:"LEARNIUM_SUPER_USER_MAIL" default:"learnium@mail.com"`
}

type PostgresConfig struct {
	PostgresUser         string `envconfig:"POSTGRES_USER" default:"postgres"`
	PostgresPassword     string `envconfig:"POSTGRES_PASSWORD" default:"postgres"`
	PostgresHost         string `envconfig:"POSTGRES_HOST" default:"localhost"`
	PostgresPort         int    `envconfig:"POSTGRES_PORT" default:"5432"`
	PostgresDatabaseName string `envconfig:"POSTGRES_DB_NAME" default:"learnium"`
	PostgresSSLMode      string `envconfig:"POSTGRES_SSL_MODE" default:"disable"`
}

type RedisConfig struct {
	RedisAddress  string `envconfig:"REDIS_ADDRESS" default:"localhost:6379"`
	RedisPassword string `envconfig:"REDIS_PASSWORD" default:"redis"`
	RedisUsername string `envconfig:"REDIS_USER" default:"user"`
}

type S3BucketConfig struct {
	S3BucketBaseURL   string `envconfig:"S3_BUCKET_BASE_URL" default:"https://your-bucket-name.s3.us-west-2.amazonaws.com/"`
	S3BucketName      string `envconfig:"S3_BUCKET_NAME" default:"S3BUCKET_NAME"`
	S3BucketRegion    string `envconfig:"S3_BUCKET_REGION" default:"us-west-2"`
	S3AccessKeyId     string `envconfig:"AWS_ACCESS_KEY_ID" default:"us-west-2"`
	S3AccessKeySecret string `envconfig:"AWS_ACCESS_KEY_SECRET" default:"us-west-2"`
}

type SentryConfig struct {
	SentryDSN              string `envconfig:"SENTRY_DSN" default:""`
	SentryEnableTracing    bool   `envconfig:"SENTRY_ENABLE_TRACING" default:"true"`
	SentryTracesSampleRate int    `envconfig:"SENTRY_TRACES_SAMPLE_RATE" default:"1"`
}

// GoogleConfig /* these are just custom settings and keys secreti  used in the project*/
type GoogleConfig struct {
	GoogleClientID         string `envconfig:"GOOGLE_CLIENT_ID" default:""`     // google auth
	GoogleClientSecret     string `envconfig:"GOOGLE_CLIENT_SECRET" default:""` // google auth secret
	GoogleCallbackUrl      string `envconfig:"GOOGLE_CALLBACK_URL" default:"http://localhost:8080/callback"`
	GoogleScopesEmailUrl   string `envconfig:"GOOGLE_SCOPES_EMAIL_URL" default:"https://www.googleapis.com/auth/userinfo.email"`
	GoogleScopesProfileUrl string `envconfig:"GOOGLE_SCOPES_PROFILE_URL" default:"https://www.googleapis.com/auth/userinfo.profile"`
	GoogleOauthUrl         string `envconfig:"GOOGLE_OAUTH_URL" default:"https://www.googleapis.com/oauth2/v2/userinfo"`
}

// FaceBookConfig /* these are just custom settings and keys secret  used in the project*/
type FaceBookConfig struct {
	FaceBookFields string `envconfig:"FaceBookFields" default:"first_name,last_name,email"`
}

// ConfigModel defines app config
type ConfigModel struct {
	EmailConfig
	LearniumCustomConfig
	PostgresConfig
	S3BucketConfig
	SentryConfig
	GoogleConfig
	FaceBookConfig
	RedisConfig
	ENVIRON             string `envconfig:"ENVIRON" default:"development"`
	PORT                string `envconfig:"PORT" default:"8005"`
	SendGridEmailAPIKey string `envconfig:"SENDGRID_EMAIL_API_KEY" default:""`
}

func Load(env_path string) ConfigModel {
	// Load environment variables from .env file
	var cfg ConfigModel
	// if err := godotenv.Load(".env"); err != nil {
	// 	log.Println("error loading config on startup or when called: ", err)
	// }
	// if err := utils.LoadEnv(env_path); err != nil {
	// 	log.Fatalln("error loading environment: ", err)
	// 	// return ConfigModel{}
	// }

	if err := godotenv.Load(); err != nil {
		log.Fatalln("error loading environment: ", err)
	}

	// Process environment variables and store them in the cfg variable
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalln("error loading environment: ", err)
		// return ConfigModel{}, err
	}

	return cfg
}

var Config = Load("/.env")
