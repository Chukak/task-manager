import React from "react";
import { Typography } from "@material-ui/core";
import { styled } from "@material-ui/core/styles";
import { compose, spacing, palette, shadows } from "@material-ui/system";

const StyledBox = styled("div")(compose(spacing, palette, shadows));
const EmptyDateTime = " 0000-00-00 00:00:00 ";

export default class DataTimeField extends React.Component {
	render() {
		var text = EmptyDateTime;
		if (this.props.date !== "") {
			text = this.props.date + " " + this.props.time;
		}

		return (
			<StyledBox
				boxShadow={2}
				bgcolor="#3f51b5"
				color="white"
				fontWeight="fontWeightBold"
				fontFamily="Monospace"
				height="100%"
				p="4px">
				<Typography variant="h6">{text}</Typography>
			</StyledBox>
		);
	}
}
