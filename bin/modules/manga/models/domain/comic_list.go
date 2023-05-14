package domain

import "komikku-api/bin/modules/chapter/models/domain"

type (
	Comic struct {
		Title		string	`json:"title"`
		Image		string	`json:"image"`
		Desc		string 	`json:"desc,omitempty"`
		Type		string 	`json:"type,omitempty"`
		Endpoint	string	`json:"endpoint"`
	}

	ComicDetail struct {
		Title				string	`json:"title"`
		FirstChapter		string 	`json:"first_chapter"`
		UrlFirstChapter		string 	`json:"url_first_chapter"`
		LatestChapter		string 	`json:"latest_chapter"`
		UrlLatestChapter 	string 	`json:"url_latest_chapter"`
		Description			string `json:"description"`
		Thumbnail			string `json:"thumbnail"`

	}

	ComicInfo struct {
		Thumbnail	string 		`json:"thumbnail"`
		Title		string 		`json:"title"`
		Type		string 		`json:"type"`
		Author		string 		`json:"author"`
		Status		string 		`json:"status"`
		Rating		string 		`json:"rating"`
		Genre		[]string 	`json:"genre"`
		ChapterList []domain.Chapter 	`json:"chapter_list"`
	}


)