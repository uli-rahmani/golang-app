package interfaces

import "saturday/domain"

type Repo interface {
	SaveComment(orgsName, comment string) error
	GetComment(orgsName string) ([]*domain.OrgsCommentBody, error)
	DeleteComment(orgsName string) error

	GetMember(orgsName string) ([]*domain.GetMemberResult, error)
}
