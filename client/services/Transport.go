package services

import (
	"context"
	"net/http"
	"strconv"
)

// 拼接url参数
func ProductDetailEncode(ctx context.Context,httpRequest *http.Request,requestParam interface{}) error {
	pdr := requestParam.(ProductDetailRequest)
	httpRequest.URL.Path += "/product/" + strconv.Itoa(pdr.ProductId)
	return nil
}
