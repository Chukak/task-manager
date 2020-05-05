import React from 'react'
import { Button } from '@material-ui/core'

export default class ActivityButton extends React.Component {
	constructor(props) {
		super(props)

		this.state = {
			active: this.props.isActive,
			color: this.props.isActive ? "primary" : "secondary"
		};

		this.changeActive = this.changeActive.bind(this)
	}

	changeActive() {
		this.setState(state => ({
			active: !state.active,
			color: !state.active ? "primary" : "secondary"
		}));
	}

	render() {
		return <Button 
			variant="contained" 
			size="large"
			color={this.state.color}
			onClick={this.changeActive}>
			{this.props.buttonName}
		</Button>
	}
}