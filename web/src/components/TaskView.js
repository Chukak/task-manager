import React from "react";
import ListTasks from "./ListTasks";
import TaskData from "./TaskData";
import { Box } from "@material-ui/core";
import { CreateNewTask, RemoveTask } from "./Request";

export default class TaskView extends React.Component {
	constructor(props) {
		super(props);

		this.listTasksRef = React.createRef();
		this.taskDataRef = React.createRef();
		this.state = {
			selectedTaskId: -1,
			selectedTaskItem: null
		};
	}

	callAction(action) {
		var handler = null;
		if (action === "create") {
			handler = CreateNewTask;
		} else if (action === "remove") {
			handler = RemoveTask;
		}

		if (handler != null) {
			handler({ id: this.state.selectedTaskId }).then(
				function (response) {
					if (this.listTasksRef.current != null) {
						this.listTasksRef.current.update();
					}
				}.bind(this)
			);
		}
	}

	render() {
		return (
			<Box display="flex" flexDirection="row">
				<Box p={2}>
					<ListTasks
						ref={this.listTasksRef}
						onSelectTask={(id, obj) => {
							if (this.state.selectedTaskItem != null) {
								this.state.selectedTaskItem.clearSelection();
							}
							this.setState({
								selectedTaskId: id,
								selectedTaskItem: obj
							});
							this.taskDataRef.current.updateComponent(id);
						}}
					/>
				</Box>
				<Box flexGrow={1}>
					<TaskData
						taskID={this.state.selectedTaskId}
						ref={this.taskDataRef}
						onUpdateItemData={(data) => {
							if (this.state.selectedTaskItem != null) {
								this.state.selectedTaskItem.updateItemData(data);
							}
						}}
					/>
				</Box>
			</Box>
		);
	}
}
