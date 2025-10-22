package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	_ "embed"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/open-policy-agent/opa/v1/rego"
)

// Core OPA policies.
var (
	//go:embed rego/authentication.rego
	regoAuthentication string

	//go:embed rego/authorization.rego
	regoAuthorization string
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// openssl genpkey -algorithm RSA -out private.pem -pkeyopt rsa_keygen_bits:2048
	// openssl rsa -pubout -in private.pem -out public.pem

	if err := genToken(); err != nil {
		return err
	}

	return nil
}

func genToken() error {
	method := jwt.GetSigningMethod(jwt.SigningMethodRS256.Name)

	type Claims struct {
		jwt.RegisteredClaims
		Roles []string `json:"roles"`
	}

	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "123456789",
			Issuer:    "ardan labs",
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(8760 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		},
		Roles: []string{"ADMIN"},
	}

	token := jwt.NewWithClaims(method, claims)
	token.Header["kid"] = "123"

	// -------------------------------------------------------------------------

	data, err := os.ReadFile("private.pem")
	if err != nil {
		return err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(data)
	if err != nil {
		return fmt.Errorf("parsing private key from PEM: %w", err)
	}

	// -------------------------------------------------------------------------

	str, err := token.SignedString(privateKey)
	if err != nil {
		return fmt.Errorf("signing token: %w", err)
	}

	fmt.Println("\nOUR TOKEN:")
	fmt.Println(str)

	// -------------------------------------------------------------------------

	// VERIFICATION SIDE

	parser := jwt.NewParser(jwt.WithValidMethods([]string{jwt.SigningMethodRS256.Name}))

	var clms Claims
	tkn, _, err := parser.ParseUnverified(str, &clms)
	if err != nil {
		return fmt.Errorf("error parsing token: %w", err)
	}

	kidRaw, exists := tkn.Header["kid"]
	if !exists {
		return fmt.Errorf("KID MISSING")
	}

	kid, ok := kidRaw.(string)
	if !ok {
		return fmt.Errorf("KID MALFORMED")
	}

	fmt.Println("KID: ", kid)

	// -------------------------------------------------------------------------

	// USED THE KID TO FIND THE PUBLIC KEY

	// OPA

	pubData, err := os.ReadFile("public.pem")
	if err != nil {
		return err
	}

	input := map[string]any{
		"Key":   pubData,
		"Token": str,
		"ISS":   "ardan labs",
	}

	query := fmt.Sprintf("x = data.%s.%s", "ardan.rego", "auth")

	ctx := context.Background()

	q, err := rego.New(
		rego.Query(query),
		rego.Module("policy.rego", regoAuthentication),
	).PrepareForEval(ctx)
	if err != nil {
		return fmt.Errorf("OPA prepare for eval failed for rule %q: %w", "auth", err)
	}

	results, err := q.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		return fmt.Errorf("OPA eval failed for rule %q: %w", "auth", err)
	}

	if len(results) == 0 {
		return fmt.Errorf("OPA policy evaluation for rule yielded no results")
	}

	result, ok := results[0].Bindings["x"].(bool)
	if !ok || !result {
		fmt.Println("OPA policy evaluation details", "rule", "auth", "results", results, "ok", ok)
		return fmt.Errorf("OPA policy rule not satisfied")
	}

	fmt.Println("TOKEN SIG VERIFIED")

	return nil
}

func genKey() error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return fmt.Errorf("generating key: %w", err)
	}

	privateFile, err := os.Create("private.pem")
	if err != nil {
		return fmt.Errorf("creating private file: %w", err)
	}
	defer privateFile.Close()

	privateBlock := pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	if err := pem.Encode(privateFile, &privateBlock); err != nil {
		return fmt.Errorf("encoding to private file: %w", err)
	}

	// -------------------------------------------------------------------------

	publicFile, err := os.Create("public.pem")
	if err != nil {
		return fmt.Errorf("creating public file: %w", err)
	}
	defer publicFile.Close()

	asn1Bytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return fmt.Errorf("marshaling public key: %w", err)
	}

	publicBlock := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	if err := pem.Encode(publicFile, &publicBlock); err != nil {
		return fmt.Errorf("encoding to public file: %w", err)
	}

	fmt.Println("private and public key files generated")

	return nil
}
