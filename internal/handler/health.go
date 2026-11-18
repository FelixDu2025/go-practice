package handler

import "net/http"

// Health 返回服务存活状态，可用于探活 / 就绪检查练习。
// TODO: 返回 HTTP 200 与 JSON {"status":"ok"}（可复用 writeJSON）。
func Health(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotImplemented, "TODO: implement health")
}
