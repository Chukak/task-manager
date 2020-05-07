-- user/password
CREATE ROLE task_manager_user WITH ENCRYPTED PASSWORD 'task_manager_password' LOGIN;
CREATE DATABASE task_manager_db OWNER task_manager_user;