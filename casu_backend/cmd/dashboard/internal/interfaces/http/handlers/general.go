package handlers

import (
	apperrors "github.com/hsynrtn/dashboard-management/pkg/errors"
	httpjson "github.com/hsynrtn/dashboard-management/pkg/http/response/json"
	"net/http"
)

type SporTypes struct {
	Footbal    string `json:"footbal"`
	Basketball string `json:"basketball"`
	Tennis     string `json:"tennis"`
}
type SporTypeResponseModel struct {
	Cities map[string]City `json:"cities"`
}

type City struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type Res struct {
	Cities map[string]City `json:"cities"`
}

// BuildListSporTypes
func BuildListSportTypes() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {

		//var i = SporTypeResponseModel{
		//	Cities: map[string]City{},
		//}
		//i.Cities["1"] = City{Name: "3333", Code: "333"}

		var datas []interface{}
		datas = append(datas, City{
			Id: 1, Name: "Denizli", Code: "https://sportshub.cbsistatic.com/i/r/2024/03/28/5152fb65-4e6a-4bbc-b003-ace0df54b759/thumbnail/1200x675/98a759414e1b665b2df7ca5aa6f1308e/rome-odunze.jpg"})
		var i = Response{Datas: datas}

		if err := httpjson.JSON(r.Context(), w, http.StatusOK, i); err != nil {
			return apperrors.Wrap(err)
		}

		return nil
	}

	return httpjson.HandlerFunc(fn)
}

type AutoGenerated struct {
	BloodTypes []BT
}

type BT struct {
	Name struct {
		Id int
	} `json:"name"`
}
type Article struct {
	Id          int    `json:"id"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	UrlToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}

type Response struct {
	Datas []interface{} `json:"data"`
}

// BuildListBloodtTypes
func BuildListBloodtTypes() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) error {
		var datas []interface{}
		datas = append(datas, Article{
			Id:          1,
			Author:      "Yazar",
			Title:       "Pete Prisco 2024 NFL Mock Draft 1.0: Jets add firepower for Aaron Rodgers; Super Bowl hopefuls address defense - CBS Sports",
			Description: "fdfd",
			Url:         "https://www.cbssports.com/nfl/draft/news/pete-prisco-2024-nfl-mock-draft-1-0-jets-add-firepower-for-aaron-rodgers-super-bowl-hopefuls-address-defense/",
			UrlToImage:  "https://sportshub.cbsistatic.com/i/r/2024/03/28/5152fb65-4e6a-4bbc-b003-ace0df54b759/thumbnail/1200x675/98a759414e1b665b2df7ca5aa6f1308e/rome-odunze.jpg",
			PublishedAt: "2024-04-02T18:12:48Z",
			Content:     "Free agency has come and gone for the most part, so mock drafts actually mean something.\\r\\nUnlike our lead NFL Draft analyst Ryan Wilson, I am late to the process when it comes to doing mock drafts. I… [+1855 chars]"})

		datas = append(datas, Article{
			Id:          1,
			Author:      "Yazar",
			Title:       "Pete Prisco 2024 NFL Mock Draft 1.0: Jets add firepower for Aaron Rodgers; Super Bowl hopefuls address defense - CBS Sports",
			Description: "fdfd",
			Url:         "https://www.cbssports.com/nfl/draft/news/pete-prisco-2024-nfl-mock-draft-1-0-jets-add-firepower-for-aaron-rodgers-super-bowl-hopefuls-address-defense/",
			UrlToImage:  "https://sportshub.cbsistatic.com/i/r/2024/03/28/5152fb65-4e6a-4bbc-b003-ace0df54b759/thumbnail/1200x675/98a759414e1b665b2df7ca5aa6f1308e/rome-odunze.jpg",
			PublishedAt: "2024-04-02T18:12:48Z",
			Content:     "Free agency has come and gone for the most part, so mock drafts actually mean something.\\r\\nUnlike our lead NFL Draft analyst Ryan Wilson, I am late to the process when it comes to doing mock drafts. I… [+1855 chars]"})

		var i = Response{Datas: datas}
		if err := httpjson.JSON(r.Context(), w, http.StatusOK, i); err != nil {
			return apperrors.Wrap(err)
		}

		return nil
	}

	return httpjson.HandlerFunc(fn)
}
