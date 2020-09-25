package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"saturday/domain"
	"saturday/utils"

	"github.com/gorilla/mux"
)

// SaveCommentHandler - for handle save comment by orgs.
func (oh OrgsHandler) SaveCommentHandler(res http.ResponseWriter, req *http.Request) {
	isValid, err := utils.ValidateToken(req.Header.Get("Authorization-ID"), req.Header.Get("Authorization"), "save-comment")
	if err != nil || !isValid {
		log.Println("SaveCommentHandler: Error matching token")
		writeResponse(res, ErrRespAuthInvalid.Response, ErrRespAuthInvalid.StatusCode, nil)
		return
	}

	orgsName := mux.Vars(req)["orgs"]

	var param domain.OrgsCommentBody
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("SaveCommentHandler: save comment body param empty")
		writeResponse(res, ErrRespBadRequest.Response, ErrRespBadRequest.StatusCode, nil)
		return
	}

	err = json.Unmarshal(reqBody, &param)
	if err != nil {
		log.Println("SaveCommentHandler: can't unmarshall req body, err: " + err.Error())
		writeResponse(res, ErrRespBadRequest.Response, ErrRespBadRequest.StatusCode, nil)
		return
	}

	if param.Comment == "" {
		log.Println("SaveCommentHandler: error comment param empty")
		writeResponse(res, ErrRespBadRequest.Response, ErrRespBadRequest.StatusCode, nil)
		return
	}

	err = oh.Repo.SaveComment(orgsName, param.Comment)
	if err != nil {
		log.Println("SaveCommentHandler: fail save comment to db, err: " + err.Error())
		writeResponse(res, ErrRespInternalServer.Response, ErrRespInternalServer.StatusCode, nil)
		return
	}

	writeResponse(
		res,
		ResponseData{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		http.StatusOK,
		nil)
	return
}

// GetCommentHandler - for handle get comment by orgs.
func (oh OrgsHandler) GetCommentHandler(res http.ResponseWriter, req *http.Request) {
	isValid, err := utils.ValidateToken(req.Header.Get("Authorization-ID"), req.Header.Get("Authorization"), "get-comment")
	if err != nil || !isValid {
		log.Println("GetCommentHandler: Error matching token")
		writeResponse(res, ErrRespAuthInvalid.Response, ErrRespAuthInvalid.StatusCode, nil)
		return
	}

	orgsName := mux.Vars(req)["orgs"]

	comments, err := oh.Repo.GetComment(orgsName)
	if err != nil {
		log.Println("GetCommentHandler: fail get comment from db, err: " + err.Error())
		writeResponse(res, ErrRespInternalServer.Response, ErrRespInternalServer.StatusCode, nil)
		return
	}

	writeResponse(
		res,
		ResponseData{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    comments,
		},
		http.StatusOK,
		nil)
	return
}

// DeleteCommentHandler - for handle delete comment by orgs.
func (oh OrgsHandler) DeleteCommentHandler(res http.ResponseWriter, req *http.Request) {
	isValid, err := utils.ValidateToken(req.Header.Get("Authorization-ID"), req.Header.Get("Authorization"), "delete-comment")
	if err != nil || !isValid {
		log.Println("DeleteCommentHandler: Error matching token")
		writeResponse(res, ErrRespAuthInvalid.Response, ErrRespAuthInvalid.StatusCode, nil)
		return
	}

	orgsName := mux.Vars(req)["orgs"]

	err = oh.Repo.DeleteComment(orgsName)
	if err != nil {
		log.Println("DeleteCommentHandler: fail delete comment from db, err: " + err.Error())
		writeResponse(res, ErrRespInternalServer.Response, ErrRespInternalServer.StatusCode, nil)
		return
	}

	writeResponse(
		res,
		ResponseData{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		http.StatusOK,
		nil)
	return
}

// GetMemberHandler - for handle get member by orgs.
func (oh OrgsHandler) GetMemberHandler(res http.ResponseWriter, req *http.Request) {
	isValid, err := utils.ValidateToken(req.Header.Get("Authorization-ID"), req.Header.Get("Authorization"), "get-member")
	if err != nil || !isValid {
		log.Println("GetMemberHandler: Error matching token")
		writeResponse(res, ErrRespAuthInvalid.Response, ErrRespAuthInvalid.StatusCode, nil)
		return
	}

	orgsName := mux.Vars(req)["orgs"]

	members, err := oh.Repo.GetMember(orgsName)
	if err != nil {
		log.Println("GetMemberHandler: fail get comment from db, err: " + err.Error())
		writeResponse(res, ErrRespInternalServer.Response, ErrRespInternalServer.StatusCode, nil)
		return
	}

	writeResponse(
		res,
		ResponseData{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data:    members,
		},
		http.StatusOK,
		nil)
	return
}
