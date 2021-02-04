package oauth2

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/store"
	oredis "github.com/go-oauth2/redis/v4"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var redisUri string
var client models.Client

func NewManager() (*manage.Manager, error) {
	if redisUri == "" {
		viper.SetConfigFile("services.env")
		if err := viper.ReadInConfig(); err != nil {
			return nil, err
		}

		redisUri = viper.GetString("REDIS_URL") + ":" + viper.GetString("REDIS_PORT")
		client = models.Client{
			ID:     viper.GetString("CLIENT_ID"),
			Secret: viper.GetString("CLIENT_SECRETE"),
		}
	}

	manager := manage.NewDefaultManager()

	manager.MapTokenStorage(oredis.NewRedisStore(&redis.Options{
		Addr:     redisUri,
		Password: "", // no password set
		DB:       0,  // use default DB
	}))
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("00000000"), jwt.SigningMethodHS512))

	clientStore := store.NewClientStore()
	clientStore.Set(client.ID, &client)
	manager.MapClientStorage(clientStore)
	return manager, nil
}
