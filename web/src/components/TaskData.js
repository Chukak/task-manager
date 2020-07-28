import React from "react";
import { Box, Typography, TextField, Button, Grid, Collapse, Container } from "@material-ui/core";
import DateTimeField from "./DateTimeField";
import ActivityButton from "./ActivityButton";
import CounterButton from "./CounterButton";
import TimerDataField from "./TimerDataField";
import { GetTaskData, UpdateTaskData, ChangeTaskStatus, ChangeTaskActivity, GetTaskTimerValues } from "./Request";

const MaxProirityValue = 5;
const MinPriorityValue = 0;

const TitleTaskPriority = "Task priority:";
const TitleTaskSaveButton = "Save";
const TitleTaskDuration = "Task duration:";
const TitleTaskStartDate = "Start date:";
const TitleTaskEndDate = "End date:";

const dateLength = 10;
const dateLengthOffset = 1;
const dateLengthOffsetTime = 8;

const FullDateLength = dateLength + dateLengthOffset;
const FullTimeLength = dateLength + dateLengthOffset + dateLengthOffsetTime;

export default class TaskData extends React.Component {
	constructor(props) {
		super(props);

		this.taskID = this.props.taskID;
		this.state = {
			title: "",
			description: "",
			priority: 0,
			open: false,
			active: false,
			startDate: "",
			endDate: "",
			startTime: "",
			endTime: "",
			timer_Days: 0,
			timer_Hours: 0,
			timer_Minutes: 0,
			timer_Seconds: 0,
			descrFieldCollaped: false
		};

		this.timer = null;
	}

	updateTaskData() {
		UpdateTaskData({
			id: this.taskID,
			title: this.state.title,
			description: this.state.description,
			priority: this.state.priority
		});
		this.props.onUpdateItemData({ title: this.state.title });
	}

	getTaskData() {
		GetTaskData({ id: this.taskID }).then(
			function (response) {
				var refreshTitle = response.data.title !== this.state.title;

				this.setState({
					title: response.data.title,
					description: response.data.description,
					priority: response.data.priority,
					open: response.data.opened,
					active: response.data.active,
					startDate: response.data.start.substring(0, dateLength),
					endDate: response.data.end.substring(0, dateLength),
					startTime: response.data.start.substring(FullDateLength, FullTimeLength),
					endTime: response.data.end.substring(FullDateLength, FullTimeLength)
				});
				if (response.data.active) {
					this.startTimer();
				} else {
					this.stopTimer();
				}

				if (refreshTitle) {
					this.props.onUpdateItemData({ title: response.data.title });
				}
			}.bind(this)
		);
	}

	updateComponent(id) {
		this.taskID = id;
		this.getTaskData();
	}

	changeTaskProperties(propertyName, value) {
		var data = { id: this.taskID };
		data[propertyName] = value;

		var handler = null;
		if (propertyName === "open") {
			handler = ChangeTaskStatus;
		} else if (propertyName === "active") {
			handler = ChangeTaskActivity;
		}
		if (handler != null) {
			handler(data).then((response) => {
				this.getTaskData();
			});
		}
	}

	startTimer() {
		this.timer = setInterval(
			function () {
				this.updateTimerValues();
			}.bind(this),
			1000
		); // 1 sec
	}

	stopTimer() {
		clearInterval(this.timer);
		this.updateTimerValues();
	}

	updateTimerValues() {
		GetTaskTimerValues({ id: this.taskID }).then((response) => {
			this.setState({
				timer_Days: response.data.duration.days,
				timer_Hours: response.data.duration.hours,
				timer_Minutes: response.data.duration.minutes,
				timer_Seconds: response.data.duration.seconds
			});
		});
	}

	render() {
		return (
			<Container mt={2}>
				<Box>
					<Box p={4} mt={2} display="flex">
						<Box mr={2}>
							<Typography variant="h6" color="textPrimary">
								{TitleTaskDuration}
							</Typography>
						</Box>
						<TimerDataField
							days={this.state.timer_Days}
							hours={this.state.timer_Hours}
							minutes={this.state.timer_Minutes}
							seconds={this.state.timer_Seconds}
							active={this.state.active}
						/>
					</Box>
					<Box p={4} mt={2} display="flex">
						<TextField
							id="standart-basic"
							label="Task title"
							value={this.state.title}
							onChange={(event) => {
								this.setState({ title: event.target.value });
							}}
						/>
					</Box>
					<Box p={4} mt={2}>
						<Grid p={4} mt={2} container justify="space-between">
							<Grid item>
								<Box mr={2}>
									<Typography variant="h6" color="textPrimary">
										{TitleTaskStartDate}
									</Typography>
								</Box>
								<DateTimeField
									date={this.state.startDate}
									time={this.state.startTime}
									InputProps={{ readonly: true }}
									colon=":"
									showSeconds
								/>
							</Grid>
							<Grid item>
								<Box mr={2}>
									<Typography variant="h6" color="textPrimary">
										{TitleTaskEndDate}
									</Typography>
								</Box>
								<DateTimeField
									date={this.state.endDate}
									time={this.state.endTime}
									InputProps={{ readonly: true }}
									colon=":"
									showSeconds
								/>
							</Grid>
						</Grid>
					</Box>
					<Box p={4} mt={2} display="flex">
						<Box mr={2}>
							<Typography variant="h6" color="textPrimary">
								{TitleTaskPriority}
							</Typography>
						</Box>
						<Box>
							<CounterButton
								max={MaxProirityValue}
								min={MinPriorityValue}
								count={this.state.priority}
								onClick={(value) => {
									this.setState({ priority: value });
								}}
							/>
						</Box>
					</Box>
					<Box p={4} mt={2} display="flex" flexDirection="row">
						<Box mr={2}>
							<ActivityButton
								isActive={this.state.open}
								buttonNameOn={"open"}
								buttonNameOff={"close"}
								onChange={(value) => {
									this.changeTaskProperties("open", value);
								}}
							/>
						</Box>
						<Box ml={2}>
							<ActivityButton
								isActive={this.state.active}
								buttonNameOn={"activate"}
								buttonNameOff={"deactivate"}
								onChange={(value) => {
									this.changeTaskProperties("active", value);
								}}
							/>
						</Box>
						<Box ml={2}>
							<Button
								variant="contained"
								size="large"
								color="default"
								onClick={() => {
									this.updateTaskData();
								}}>
								{TitleTaskSaveButton}
							</Button>
						</Box>
						<Box ml={2}>
							<ActivityButton
								isActive={this.state.descrFieldCollaped}
								buttonNameOn={"show description"}
								buttonNameOff={"hide description"}
								colorChange="no"
								onChange={(value) => {
									this.setState({ descrFieldCollaped: value });
								}}
							/>
						</Box>
					</Box>
					<Box p={4} mt={2}>
						<Box p={4}>
							<Collapse timeout="auto" unmountOnExit in={this.state.descrFieldCollaped}>
								<TextField
									id="outlined-multiline-static"
									rows={30}
									variant="outlined"
									value={this.state.description}
									multiline
									fullWidth
									onChange={(event) => {
										this.setState({ description: event.target.value });
									}}
								/>
							</Collapse>
						</Box>
					</Box>
				</Box>
			</Container>
		);
	}
}
