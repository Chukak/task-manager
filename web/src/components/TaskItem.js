import React from 'react';
import { ListItem, MenuItem, ListItemText, Box } from '@material-ui/core';

const UnknowTitle = "<Unknown>"
const SetTitle = (title) => { return title === "" ? UnknowTitle : title; };

export default class TaskItem extends React.Component {
	constructor(props) {
		super(props);

		this.state = {
			isSelected: false,
			taskTitle: SetTitle(this.props.taskTitle)
		};

		this.taskID = props.taskID;
	}

	clearSelection() {
		this.setState({isSelected: false});
	}

	updateItemData(data) {
		if (data.title != undefined) {
			this.setState({
				taskTitle: SetTitle(data.title)
			});
		}
	}

	render() {
		/** todo: subtasks */
		return <Box mb={1.5} mt={0.3} border={1} borderColor="primary.main" borderRadius={6}>
			<MenuItem button key={this.props.ikey} 
				onClick={() => { 
					this.props.onClickItem(this.taskID, this); 
					this.setState({isSelected: true});
				}}
				selected={this.state.isSelected}>
				<ListItemText primary={this.state.taskTitle} />
			</MenuItem>
		</Box>
	}
};
