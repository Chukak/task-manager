import React from 'react'
import { Typography, IconButton, Popper, MenuList, 
	MenuItem, ListItemIcon, Box } from '@material-ui/core'
import { Menu, Add, Delete } from '@material-ui/icons'
import { makeStyles } from "@material-ui/core/styles";

const userStyles = makeStyles(theme => ({
  paper: {
    padding: theme.spacing(1),
    backgroundColor: theme.palette.background.paper
  }
}));

export default function TaskMenu() {
	const classes = userStyles()
	const [anchorEl, setAnchorEl] = React.useState(null);

	const handleClick = event => {
    setAnchorEl(anchorEl ? null : event.currentTarget);
	};
	
	const open = Boolean(anchorEl);

	return (
		<div>
			<Box>
				<IconButton aria-label="menu" edge="start" onClick={handleClick}>
					<Menu />
				</IconButton>
			</Box>
			<Popper 
				open={open}
				anchorEl={anchorEl}
				placement="bottom">
				<div className={classes.paper}>
					<MenuList>
						<MenuItem>
							<ListItemIcon>
								<Add />
								<Box pr={3}>
									<Typography variant="subtitle1">
										New subtask
									</Typography>
								</Box>
							</ListItemIcon>
						</MenuItem>
						<MenuItem>
							<ListItemIcon>
								<Delete />
								<Box pr={3}>
									<Typography variant="subtitle1">
										Remove task
									</Typography>
								</Box>
							</ListItemIcon>
						</MenuItem>
					</MenuList>
				</div>
			</Popper>
		</div>
	);
}