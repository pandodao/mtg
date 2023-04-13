package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pandodao/mtg/mtgpack"
	"github.com/pandodao/mtg/protocol"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			Memo []byte `json:"memo"`
		}

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "decode body", err)
			return
		}

		d := mtgpack.NewDecoder(body.Memo)

		var h protocol.Header
		if err := mtgpack.DecodeValue(d, &h); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "decode header", err)
			return
		}

		var receiver protocol.MultisigReceiver
		if err := mtgpack.DecodeValue(d, &receiver); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "decode receiver", err)
			return
		}

		fillAssetID, err := d.DecodeUUID()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "decode fill asset id", err)
			return
		}

		min, err := d.DecodeDecimal()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "decode min", err)
			return
		}

		route, err := d.DecodeString()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "decode route", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"header":        h,
			"receiver":      receiver,
			"fill_asset_id": fillAssetID,
			"min":           min,
			"route":         route,
		})
	})

	http.ListenAndServe(":8080", nil)
}
