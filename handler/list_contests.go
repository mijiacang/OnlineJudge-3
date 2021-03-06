package handler

import (
	"OnlineJudge/pbgen/api"
)

func (this *Handler) ListContests(response *api.ListContestsResponse, req *api.ListContestsRequest) {
	if err := this.OpenDB(); err != nil {
		MakeResponseError(response, this.debug, PBInternalError, err)
		return
	}
	defer this.CloseDB()
}
