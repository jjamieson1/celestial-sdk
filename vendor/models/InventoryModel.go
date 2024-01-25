package models

type InventoryRequest struct {
	InventoryItems 	[]InventoryItem `json:"inventoryItems"`
}

type InventoryResponse struct {
	InventoryItems 	[]InventoryItem `json:"inventoryItems"`
}

type InventoryItem struct {
	ItemId 				string 			`json:"itemId"`
	Quantity 			int 			`json:"quantity"`
	CatalogProvider 	string			`json:"catalogProvider"`
	CanFullFill 		bool 			`json:"canFullFill"`
}

