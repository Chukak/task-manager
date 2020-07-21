import React from 'react';
import { ButtonGroup, IconButton, Button, Typography, 
	Box } from '@material-ui/core';
import { ExposureNeg1, ExposurePlus1 } from '@material-ui/icons';

export default class CounterButton extends React.Component {
	render() {
		return <ButtonGroup size="small" aria-label="small button outlined group">
			<IconButton 
				onClick={() => { 
					if (this.props.max > this.props.count) {
						this.props.onClick(this.props.count + 1);
					}  
				}}>
					<ExposurePlus1 />
 			</IconButton>
			<Box ml={2} mr={2}>
			<Button variant="outlined">
				<Typography variant="h6">
					{this.props.count}
				</Typography>
			</Button>
			</Box>
			<IconButton 
				onClick={() => {
					if (this.props.min < this.props.count) {
						this.props.onClick(this.props.count - 1);
					}
				}}>
				<ExposureNeg1 />
			</IconButton>
		</ButtonGroup>
	}
}