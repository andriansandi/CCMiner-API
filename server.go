package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

type Hashrate struct {
	Value float64 `json:"hashrate"`
	Unit  string  `json:"unit"`
}

func getHashrate() Hashrate {
	// Corrected escaping for grep regex & removed ANSI color codes
	cmd := `grep -E '([0-9]+\.?[0-9]*)\s?(kH/s|KH/s|MH/s)' /tmp/ccminer.log | tail -n 1 | sed 's/\x1B\[[0-9;]*[mK]//g' | awk '{for(i=1;i<=NF;i++) if ($i ~ /^[0-9]+\.?[0-9]*$/ && $(i+1) ~ /(kH\/s|KH\/s|MH\/s)/) print $i, $(i+1)}'`
	out, err := exec.Command("bash", "-c", cmd).Output()

	if err != nil || len(out) == 0 {
		return Hashrate{Value: 0, Unit: "N/A"}
	}

	cleanOutput := strings.TrimSpace(string(out)) // Remove extra spaces/newlines
	parts := strings.Fields(cleanOutput)         // Split into words

	if len(parts) != 2 {
		return Hashrate{Value: 0, Unit: "N/A"}
	}

	var value float64
	fmt.Sscanf(parts[0], "%f", &value)

	return Hashrate{Value: value, Unit: parts[1]}
}

func hashrateHandler(w http.ResponseWriter, r *http.Request) {
	hashrate := getHashrate()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hashrate)
}

func main() {
	http.HandleFunc("/hashrate", hashrateHandler)
	fmt.Println("Server running on port 5000...")
	http.ListenAndServe(":5000", nil)
}
