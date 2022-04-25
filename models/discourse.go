package models

import "time"

type DiscourseSearchResponse struct {
	Posts []DiscourseSearchResponsePost `json:"posts"`
}

type DiscourseSearchResponsePost struct {
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

type DiscoursePostResponse struct {
	PostStream struct {
		Posts []struct {
			Id                int         `json:"id"`
			Name              string      `json:"name"`
			Username          string      `json:"username"`
			AvatarTemplate    string      `json:"avatar_template"`
			CreatedAt         time.Time   `json:"created_at"`
			Cooked            string      `json:"cooked"`
			PostNumber        int         `json:"post_number"`
			PostType          int         `json:"post_type"`
			UpdatedAt         time.Time   `json:"updated_at"`
			ReplyCount        int         `json:"reply_count"`
			ReplyToPostNumber interface{} `json:"reply_to_post_number"`
			QuoteCount        int         `json:"quote_count"`
			IncomingLinkCount int         `json:"incoming_link_count"`
			Reads             int         `json:"reads"`
			ReadersCount      int         `json:"readers_count"`
			Score             float64     `json:"score"`
			Yours             bool        `json:"yours"`
			TopicId           int         `json:"topic_id"`
			TopicSlug         string      `json:"topic_slug"`
			DisplayUsername   string      `json:"display_username"`
			PrimaryGroupName  interface{} `json:"primary_group_name"`
			FlairName         interface{} `json:"flair_name"`
			FlairUrl          interface{} `json:"flair_url"`
			FlairBgColor      interface{} `json:"flair_bg_color"`
			FlairColor        interface{} `json:"flair_color"`
			Version           int         `json:"version"`
			CanEdit           bool        `json:"can_edit"`
			CanDelete         bool        `json:"can_delete"`
			CanRecover        bool        `json:"can_recover"`
			CanWiki           bool        `json:"can_wiki"`
			Read              bool        `json:"read"`
			UserTitle         interface{} `json:"user_title"`
			Bookmarked        bool        `json:"bookmarked"`
			ActionsSummary    []struct {
				Id    int `json:"id"`
				Count int `json:"count"`
			} `json:"actions_summary"`
			Moderator             bool        `json:"moderator"`
			Admin                 bool        `json:"admin"`
			Staff                 bool        `json:"staff"`
			UserId                int         `json:"user_id"`
			Hidden                bool        `json:"hidden"`
			TrustLevel            int         `json:"trust_level"`
			DeletedAt             interface{} `json:"deleted_at"`
			UserDeleted           bool        `json:"user_deleted"`
			EditReason            interface{} `json:"edit_reason"`
			CanViewEditHistory    bool        `json:"can_view_edit_history"`
			Wiki                  bool        `json:"wiki"`
			CustomerFlairCustomer interface{} `json:"customer_flair_customer"`
			CanAcceptAnswer       bool        `json:"can_accept_answer"`
			CanUnacceptAnswer     bool        `json:"can_unaccept_answer"`
			AcceptedAnswer        bool        `json:"accepted_answer"`
		} `json:"posts"`
	} `json:"post_stream"`
	Id int `json:"id"`
}
