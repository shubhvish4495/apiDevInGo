package jsonAPI

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"

	"github.com/shubhvish4495/apiDevInGo/pkg/helper"
)

//UserDataHandler returns user data with his related posts
//and all his comments
func UserDataHandler(rw http.ResponseWriter, r *http.Request) {

	userID, _ := strconv.Atoi(mux.Vars(r)["userID"])
	var userDataResp UserData

	//if userID not found return
	if userID == 0 {
		helper.ReturnErrResponse(errors.New("userId param is required"), rw)
		return
	}

	// fetch user data
	if userData, err := getUserData(userID); err != nil {
		userDataResp.UserData.Name = userData.Name
		userDataResp.UserData.ID = userData.ID
		userDataResp.UserData.Email = userData.Email
	} else {
		helper.ReturnErrResponse(err, rw)
	}

	//fetch all the posts
	postList, err := listPosts()
	if err != nil {
		helper.ReturnErrResponse(err, rw)
		return
	}

	relatedPosts := make([]Post, 0)

	//loop over post to fetch all the related posts
	for _, postData := range postList {
		if postData.UserID == userID {
			relatedPosts = append(relatedPosts, postData)
		}
	}

	//list all comments and then use go routine to fetch
	//each comment. Just wanted to use and practice go routines
	//in a rest repository
	allComments, err := listComments()
	if err != nil {
		helper.ReturnErrResponse(err, rw)
	}

	relatedCommentIDList := make([]int, 0)

	//range over all comments to load realted comment IDs
	// to fetch their data from another call
	for _, comment := range allComments {
		if validComment(relatedPosts, comment) {
			relatedCommentIDList = append(relatedCommentIDList, comment.ID)
		}
	}

	//create a channel to store data
	commChan := make(chan *Comment)
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for _, commentID := range relatedCommentIDList {
		wg.Add(1)
		go func(commentID int) {
			defer wg.Done()
			commentData, err := getComment(commentID)
			if err != nil {
				log.Println("Error while fetching comments ", err.Error())
			}
			commChan <- commentData
		}(commentID)
	}

	channelOutComments := make([]Comment, 0)

	go func() {
		wg.Wait()
		close(commChan)
	}()

	for data := range commChan {
		mutex.Lock()
		channelOutComments = append(channelOutComments, *data)
		mutex.Unlock()
	}

	helper.ReturnJSONResponse(rw, channelOutComments)

}

func validComment(allPostsForUser []Post, comment Comment) bool {

	for _, post := range allPostsForUser {
		if post.ID == comment.ID {
			return true
		}
	}

	return false
}
