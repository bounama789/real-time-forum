package repositories

import (
	"fmt"
	"forum/config"
	db "forum/database"
	q "forum/database/query"
	"forum/models"
	"time"
)

type CommentRepository struct {
	BaseRepo
}

func (r *CommentRepository) init() {
	r.DB = db.DB
	r.TableName = db.COMMENTS_TABLE
}
func (r *CommentRepository) SaveComment(comment models.Comment) error {
	comment.CreatedAt =  time.Now().Format(string(config.Get("TIME_FORMAT").ToString()))
	err := r.DB.Insert(r.TableName, comment)
	if err != nil {
		return err
	}
	return nil
}
func (r *CommentRepository) DeleteComment(commentId string) error {
	err := r.DB.Delete(r.TableName, q.WhereOption{"comment_id": commentId})
	if err != nil {
		return err
	}
	return nil
}
func (r *CommentRepository) RemoveComment(commentId string) error {
	err := r.DB.Delete(r.TableName, q.WhereOption{"comment_id": commentId})
	if err != nil {
		return err
	}
	return nil
}
func (r *CommentRepository) GetComment(commentId string) (comment models.Comment, err error) {
	row,err := r.DB.GetOneFrom(r.TableName, q.WhereOption{"comment_id": commentId})
	if err != nil {
		return comment, err
	}
	err = row.Scan(&comment.CommentId, &comment.UserId,&comment.PostId, &comment.Body,&comment.Username,&comment.CreatedAt)
	if err != nil {
		return comment, fmt.Errorf("no value found")
	}
	return comment, nil
}

func (r *CommentRepository) UpdateComment(comment models.Comment) error {
	err := r.DB.Delete(r.TableName, q.WhereOption{"comment_id": comment.CommentId})
	if err != nil {
		return err
	}
	return nil
}

func (r *CommentRepository) GetPostComments(postId string) (comments []models.Comment, err error) {
	rows,err := r.DB.GetAllFrom(r.TableName, q.WhereOption{"pst_id": postId},"created_at DESC")
	if err != nil {
		return comments, err
	}
	var comment models.Comment
	for rows.Next() {
		err := rows.Scan(&comment.CommentId, &comment.UserId, &comment.PostId, &comment.Body,&comment.Username,&comment.CreatedAt)
		comments = append(comments, comment)
		fmt.Println(err)
	}

	return comments, nil
}

func (r *CommentRepository) GetCommentsCount(postId string) (int, error) {
	var count int
		row, err := r.DB.GetCount(r.TableName, q.WhereOption{"pst_id": postId})
		if err != nil {
			return 0, err
		}
		row.Scan(&count)

	return count, nil
}

