import React from "react";
import {
	AppBar,
	Toolbar,
	Typography,
	IconButton,
	Popper,
	MenuList,
	MenuItem,
	ListItemIcon,
	Box
} from "@material-ui/core";
import { Menu, Add, Remove } from "@material-ui/icons";
import { makeStyles } from "@material-ui/core/styles";

const userStyles = makeStyles((theme) => ({
	paper: {
		padding: theme.spacing(1),
		backgroundColor: theme.palette.background.paper
	}
}));

export default function MenuBar(props) {
	const classes = userStyles();
	const [anchorEl, setAnchorEl] = React.useState(null);

	const handleClick = (event) => {
		setAnchorEl(anchorEl ? null : event.currentTarget);
	};

	const open = Boolean(anchorEl);

	return (
		<div>
			<AppBar id="menuBar" position="static">
				<Toolbar>
					<IconButton aria-label="menu" edge="start" onClick={handleClick}>
						<Menu />
					</IconButton>
					<Typography variant="h6">Menu</Typography>
				</Toolbar>
			</AppBar>
			<Popper open={open} anchorEl={anchorEl} placement="bottom">
				<div className={classes.paper}>
					<MenuList>
						<MenuItem
							onClick={() => {
								props.onCallAction("create");
							}}>
							<ListItemIcon>
								<Add />
								<Box pr={3}>
									<Typography variant="subtitle1">Create task</Typography>
								</Box>
							</ListItemIcon>
						</MenuItem>
						<MenuItem
							onClick={() => {
								props.onCallAction("remove");
							}}>
							<ListItemIcon>
								<Remove />
								<Box pr={3}>
									<Typography variant="subtitle1">Delete task</Typography>
								</Box>
							</ListItemIcon>
						</MenuItem>
					</MenuList>
				</div>
			</Popper>
		</div>
	);
}
