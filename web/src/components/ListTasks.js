import React from "react";
import Task from "./Task";
import "./ListTasks.css";

export default class ListTasks extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			taskCount: 0,
			listTasks: []
		};
	}

	componentDidMount() {
		fetch('/api/task/all')
			.then(response => response.json())
				.then(json => this.setState({
					taskCount: json.taskCount,
					listTasks: json.listTasks
				}));
	}

	render() {
		console.log(this.state.listTasks)
		var tasks = this.state.listTasks.map((t) =>
			<Task title={t.title} 
				description={t.description} 
				/>
		);

		return <div>{tasks}</div>
	}

	componentWillUnmount() {
		this.state.listTasks = [];
		this.state.taskCout = 0;
	}
}