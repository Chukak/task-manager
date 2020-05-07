import React from 'react';
import SwipeableViews from 'react-swipeable-views';
import { AppBar, Tabs, Tab, Box, Typography, 
	TextField } from '@material-ui/core'
import PriorityButton from './PriorityButton'
import DateTimeField from './DateTimeField'
import ActivityButton from './ActivityButton'
import ListTasks from './ListTasks' 

class TabPanelDescription extends React.Component {
	render() {
		return <div
			role="tabpanel"
			hidden={this.props.value !== this.props.index}
			id={`full-width-tabpanel-${this.props.index}`}
			aria-labelledby={`full-width-tab-${this.props.index}`}>
				{this.props.value === this.props.index && (
					<Box p={4}>
						<TextField
							multiline
							fullWidth
							defaultValue={this.props.taskData.description}
							id="outlined-multiline-flexible"
							rows={20}
							variant="outlined" />
					</Box>
				)}
		</div>
	}
}

class TabPanelExtraInfo extends React.Component {
	render() {
		return <div
			role="tabpanel"
			hidden={this.props.value !== this.props.index}
			id={`full-width-tabpanel-${this.props.index}`}
			aria-labelledby={`full-width-tab-${this.props.index}`}>
				{this.props.value === this.props.index && (
					<div>
						<Box p={4} mt={2} display="flex">
							<TextField 
								id="outlined-basic" 
								label="Title" 
								variant="outlined"
								defaultValue={this.props.taskData.title} 
								inputProps={{
									maxLength: 10,
								}} />
						</Box>
						<Box p={4} mt={2} display="flex">
							<Box flexGrow={1}>
								<Typography variant="h6" color="textPrimary">Task priority: </Typography>
							</Box>
							<Box>
								<PriorityButton priority={this.props.taskData.priority} />
							</Box>
						</Box>
						
						<Box p={4} mt={2} display="flex">
							<Box flexGrow={1}>
								<Typography variant="h6" color="textPrimary">Start date: </Typography>
							</Box>
							<Box>
								<DateTimeField dateTime={this.props.taskData.start} />
							</Box>
						</Box>
						<Box p={4} mt={2} display="flex">
							<Box flexGrow={1}>
								<Typography variant="h6" color="textPrimary">End date: </Typography>
							</Box>
							<Box>
								<DateTimeField dateTime={this.props.taskData.start} />
							</Box>
						</Box>
						<Box p={4} mt={2} display="flex" flexDirection="row"> 
							<Box mr={2}>
								<ActivityButton buttonName="Open"  isActive={this.props.taskData.opened}/>
							</Box>
							<Box ml={2}>
								<ActivityButton buttonName="Active" isActive={this.props.taskData.active}/>
							</Box>
						</Box>
					</div>
				)}
		</div>
	}
}

class TabPanelSubtasks extends React.Component {
	render() {
		return <div
			role="tabpanel"
			hidden={this.props.value !== this.props.index}
			id={`full-width-tabpanel-${this.props.index}`}
			aria-labelledby={`full-width-tab-${this.props.index}`}>
				{this.props.value === this.props.index && (
					<Box pt={2} ml={3}>
						<ListTasks hasSubtasks={true} listTasks={this.props.subtasks} />
					</Box>
				)}
		</div>
	}
}

export default class TaskInfoComponent extends React.Component {
	constructor(props) {
		super(props)

		this.state = {
			index: 0
		}
		this.changeTabHandler = this.changeTabHandler.bind(this)
		this.changeTabIndex = this.changeTabIndex.bind(this)
	} 
	
	changeTabHandler(event, value) {
		this.setState({
			index: value
		})
	}

	changeTabIndex(value) {
		this.setState({
			index: value
		})
	}

	render() {
		return <div>
			<AppBar position="static" color="default">
				<Tabs
					value={this.state.index}
					onChange={this.changeTabHandler}
					indicatorColor="primary"
					textColor="primary"
					variant="fullWidth"
					aria-label="full width tabs example">
					<Tab label="Description" id="full-width-tab-0" aria-controls="full-width-tabpanel-0" />
					<Tab label="Information" id="full-width-tab-1" aria-controls="full-width-tabpanel-1" />
					<Tab label="Subtasks" id="full-width-tab-2" aria-controls="full-width-tabpanel-2" />
				</Tabs>
			</AppBar>
			<SwipeableViews
				axis='x'
				index={this.state.index}
				onChangeIndex={this.changeTabIndex} >
				<TabPanelDescription 
					value={this.state.index} 
					index={0} 
					dir='ltr'
					taskData={this.props.taskData} />
				<TabPanelExtraInfo 
					value={this.state.index} 
					index={1} 
					dir='ltr'
					taskData={this.props.taskData} />
				<TabPanelSubtasks
					value={this.state.index} 
					index={2} 
					dir='ltr'
					subtasks={this.props.taskData.subtasks} />
			</SwipeableViews>
		</div>;
	}
}
