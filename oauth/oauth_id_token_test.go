package google

import (
	"testing"

	assert "github.com/blend/go-sdk/assert"
)

const (
	testToken = `eyJhbGciOiJSUzI1NiIsImtpZCI6IjM3NmVhMWUyZjRjOTM3YzMzM2QxZTI0YjU2NDczOGZjMDRjOTkwMDkifQ.eyJhenAiOiIzNjgxOTQ4MjIxNTctYmI4NDlsZ2VudWJncjFsYXFyMmkwbDRtb3RoYnByOGQuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiIzNjgxOTQ4MjIxNTctYmI4NDlsZ2VudWJncjFsYXFyMmkwbDRtb3RoYnByOGQuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzdWIiOiIxMDg2OTE2MDg3MjA3NzE1ODY3MjgiLCJoZCI6ImJsZW5kLmNvbSIsImVtYWlsIjoid2lsbEBibGVuZC5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiYXRfaGFzaCI6Il9XaUFjWXJUOF82RmNnY2JXZHJHZVEiLCJub25jZSI6ImQzOWEzMzI2ZjYyNTQ5YTNhY2MzMDAwNDhkOWI5MzRmIiwiZXhwIjoxNTIxNzUxMDYzLCJpc3MiOiJhY2NvdW50cy5nb29nbGUuY29tIiwiaWF0IjoxNTIxNzQ3NDYzfQ.ZlCdvkzcb9BKuqNtHFpbyKTWChhgIOzlwNZPHrjGLvaqVxA02hsh9iraQAMFnhjRQs-j_EXsSd2343RfhUQbUaIE3nTVqFQ1CUOaMUSKRuLm4WuaJAoBYVs-fQCjBKFkN7ugbDGgadjJuSJKMjeGpwQV4-LMIA_Ud46DGXmelnP3CiMQIWwWTeous4TCiOadIb5FLnE-rnX-xOs-KrZ3pR5HjvUIWW4XPvvgVPeqbU--gzwO7S5ej25FrTVqYp0HXvELIiqF0xf-Fr6mlA98VylFkkWbKln1VbNjFD4Fiq4T_V3m4VhImigi7UbyUP52HF09ep2jd2L5fczuhTbMvQ`
)

func TestDeserializeJWTToken(t *testing.T) {
	assert := assert.New(t)

	jwt, err := DeserializeJWT(testToken)
	assert.Nil(err)
	assert.NotNil(jwt)
}
