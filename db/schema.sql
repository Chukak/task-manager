CREATE SEQUENCE task_duration_sq 
	START WITH 1
	INCREMENT BY 1
	NO MINVALUE
	NO MAXVALUE
	CACHE 1;

CREATE TABLE task_duration (
	id BIGINT PRIMARY KEY DEFAULT nextval('task_duration_sq'::regclass) NOT NULL,
	second SMALLINT DEFAULT 0,
	minute SMALLINT DEFAULT 0,
	hour SMALLINT DEFAULT 0,
	day SMALLINT DEFAULT 0
);

CREATE SEQUENCE tasks_sq 
	START WITH 1
	INCREMENT BY 1
	NO MINVALUE
	NO MAXVALUE
	CACHE 1;

CREATE TABLE tasks (
	id BIGINT PRIMARY KEY DEFAULT nextval('tasks_sq'::regclass) NOT NULL,
	parent_id BIGINT DEFAULT -1,
	start_time TIMESTAMP WITH TIME ZONE DEFAULT TO_TIMESTAMP(0),
	end_time TIMESTAMP WITH TIME ZONE DEFAULT TO_TIMESTAMP(0),
	duration_id BIGINT, 
	is_open BOOLEAN DEFAULT false,
	is_active BOOLEAN DEFAULT false,
	title CHARACTER VARYING DEFAULT '',
	descr CHARACTER VARYING DEFAULT '',
	priority SMALLINT DEFAULT 0,
	FOREIGN KEY (duration_id) REFERENCES task_duration (id) ON DELETE CASCADE
);

