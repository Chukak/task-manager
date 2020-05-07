CREATE SEQUENCE task_duration_sq 
	START WITH 1
	INCREMENT BY 1
	NO MINVALUE
	NO MAXVALUE
	CACHE 1;

CREATE TABLE task_duration (
	id BIGINT PRIMARY KEY DEFAULT nextval('task_duration_sq'::regclass) NOT NULL,
	second smallint,
	minute smallint,
	hour smallint,
	day smallint
);

CREATE SEQUENCE tasks_sq 
	START WITH 1
	INCREMENT BY 1
	NO MINVALUE
	NO MAXVALUE
	CACHE 1;

CREATE TABLE tasks (
	id BIGINT PRIMARY KEY DEFAULT nextval('tasks_sq'::regclass) NOT NULL,
	start_time TIMESTAMP DEFAULT TO_TIMESTAMP(0),
	end_date TIMESTAMP DEFAULT TO_TIMESTAMP(0),
	duration_id BIGINT, 
	is_open BOOLEAN DEFAULT false,
	is_active BOOLEAN DEFAULT false,
	title CHARACTER VARYING DEFAULT '',
	descr CHARACTER VARYING DEFAULT '',
	priority smallint DEFAULT 0,
	FOREIGN KEY (duration_id) REFERENCES task_duration (id) ON DELETE CASCADE
);

