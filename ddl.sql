CREATE SEQUENCE comments_pk_seq
	INCREMENT BY 1
	MINVALUE 0
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE comments(
	comment_id bigint NOT NULL DEFAULT nextval('comments_pk_seq'::regclass),
	org_name VARCHAR(225) NOT NULL,
	comment VARCHAR(225) NOT NULL,
	status VARCHAR(20) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NULL,
	CONSTRAINT pk_comment_id PRIMARY KEY (comment_id)
);

CREATE INDEX idx_comments_org_name ON comments (org_name);
CREATE INDEX idx_comments_comment ON comments (comment);

CREATE SEQUENCE members_pk_seq
	INCREMENT BY 1
	MINVALUE 0
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE members(
	member_id bigint NOT NULL DEFAULT nextval('members_pk_seq'::regclass),
	org_name VARCHAR(225) NOT NULL,
	avatar_url VARCHAR(225) NOT NULL,
	followers int NOT NULL,
	following int NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NULL,
	CONSTRAINT pk_member_id PRIMARY KEY (member_id)
);

CREATE INDEX idx_members_org_name ON comments (org_name);

