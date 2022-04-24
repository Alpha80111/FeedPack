package models

import "time"

type DiscourseSearchResponse struct {
	Posts []DiscoursePost `json:"posts"`
}

type DiscoursePost struct {
	Id                 int       `json:"id"`
	Name               string    `json:"name"`
	Username           string    `json:"username"`
	AvatarTemplate     string    `json:"avatar_template"`
	CreatedAt          time.Time `json:"created_at"`
	LikeCount          int       `json:"like_count"`
	Blurb              string    `json:"blurb"`
	PostNumber         int       `json:"post_number"`
	TopicTitleHeadline string    `json:"topic_title_headline"`
	TopicId            int       `json:"topic_id"`
}

type T struct {
	Posts []struct {
		Id                 int       `json:"id"`
		Name               *string   `json:"name"`
		Username           string    `json:"username"`
		AvatarTemplate     string    `json:"avatar_template"`
		CreatedAt          time.Time `json:"created_at"`
		LikeCount          int       `json:"like_count"`
		Blurb              string    `json:"blurb"`
		PostNumber         int       `json:"post_number"`
		TopicTitleHeadline string    `json:"topic_title_headline"`
		TopicId            int       `json:"topic_id"`
	} `json:"posts"`
	Topics []struct {
		Id                int         `json:"id"`
		Title             string      `json:"title"`
		FancyTitle        string      `json:"fancy_title"`
		Slug              string      `json:"slug"`
		PostsCount        int         `json:"posts_count"`
		ReplyCount        int         `json:"reply_count"`
		HighestPostNumber int         `json:"highest_post_number"`
		CreatedAt         time.Time   `json:"created_at"`
		LastPostedAt      time.Time   `json:"last_posted_at"`
		Bumped            bool        `json:"bumped"`
		BumpedAt          time.Time   `json:"bumped_at"`
		Archetype         string      `json:"archetype"`
		Unseen            bool        `json:"unseen"`
		Pinned            bool        `json:"pinned"`
		Unpinned          interface{} `json:"unpinned"`
		Visible           bool        `json:"visible"`
		Closed            bool        `json:"closed"`
		Archived          bool        `json:"archived"`
		Bookmarked        interface{} `json:"bookmarked"`
		Liked             interface{} `json:"liked"`
		Tags              []string    `json:"tags"`
		TagsDescriptions  struct {
			DevInstall string `json:"dev-install,omitempty"`
		} `json:"tags_descriptions"`
		CategoryId        int    `json:"category_id"`
		HasAcceptedAnswer bool   `json:"has_accepted_answer"`
		UnicodeTitle      string `json:"unicode_title,omitempty"`
	} `json:"topics"`
	Users               []interface{} `json:"users"`
	Categories          []interface{} `json:"categories"`
	Tags                []interface{} `json:"tags"`
	Groups              []interface{} `json:"groups"`
	GroupedSearchResult struct {
		MorePosts           interface{}   `json:"more_posts"`
		MoreUsers           interface{}   `json:"more_users"`
		MoreCategories      interface{}   `json:"more_categories"`
		Term                string        `json:"term"`
		SearchLogId         int           `json:"search_log_id"`
		MoreFullPageResults bool          `json:"more_full_page_results"`
		CanCreateTopic      bool          `json:"can_create_topic"`
		Error               interface{}   `json:"error"`
		PostIds             []int         `json:"post_ids"`
		UserIds             []interface{} `json:"user_ids"`
		CategoryIds         []interface{} `json:"category_ids"`
		TagIds              []interface{} `json:"tag_ids"`
		GroupIds            []interface{} `json:"group_ids"`
	} `json:"grouped_search_result"`
}
