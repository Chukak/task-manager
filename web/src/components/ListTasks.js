import React from "react";
import TaskItem from "./TaskItem";
import { List, Paper } from "@material-ui/core";
import { GetAllTasks } from "./Request";

export default class ListTasks extends React.Component {
	constructor(props) {
		super(props);

		this.state = {
			listTasks: []
		};
	}

	update() {
		GetAllTasks().then(
			function (response) {
				var reloadPage = response.data.listTasks.length < this.state.listTasks.length;
				this.setState({
					listTasks: response.data.listTasks
				});
				if (reloadPage) {
					window.location.reload(false);
				}
			}.bind(this)
		);
	}

	componentDidMount() {
		this.update();
	}

	render() {
		return (
			<Paper style={{ minWidth: 300, maxHeight: 600, overflow: "auto" }}>
				<List>
					{this.state.listTasks.map((t, index) => {
						return (
							<TaskItem
								key={index}
								ikey={index}
								taskTitle={t.title}
								taskID={t.taskID}
								onClickItem={(id, obj) => {
									this.props.onSelectTask(id, obj);
								}}
							/>
						);
					})}
				</List>
			</Paper>
		);
	}
}
