package discourse

import (
	"encoding/json"
	"enterpret/dataaccess"
	"enterpret/models"
	sourceInterface "enterpret/sources/interface"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	source = "discourse"
)

type discourseFeedbackProcessor struct {
	store  dataaccess.DataStore
	logger *log.Logger
	client http.Client
}

func NewDiscourseFeedbackProcessor(store dataaccess.DataStore) sourceInterface.FeedbackProcessor {
	return &discourseFeedbackProcessor{store: store,
		logger: log.New(os.Stdout, "discourseFeedbackProcessor: ", 1),
		client: http.Client{
			Transport: http.DefaultTransport,
			Timeout:   time.Minute,
		}}
}

func (p *discourseFeedbackProcessor) fetchFeedbacks(params models.Params, page int) (models.DiscourseSearchResponse, error) {

	urlString := fmt.Sprintf("https://meta.discourse.org/search.json?page=%d&q=%s", page, params.SearchQuery)

	if params.SearchQuery != "" {
		urlString += "+"
	}

	if params.Before != nil && params.Since != nil {
		yearB, monthB, dayB := params.Before.Date()
		yearA, monthA, dayA := params.Since.Date()
		urlString += "after%3A" + fmt.Sprintf("%d-%d-%d", yearA, monthA, dayA) +
			"+before%3A" + fmt.Sprintf("%d-%d-%d", yearB, monthB, dayB)
	} else if params.Before != nil {
		year, month, day := params.Before.Date()
		urlString += "before%3A" + fmt.Sprintf("%d-%d-%d", year, month, day)
	} else if params.Since != nil {
		year, month, day := params.Since.Date()
		urlString += "after%3A" + fmt.Sprintf("%d-%d-%d", year, month, day)
	}

	parse, err := url.Parse(urlString)
	if err != nil {
		p.logger.Println("Error: ", err.Error())
		return models.DiscourseSearchResponse{}, err
	}

	do, err := p.client.Do(&http.Request{
		Method: http.MethodGet,
		URL:    parse,
	})
	if err != nil {
		return models.DiscourseSearchResponse{}, err
	}

	var searchResp = models.DiscourseSearchResponse{}
	err = json.NewDecoder(do.Body).Decode(&searchResp)
	if err != nil {
		p.logger.Println(err.Error())
		return models.DiscourseSearchResponse{}, err
	}

	return searchResp, nil
}

func (p *discourseFeedbackProcessor) fetchAndStoreIndividualFeedback(id int, topicId int, tenant string) (models.FeedbackIngest, error) {

	urlString := "https://meta.discourse.org/t/:topic_id/posts.json?post_ids%5B%5D=:post_id"
	urlString = strings.ReplaceAll(urlString, ":topic_id", fmt.Sprint(topicId))
	urlString = strings.ReplaceAll(urlString, ":post_id", fmt.Sprint(id))

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, urlString, nil)
	if err != nil {
		p.logger.Printf("Error: error creating request for post %d in topic %d: "+err.Error(), id, topicId)
		return models.FeedbackIngest{}, err
	}

	res, err := client.Do(req)
	if err != nil {
		p.logger.Println("Error: ", err.Error())
		return models.FeedbackIngest{}, err
	}

	defer func(Body io.ReadCloser) {
		errClose := Body.Close()
		if errClose != nil {
			p.logger.Println("Error: failed closing the body with the following error: ", errClose.Error())
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		p.logger.Println("Error: Failed reading the body with the following error: ", err.Error())
		return models.FeedbackIngest{}, err
	}
	return p.IngestAndStoreFeedback(body, tenant)
}

//FetchAndStoreFeedbacks fetches feedback from Discourse, processes it and stores them
func (p *discourseFeedbackProcessor) FetchAndStoreFeedbacks(params models.Params, tenant string) ([]models.FeedbackIngest, error) {
	page := 1
	var messages []models.FeedbackIngest

	for true {
		var searchResp models.DiscourseSearchResponse
		var err error
		searchResp, err = p.fetchFeedbacks(params, page)
		if err != nil {
			return nil, err
		}

		if len(searchResp.Posts) == 0 {
			return messages, nil
		}

		for _, post := range searchResp.Posts {
			mI, err := p.fetchAndStoreIndividualFeedback(post.Id, post.TopicId, tenant)
			if err != nil {
				p.logger.Println("Error: ", err.Error())
			}

			messages = append(messages, mI)
		}
		page++
	}

	return messages, nil
}

//IngestAndStoreFeedback processes the feedback and stores it
func (p *discourseFeedbackProcessor) IngestAndStoreFeedback(blob []byte, tenant string) (models.FeedbackIngest, error) {
	var postResp = models.DiscoursePostResponse{}
	err := json.Unmarshal(blob, &postResp)
	if err != nil {
		p.logger.Println(err.Error())
		return models.FeedbackIngest{}, err
	}

	for _, post := range postResp.PostStream.Posts {
		mI := models.FeedbackIngest{
			Meta: models.Meta{
				Tenant:       tenant,
				Source:       source,
				CreationTime: post.CreatedAt,
				User:         post.Username,
				ID:           fmt.Sprint(post.Id),
				Language:     "English",
				Attributes: map[string]interface{}{
					"id":                      post.Id,
					"name":                    post.Name,
					"username":                post.Username,
					"avatar_template":         post.AvatarTemplate,
					"created_at":              post.CreatedAt,
					"post_number":             post.PostNumber,
					"post_type":               post.PostType,
					"updated_at":              post.UpdatedAt,
					"reply_count":             post.ReplyCount,
					"reply_to_post_number":    post.ReplyToPostNumber,
					"quote_count":             post.QuoteCount,
					"incoming_link_count":     post.IncomingLinkCount,
					"reads":                   post.Reads,
					"readers_count":           post.ReadersCount,
					"score":                   post.Score,
					"yours":                   post.Yours,
					"topic_id":                post.TopicId,
					"topic_slug":              post.TopicSlug,
					"display_username":        post.DisplayUsername,
					"primary_group_name":      post.PrimaryGroupName,
					"flair_name":              post.FlairName,
					"flair_url":               post.FlairUrl,
					"flair_bg_color":          post.FlairBgColor,
					"flair_color":             post.FlairColor,
					"version":                 post.Version,
					"can_edit":                post.CanEdit,
					"can_delete":              post.CanDelete,
					"can_recover":             post.CanRecover,
					"can_wiki":                post.CanWiki,
					"read":                    post.Read,
					"user_title":              post.UserTitle,
					"bookmarked":              post.Bookmarked,
					"moderator":               post.Moderator,
					"admin":                   post.Admin,
					"staff":                   post.Staff,
					"user_id":                 post.UserId,
					"hidden":                  post.Hidden,
					"trust_level":             post.TrustLevel,
					"deleted_at":              post.DeletedAt,
					"user_deleted":            post.UserDeleted,
					"edit_reason":             post.EditReason,
					"can_view_edit_history":   post.CanViewEditHistory,
					"wiki":                    post.Wiki,
					"customer_flair_customer": post.CustomerFlairCustomer,
					"can_accept_answer":       post.CanAcceptAnswer,
					"can_unaccept_answer":     post.CanUnacceptAnswer,
					"accepted_answer":         post.AcceptedAnswer,
				},
			},
			Data: models.Data{
				Message: post.Cooked,
			},
		}

		err := p.store.Store(mI)
		if err != nil {
			p.logger.Println("Error: ", err.Error())
		}
	}

	return models.FeedbackIngest{}, nil
}
