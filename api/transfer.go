package api

import (
	"net/http"

	db "github.com/ariefro/simple-bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type transferRequest struct {
	FromAccountId int64 `json:"from_account_id" binding:"required"`
	ToAccountId int64 `json:"to_account_id" binding:"required"`
	Amount int64 `json:"amount" binding:"required,gt=0"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR SGD"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountId,
		ToAccountID: req.ToAccountId,
		Amount: req.Amount,
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, result)
}