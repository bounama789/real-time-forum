package repositories

import (
	"database/sql"
	"fmt"
	db "forum/backend/database"
	q "forum/backend/database/query"
	"forum/backend/models"
)

// Todo: Implement reaction repository's functions here
type ReactionRepository struct {
	BaseRepo
}

func (r *ReactionRepository) init() {
	r.DB = db.DB
	r.TableName = db.REACTIONS_TABLE
}

func (r *ReactionRepository) GetReactById(reactId string) (reaction models.Reaction, err error) {
	row, err := r.DB.GetOneFrom(r.TableName, q.WhereOption{"react_id": reactId})
	if err != nil {
		return reaction, err
	}
	err = row.Scan(&reaction.ReactId, &reaction.PostId, &reaction.CommentId, &reaction.UserId, &reaction.Reactions, &reaction.ReacType, &reaction.CreatedAt, &reaction.UpdatedAt)
	if err != nil {
		return reaction, err
	}
	return reaction, nil
}

func (r *ReactionRepository) GetReactByUser(userId, id, reactType string) (reaction models.Reaction, err error) {
	switch reactType {
	case "POST":
		row, err := r.DB.GetOneFrom(r.TableName, q.WhereOption{"usr_id": userId, "pst_id": id, "react_type": reactType})
		if err != nil {
			return reaction, err
		}
		err = row.Scan(&reaction.ReactId, &reaction.PostId, &reaction.CommentId, &reaction.UserId, &reaction.Reactions, &reaction.ReacType, &reaction.CreatedAt, &reaction.UpdatedAt)
		if err != nil && err != sql.ErrNoRows {
			fmt.Println(err)
			return reaction, err
		}
		return reaction, nil
	case "COMMENT":
		row, err := r.DB.GetOneFrom(r.TableName, q.WhereOption{"usr_id": userId, "comment_id": id, "react_type": reactType})
		if err != nil {
			return reaction, err
		}
		err = row.Scan(&reaction.ReactId, &reaction.PostId, &reaction.CommentId, &reaction.UserId, &reaction.Reactions, &reaction.ReacType, &reaction.CreatedAt, &reaction.UpdatedAt)
		if err != nil && err != sql.ErrNoRows {
			fmt.Println(err)

			return reaction, err
		}
		return reaction, nil
	}
	return models.Reaction{}, fmt.Errorf("unknown reaction type")
}

func (r *ReactionRepository) SaveReaction(react models.Reaction) error {
	err := r.DB.Insert(r.TableName, react)
	if err != nil {
		return err
	}
	return nil
}

func (r *ReactionRepository) UpdateReaction(react models.Reaction) error {
	err := r.DB.Update(r.TableName, react, q.WhereOption{"react_id": react.ReactId})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (r *ReactionRepository) DeleteReaction(reactId string) error {
	err := r.DB.Delete(r.TableName, q.WhereOption{"react_id": reactId})
	if err != nil {
		return err
	}
	return nil
}

func (r *ReactionRepository) DeleteCommentReact(userId, comment_id string) error {
	err := r.DB.Delete(r.TableName, q.WhereOption{"usr_id": userId, "pst_id": comment_id})
	if err != nil {
		return err
	}
	return nil
}

func (r *ReactionRepository) GetVotes(reactType, id string) (int, error) {
	var posCount int
	switch reactType {
	case "POST":
		row, err := r.DB.GetCount(r.TableName, q.WhereOption{"react_type": reactType, "pst_id": id, "reactions": "LIKE"})
		if err != nil {
			return 0, err
		}
		row.Scan(&posCount)
	case "COMMENT":
		row, err := r.DB.GetCount(r.TableName, q.WhereOption{"react_type": reactType, "comment_id": id, "reactions": "LIKE"})
		if err != nil {
			return 0, err
		}
		row.Scan(&posCount)
	}
	return posCount, nil
}
