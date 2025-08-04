/***********************
 *        Users        *
 ***********************/

CREATE TABLE IF NOT EXISTS users (
	           id INTEGER NOT NULL,
	   created_at NUMERIC NOT NULL DEFAULT unixepoch( 'now' ),
	   updated_at NUMERIC NOT NULL DEFAULT unixepoch( 'now' ),
	         uuid    TEXT NOT NULL,
	 display_name    TEXT NOT NULL,
	   first_name    TEXT NOT NULL,
	    last_name    TEXT NOT NULL,
	password_hash    TEXT NOT NULL,
	password_salt    TEXT NOT NULL,
	          bio    TEXT     NULL,
	   avatar_url    TEXT     NULL,
	  totp_secret    TEXT     NULL,
	  refresh_jti    TEXT     NULL,

	PRIMARY KEY ( id ),
	     UNIQUE ( display_name )
);

CREATE INDEX IF NOT EXISTS         users_by_date ON users ( created_at ASC );
CREATE INDEX IF NOT EXISTS         users_by_uuid ON users ( uuid ASC );
CREATE INDEX IF NOT EXISTS users_by_display_name ON users ( display_name ASC );
CREATE INDEX IF NOT EXISTS    users_by_full_name ON users ( last_name ASC, first_name ASC );

/*************************
 *        Friends        *
 *************************/

CREATE TABLE IF NOT EXISTS friends (
	  user_id INTEGER NOT NULL,
	friend_id INTEGER NOT NULL,

	PRIMARY KEY ( user_id, friend_id ),
	FOREIGN KEY ( user_id ) REFERENCES users ( id )
		ON DELETE CASCADE
		ON UPDATE RESTRICT,
	FOREIGN KEY ( friend_id ) REFERENCES users ( id )
		ON DELETE CASCADE
		ON UPDATE RESTRICT
);

/***********************
 *        Media        *
 ***********************/

CREATE TABLE IF NOT EXISTS media (
	        id INTEGER NOT NULL,
	created_at NUMERIC NOT NULL DEFAULT unixepoch( 'now' ),
	updated_at NUMERIC NOT NULL DEFAULT unixepoch( 'now' ),
	      uuid    TEXT NOT NULL,
	     title    TEXT NOT NULL,
	      desc    TEXT     NULL,
	   user_id INTEGER NOT NULL,
	      type    TEXT NOT NULL, -- MIME Media Type: [RFC2046](https://www.rfc-editor.org/rfc/rfc2046.html)
	    format    TEXT NOT NULL, -- MIME Media Subtype: [RFC2046](https://www.rfc-editor.org/rfc/rfc2046.html)
	       md5    TEXT NOT NULL,

	PRIMARY KEY ( id ),
	FOREIGN KEY ( user_id ) REFERENCES users ( id )
		ON DELETE CASCADE
		ON UPDATE RESTRICT
);

CREATE INDEX IF NOT EXISTS  media_by_date ON media ( created_at DESC );
CREATE INDEX IF NOT EXISTS  media_by_uuid ON media ( uuid ASC );
CREATE INDEX IF NOT EXISTS media_by_title ON media ( title ASC, created_at DESC );
CREATE INDEX IF NOT EXISTS  media_by_user ON media ( user_id ASC, created_at DESC );
CREATE INDEX IF NOT EXISTS  media_by_mime ON media ( type ASC, format ASC, created_at DESC );

/**********************
 *        Tags        *
 **********************/

CREATE TABLE IF NOT EXISTS tags (
	        id INTEGER NOT NULL,
	created_at NUMERIC NOT NULL DEFAULT unixepoch( 'now' ),
	   content    TEXT NOT NULL,

	PRIMARY KEY ( id ),
	     UNIQUE ( content )
);

CREATE INDEX IF NOT EXISTS    tags_by_date ON tags ( created_at DESC );
CREATE INDEX IF NOT EXISTS tags_by_content ON tags ( content ASC, created_at DESC );

/****************************
 *        Media Tags        *
 ****************************/

CREATE TABLE IF NOT EXISTS media_tags (
	created_at INTEGER NOT NULL DEFAULT unixepoch( 'now' ),
	  media_id INTEGER NOT NULL,
	    tag_id INTEGER NOT NULL,

	PRIMARY KEY ( media_id, tag_id ),

	FOREIGN KEY ( media_id ) REFERENCES media ( id )
		ON DELETE CASCADE
		ON UPDATE RESTRICT,

	FOREIGN KEY ( tag_id ) REFERENCES tags ( id )
		ON DELETE CASCADE
		ON UPDATE RESTRICT
);

CREATE INDEX media_tags_by_media ON media_tags ( media_id ASC, created_at DESC );
CREATE INDEX   media_tags_by_tag ON media_tags ( tag_id ASC, created_at DESC );

/**************************
 *        Comments        *
 **************************/

CREATE TABLE IF NOT EXISTS comments (
	        id INTEGER NOT NULL,
	created_at INTEGER NOT NULL DEFAULT unixepoch( 'now' ),
	updated_at INTEGER NOT NULL DEFAULT unixepoch( 'now' ),
	  media_id INTEGER NOT NULL,
	   user_id INTEGER NOT NULL,
	 parent_id INTEGER     NULL,
	   content    TEXT NOT NULL,

	PRIMARY KEY ( id ),
	FOREIGN KEY ( media_id ) REFERENCES media ( id )
		ON DELETE CASCADE
		ON UPDATE RESTRICT,
	FOREIGN KEY ( user_id ) REFERENCES users ( id )
		ON DELETE CASCADE
		ON UPDATE RESTRICT,
	FOREIGN KEY ( parent_id ) REFERENCES comments ( id )
		ON DELETE CASCADE
		ON UPDATE RESTRICT
);

CREATE INDEX  comments_by_media ON comments ( media_id ASC, user_id ASC, created_at DESC );
CREATE INDEX   comments_by_user ON comments ( user_id ASC, created_at DESC );
CREATE INDEX comments_by_parent ON comments ( parent_id ASC, created_at DESC );
