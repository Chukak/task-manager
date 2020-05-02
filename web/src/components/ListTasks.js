import React from "react";
import Task from "./Task";
import "./ListTasks.css";
import { List } from '@material-ui/core'

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
		return <div>
			<List>
				{this.state.listTasks.map((t, index) => {
					return (<Task key={index} taskData={t} />);
				})}
			</List>
		</div>
	}

	componentWillUnmount() {
	}
}