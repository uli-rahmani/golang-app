package constants

// Query comments.
const (
	QueryGetComment string = `
		SELECT 
			comment
		FROM comments
		WHERE
			org_name = ? and status = 'active'
	`

	QueryInsertComment string = `
		INSERT INTO comments 
			(org_name, comment, status, created_at)
		VALUES 
			(?, ?, 'active', NOW())
	`

	QueryDeleteComment string = `
		UPDATE comments SET 
			status = 'deleted', updated_at = NOW()
		WHERE org_name = ?
	`

	QueryGetMember string = `
		SELECT 
			org_name, avatar_url, followers, following
		FROM members
		WHERE
			org_name = ? ORDER BY followers DESC
	`
)
