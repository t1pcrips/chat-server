package helpers

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/t1pcrips/chat-service/internal/service/mocks"
	dbMock "github.com/t1pcrips/platform-pkg/pkg/database/mocks"
	"github.com/t1pcrips/platform-pkg/pkg/database/postgres"
	"math/rand"
	"testing"
)

const (
	SuccessDelete = "success delete test"
	FailedDelete  = "failed delete test"
	SuccessCreate = "success create test"
	FailedCreate  = "failed create test"
	SuccessSend   = "success send test"
	FailedSend    = "failed send test"
)

func Names(numNames int) []string {
	// Массив имен
	names := []string{"Alice", "Bob", "Charlie", "David", "Eve", "Frank", "Grace", "Heidi", "Tom", "Keyle"}

	// Генерация массива случайных имен
	randomNames := make([]string, numNames)
	for i := 0; i < numNames; i++ {
		randomNames[i] = names[rand.Intn(len(names))]
	}
	return randomNames
}

func TxFakerAndCtxWithSetup(ctx context.Context, t *testing.T, successTx bool) (*mocks.FakeTxMock, context.Context) {
	t.Helper()

	txFaker := mocks.NewFakeTxMock(t)
	ctxWithTx := postgres.MakeContextTx(ctx, txFaker)

	if successTx {
		txFaker.CommitMock.Expect(ctxWithTx).Return(nil)
	} else {
		txFaker.RollbackMock.Expect(ctxWithTx).Return(nil)
	}

	return txFaker, ctxWithTx
}

func FakeTransactorMock(ctx context.Context, t *testing.T, fakeTx *mocks.FakeTxMock) *dbMock.TransactorMock {
	txOpts := pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
	transactorMock := dbMock.NewTransactorMock(t)
	transactorMock.BeginTxMock.Expect(ctx, txOpts).Return(fakeTx, nil)

	return transactorMock
}
