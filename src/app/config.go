package app

import (
	"context"
	"os"
	"time"

	"esb-test/library/logger"

	"github.com/go-playground/validator/v10"

	"github.com/spf13/viper"
)

/*
	All config should be required.
	Optional only allowed if zero value of the type is expected being the default value.
	time.Duration units are “ns”, “us” (or “µs”), “ms”, “s”, “m”, “h”. as in time.ParseDuration().
*/

type (
	MySQL struct {
		ConnURI            string        `mapstructure:"MS_CONN_URI" validate:"required"`
		MaxPoolSize        int           `mapstructure:"MS_MAX_POOL_SZE"` //Optional, default to 0 (zero value of int)
		MaxIdleConnections int           `mapstructure:"MS_MAX_IDLE_CONNECTIONS"`
		MaxIdleTime        time.Duration `mapstructure:"MS_MAX_IDLE_TIME"` //Optional, default to '0s' (zero value of time.Duration)
		MaxLifeTime        time.Duration `mapstructure:"MS_MAX_IDLE_TIME"` //Optional, default to '0s' (zero value of time.Duration)
	}
	AccessToken struct {
		ValidFor time.Duration `mapstructure:"ACCESS_TOKEN_VALID_FOR" validate:"required"`
	}
	RefreshToken struct {
		ValidFor time.Duration `mapstructure:"REFRESH_TOKEN_VALID_FOR" validate:"required"`
	}

	// JWTKey struct {
	// 	KeyId     string `mapstructure:"JWK_KID" validate:"required"`
	// 	SignKey   string `mapstructure:"ACCESS_TOKEN_RSA256_PRIVATE_KEY" validate:"required"` //RSA Private Key in PEM
	// 	VerifyKey string `mapstructure:"ACCESS_TOKEN_RSA256_PUBLIC_KEY" validate:"required"`  //RSA Public Key in PEM
	// }

	Ftp struct {
		Host       string `mapstructure:"FTP_HOST" validate:"required"`
		Username   string `mapstructure:"FTP_USERNAME" validate:"required"`
		Password   string `mapstructure:"FTP_PASSWORD" validate:"required"`
		Root       string `mapstructure:"FTP_ROOT" validate:"required"`
		BaseDomain string `mapstructure:"FTP_BASE_DOMAIN_IMAGE" validate:"required"`
	}

	Configuration struct {
		ServiceName string      `mapstructure:"SERVICE_NAME"`
		MySQL       MySQL       `mapstructure:",squash"`
		Translation Translation `mapstructure:",squash"`
		// AccessToken  AccessToken   `mapstructure:",squash"`
		// RefreshToken RefreshToken  `mapstructure:",squash"`
		// TokenIssuer  string        `mapstructure:"TOKEN_ISSUER" validate:"required"`
		IatLeeway time.Duration `mapstructure:"IAT_LEEWAY" validate:"required"` //Leeway time for iat to accommodate server time discrepancy
		// JWTKey       JWTKey        `mapstructure:",squash"`
		Ftp         Ftp `mapstructure:",squash"`
		BindAddress int `mapstructure:"BIND_ADDRESS" validate:"required"`
		LogLevel    int `mapstructure:"LOG_LEVEL" validate:"required"`
	}
)

func InitConfig(ctx context.Context) (*Configuration, error) {
	var cfg Configuration

	viper.SetConfigType("env")
	envFile := os.Getenv("ENV_FILE")
	if envFile == "" {
		envFile = ".env"
	}

	_, err := os.Stat(envFile)
	if !os.IsNotExist(err) {
		viper.SetConfigFile(envFile)

		if err := viper.ReadInConfig(); err != nil {
			logger.GetLogger(ctx).Errorf("failed to read config:%v", err)
			return nil, err
		}
	}

	viper.AutomaticEnv()

	if err := viper.Unmarshal(&cfg); err != nil {
		logger.GetLogger(ctx).Errorf("failed to bind config:%v", err)
		return nil, err
	}

	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			logger.GetLogger(ctx).Errorf("invalid config:%v", err)
		}
		logger.GetLogger(ctx).Errorf("failed to load config")
		return nil, err
	}

	logger.GetLogger(ctx).Infof("Config loaded: %+v", cfg)
	return &cfg, nil
}
