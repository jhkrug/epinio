package v1

import (
	"encoding/json"
	"net/http"

	"github.com/suse/carrier/version"
)

type InfoController struct {
}

func (hc InfoController) Info(w http.ResponseWriter, r *http.Request) {
	info := struct {
		Version string
	}{
		Version: version.Version,
	}
	js, err := json.Marshal(info)
	if handleError(w, err, 500) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
