package sweb

import (
	"encoding/json"
	"net/http"
)

func notFound(res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusNotFound)
	info := Response{"info": "no handler found.", "status": -1}
	b, _ := json.Marshal(info)
	res.Write(b)
}
