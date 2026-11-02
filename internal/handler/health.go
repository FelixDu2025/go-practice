package handler

import "net/http"

// Health 返回服务存活状态，可用于探活 / 就绪检查练习。
func Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
