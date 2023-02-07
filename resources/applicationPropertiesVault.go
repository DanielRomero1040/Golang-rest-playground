package resources

import (
	"context"
	"fmt"
	"log"

	vault "github.com/hashicorp/vault/api"
	"github.com/magiconair/properties"
)

var (
	DbHost     string
	DbName     string
	DbPassword string
	DbPort     string
	DbUser     string
	JwtSecret  string
)

type Config struct {
	Address string `properties:"vault.address"`
	Token   string `properties:"vault.token"`
	Secret  string `properties:"vault.secret"`
	Folder  string `properties:"vault.folder"`
}

func VaultConfig() {
	p := properties.MustLoadFile("resources/application.properties.conf", properties.UTF8)
	config := vault.DefaultConfig()

	var cfg Config
	if err := p.Decode(&cfg); err != nil {
		log.Fatal(err)
	}

	config.Address = cfg.Address

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("unable to initialize Vault client: %v", err)
	}

	// Authenticate
	client.SetToken(cfg.Token)

	// Read a secret from the default mount path for KV v2 in dev mode, "secret"
	secret, err := client.KVv2(cfg.Secret).Get(context.Background(), cfg.Folder)
	if err != nil {
		log.Fatalf("unable to read secret: %v", err)
	}

	setSecretString("jwt-secret", secret, &JwtSecret)
	setSecretString("db-host", secret, &DbHost)
	setSecretString("db-name", secret, &DbName)
	setSecretString("db-password", secret, &DbPassword)
	setSecretString("db-port", secret, &DbPort)
	setSecretString("db-user", secret, &DbUser)
}

func setSecretString(secretKeyName string, secret *vault.KVSecret, propertyToSet *string) {
	value, ok := secret.Data[secretKeyName].(string)

	if !ok {
		fmt.Printf("No se encuentra la propiedad %s en Vault, verifica el KeyName ", secretKeyName)
	}

	*propertyToSet = value
}
