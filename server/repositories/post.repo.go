package repositories

import (
	"database/sql"
	"fmt"
	db "forum/database"
	q "forum/database/query"
	"forum/models"
	"strconv"
	"strings"

	"github.com/gofrs/uuid/v5"
)

type PostRepository struct {
	BaseRepo
}

func (r *PostRepository) init() {
	r.DB = db.DB
	r.TableName = db.POSTS_TABLE
}

type PostCatLink struct {
	PostId     uuid.UUID `json:"pst_id"`
	CategoryId int       `json:"cat_id"`
}

func (r *PostRepository) SavePost(post models.Post, categories []int) error {
	err := r.DB.Insert(r.TableName, post)
	if err != nil {
		return err
	}
	for _, id := range categories {
		link := PostCatLink{
			PostId:     post.PostId,
			CategoryId: id,
		}
		err := r.DB.Insert(db.CAT_POST_TABLE, link)
		if err != nil {
			r.DeletePost(post.PostId.String())
			return err
		}
	}

	return nil
}

func (r *PostRepository) DeletePost(postId string) error {
	err := r.DB.Delete(r.TableName, q.WhereOption{"post_id": postId})
	if err != nil {
		return err
	}
	return nil
}

func (r *PostRepository) UpdatePost(post models.Post) error {
	err := r.DB.Delete(r.TableName, q.WhereOption{"post_id": post.PostId})
	if err != nil {
		return err
	}
	return nil
}

func (r *PostRepository) GetPost(postId string) (post models.Post, err error) {
	row, err := r.DB.GetOneFrom(r.TableName, q.WhereOption{"post_id": postId})
	if err != nil {
		return post, err
	}
	err = row.Scan(&post.PostId, &post.Title, &post.Body, &post.Username, &post.UserId, &post.Status, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		fmt.Println(err)
		return post, fmt.Errorf("no value found")
	}
	return post, nil
}

func (r *PostRepository) GetPostByUser(userId string) (posts []models.Post, err error) {
	var post models.Post
	rows, err := r.DB.GetAllFrom(r.TableName, q.WhereOption{"user_id": userId}, "")
	if err != nil {
		return posts, err
	}
	for rows.Next() {
		rows.Scan(&post.PostId, &post.Title, &post.Body, &post.Username, &post.UserId, &post.Status, &post.CreatedAt, &post.UpdatedAt)
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *PostRepository) GetPostByCategory(categoryId int) (posts []models.Post, err error) {
	var post models.Post
	rows, err := r.DB.GetAllFrom(r.TableName, q.WhereOption{"category_id": categoryId}, "")
	if err != nil {
		return posts, err
	}
	for rows.Next() {
		rows.Scan(&post.PostId, &post.Title, &post.Body, &post.Username, &post.UserId, &post.Status, &post.CreatedAt, &post.UpdatedAt)
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *PostRepository) GetPostByFollow(userId string) (posts []models.Post, err error) {
	var post models.Post
	joinCond := []q.JoinCondition{
		{Table: db.USERS_TABLE, ForeignKey: "following_user_id", Reference: "user_id"},
		{Table: db.FOLLOWS_TABLE, ForeignKey: "followed_user_id", Reference: "user_id"},
	}
	rows, err := r.DB.GetAllAndJoin(r.TableName, joinCond, q.WhereOption{"user_id": userId}, "")
	if err != nil {
		return posts, err
	}
	for rows.Next() {
		rows.Scan(&post.PostId, &post.Title, &post.Body, &post.Username, &post.UserId, &post.Status, &post.CreatedAt, &post.UpdatedAt)
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *PostRepository) GetPosts(t models.TokenData, options map[string]string) (posts []models.Post, err error) {
	var post models.Post
	var wh = make(q.WhereOption)
	var joinConds []q.JoinCondition
	var orderBy string
	if options["liked"] == "1" {
		joinConds = append(joinConds, []q.JoinCondition{
			{Table: db.REACTIONS_TABLE,
				ForeignKey: "pst_id",
				Reference:  "post_id",
			},
		}...)
		wh["usr_id"] = t.UserId
		wh["reactions"] = "LIKE"
		wh["react_type"] = "POST"

	}

	if options["commented"] == "1" {
		joinConds = append(joinConds, []q.JoinCondition{
			{Table: db.COMMENTS_TABLE,
				ForeignKey: "pst_id",
				Reference:  "post_id",
			},
		}...)
		wh["usr_id"] = t.UserId
	}

	if categoryId, err := strconv.Atoi(options["category"]); err == nil {
		joinConds = append(joinConds, []q.JoinCondition{
			{Table: db.CAT_POST_TABLE,
				ForeignKey: "pst_id",
				Reference:  "post_id",
			},
			{Table: db.CATEGORIES_TABLE,
				ForeignKey: "cat_id",
				Reference:  "category_id",
			},
		}...)
		wh["category_id"] = categoryId
	}
	if options["created"] == "1" {
		wh["user_id"] = t.UserId
	}
	if orderBy != "TIME-ASC" && orderBy != "TIME-DESC" {
		orderBy = ""
	} else {
		orderBy = "created_at " + strings.Split(orderBy, "-")[1]
	}

	var rows *sql.Rows
	if len(joinConds) > 0 {
		rows, err = r.DB.GetAllAndJoin(r.TableName, joinConds, wh, orderBy)
		if err != nil {
			return posts, err
		}
	} else {
		n := len(wh)
		if n > 0 {
			rows, err = r.DB.GetAllFrom(r.TableName, wh, "created_at DESC")
		} else {
			rows, err = r.DB.GetAllFrom(r.TableName, nil, "created_at DESC")

		}
		if err != nil {
			return posts, err
		}

	}

	for rows.Next() {
		err := rows.Scan(&post.PostId, &post.Title, &post.Body, &post.Username, &post.UserId, &post.Status, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			fmt.Println(err)
		}
		posts = append(posts, post)
	}
	return posts, err
}

func (r *PostRepository) GetPostCategories(postId string) (cats []models.Category, err error) {
	var joinCond = []q.JoinCondition{
		{
			Table:      db.CAT_POST_TABLE,
			ForeignKey: "cat_id",
			Reference:  "category_id",
		},
		{
			Table:      r.TableName,
			ForeignKey: "post_id",
			Reference:  "pst_id",
		},
	}
	var wh = q.WhereOption{
		"pst_id": postId,
	}
	rows, err := r.DB.GetAllAndJoin(db.CATEGORIES_TABLE, joinCond, wh, "name ASC")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var category models.Category
	for rows.Next() {
		rows.Scan(&category.CategoryId, &category.Name, &category.Color)
		cats = append(cats, category)
	}
	return cats, nil
}

func (r *PostRepository) SearchSuggestions(keywords []string) (rows *sql.Rows) {
	query := q.SearchPostSuggestionQuery(keywords)
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, err = stmt.Query()
	if err != nil {
		fmt.Println(err)
		return
	}
	return rows
}
