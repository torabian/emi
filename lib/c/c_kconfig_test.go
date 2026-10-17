package c

import (
	"strings"
	"testing"

	"github.com/torabian/emi/lib/core"
)

func TestCKconfigVirtualFile(t *testing.T) {
	module := &core.Emi{
		Name: "tinkerlane",
		Config: []core.EmiConfig{
			{Name: "sqlite", Type: "bool", Default: "true", Description: "Store data in SQLite\nOff: no database.\n\nSaves flash."},
			{Name: "listenPort", Type: "int", Default: "8000", Min: "1", Max: "65535", Description: `HTTP "port"`},
			{Name: "keepReadings", Type: "int", Default: "500", Min: "1", Max: "100", Depends: "sqlite && !SOC_X"},
			{Name: "apSsid", Default: `Tinker"lane`, Env: "TINKERLANE_AP_SSID"},
			{Name: "wifi", Description: "Wi-Fi", Fields: []core.EmiConfig{{Name: "ssid"}}},
		},
	}
	f, err := CKconfigVirtualFile(module)
	if err != nil {
		t.Fatal(err)
	}
	if f.Name+f.Extension != "Kconfig.tinkerlane" {
		t.Fatalf("file name %q", f.Name+f.Extension)
	}
	for _, want := range []string{
		"config TINKERLANE_SQLITE\n    bool \"Store data in SQLite\"\n    default y\n",
		"    help\n        Off: no database.\n\n        Saves flash.\n",
		"config TINKERLANE_LISTEN_PORT\n    int \"HTTP \\\"port\\\"\"\n    default 8000\n    range 1 65535\n",
		"    depends on TINKERLANE_SQLITE && !SOC_X\n",
		"config TINKERLANE_AP_SSID\n    string \"apSsid\"\n    default \"Tinker\\\"lane\"\n",
		"menu \"Wi-Fi\"\n\nconfig TINKERLANE_WIFI_SSID\n",
		"endmenu\n",
	} {
		if !strings.Contains(f.ActualScript, want) {
			t.Errorf("missing %q in:\n%s", want, f.ActualScript)
		}
	}

	if _, err := CKconfigVirtualFile(&core.Emi{Name: "x", Config: []core.EmiConfig{{Name: "f", Type: "float64"}}}); err == nil {
		t.Error("float must be rejected")
	}
	if _, err := CKconfigVirtualFile(&core.Emi{Name: "x", Config: []core.EmiConfig{{Name: "f", Type: "int", Min: "1"}}}); err == nil {
		t.Error("min without max must be rejected")
	}
}
