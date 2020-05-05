import React from "react";
import Task from "./Task";
import { List } from '@material-ui/core'

const axios = require('axios');

export default class ListTasks extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			taskCount: 0,
			listTasks: []
		};
	}

	componentDidMount() {
		axios.get('/api/task/all')
			.then(function(response) {
				console.log(response);
				this.setState({
					taskCount: response.data.taskCount,
					listTasks: response.data.listTasks
				});
			}.bind(this));
	}

	render() {
		return <div>
			<List>
				{this.state.listTasks.map((t, index) => {
					return (<Task key={index} ikey={index} taskData={t} />);
				})}
			</List>
		</div>
	}

	componentWillUnmount() {
	}
}