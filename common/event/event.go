package event

import "time"

// Base 所有事件通用字段
type Base struct {
	EventID    string    `json:"event_id"`
	EventType  string    `json:"event_type"`
	OccurredAt time.Time `json:"occurred_at"`
	Producer   string    `json:"producer"`
}

// Payload 接口，具体业务实现
type Payload interface{}

// Envelope 事件信封
type Envelope struct {
	Base
	Payload Payload `json:"payload"`
}

// --- 事件类型常量（topic 名）---
const (
	TopicSalesOrderShipped       = "sales.order.shipped"
	TopicPurchaseWarehousingCreated = "purchase.warehousing.created"
	TopicPurchaseReturnCreated   = "purchase.return.created"
	TopicProductionWarehousingCreated = "production.warehousing.created"
	TopicQualityInspectionPassed = "quality.inspection.passed"
	TopicQualityInspectionFailed = "quality.inspection.failed"
	TopicPaymentReceived         = "payment.received"
)

// --- 各 topic 的 Payload 结构体 ---

// SalesOrderShippedPayload 发货完成
type SalesOrderShippedPayload struct {
	OrderID      int64   `json:"order_id"`
	OrderNumber  string  `json:"order_number"`
	CustomerID   int64   `json:"customer_id"`
	TotalAmount  string  `json:"total_amount"` // decimal string
	ShippedAt    string  `json:"shipped_at"`
}

// PurchaseWarehousingCreatedPayload 采购入库
type PurchaseWarehousingCreatedPayload struct {
	WarehousingID   int64  `json:"warehousing_id"`
	WarehousingNo   string `json:"warehousing_no"`
	PurchaseOrderID int64  `json:"purchase_order_id"`
	WarehouseID     int64  `json:"warehouse_id"`
	Items           []PurchaseWarehousingItemPayload `json:"items"`
	CreatedAt       string `json:"created_at"`
}

type PurchaseWarehousingItemPayload struct {
	ProductID   int64  `json:"product_id"`
	Quantity    int64  `json:"quantity"`
	LocationID  int64  `json:"location_id,optional"`
}

// PurchaseReturnCreatedPayload 采购退货
type PurchaseReturnCreatedPayload struct {
	ReturnOrderID int64  `json:"return_order_id"`
	ReturnNo      string `json:"return_no"`
	Items         []PurchaseReturnItemPayload `json:"items"`
	CreatedAt     string `json:"created_at"`
}

type PurchaseReturnItemPayload struct {
	ProductID int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
}

// ProductionWarehousingCreatedPayload 生产入库
type ProductionWarehousingCreatedPayload struct {
	ProductionOrderID int64   `json:"production_order_id"`
	Items             []ProductionWarehousingItemPayload `json:"items"`
	CreatedAt         string  `json:"created_at"`
}

type ProductionWarehousingItemPayload struct {
	ProductID  int64 `json:"product_id"`
	Quantity   int64 `json:"quantity"`
	LocationID int64 `json:"location_id,optional"`
}

// QualityInspectionPayload 质检结果
type QualityInspectionPayload struct {
	InspectionID   int64  `json:"inspection_id"`
	InspectionType string `json:"inspection_type"` // shipping/production/incoming/first_article
	Passed         bool   `json:"passed"`
	Remarks        string `json:"remarks,optional"`
}

// PaymentReceivedPayload 收款
type PaymentReceivedPayload struct {
	ReceivableID int64  `json:"receivable_id"`
	Amount       string `json:"amount"`
	PaidAt       string `json:"paid_at"`
}
