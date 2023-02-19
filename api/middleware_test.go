package api

import (
	"fmt"
	"github.com/VL-037/go-bank/token"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const USERNAME = "username"
const DURATION = time.Minute

func addAuthorization(
	t *testing.T,
	req *http.Request,
	tokenMaker token.Maker,
	authorizationType string,
	username string,
	duration time.Duration) {
	token, err := tokenMaker.CreateToken(username, duration)
	require.NoError(t, err)

	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, token)
	req.Header.Set(AUTHORIZATION_HEADER_KEY, authorizationHeader)
}

func TestAuthMiddleware(t *testing.T) {
	testCases := []struct{
		name string
		setupAuth func(t *testing.T, req *http.Request, tokenMaker token.Maker)
		checkResposne func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:          "OK",
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, req, tokenMaker, AUTHORIZATION_TYPE_BEARER, USERNAME, DURATION)
			},
			checkResposne: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:          "UNAUTHORIZED - Invalid Header",
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {},
			checkResposne: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name:          "UNAUTHORIZED - Unsupported Authorization Type",
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, req, tokenMaker, "INVALID_AUTHORIZATION_TYPE", USERNAME, DURATION)
			},
			checkResposne: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name:          "UNAUTHORIZED - Invalid  Authorization Format",
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, req, tokenMaker, "", USERNAME, DURATION)
			},
			checkResposne: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name:          "UNAUTHORIZED - Expired Token",
			setupAuth: func(t *testing.T, req *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, req, tokenMaker, AUTHORIZATION_TYPE_BEARER, USERNAME, -DURATION)
			},
			checkResposne: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			server := newTestServer(t, nil)

			authPath := "/auth"
			server.router.GET(
				authPath,
				authMiddleware(server.tokenMaker),
				func(ctx *gin.Context) {
					ctx.JSON(http.StatusOK, gin.H{})
				})

			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, authPath, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
		})
	}
}
