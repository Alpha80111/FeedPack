package server

import (
	"enterpret/dataaccess"
	"enterpret/models"
	mock2 "enterpret/sources/mock"
	mock3 "enterpret/sources/sourceinterface/mock"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestNewServer(t *testing.T) {
	go func() {
		ctrl := gomock.NewController(t)
		mSP := mock2.NewMockSourceProcessor(ctrl)
		fP := mock3.NewMockFeedbackProcessor(ctrl)
		fP.EXPECT().FetchAndStoreFeedbacks(gomock.Any(), gomock.Any()).Return([]models.FeedbackIngest{}, nil).AnyTimes()
		fP.EXPECT().IngestAndStoreFeedback(gomock.Any(), gomock.Any()).Return(models.FeedbackIngest{}, nil).AnyTimes()
		mSP.EXPECT().GetProcessor(gomock.Any()).Return(fP, nil).AnyTimes()
		err := NewServer(dataaccess.NewDataStore(), mSP)
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	url := "http://localhost:8088/push/feedback"
	payload := strings.NewReader(`{
    "tenant": "zoom.us",
    "source": "discourse",
    "post_stream": {
        "posts": [
            {
                "id": 844350,
                "name": "",
                "username": "bekircem",
                "avatar_template": "/user_avatar/meta.discourse.org/bekircem/{size}/44582_2.png",
                "created_at": "2020-11-09T18:05:21.274Z",
                "cooked": "<p>Is it possible to add some basic dropdown menu to an any item?</p>\n<p>I couldn’t create a dropdown menu with “Custom header links”. It seems <a href=\"https://devforum.zoom.us/\" rel=\"noopener nofollow ugc\">Zoom</a> did that. I reviewed their dropdown menu via console, but I couldn’t figure out how they interfere with the html of this component for adding dropdown to any item.</p>\n<p><div class=\"lightbox-wrapper\"><a class=\"lightbox\" href=\"https://d11a6trkgmumsb.cloudfront.net/original/3X/6/5/656b022dc6ddb150c1c72d63ee88b70f78fc7980.png\" data-download-href=\"/uploads/short-url/etbukslmgBvGf4dH09n0jcOJ02c.png?dl=1\" title=\"image\" rel=\"noopener nofollow ugc\"><img src=\"https://d11a6trkgmumsb.cloudfront.net/optimized/3X/6/5/656b022dc6ddb150c1c72d63ee88b70f78fc7980_2_690x413.png\" alt=\"image\" data-base62-sha1=\"etbukslmgBvGf4dH09n0jcOJ02c\" width=\"690\" height=\"413\" srcset=\"https://d11a6trkgmumsb.cloudfront.net/optimized/3X/6/5/656b022dc6ddb150c1c72d63ee88b70f78fc7980_2_690x413.png, https://d11a6trkgmumsb.cloudfront.net/original/3X/6/5/656b022dc6ddb150c1c72d63ee88b70f78fc7980.png 1.5x, https://d11a6trkgmumsb.cloudfront.net/original/3X/6/5/656b022dc6ddb150c1c72d63ee88b70f78fc7980.png 2x\" data-small-upload=\"https://d11a6trkgmumsb.cloudfront.net/optimized/3X/6/5/656b022dc6ddb150c1c72d63ee88b70f78fc7980_2_10x10.png\"><div class=\"meta\"><svg class=\"fa d-icon d-icon-far-image svg-icon\" aria-hidden=\"true\"><use xlink:href=\"#far-image\"></use></svg><span class=\"filename\">image</span><span class=\"informations\">894×536 37.8 KB</span><svg class=\"fa d-icon d-icon-discourse-expand svg-icon\" aria-hidden=\"true\"><use xlink:href=\"#discourse-expand\"></use></svg></div></a></div></p>\n<p>Is there a way to add this dropdown to item? <a class=\"mention\" href=\"/u/johani\">@Johani</a></p>\n<pre><code>&lt;div id=\"dropdown\"&gt;\n  &lt;a title=\"Zoom Developer Documentation\" href=\"https://marketplace.zoom.us/docs\" target=\"_blank\"&gt;Developer&lt;/a&gt;\n  &lt;span class=\"caret\"&gt;&lt;/span&gt;\n  &lt;div class=\"dropdown-content\"&gt;\n    &lt;a title=\"Zoom API Docs\" href=\"https://marketplace.zoom.us/docs/api-reference/introduction\" target=\"_blank\"&gt;API&lt;/a&gt;\n    &lt;a title=\"Zoom SDK Docs\" href=\"https://marketplace.zoom.us/docs/sdk/native-sdks/introduction\" target=\"_blank\"&gt;SDK&lt;/a&gt;\n    &lt;a title=\"Zoom Developer Blog\" href=\"https://medium.com/zoom-developer-blog\" target=\"_blank\"&gt;Blog&lt;/a&gt;\n    &lt;a title=\"Zoom Developer Changelog\" href=\"https://marketplace.zoom.us/docs/changelog\" target=\"_blank\"&gt;Changelog&lt;/a&gt;\n    &lt;a title=\"Zoom Developer Survey\" href=\"https://docs.google.com/forms/d/e/1FAIpQLSeJPLhNuxjtkxyyV276R8S_nYz99fpMbbS8VWkC8Hwi7-2Byg/viewform\" target=\"_blank\"&gt;Survey&lt;/a&gt;\n  &lt;/div&gt;\n&lt;/div&gt;\n</code></pre>",
                "post_number": 106,
                "post_type": 1,
                "updated_at": "2020-11-16T13:57:39.065Z",
                "reply_count": 2,
                "reply_to_post_number": null,
                "quote_count": 0,
                "incoming_link_count": 2,
                "reads": 125,
                "readers_count": 124,
                "score": 105.0,
                "yours": false,
                "topic_id": 90588,
                "topic_slug": "custom-header-links",
                "display_username": "",
                "primary_group_name": null,
                "flair_name": null,
                "flair_url": null,
                "flair_bg_color": null,
                "flair_color": null,
                "version": 4,
                "can_edit": false,
                "can_delete": false,
                "can_recover": false,
                "can_wiki": false,
                "link_counts": [
                    {
                        "url": "https://devforum.zoom.us/",
                        "internal": false,
                        "reflection": false,
                        "title": "Zoom Developer Forum",
                        "clicks": 4
                    },
                    {
                        "url": "https://d11a6trkgmumsb.cloudfront.net/original/3X/6/5/656b022dc6ddb150c1c72d63ee88b70f78fc7980.png",
                        "internal": false,
                        "reflection": false,
                        "title": "656b022dc6ddb150c1c72d63ee88b70f78fc7980.png",
                        "clicks": 0
                    }
                ],
                "read": true,
                "user_title": null,
                "bookmarked": false,
                "actions_summary": [
                    {
                        "id": 2,
                        "count": 4
                    }
                ],
                "moderator": false,
                "admin": false,
                "staff": false,
                "user_id": 16290,
                "hidden": false,
                "trust_level": 2,
                "deleted_at": null,
                "user_deleted": false,
                "edit_reason": null,
                "can_view_edit_history": true,
                "wiki": false,
                "customer_flair_customer": null,
                "can_accept_answer": false,
                "can_unaccept_answer": false,
                "accepted_answer": false
            },
             {
                "id": 844352,
                "name": "",
                "username": "bekircem",
                "avatar_template": "/user_avatar/meta.discourse.org/bekircem/{size}/44582_2.png",
                "created_at": "2020-11-09T18:05:21.274Z",
                "cooked": "<p>Is it possible to add some basic dropdown menu to an any item?</p>\n<p>I couldn’t create a dropdown menu with “Custom header links”. It seems <a href=\"https://devforum.zoom.us/\" rel=\"noopener nofollow ugc\">Zoom</a> did that. I reviewed their dropdown menu via console, but I couldn’t figure out how they interfere with the html of this component for adding dropdown to any item.</p>\n<p><div class=\"lightbox-wrapper\"><a class=\"lightbox\" href=\"https://d11a6trkgmumsb.cloudfront.net/original/3X/6/5/656b022dc6ddb150c1c72d63ee88b70f78fc7980.png\" data-download-href=\"/uploads/short-url/etbukslmgBvGf4dH09n0jcOJ02c.png?dl=1\" title=\"image\" rel=\"noopener nofollow ugc\"><img src=\"https://d11a6trkgmumsb.cloudfront.net/optimized/3X/6/5/656b022dc6ddb150c1c72d63ee88b70f78fc7980_2_690x413.png\" alt=\"image\" data-base62-sha1=\"etbukslmgBvGf4dH09n0jcOJ02c\" width=\"690\" height=\"413\" srcset=\"https://d11a6trkgmumsb.cloudfront.net/optimized/3X/6/5/656b022dc6ddb150c1c72d63ee88b70f78fc7980_2_690x413.png, https://d11a6trkgmumsb.cloudfront.net/original/3X/6/5/656b022dc6ddb150c1c72d63ee88b70f78fc7980.png 1.5x, https://d11a6trkgmumsb.cloudfront.net/original/3X/6/5/656b022dc6ddb150c1c72d63ee88b70f78fc7980.png 2x\" data-small-upload=\"https://d11a6trkgmumsb.cloudfront.net/optimized/3X/6/5/656b022dc6ddb150c1c72d63ee88b70f78fc7980_2_10x10.png\"><div class=\"meta\"><svg class=\"fa d-icon d-icon-far-image svg-icon\" aria-hidden=\"true\"><use xlink:href=\"#far-image\"></use></svg><span class=\"filename\">image</span><span class=\"informations\">894×536 37.8 KB</span><svg class=\"fa d-icon d-icon-discourse-expand svg-icon\" aria-hidden=\"true\"><use xlink:href=\"#discourse-expand\"></use></svg></div></a></div></p>\n<p>Is there a way to add this dropdown to item? <a class=\"mention\" href=\"/u/johani\">@Johani</a></p>\n<pre><code>&lt;div id=\"dropdown\"&gt;\n  &lt;a title=\"Zoom Developer Documentation\" href=\"https://marketplace.zoom.us/docs\" target=\"_blank\"&gt;Developer&lt;/a&gt;\n  &lt;span class=\"caret\"&gt;&lt;/span&gt;\n  &lt;div class=\"dropdown-content\"&gt;\n    &lt;a title=\"Zoom API Docs\" href=\"https://marketplace.zoom.us/docs/api-reference/introduction\" target=\"_blank\"&gt;API&lt;/a&gt;\n    &lt;a title=\"Zoom SDK Docs\" href=\"https://marketplace.zoom.us/docs/sdk/native-sdks/introduction\" target=\"_blank\"&gt;SDK&lt;/a&gt;\n    &lt;a title=\"Zoom Developer Blog\" href=\"https://medium.com/zoom-developer-blog\" target=\"_blank\"&gt;Blog&lt;/a&gt;\n    &lt;a title=\"Zoom Developer Changelog\" href=\"https://marketplace.zoom.us/docs/changelog\" target=\"_blank\"&gt;Changelog&lt;/a&gt;\n    &lt;a title=\"Zoom Developer Survey\" href=\"https://docs.google.com/forms/d/e/1FAIpQLSeJPLhNuxjtkxyyV276R8S_nYz99fpMbbS8VWkC8Hwi7-2Byg/viewform\" target=\"_blank\"&gt;Survey&lt;/a&gt;\n  &lt;/div&gt;\n&lt;/div&gt;\n</code></pre>",
                "post_number": 106,
                "post_type": 1,
                "updated_at": "2020-11-16T13:57:39.065Z",
                "reply_count": 2,
                "reply_to_post_number": null,
                "quote_count": 0,
                "incoming_link_count": 2,
                "reads": 125,
                "readers_count": 124,
                "score": 105.0,
                "yours": false,
                "topic_id": 90588,
                "topic_slug": "custom-header-links",
                "display_username": "",
                "primary_group_name": null,
                "flair_name": null,
                "flair_url": null,
                "flair_bg_color": null,
                "flair_color": null,
                "version": 4,
                "can_edit": false,
                "can_delete": false,
                "can_recover": false,
                "can_wiki": false,
                "link_counts": [
                    {
                        "url": "https://devforum.zoom.us/",
                        "internal": false,
                        "reflection": false,
                        "title": "Zoom Developer Forum",
                        "clicks": 4
                    },
                    {
                        "url": "https://d11a6trkgmumsb.cloudfront.net/original/3X/6/5/656b022dc6ddb150c1c72d63ee88b70f78fc7980.png",
                        "internal": false,
                        "reflection": false,
                        "title": "656b022dc6ddb150c1c72d63ee88b70f78fc7980.png",
                        "clicks": 0
                    }
                ],
                "read": true,
                "user_title": null,
                "bookmarked": false,
                "actions_summary": [
                    {
                        "id": 2,
                        "count": 4
                    }
                ],
                "moderator": false,
                "admin": false,
                "staff": false,
                "user_id": 16290,
                "hidden": false,
                "trust_level": 2,
                "deleted_at": null,
                "user_deleted": false,
                "edit_reason": null,
                "can_view_edit_history": true,
                "wiki": false,
                "customer_flair_customer": null,
                "can_accept_answer": false,
                "can_unaccept_answer": false,
                "accepted_answer": false
            }
        ]
    },
    "id": 90588
}`)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, payload)
	assert.Nil(t, err)

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	assert.Nil(t, err)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, string(body), "Successfully processed message")

	url = "http://localhost:8088/fetch/feedbacks"
	payload = strings.NewReader(`{
    "sources": [
        "discourse"
    ],
    "tenant": "zoom.us",
    "page": 1,
    "size": 50
}`)

	req, err = http.NewRequest(http.MethodPost, url, payload)
	assert.Nil(t, err)

	req.Header.Add("Content-Type", "application/json")

	res, err = client.Do(req)
	assert.Nil(t, err)

	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	assert.Nil(t, err)

	url = "http://localhost:8088/pull/feedback"
	payload = strings.NewReader(`{
    "source": "discourse",
    "tenant": "test",
    "params": {
        "searchQuery": "test"
    }
}`)

	req, err = http.NewRequest(http.MethodPost, url, payload)
	assert.Nil(t, err)

	req.Header.Add("Content-Type", "application/json")

	res, err = client.Do(req)
	assert.Nil(t, err)

	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
}
