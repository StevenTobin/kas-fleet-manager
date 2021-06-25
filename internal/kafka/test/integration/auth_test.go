package integration

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/api/openapi"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/auth"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/errors"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/test/mocks"

	"github.com/bxcodec/faker/v3"
	"github.com/dgrijalva/jwt-go"
	. "github.com/onsi/gomega"
	"gopkg.in/resty.v1"
)

func TestAuth_success(t *testing.T) {
	ocmServer := mocks.NewMockConfigurableServerBuilder().Build()
	defer ocmServer.Close()

	// setup the test environment, if OCM_ENV=integration then the ocmServer provided will be used instead of actual
	// ocm
	h, _, teardown := NewKafkaHelper(t, ocmServer)
	defer teardown()

	serviceAccount := h.NewAccount(h.NewID(), faker.Name(), faker.Email(), "13640203")
	restyResp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(h.CreateJWTString(serviceAccount)).
		Get(h.RestURL("/"))

	Expect(err).NotTo(HaveOccurred())
	Expect(restyResp.StatusCode()).To(Equal(http.StatusOK))

}

func TestAuthSucess_publicUrls(t *testing.T) {
	ocmServer := mocks.NewMockConfigurableServerBuilder().Build()
	defer ocmServer.Close()

	h, _, teardown := NewKafkaHelper(t, ocmServer)
	defer teardown()
	restyResp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		Get(h.RestURL("/"))
	Expect(err).To(BeNil())
	Expect(restyResp.StatusCode()).To(Equal(http.StatusOK))
}

func TestAuthSuccess_usingSSORHToken(t *testing.T) {
	ocmServer := mocks.NewMockConfigurableServerBuilder().Build()
	defer ocmServer.Close()

	h, _, teardown := NewKafkaHelper(t, ocmServer)
	serviceAccount := h.NewAccount(h.NewID(), faker.Name(), faker.Email(), "13640203")
	defer teardown()
	claims := jwt.MapClaims{
		"username":           nil,
		"preferred_username": serviceAccount.Username(),
	}

	token := h.CreateJWTStringWithClaim(serviceAccount, claims)

	restyResp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(token).
		Get(h.RestURL("/kafkas"))
	Expect(err).NotTo(HaveOccurred())
	Expect(restyResp.StatusCode()).To(Equal(http.StatusOK))
}

func TestAuthFailure_withoutToken(t *testing.T) {
	ocmServer := mocks.NewMockConfigurableServerBuilder().Build()
	defer ocmServer.Close()

	h, _, teardown := NewKafkaHelper(t, ocmServer)
	defer teardown()

	restyResp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		Get(h.RestURL("/kafkas"))
	Expect(err).To(BeNil())
	re := parseResponse(restyResp)
	Expect(re.Code).To(Equal(fmt.Sprintf("%s-%d", errors.ERROR_CODE_PREFIX, errors.ErrorUnauthenticated)))
	Expect(re.Reason).To(Equal("Request doesn't contain the 'Authorization' header or the 'cs_jwt' cookie"))
	Expect(restyResp.StatusCode()).To(Equal(http.StatusUnauthorized))
}

func TestAuthFailure_invalidTokenWithTypMissing(t *testing.T) {
	ocmServer := mocks.NewMockConfigurableServerBuilder().Build()
	defer ocmServer.Close()

	h, _, teardown := NewKafkaHelper(t, ocmServer)
	serviceAccount := h.NewAccount(h.NewID(), faker.Name(), faker.Email(), "13640203")
	defer teardown()
	claims := jwt.MapClaims{
		"typ": nil,
	}

	token := h.CreateJWTStringWithClaim(serviceAccount, claims)

	restyResp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(token).
		Get(h.RestURL("/kafkas"))
	Expect(err).To(BeNil())
	re := parseResponse(restyResp)
	Expect(re.Code).To(Equal(fmt.Sprintf("%s-%d", errors.ERROR_CODE_PREFIX, errors.ErrorUnauthenticated)))
	Expect(re.Reason).To(Equal("Bearer token doesn't contain required claim 'typ'"))
	Expect(restyResp.StatusCode()).To(Equal(http.StatusUnauthorized))
}

func TestAuthFailure_ExpiredToken(t *testing.T) {
	ocmServer := mocks.NewMockConfigurableServerBuilder().Build()
	defer ocmServer.Close()

	h, _, teardown := NewKafkaHelper(t, ocmServer)
	serviceAccount := h.NewAccount(h.NewID(), faker.Name(), faker.Email(), "13640203")
	defer teardown()
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Duration(-15) * time.Minute).Unix(),
	}

	token := h.CreateJWTStringWithClaim(serviceAccount, claims)
	restyResp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(token).
		Get(h.RestURL("/kafkas"))
	Expect(err).To(BeNil())
	re := parseResponse(restyResp)
	Expect(re.Code).To(Equal(fmt.Sprintf("%s-%d", errors.ERROR_CODE_PREFIX, errors.ErrorUnauthenticated)))
	Expect(re.Reason).To(Equal("Bearer token is expired"))
	Expect(restyResp.StatusCode()).To(Equal(http.StatusUnauthorized))
}

func TestAuthFailure_invalidTokenMissingIat(t *testing.T) {
	ocmServer := mocks.NewMockConfigurableServerBuilder().Build()
	defer ocmServer.Close()

	h, _, teardown := NewKafkaHelper(t, ocmServer)
	serviceAccount := h.NewAccount(h.NewID(), faker.Name(), faker.Email(), "13640203")
	defer teardown()
	claims := jwt.MapClaims{
		"iat": nil,
	}

	token := h.CreateJWTStringWithClaim(serviceAccount, claims)
	restyResp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(token).
		Get(h.RestURL("/kafkas"))
	Expect(err).To(BeNil())
	re := parseResponse(restyResp)
	Expect(re.Code).To(Equal(fmt.Sprintf("%s-%d", errors.ERROR_CODE_PREFIX, errors.ErrorUnauthenticated)))
	Expect(re.Reason).To(Equal("Bearer token doesn't contain required claim 'iat'"))
	Expect(restyResp.StatusCode()).To(Equal(http.StatusUnauthorized))
}

func TestAuthFailure_invalidTokenMissingAlgHeader(t *testing.T) {
	ocmServer := mocks.NewMockConfigurableServerBuilder().Build()
	defer ocmServer.Close()

	h, _, teardown := NewKafkaHelper(t, ocmServer)
	serviceAccount := h.NewAccount(h.NewID(), faker.Name(), faker.Email(), "13640203")
	defer teardown()
	claims := jwt.MapClaims{
		"username":   serviceAccount.Username(),
		"first_name": serviceAccount.FirstName(),
		"last_name":  serviceAccount.LastName(),
		"typ":        "Bearer",
		"iat":        time.Now().Unix(),
		"exp":        time.Now().Add(time.Minute * time.Duration(auth.TokenExpMin)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Header["kid"] = auth.JwkKID
	token.Header["alg"] = ""
	strToken, _ := token.SignedString(h.JWTPrivateKey)

	restyResp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(strToken).
		Get(h.RestURL("/kafkas"))
	Expect(err).To(BeNil())

	re := parseResponse(restyResp)
	Expect(re.Code).To(Equal(fmt.Sprintf("%s-%d", errors.ERROR_CODE_PREFIX, errors.ErrorUnauthenticated)))
	Expect(re.Reason).To(Equal("Bearer token can't be verified"))
	Expect(restyResp.StatusCode()).To(Equal(http.StatusUnauthorized))
}

func TestAuthFailure_invalidTokenUnsigned(t *testing.T) {
	ocmServer := mocks.NewMockConfigurableServerBuilder().Build()
	defer ocmServer.Close()

	h, _, teardown := NewKafkaHelper(t, ocmServer)
	serviceAccount := h.NewAccount(h.NewID(), faker.Name(), faker.Email(), "13640203")
	defer teardown()
	claims := jwt.MapClaims{
		"username":   serviceAccount.Username(),
		"first_name": serviceAccount.FirstName(),
		"last_name":  serviceAccount.LastName(),
		"typ":        "Bearer",
		"iat":        time.Now().Unix(),
		"exp":        time.Now().Add(time.Minute * time.Duration(auth.TokenExpMin)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Header["kid"] = auth.JwkKID
	strToken := token.Raw

	restyResp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetAuthToken(strToken).
		Get(h.RestURL("/kafkas"))
	Expect(err).To(BeNil())
	re := parseResponse(restyResp)
	Expect(re.Code).To(Equal(fmt.Sprintf("%s-%d", errors.ERROR_CODE_PREFIX, errors.ErrorUnauthenticated)))
	Expect(re.Reason).To(Equal("Request doesn't contain the 'Authorization' header or the 'cs_jwt' cookie"))
	Expect(restyResp.StatusCode()).To(Equal(http.StatusUnauthorized))
}

func parseResponse(restyResp *resty.Response) openapi.Error {
	var re openapi.Error
	if err := json.Unmarshal(restyResp.Body(), &re); err != nil {
		panic(err)
	}
	return re
}