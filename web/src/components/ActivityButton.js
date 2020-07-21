import React from 'react';
import { Button } from '@material-ui/core'

const PrimaryColor = "primary";
const SecondaryColor = "secondary";
const DefaultColor = "default";

export default class ActivityButton extends React.Component {
	render() {
		var firstColor = PrimaryColor;
		var secondColor = SecondaryColor;

		if (this.props.colorChange === "no") {
			firstColor = DefaultColor;
			secondColor = DefaultColor;
		}

		return <Button variant="contained" size="large"
			color={this.props.isActive ? firstColor : secondColor} 
			onClick={() => { this.props.onChange(!this.props.isActive); }}>
			{this.props.isActive ? this.props.buttonNameOff : this.props.buttonNameOn}
		</Button>
	}
}