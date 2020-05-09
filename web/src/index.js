import ReactDOM from 'react-dom';
import React from 'react';
import ListTasks from './components/ListTasks';
import MenuBar from './components/MenuBar'
import { Container } from '@material-ui/core'

class Main extends React.Component {
	render() {
		return (
			<div>
				<MenuBar />
				<Container fixed>
					<ListTasks />
				</Container>
			</div>)
	}
};

ReactDOM.render(
	<Main />,
	document.getElementById('root')
);