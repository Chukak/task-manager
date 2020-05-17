const axios = require('axios');

const mainURL = "/api/task"

function postRequestTo(url, data) {
	return axios.post(url, data)
}

function getRequestTo(url) {
	return axios.get(url)
}

export function GetAllTasks() {
	return getRequestTo(mainURL + "/all")
}

export function CreateNewTask(data) {
	return postRequestTo(mainURL + "/new", data)
}

export function UpdateTaskData(taskData) {
	return postRequestTo(mainURL + "/update", taskData)
}

export function GetTaskData(data) {
	return postRequestTo(mainURL + "/get", data)
}

export function RemoveTask(data) {
	return postRequestTo(mainURL + "/remove", data)
}

export function ChangeTaskStatus(data) {
	console.log("OPENED: ", data)
	return postRequestTo(mainURL + "/open", data)
}

export function ChangeTaskActivity(data) {
	return postRequestTo(mainURL + "/active", data)
}