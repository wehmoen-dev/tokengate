package tokengate

import (
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

var testClient = New()

var testAddress = common.HexToAddress("0x2d62c27ce2e9e66bb8a667ce1b60f7cb02fa9810")

func TestClient_HasAxieFromCollection(t *testing.T) {
	hasMystic, err := testClient.HasAxieFromCollection(testAddress, MysticAxie)
	if err != nil {
		t.Error(err)
		return
	}

	if !hasMystic {
		t.Error("expected jihoz to have a mystic axie, got false")
	}
}

func TestClient_GetAxieSummary(t *testing.T) {
	summary, err := testClient.GetAxieSummary(testAddress)
	if err != nil {
		t.Error(err)
		return
	}
	if summary[MysticAxie] == 0 {
		t.Error("expected jihoz to have a mystic axie, got 0")
	}
}

func TestClient_GetLandSummary(t *testing.T) {
	_, err := testClient.GetLandSummary(testAddress)
	if err == nil {
		t.Error("expected error - not implemented, got nil")
	}
}

func TestClient_HasLandFromCollection(t *testing.T) {
	_, err := testClient.HasLandFromCollection(testAddress, MysticLand)
	if err == nil {
		t.Error("expected error - not implemented, got nil")
	}
}
