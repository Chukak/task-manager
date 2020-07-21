import ReactDOM from 'react-dom';
import React from 'react';
import MenuBar from './components/MenuBar';
import TaskView from './components/TaskView';
import { Container } from '@material-ui/core';

class Main extends React.Component {
	constructor(props) {
		super(props);

		this.taskViewRef = React.createRef();
	}

	render() {
		return (
			<div>
				<MenuBar onCallAction={(action) => {
					if (this.taskViewRef) {
						this.taskViewRef.current.callAction(action);
					}
				}}/>
				<Container fixed>
					<TaskView ref={this.taskViewRef} />
				</Container>
			</div>)
	}
};

ReactDOM.render(
	<Main />,
	document.getElementById('root')
);