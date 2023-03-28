import { AppBar, Link, Toolbar, Typography } from "@mui/material";

function Footer() {
    return (
        <div>
            <AppBar position="fixed" color="secondary" sx={{ top: 'auto', bottom: 0 }}>
                <Toolbar>
                    <Typography>
                        <Link href='https://github.com/sebastianaldi17/DockerVue-Todo' color="inherit" underline="none">
                            Source code
                        </Link>
                    </Typography>
                </Toolbar>
            </AppBar>
            <Toolbar />
        </div>
    )
}

export default Footer;