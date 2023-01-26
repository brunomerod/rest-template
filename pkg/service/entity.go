package service

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math"
)

type Account struct {
	AccountId      primitive.ObjectID `bson:"_id,omitempty" json:"accountId"`
	DocumentNumber string             `bson:"document_number" json:"document_number"`
}

type OperationType struct {
	OperationTypeId int    `bson:"_id" json:"operationTypeId"`
	Description     string `bson:"description" json:"description"`
	OperationValue  string `bson:"operationValue" json:"operationValue"`
}

type Transaction struct {
	TransactionId   int                `bson:"_id,omitempty" json:"transaction_id"`
	AccountId       primitive.ObjectID `bson:"account_id" json:"account_id"`
	OperationTypeId int                `bson:"operation_type_id" json:"operation_type_id"`
	Amount          float64            `bson:"amount" json:"amount"`
	EventDate       string             `bson:"event_date" json:"event_date"`
}

func registeredOperations() []OperationType {
	var registeredOperation []OperationType

	registeredOperation = append(registeredOperation, OperationType{
		OperationTypeId: 1,
		Description:     "COMPRA A VISTA",
		OperationValue:  "-",
	})

	registeredOperation = append(registeredOperation, OperationType{
		OperationTypeId: 2,
		Description:     "COMPRA PARCELADA",
		OperationValue:  "-",
	})

	registeredOperation = append(registeredOperation, OperationType{
		OperationTypeId: 3,
		Description:     "SAQUE",
		OperationValue:  "-",
	})

	registeredOperation = append(registeredOperation, OperationType{
		OperationTypeId: 4,
		Description:     "PAGAMENTO",
		OperationValue:  "",
	})

	return registeredOperation
}

func transactionExists(operationTypeId int) bool {
	for _, v := range registeredOperations() {
		if v.OperationTypeId == operationTypeId {
			return true
		}
	}

	return false
}

func GetTransactionValue(operationTypeId int, amount float64) (float64, error) {
	// This guarantees that the value is always tied to the RegisteredOperations
	// And not to the value sent, as it's the domain of this service to register the transaction
	if math.Signbit(amount) {
		amount = amount * -1
	}

	if !transactionExists(operationTypeId) {
		return 0.0, errors.New("transaction type not found")
	}

	for _, v := range registeredOperations() {
		if v.OperationTypeId == operationTypeId && v.OperationValue == "-" {
			return amount * -1.0, nil
		}
	}

	return amount, nil
}
