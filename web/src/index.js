import React from "react"
import ReactDOM from "react-dom"
import "./index.css"

// test timer class
class Time extends React.Component {
	constructor(props) {
		super(props);
		this.state = {current: ""};
	}

	render() {
		return (<div>
			<p><strong>Current time: {this.state.current}.</strong></p>
		</div>);
	}

	componentDidMount() {
		this.timer = setInterval(
			() => { this.updateClock() },
			1000);
	}

	componentWillUnmount() {
		clearInterval(this.timer);
	}

	updateClock() {
		fetch('/api/time')
			.then(response => response.json())
				.then(response => this.setState({current: response.time}))
		
	}
}

ReactDOM.render(
	<Time />,
	document.getElementById("root")
);