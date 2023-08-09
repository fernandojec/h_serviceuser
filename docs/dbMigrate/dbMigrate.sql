CREATE TABLE public.users (
	user_id varchar(10) NOT NULL,
	first_name varchar(100) NULL,
	last_name varchar(100) NULL,
	"password" varchar(100) NULL,
	email varchar(100) NULL,
	is_active bool NULL,
	user_create varchar(10) NULL,
	create_at timestamp NULL,
	CONSTRAINT users_pkey PRIMARY KEY (user_id)
);