package esi_test

import (
	"crypto/ecdsa"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-jose/go-jose"
	"github.com/maddiesch/esi"
	"github.com/planningcenter/signature"
)

var rootPath string
var signingKey *ecdsa.PrivateKey

var app = &esi.App{
	ClientID:  "foo",
	SecretKey: "bar",
	CallbackURL: &url.URL{
		Scheme: "http",
		Host:   "localhost:3000",
		Path:   "/oauth/esi/callback",
	},
}

func init() {
	_, f, _, _ := runtime.Caller(0)

	rootPath = filepath.Dir(f)

	pem, err := ioutil.ReadFile(filepath.Join(rootPath, "test-key.pem"))
	if err != nil {
		panic(err)
	}

	key, err := signature.UnmarshalPrivateKeyPem(pem, signature.KeyFormatASN1)
	if err != nil {
		panic(err)
	}

	signingKey = key
}

func TestMain(m *testing.M) {
	os.Exit(run(m))
}

func run(m *testing.M) int {
	server := createAuthorizationServer()
	defer server.Close()

	testURL, _ := url.Parse(server.URL)

	esi.HTTPClient = server.Client()
	esi.AuthenticationHost = testURL.Host

	return m.Run()
}

func createAuthorizationServer() *httptest.Server {
	gin.SetMode(gin.ReleaseMode)

	app := gin.New()

	app.Handle("GET", "/oauth/jwks", func(c *gin.Context) {
		key := jose.JSONWebKey{
			KeyID: "b97d26e6-f546-44ca-bd7b-15f31c61d853",
			Key:   signingKey,
		}

		c.JSON(http.StatusOK, gin.H{
			"keys": []interface{}{
				key.Public(),
			},
		})
	})

	return httptest.NewTLSServer(app)
}
