package discourse

import (
	"enterpret/dataaccess/mock"
	"enterpret/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestDiscourseFeedbackProcessor_FetchAndStoreFeedbacks(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDS := mock.NewMockDataStore(ctrl)

	mockDS.EXPECT().Store(gomock.Any()).Return(nil).AnyTimes()

	dFP := NewDiscourseFeedbackProcessor(mockDS)

	s1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{
    "posts": [
        {
            "id": 844350,
            "name": "",
            "username": "bekircem",
            "avatar_template": "/user_avatar/meta.discourse.org/bekircem/{size}/44582_2.png",
            "created_at": "2020-11-09T18:05:21.274Z",
            "like_count": 4,
            "blurb": "Is it possible to add some basic dropdown menu to an any item? I couldn’t create a dropdown menu with “Custom header links”. It seems https://devforum.zoom.us/ Zoom did that. I reviewed their dropdown...",
            "post_number": 106,
            "topic_title_headline": "Custom Header Links",
            "topic_id": 90588
        },
        {
            "id": 720158,
            "name": "David Kingham",
            "username": "davidkingham",
            "avatar_template": "/user_avatar/meta.discourse.org/davidkingham/{size}/119528_2.png",
            "created_at": "2020-03-24T01:55:57.390Z",
            "like_count": 0,
            "blurb": "pmusaraj: Did you complete the event subscription steps? That’s a webhook from the Zoom API, and it’s used to switch the Registered button to a “join now” button as soon as an event starts. I did all ...",
            "post_number": 8,
            "topic_title_headline": "Zoom Webinars Plugin",
            "topic_id": 142711
        },
        {
            "id": 1057091,
            "name": "Angus McLeod",
            "username": "angus",
            "avatar_template": "/user_avatar/meta.discourse.org/angus/{size}/247364_2.png",
            "created_at": "2022-02-09T23:33:25.111Z",
            "like_count": 9,
            "blurb": "...that will allow online communities to easily transfer calendar event data between popular event management platforms (initially http://eventbrite.com Eventbrite , http://meetup.com Meetup and http://z...",
            "post_number": 2,
            "topic_title_headline": "Calendar plugin features to make it really useful for us",
            "topic_id": 210333
        },
        {
            "id": 847368,
            "name": "",
            "username": "bekircem",
            "avatar_template": "/user_avatar/meta.discourse.org/bekircem/{size}/44582_2.png",
            "created_at": "2020-11-16T13:38:14.414Z",
            "like_count": 2,
            "blurb": "Mark_Schmucker: on’t need dropdowns- less drama and feels better integrated. I don’t know the answer to your actual question, but I’m I am currently using this plugin but since there is not enough spa...",
            "post_number": 107,
            "topic_title_headline": "Header submenus",
            "topic_id": 94584
        }
    ],
    "topics": [
        {
            "id": 90588,
            "title": "Custom Header Links",
            "fancy_title": "Custom Header Links",
            "slug": "custom-header-links",
            "posts_count": 104,
            "reply_count": 78,
            "highest_post_number": 140,
            "created_at": "2018-06-24T10:22:32.915Z",
            "last_posted_at": "2022-04-06T00:19:38.129Z",
            "bumped": true,
            "bumped_at": "2022-04-06T00:19:38.129Z",
            "archetype": "regular",
            "unseen": false,
            "pinned": false,
            "unpinned": null,
            "visible": true,
            "closed": false,
            "archived": false,
            "bookmarked": null,
            "liked": null,
            "tags": [
                "theme-component"
            ],
            "tags_descriptions": {},
            "category_id": 61,
            "has_accepted_answer": false
        },
        {
            "id": 142711,
            "title": "Zoom Webinars Plugin",
            "fancy_title": "Zoom Webinars Plugin",
            "slug": "zoom-webinars-plugin",
            "posts_count": 65,
            "reply_count": 38,
            "highest_post_number": 77,
            "created_at": "2020-02-26T19:18:08.506Z",
            "last_posted_at": "2021-09-17T18:28:43.618Z",
            "bumped": true,
            "bumped_at": "2021-09-17T18:28:43.618Z",
            "archetype": "regular",
            "unseen": false,
            "pinned": false,
            "unpinned": null,
            "visible": true,
            "closed": false,
            "archived": false,
            "bookmarked": null,
            "liked": null,
            "tags": [],
            "tags_descriptions": {},
            "category_id": 22,
            "has_accepted_answer": false
        },
        {
            "id": 210333,
            "title": "Calendar plugin features to make it really useful for us",
            "fancy_title": "Calendar plugin features to make it really useful for us",
            "slug": "calendar-plugin-features-to-make-it-really-useful-for-us",
            "posts_count": 3,
            "reply_count": 1,
            "highest_post_number": 3,
            "created_at": "2021-11-28T17:45:49.543Z",
            "last_posted_at": "2022-02-10T05:46:31.979Z",
            "bumped": true,
            "bumped_at": "2022-02-10T05:46:31.979Z",
            "archetype": "regular",
            "unseen": false,
            "pinned": false,
            "unpinned": null,
            "visible": true,
            "closed": false,
            "archived": false,
            "bookmarked": null,
            "liked": null,
            "tags": [],
            "tags_descriptions": {},
            "category_id": 2,
            "has_accepted_answer": false
        },
        {
            "id": 94584,
            "title": "Header submenus",
            "fancy_title": "Header submenus",
            "slug": "header-submenus",
            "posts_count": 54,
            "reply_count": 52,
            "highest_post_number": 125,
            "created_at": "2018-08-13T09:50:07.691Z",
            "last_posted_at": "2022-04-25T07:23:10.235Z",
            "bumped": true,
            "bumped_at": "2022-04-25T07:23:10.235Z",
            "archetype": "regular",
            "unseen": false,
            "pinned": false,
            "unpinned": null,
            "visible": true,
            "closed": false,
            "archived": false,
            "bookmarked": null,
            "liked": null,
            "tags": [
                "theme-component"
            ],
            "tags_descriptions": {},
            "category_id": 61,
            "has_accepted_answer": false
        }
    ],
    "users": [],
    "categories": [],
    "tags": [],
    "groups": [],
    "grouped_search_result": {
        "more_posts": null,
        "more_users": null,
        "more_categories": null,
        "term": "zoom.us after:2010-01-01 before:2023-01-01'",
        "search_log_id": 1516455,
        "more_full_page_results": null,
        "can_create_topic": false,
        "error": null,
        "post_ids": [
            844350,
            720158,
            1057091,
            847368
        ],
        "user_ids": [],
        "category_ids": [],
        "tag_ids": [],
        "group_ids": []
    }
}`))
	}))
	defer s1.Close()

	dFP.(*discourseFeedbackProcessor).feedbackSearchUrl = s1.URL + "?/%s/%s"

	s2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{
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
            }
        ]
    },
    "id": 90588
}`))
	}))
	defer s2.Close()

	dFP.(*discourseFeedbackProcessor).feedbackFetchUrl = s2.URL

	time.Sleep(5 * time.Second)

	//Ideally should test by mocking a discourse server
	_, err := dFP.FetchAndStoreFeedbacks(models.Params{
		Since:                    nil,
		Before:                   nil,
		SearchQuery:              "search-query",
		SourceSpecificParameters: nil,
	}, "tenant")

	assert.Nil(t, err)
}

func TestDiscourseFeedbackProcessor_IngestAndStoreFeedback(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDS := mock.NewMockDataStore(ctrl)

	mockDS.EXPECT().Store(gomock.Any()).Return(nil).AnyTimes()

	dFP := NewDiscourseFeedbackProcessor(mockDS)

	_, err := dFP.IngestAndStoreFeedback([]byte(`{
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
}`), "tenant")

	assert.Nil(t, err)
}

func TestNewDiscourseFeedbackProcessor(t *testing.T) {
	NewDiscourseFeedbackProcessor(nil)
}
