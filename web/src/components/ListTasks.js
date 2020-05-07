import React from "react";
import Task from "./Task";
import { List } from '@material-ui/core'

const axios = require('axios');

export default class ListTasks extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			listTasks: []
		};
	}

	componentDidMount() {
		if (this.props.hasSubtasks) {
			this.setState({
				listTasks: this.props.listTasks
			});
		} else {
			axios.get('/api/task/all')
				.then(function(response) {
					console.log(response);
					this.setState({
						listTasks: response.data.listTasks
					});
				}.bind(this));
		}
	}

	render() {
		return <div>
			<List>
				{this.state.listTasks.map((t, index) => {
					return <Task key={index} ikey={index} taskData={t} />
				})}
			</List>
		</div>
	}

	componentWillUnmount() {
	}
}