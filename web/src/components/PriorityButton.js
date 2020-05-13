import React from 'react';
import { ButtonGroup, Button } from '@material-ui/core'

export default class PriorityButton extends React.Component {
	constructor(props) {
		super(props);

		this.state = {
			counter: this.props.priority
		};

		this.increment = this.increment.bind(this);
		this.decrement = this.decrement.bind(this);
		this.getCounter = this.props.counterHandler
	}

	increment() {
		// maximum limit priority
		if (this.state.counter < 5) {
			this.setState(state => ({
				counter: state.counter + 1
			}));
			this.getCounter(this.state.counter)
		}
	}

	decrement() {
		// minimum limit priority
		if (this.state.counter > 0) {
			this.setState(state => ({
				counter: state.counter - 1
			}));
			this.getCounter(this.state.counter)
		}
	}

	render() {
		return <ButtonGroup size="small" aria-label="small button outlined group">
			<Button onClick={this.increment}>+</Button>
			<Button disabled>{this.state.counter}</Button>
			<Button onClick={this.decrement}>-</Button>
		</ButtonGroup>
	}
}
