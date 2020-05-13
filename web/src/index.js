import ReactDOM from 'react-dom';
import React from 'react';
import ListTasks from './components/ListTasks';
import MenuBar from './components/MenuBar'
import { Container } from '@material-ui/core'

class Main extends React.Component {
	constructor(props) {
		super(props)

		this.state = {
			listTaskRef: React.createRef()
		}
	}

	render() {
		var listTasks = <ListTasks ref={this.state.listTaskRef} />

		return (
			<div>
				<MenuBar listTask={this.state.listTaskRef} />
				<Container  fixed>
					{listTasks}
				</Container>
			</div>)
	}
};

ReactDOM.render(
	<Main />,
	document.getElementById('root')
);