package slack

type Event struct {
	Token    string `json:"token"`
	TeamID   string `json:"team_id"`
	APIAppID string `json:"api_app_id"`
	Event    struct {
		Type string `json:"type"`
		User struct {
			ID       string `json:"id"`
			TeamID   string `json:"team_id"`
			Name     string `json:"name"`
			Deleted  bool   `json:"deleted"`
			Color    string `json:"color"`
			RealName string `json:"real_name"`
			Tz       string `json:"tz"`
			TzLabel  string `json:"tz_label"`
			TzOffset int    `json:"tz_offset"`
			Profile  struct {
				Title                 string      `json:"title"`
				Phone                 string      `json:"phone"`
				Skype                 string      `json:"skype"`
				RealName              string      `json:"real_name"`
				RealNameNormalized    string      `json:"real_name_normalized"`
				DisplayName           string      `json:"display_name"`
				DisplayNameNormalized string      `json:"display_name_normalized"`
				Fields                interface{} `json:"fields"`
				StatusText            string      `json:"status_text"`
				StatusEmoji           string      `json:"status_emoji"`
				StatusExpiration      int         `json:"status_expiration"`
				AvatarHash            string      `json:"avatar_hash"`
				Email                 string      `json:"email"`
				Image24               string      `json:"image_24"`
				Image32               string      `json:"image_32"`
				Image48               string      `json:"image_48"`
				Image72               string      `json:"image_72"`
				Image192              string      `json:"image_192"`
				Image512              string      `json:"image_512"`
				StatusTextCanonical   string      `json:"status_text_canonical"`
				Team                  string      `json:"team"`
			} `json:"profile"`
			IsAdmin           bool   `json:"is_admin"`
			IsOwner           bool   `json:"is_owner"`
			IsPrimaryOwner    bool   `json:"is_primary_owner"`
			IsRestricted      bool   `json:"is_restricted"`
			IsUltraRestricted bool   `json:"is_ultra_restricted"`
			IsBot             bool   `json:"is_bot"`
			IsAppUser         bool   `json:"is_app_user"`
			Updated           int    `json:"updated"`
			Presence          string `json:"presence"`
		} `json:"user"`
		CacheTs int    `json:"cache_ts"`
		EventTs string `json:"event_ts"`
	} `json:"event"`
	Type        string   `json:"type"`
	EventID     string   `json:"event_id"`
	EventTime   int      `json:"event_time"`
	AuthedUsers []string `json:"authed_users"`
	Challenge   string   `json:"challenge"`
}

type Verify struct {
	Challenge string `json:"challenge"`
}
