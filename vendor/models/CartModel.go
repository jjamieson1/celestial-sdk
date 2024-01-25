package models

import "time"


type CartItem struct {
  CartItemId    string      `json:"cartItemId"`
  CartId        string      `json:"cartId"`
  UserId        string      `json:"userId"`
  ProductId     string      `json:"productId"`
  ProductName   string      `json:"productName"`
  ProductImage  string      `json:"productImage"`
  Quantity      int         `json:"quantity"`
  ProductPrice  int64       `json:"productPrice"`
  Created       time.Time   `json:"created"`
  Updated       time.Time   `json:"updated"`
  MeasurementType string    `json:"measurementType"`
  MeasurementValue float64  `json:"measurementValue"`

}

type Cart struct {
  CartId                    string     `json:"cartId"`
  CartItems                 []CartItem `json:"cartItems"`
  Total                     int64      `json:"total"`
  GstTax                    int64       `json:"gstTax"`
  PstTax                    int64       `json:"pstTax"`
  TotalWithDelivery         int64        `json:"totalWithDelivery"`
  TotalWithTaxAndDelivery   int64        `json:"totalWithTaxAndDelivery"`
  CartLimitTotal            float64      `json:"cartLimitTotal"`
  CartLimitValue           float64       `json:"cartLimitValue"`
  Complete                  bool        `json:"complete"`
  UserId                    string      `json:"userId"`
  StoreId                   int         `json:"storeId"`
  TenantId                  string      `json:"tenantId"`
}