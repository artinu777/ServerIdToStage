package apiv1

import (
	"github.com/go-chi/chi"
	"github.com/sungora/app/connect"
	"github.com/sungora/app/request"
	"github.com/sungora/sample/internal/stage"
	"net/http"
	"strconv"
)

// handlerPing проверка доступности сервера
// @Tags General
// @Router /server/{id}/stage/{stage_id} [get]
// @Param id path int true "id valid int"
// @Param stage_id path int true "stage_id valid int"
// @Summary проверка доступности сервера
// @Success 200 {string} string
// @Failure 500 {string} string
func Stage1Handler(w http.ResponseWriter, r *http.Request) {
	idFrom := chi.URLParam(r, "id")
	stageId := chi.URLParam(r, "stage_id")

	idToDatabase, err := strconv.ParseInt(idFrom, 10, 64)
	if nil != err {
		request.NewIn(w, r).Json("id dont init")
		return
	}
	StageidToDatabase, err := strconv.ParseInt(stageId, 10, 64)
	if nil != err {
		request.NewIn(w, r).Json("stage dont init")
		return
	}
	server := stage.Stage{
		ServerID: idToDatabase,
		Stage:    StageidToDatabase,
	}

	if err := connect.GetDB().Where("server_id = ?", idFrom).Assign(server).FirstOrCreate(&server); nil != err {
		//request.NewIn(w, r).Json(err)
	}
	request.NewIn(w, r).Json(server)
}
