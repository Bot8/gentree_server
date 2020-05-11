package services

import (
	"artarn/gentree/domain/user"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"strconv"
)

type (
	JWTService struct {
		encryptionKey []byte
	}
	ParsedToken struct {
		UserId int
	}
)

func (service JWTService) GetAuthToken(user *user.User) string {
	signer, err := service.getSigner()
	encrypter, err := service.getEncrypter()

	if err != nil {
		panic(err)
	}

	claims := getClaims(user)

	raw, err := jwt.SignedAndEncrypted(signer, encrypter).Claims(claims).CompactSerialize()
	if err != nil {
		panic(err)
	}

	return raw
}

func (service JWTService) ParseAuthToken(authToken string) (*ParsedToken, error) {
	token, err := jwt.ParseSignedAndEncrypted(authToken)

	if err != nil {
		return nil, err
	}

	decrypt, err := token.Decrypt(service.encryptionKey)

	if err != nil {
		return nil, err
	}

	claims := jwt.Claims{}

	err = decrypt.Claims(service.encryptionKey, &claims)

	if err != nil {
		return nil, err
	}

	userId, err := strconv.Atoi(claims.Subject)

	if err != nil {
		return nil, err
	}

	return &ParsedToken{UserId: userId}, nil
}

func getClaims(user *user.User) jwt.Claims {
	return jwt.Claims{
		Subject: strconv.Itoa(user.Id),
	}
}

func (service JWTService) getSigner() (jose.Signer, error) {
	signKey := jose.SigningKey{Algorithm: jose.HS256, Key: service.encryptionKey}
	signerOptions := (&jose.SignerOptions{}).WithType("JWT").WithContentType("JWT")

	return jose.NewSigner(signKey, signerOptions)
}

func (service JWTService) getEncrypter() (jose.Encrypter, error) {
	encryption := jose.A128GCM
	recipient := jose.Recipient{
		Algorithm: jose.DIRECT,
		Key:       service.encryptionKey,
	}
	encrypterOptions := (&jose.EncrypterOptions{}).WithType("JWT").WithContentType("JWT")
	return jose.NewEncrypter(encryption, recipient, encrypterOptions)
}

func NewJWTService(encryptionKey string) *JWTService {
	return &JWTService{encryptionKey: []byte(encryptionKey)}
}
