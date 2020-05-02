import ReactDOM from 'react-dom';
import React from 'react';
import './index.css';
import ListTasks from './components/ListTasks';
import { Container } from '@material-ui/core'

class Main extends React.Component {
	render() {
		return <Container fixed>
				<ListTasks />
		</Container>
	}
};

ReactDOM.render(
	<Main />,
	document.getElementById('root')
);