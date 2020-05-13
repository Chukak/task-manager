import React from "react";
import Task from "./Task";
import { List } from '@material-ui/core'
import { GetAllTasks } from './Request'

const axios = require('axios');

export default class ListTasks extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			listTasks: []
		};
	}

	update() {
		GetAllTasks()
			.then(function(response) {
				var l = response.data.listTasks
				this.setState({
					listTasks: l === null ? [] : l
				});
			}.bind(this))
	}

	componentDidMount() {
		if (this.props.hasSubtasks) {
			this.setState({
				listTasks: this.props.listTasks
			});
		} else {
			this.update()
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
}