import React from 'react'
import { TextField, Box } from '@material-ui/core'

const dateLength = 10;

export default class DataTimeField extends React.Component {
	constructor(props) {
		super(props)

		let dt = props.dateTime;
		this.date = dt.substring(0, dateLength);
		this.time = dt.substring(dateLength + 1);
	}

	render() {
		return <Box p={2}>
			<TextField 
				id="date" 
				defaultValue={this.date} 
				type="date" 
				InputLabelProps={{
					shrink: true,
				}} />
			<TextField 
				id="time"
				defaultValue={this.time}
				type="time" 
				InputLabelProps={{
					shrink: true,
				}} />
		</Box>
	}
}