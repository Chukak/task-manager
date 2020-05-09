import React from 'react'
import { AppBar, Toolbar, Typography, IconButton,
	Popper, MenuList, MenuItem, ListItemIcon, 
	Box } from '@material-ui/core'
import { Menu, Add } from '@material-ui/icons'
import { makeStyles } from "@material-ui/core/styles";


const userStyles = makeStyles(theme => ({
  paper: {
    padding: theme.spacing(1),
    backgroundColor: theme.palette.background.paper
  }
}));

export default function MenuBar() {
  const classes = userStyles();
  const [anchorEl, setAnchorEl] = React.useState(null);

  const handleClick = event => {
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
					<Typography variant="h6">
						Menu
					</Typography>
				</Toolbar>
			</AppBar>
			<Popper 
				open={open}
				anchorEl={anchorEl}
				placement="bottom">
				<div className={classes.paper}>
					<MenuList>
						<MenuItem>
							<ListItemIcon >
								<Add />
								<Box pr={3}>
									<Typography variant="subtitle1">
										Create task
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

/*
const styles = {
  root: {
		background: 'linear-gradient(45deg, #FE6B8B 30%, #FF8E53 90%)',
	}
}

export default class MenuBar extends React.Component {
	constructor(props) {
		super(props)

		this.bar = React.createRef()
		this.state = {
			open: false,
			menuTransform: "translate3d(0px, 0px, 0px)"
		};

		this.handleClick = this.handleClick.bind(this);
	}

	handleClick(e) {
		console.log(e.currentTarget, this.state.open)
		this.setState(prev => ({
			open: !prev.open,
		}));
	}

	componentDidMount() {
		const height = document.getElementById("menuBar").clientHeight;
		this.setState({ 
			menuTransform: "translate3d(" + 0 + "px, " + height + "px, 0px)"
		})
	}

	render() {
		return <div>
			<AppBar id="menuBar" position="static" ref={this.bar}>
				<Toolbar>
					<IconButton aria-label="menu" edge="start" onClick={this.handleClick}>
						<Menu />
					</IconButton>
					<Typography variant="h6">
						Menu
					</Typography>
				</Toolbar>
			</AppBar>
			<Popper 
						open={this.state.open} 
						anchorEl={this.bar}
						placement="bottom">
							<MenuList>
								<MenuItem style={{ transform: this.state.menuTransform }}>
									<ListItemIcon>
										<Add />
										<Typography variant="subtitle1">
											New...
										</Typography>
									</ListItemIcon>
								</MenuItem>
							</MenuList>
					</Popper>
		</div>
	}
} */
