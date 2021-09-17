package bridges

import (
	"github.com/gregpr07/bridge-decoder/utils"
	"gorm.io/gorm"
)

type PolygonPOSBridgeTx struct {
	gorm.Model
	Successful			bool
	Type			string
	Method			string
	OriginFrom		string
	OriginTo		string
	Bridge          string
	From			string
	To              string
	Amount			string
	BaseToken		string
	Input           string
	GasUsed         string
	BlockNumber     int
	TransactionHash string
}

func (PolygonPOSBridgeTx) TableName() string {
	return utils.SCHEMA + ".PolygonPOSBridgeDepositTx"
}