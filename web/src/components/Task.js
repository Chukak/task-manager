import React from "react";
import "./Task.css";
import { ListItem, ListItemText, ListItemAvatar, Collapse, 
	Box } from '@material-ui/core'
import { Assignment, ExpandLess, ExpandMore } from '@material-ui/icons'
import TaskInfo from './TaskInfo'

export default class Task extends React.Component {
	constructor(props) {
		super(props)

		this.state = {
			isCollapsed: false
		}
		this.collapseHandler = this.collapseHandler.bind(this)
	}

	collapseHandler() {
		this.setState(prev => ({
			isCollapsed: !prev.isCollapsed
		}));
	}

	render() {
		return <div>
			<Box border={1} borderColor="primary.main" borderRadius={6}>
			<ListItem button key={this.props.ikey} onClick={this.collapseHandler}>
				<ListItemAvatar>
					<Assignment />
				</ListItemAvatar>
				<ListItemText primary={this.props.taskData.title} 
					second={this.props.taskData.priority}/>
					{this.state.isCollapsed ? <ExpandLess /> : <ExpandMore />}
			</ListItem>
			</Box>
			<Collapse 
				key={this.props.ikey}
				in={this.state.isCollapsed}
				timeout='auto'
				unmountOnExit>
					<TaskInfo taskData={this.props.taskData}></TaskInfo>
			</Collapse>
		</div>
	}
}
