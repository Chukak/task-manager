import React from 'react';
import SwipeableViews from 'react-swipeable-views';
import { AppBar, Toolbar, Tabs, Tab, Box, Typography, 
	TextField, Grid } from '@material-ui/core'
import PriorityButton from './PriorityButton'
import DateTimeField from './DateTimeField'
import ActivityButton from './ActivityButton'
import ListTasks from './ListTasks' 
import TaskMenu from './TaskMenu'
import { UpdateTaskData, GetTaskData, ChangeTaskStatus,
	ChangeTaskActivity } from './Request'

class TabPanelDescription extends React.Component {
	constructor(props) {
		super(props)

		this.updateDescription = this.updateDescription.bind(this)
	}

	updateDescription(event) {
		UpdateTaskData({
			"id": this.props.taskData.taskID,
			"description": event.target.value
		})
	}

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
							variant="outlined" 
							onChange={this.updateDescription} />
					</Box>
				)}
		</div>
	}
}

class TabPanelExtraInfo extends React.Component {
	constructor(props) {
		super(props)

		this.state = {
			title: this.props.taskData.title,
			priority: this.props.taskData.priority,
			startTime: this.props.taskData.start,
			endTime: this.props.taskData.end,
			isOpen: this.props.taskData.opened,
			isActive: this.props.taskData.active
		}

		this.updateTitle = this.updateTitle.bind(this)
		this.updatePriority = this.updatePriority.bind(this)
		this.update = this.update.bind(this)
		this.startDate = React.createRef()
		this.endDate = React.createRef()
	}

	update() {
		GetTaskData({
			"id": this.props.taskData.taskID
		}).then(function(response) {
			this.setState({
				title: response.data.title,
				priority: response.data.priority,
				startTime: response.data.start,
				endTime: response.data.end,
				isOpen: response.data.opened,
				isActive: response.data.active
			})
			this.startDate.current.update({
				dateTime: response.data.start
			})
			this.endDate.current.update({
				dateTime: response.data.end
			})
		}.bind(this))
	}

	updateTitle(target) {
		UpdateTaskData({
			"id": this.props.taskData.taskID,
			"title": target.value
		})
		this.props.onUpdate({
			"pText": target.value,
			"sText": ""
		})
		this.setState({
			title: target.value
		})
	}

	updatePriority(value) {
		UpdateTaskData({
			"id": this.props.taskData.taskID,
			"priority": value
		})
	}

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
								defaultValue={this.state.title} 
								inputProps={{
									maxLength: 10,
								}}
								onKeyUp={(e) => {
									console.log("KEY: ", e)
									if (e.keyCode === 13) {
										this.updateTitle(e.target)
									}
								}}
								ref={this.titleField}/>
						</Box>
						<Box p={4} mt={2} display="flex">
							<Box flexGrow={1}>
								<Typography variant="h6" color="textPrimary">Task priority: </Typography>
							</Box>
							<Box>
								<PriorityButton 
									priority={this.state.priority}
									counterHandler={this.updatePriority} />
							</Box>
						</Box>
						
						<Box p={4} mt={2} display="flex">
							<Box flexGrow={1}>
								<Typography variant="h6" color="textPrimary">Start date: </Typography>
							</Box>
							<Box>
								<DateTimeField 
									dateTime={this.state.startTime} 
									ref={this.startDate}
									InputProps={{
										readOnly: true,
									}} 
									colon=":" 
									showSeconds />
							</Box>
						</Box>
						<Box p={4} mt={2} display="flex">
							<Box flexGrow={1}>
								<Typography variant="h6" color="textPrimary">End date: </Typography>
							</Box>
							<Box>
								<DateTimeField 
									dateTime={this.state.endTime} 
									ref={this.endDate}
									InputProps={{
										readOnly: true,
									}} 
									colon=":" 
									showSeconds />
							</Box>
						</Box>
						<Box p={4} mt={2} display="flex" flexDirection="row"> 
							<Box mr={2}>
								<ActivityButton 
									buttonName="Open"  
									isActive={this.state.isOpen} 
									onChangeValueHandler={ChangeTaskStatus}
									nameChangedProperty="open"
									onClickHandler={this.update} 
									TaskID={this.props.taskData.taskID} />
							</Box>
							<Box ml={2}>
								<ActivityButton 
									buttonName="Active" 
									isActive={this.state.isActive} 
									onChangeValueHandler={ChangeTaskActivity}
									nameChangedProperty="active"
									onClickHandler={this.update}
									TaskID={this.props.taskData.taskID} />
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
						<ListTasks hasSubtasks={true} listTasks={this.props.taskData.subtasks} />
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
		this.update = this.update.bind(this)
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

	update(data) {
		this.props.updateItem(data)
	}

	render() {
		return <div>
			<AppBar position="static" color="default">
				<Toolbar>
					<Grid justify={"space-between"} container>
						<Grid xs={10} item>
							<Tabs
								value={this.state.index}
								onChange={this.changeTabHandler}
								indicatorColor="primary"
								textColor="primary"
								variant="standard"
								aria-label="full width tabs example">
								<Tab label="Description" id="full-width-tab-0" aria-controls="full-width-tabpanel-0" />
								<Tab label="Information" id="full-width-tab-1" aria-controls="full-width-tabpanel-1" />
								<Tab label="Subtasks" id="full-width-tab-2" aria-controls="full-width-tabpanel-2" />
							</Tabs>
						</Grid>
						<Grid xs={1} item/>
						<Grid xs={1} item>
							<TaskMenu />
						</Grid>
					</Grid>
				</Toolbar>
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
					taskData={this.props.taskData}
					onUpdate={this.update} />
				<TabPanelSubtasks
					value={this.state.index} 
					index={2} 
					dir='ltr'
					taskData={this.props.taskData} />
			</SwipeableViews>
		</div>;
	}
}
