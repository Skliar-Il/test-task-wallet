package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Skliar-Il/test-task-wallet/internal/dto"
	"github.com/google/uuid"
	"net/http"
	"testing"
)

func sendWalletUpdateRequest(t *testing.T, update dto.UpdateWalletDTO, expectedStatus int) dto.WalletDTO {
	bodyBytes, err := json.Marshal(update)
	if err != nil {
		t.Fatalf("marshal request body error: %v", err)
	}

	resp, err := http.Post(fmt.Sprintf("%s/wallet", baseURL), "application/json", bytes.NewReader(bodyBytes))
	if err != nil {
		t.Fatalf("send POST request error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != expectedStatus {
		t.Fatalf("expected status %d, got %d", expectedStatus, resp.StatusCode)
	}

	var wallet dto.WalletDTO
	if expectedStatus == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(&wallet); err != nil {
			t.Fatalf("decode response error: %v", err)
		}
	}

	return wallet
}

func TestWalletDepositWithdraw(t *testing.T) {
	walletID := uuid.New()

	deposit := dto.UpdateWalletDTO{
		WalletId:      walletID,
		OperationType: "DEPOSIT",
		Amount:        1000,
	}
	wallet := sendWalletUpdateRequest(t, deposit, http.StatusOK)
	if wallet.Amount != 1000 {
		t.Errorf("expected amount 1000 after deposit, got %d", wallet.Amount)
	}

	withdraw := dto.UpdateWalletDTO{
		WalletId:      walletID,
		OperationType: "WITHDRAW",
		Amount:        500,
	}
	wallet = sendWalletUpdateRequest(t, withdraw, http.StatusOK)
	if wallet.Amount != 500 {
		t.Errorf("expected amount 500 after withdraw, got %d", wallet.Amount)
	}

}

func TestWalletGet(t *testing.T) {
	walletID := uuid.New()

	resp, err := http.Get(fmt.Sprintf("%s/wallets/%s", baseURL, walletID))
	if err != nil {
		t.Fatalf("send GET request error: %v", err)
	}
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("expected 404 for missing wallet, got %d", resp.StatusCode)
	}
	resp.Body.Close()

	deposit := dto.UpdateWalletDTO{
		WalletId:      walletID,
		OperationType: "DEPOSIT",
		Amount:        750,
	}
	sendWalletUpdateRequest(t, deposit, http.StatusOK)

	resp, err = http.Get(fmt.Sprintf("%s/wallets/%s", baseURL, walletID))
	if err != nil {
		t.Fatalf("send GET request error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 on get wallet, got %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	var wallet dto.WalletDTO
	if err := json.NewDecoder(resp.Body).Decode(&wallet); err != nil {
		t.Fatalf("decode GET response error: %v", err)
	}

	if wallet.Amount != 750 {
		t.Errorf("expected amount 750, got %d", wallet.Amount)
	}
}
