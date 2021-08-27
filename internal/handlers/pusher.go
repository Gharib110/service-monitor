package handlers

import (
	"github.com/pusher/pusher-http-go"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (repo *DBRepo) PusherAuth(w http.ResponseWriter, r *http.Request) {
	userID := repo.App.Session.GetInt(r.Context(), "userID")

	u, _ := repo.DB.GetUserById(userID)
	params, _ := ioutil.ReadAll(r.Body)
	presenceData := pusher.MemberData{
		UserID: strconv.Itoa(userID),
		UserInfo: map[string]string{
			"name": u.FirstName,
			"id":   strconv.Itoa(userID),
		},
	}

	resp, err := app.WsClient.AuthenticatePresenceChannel(params, presenceData)
	if err != nil {
		log.Error().Msg(err.Error() + "; in authenticating presence channel")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(resp)
	return
}
