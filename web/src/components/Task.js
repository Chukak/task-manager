import React from "react";
import { ListItem, ListItemText, ListItemAvatar, Collapse, 
	Box } from '@material-ui/core'
import { Assignment, ExpandLess, ExpandMore } from '@material-ui/icons'
import TaskInfoComponent from './TaskInfo'

export default class Task extends React.Component {
	constructor(props) {
		super(props)

		this.state = {
			isCollapsed: false,
			primaryText: props.taskData.title,
			secondText: props.taskData.priority
		}
		this.collapseHandler = this.collapseHandler.bind(this)
		this.updateItem = this.updateItem.bind(this)
	}

	collapseHandler() {
		this.setState(prev => ({
			isCollapsed: !prev.isCollapsed
		}));
	}

	updateItem(data) {
		this.setState({
			primaryText: data.pText,
			secondText: data.sText
		})
	}

	render() {
		return <div>
			<Box mb={1.5} mt={0.3} border={1} borderColor="primary.main" borderRadius={6}>
			<ListItem button key={this.props.ikey} onClick={this.collapseHandler}>
				<ListItemAvatar>
					<Assignment />
				</ListItemAvatar>
				<ListItemText primary={this.state.primaryText} 
					second={this.state.secondText}/>
					{this.state.isCollapsed ? <ExpandLess /> : <ExpandMore />}
			</ListItem>
			</Box>
			<Collapse 
				key={this.props.ikey}
				in={this.state.isCollapsed}
				timeout='auto'
				unmountOnExit>
					<TaskInfoComponent
						taskData={this.props.taskData}
						updateItem={this.updateItem} />
			</Collapse>
		</div>
	}
}
