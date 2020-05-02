import React from 'react';
import SwipeableViews from 'react-swipeable-views';
import './TaskInfo.css'
import { AppBar, Tabs, Tab, Box, Typography } from '@material-ui/core'

class TabPanelDescription extends React.Component {
	render() {
		return <div
			role="tabpanel"
			hidden={this.props.value !== this.props.index}
			id={`full-width-tabpanel-${this.props.index}`}
			aria-labelledby={`full-width-tab-${this.props.index}`}>
				{this.props.value === this.props.index && (
					<Box p={3}>
						<Typography variant="body1">
							{this.props.description}
						</Typography>
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
					<Box p={3}>
						<div></div>
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
				</Tabs>
			</AppBar>
			<SwipeableViews
				axis='x'
				index={this.state.index}
				onChangeIndex={this.changeTabIndex} >
				<TabPanelDescription value={this.state.index} 
					index={0} 
					dir='ltr'
					description={this.props.taskData.description}>
				</TabPanelDescription>
				<TabPanelExtraInfo value={this.state.index} index={1} dir='ltr'>
				</TabPanelExtraInfo>
			</SwipeableViews>
		</div>;
	}
}
