import * as React from "react";
import Container from "@mui/material/Container";
import Typography from "@mui/material/Typography";
import Box from "@mui/material/Box";
import Link from "@mui/material/Link";
import ProTip from "./ProTip";

function Copyright() {
    return (
        <Typography
            variant="body2"
            align="center"
            sx={{
                color: "text.secondary",
            }}
        >
            {"Copyright © "}
            <Link color="inherit" href="https://mui.com/">
                Bitcoin Pulse
            </Link>{" "}
            {new Date().getFullYear()}.
        </Typography>
    );
}

function App() {
    return (
        <Container maxWidth="sm">
            <Box sx={{ my: 4 }}>
                <Typography variant="h4" component="h1" sx={{ mb: 2 }}>
                    Material UI Create React App example in TypeScript
                </Typography>
                <ProTip />
                <Copyright />
            </Box>
        </Container>
    );
}

export default App;