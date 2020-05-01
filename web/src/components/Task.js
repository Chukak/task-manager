import React from "react";
import "./Task.css";

export default class Task extends React.Component {
	constructor(props) {
		super(props)
		this.title = ""
		this.description = ""
		this.priority = 0
		this.isOpened = false
		this.isActive = false
		this.startTime = ""
		this.endTime = ""
		this.duration = "00:00:00"
	}

	render() {
		return <div>
				<div>{this.props.title}</div>
				<div>{this.props.description}</div>
			</div>
	}
}
