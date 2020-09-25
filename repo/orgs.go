package repo

import (
	"errors"
	"log"
	"saturday/constants"
	"saturday/domain"
)

// SaveComment - function for insert comment by orgs on db.
func (or *OrgsRepo) SaveComment(orgsName, comment string) error {
	query, args, err := or.DB.In(constants.QueryInsertComment, orgsName, comment)
	if err != nil {
		log.Println("Repo | SaveComment | error build param query, err: " + err.Error())
		return err
	}

	query = or.DB.Rebind(query)
	result, err := or.DB.Exec(query, args...)
	if err != nil {
		log.Println("Repo | SaveComment | error exec query, err: " + err.Error())
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Println("Repo | SaveComment | error get row affected, err: " + err.Error())
		return err
	}

	if rows == 0 {
		log.Println("Repo | SaveComment | no rows affected")
		return errors.New("no rows affected")
	}

	return nil
}

// GetComment - function for get comment by orgs on db.
func (or *OrgsRepo) GetComment(orgsName string) ([]*domain.OrgsCommentBody, error) {
	comments := make([]*domain.OrgsCommentBody, 0)

	query, args, err := or.DB.In(constants.QueryGetComment, orgsName)
	if err != nil {
		log.Println("Repo | GetComment | error build param query, err: " + err.Error())
		return comments, err
	}

	query = or.DB.Rebind(query)
	err = or.DB.Select(&comments, query, args...)
	if err != nil {
		log.Println("Repo | GetComment | error select query, err: " + err.Error())
		return comments, err
	}

	return comments, nil
}

// DeleteComment - function for delete comment by orgs on db (update status to deleted).
func (or *OrgsRepo) DeleteComment(orgsName string) error {
	query, args, err := or.DB.In(constants.QueryDeleteComment, orgsName)
	if err != nil {
		log.Println("Repo | DeleteComment | error build param query, err: " + err.Error())
		return err
	}

	query = or.DB.Rebind(query)
	result, err := or.DB.Exec(query, args...)
	if err != nil {
		log.Println("Repo | DeleteComment | error exec query, err: " + err.Error())
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Println("Repo | DeleteComment | error get row affected, err: " + err.Error())
		return err
	}

	if rows == 0 {
		log.Println("Repo | DeleteComment | no rows affected")
	}

	return nil
}

// GetMember - function for get member by orgs on db.
func (or *OrgsRepo) GetMember(orgsName string) ([]*domain.GetMemberResult, error) {
	members := make([]*domain.GetMemberResult, 0)

	query, args, err := or.DB.In(constants.QueryGetMember, orgsName)
	if err != nil {
		log.Println("Repo | GetMember | error build param query, err: " + err.Error())
		return members, err
	}

	query = or.DB.Rebind(query)
	err = or.DB.Select(&members, query, args...)
	if err != nil {
		log.Println("Repo | GetMember | error select query, err: " + err.Error())
		return members, err
	}

	return members, nil
}
