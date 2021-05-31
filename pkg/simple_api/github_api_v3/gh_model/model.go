package gh_model

import (
	"github.com/google/go-github/v35/github"
	"github.com/xunull/helper-github/pkg/local_db/custom"
)

type Repository struct {
	ID           int64                  `json:"id,omitempty" gorm:"uniqueIndex"`
	Name         string                 `json:"name,omitempty"`
	FullName     string                 `json:"full_name,omitempty" gorm:"uniqueIndex"`
	Description  string                 `json:"description,omitempty"`
	Language     string                 `json:"language,omitempty"`
	Visibility   string                 `json:"visibility,omitempty"`
	Topics       custom.StringArrayJson `json:"topics,omitempty" gorm:"type:text"`
	LabelsURL    string                 `json:"labels_url,omitempty"`
	LanguagesURL string                 `json:"languages_url,omitempty"`
	CreatedAt    *github.Timestamp
}

type User struct {
	ID          int64  `json:"id,omitempty"`
	NodeID      string `json:"node_id,omitempty"`
	AvatarURL   string `json:"avatar_url,omitempty"`
	Name        string `json:"name,omitempty"`
	Company     string `json:"company,omitempty"`
	Blog        string `json:"blog,omitempty"`
	Location    string `json:"location,omitempty"`
	Email       string `json:"email,omitempty"`
	PublicRepos int    `json:"public_repos,omitempty"`
	Followers   int    `json:"followers,omitempty"`
	Following   int    `json:"following,omitempty"`
	//CreatedAt   github.Timestamp `json:"created_at,omitempty"`
	//UpdatedAt   github.Timestamp `json:"updated_at,omitempty"`
	//SuspendedAt github.Timestamp `json:"suspended_at,omitempty"`
	Starred []Repository `gorm:"many2many:user_starred;"`
}
