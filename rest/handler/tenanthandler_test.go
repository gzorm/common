package handler

import (
	"context"
	"github.com/gzorm/commons/rest/enum"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTenantHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	req.Header.Set("Tenant-Id", "tenant-1")
	handler := TenantHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// add tenant to context
		if tenantId := r.Header.Get(enum.TENANT_ID_HEADER_KEY); tenantId == "" {
			ctx = context.WithValue(ctx, enum.TENANT_ID_CTX_KEY, "1")
			ctx = metadata.AppendToOutgoingContext(ctx, enum.TENANT_ID_CTX_KEY, "1")
		} else {
			ctx = context.WithValue(ctx, enum.TENANT_ID_CTX_KEY, tenantId)
			ctx = metadata.AppendToOutgoingContext(ctx, enum.TENANT_ID_CTX_KEY, tenantId)
		}

		assert.Equal(t, "tenant-1", ctx.Value(enum.TENANT_ID_CTX_KEY).(string))
	}))

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestTenantHandlerByDefault(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost", http.NoBody)
	handler := TenantHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// add tenant to context
		if tenantId := r.Header.Get(enum.TENANT_ID_HEADER_KEY); tenantId == "" {
			ctx = context.WithValue(ctx, enum.TENANT_ID_CTX_KEY, "1")
			ctx = metadata.AppendToOutgoingContext(ctx, enum.TENANT_ID_CTX_KEY, "1")
		} else {
			ctx = context.WithValue(ctx, enum.TENANT_ID_CTX_KEY, tenantId)
			ctx = metadata.AppendToOutgoingContext(ctx, enum.TENANT_ID_CTX_KEY, tenantId)
		}

		assert.Equal(t, "1", ctx.Value(enum.TENANT_ID_CTX_KEY).(string))
	}))

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}
