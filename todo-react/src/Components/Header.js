import { AppBar, Link, Toolbar, Typography } from "@mui/material";

function Header() {
    return (
        <div>
            <AppBar>
                <Toolbar>
                    <Typography>
                        <Link href='/' color="inherit" underline="none">
                            Home
                        </Link>
                    </Typography>
                </Toolbar>
            </AppBar>
            <Toolbar />
        </div>
    )
}

export default Header;