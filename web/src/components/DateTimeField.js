import React from 'react'
import { TextField, Box } from '@material-ui/core'

const dateLength = 10;

export default class DataTimeField extends React.Component {
	constructor(props) {
		super(props)

		let dt = props.dateTime;
		this.state =  {
			date: dt.substring(0, dateLength),
			time: dt.substring(dateLength + 1, dateLength + 1 + 8)
		}
	}

	update(data) {
		let dt = data.dateTime;
		this.setState({
			date: dt.substring(0, dateLength),
			time: dt.substring(dateLength + 1, dateLength + 1 + 8)
		})
	}

	render() {
		return <Box p={2}>
			<TextField 
				id="date" 
				value={this.state.date} 
				type="date" 
				InputLabelProps={{
					shrink: true,
				}} 
				InputProps={{
					disableUnderline: true,
				 }} />
			<TextField 
				id="time"
				value={this.state.time}
				type="time" 
				InputLabelProps={{
					shrink: true,
				}} 
				InputProps={{
					disableUnderline: true,
				}} />
		</Box>
	}
}