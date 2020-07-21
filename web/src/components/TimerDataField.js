import React from 'react';
import { Typography } from '@material-ui/core';
import { styled } from '@material-ui/core/styles';
import { compose, spacing, palette, shadows } from '@material-ui/system';

const StyledBox = styled('div')(compose(spacing, palette, shadows));

export default function TimerDataField(props) {
	var text = props.days + ":" + props.hours + ":" +
		+ props.minutes + ":" + props.seconds;

	return <StyledBox boxShadow={2} bgcolor="#3f51b5" color="white"
			fontWeight="fontWeightBold" fontFamily="Monospace" 
			height="100%" p="4px">
			<Typography variant="h4">
					{text}
			</Typography>
		</StyledBox>
};
