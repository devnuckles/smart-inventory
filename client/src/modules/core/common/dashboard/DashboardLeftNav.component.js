import * as React from "react";
import PropTypes from "prop-types";
import Tabs from "@mui/material/Tabs";
import Tab from "@mui/material/Tab";
import Typography from "@mui/material/Typography";
import Box from "@mui/material/Box";

function TabPanel(props) {
    const { children, value, index, ...other } = props;

    return (
        <div
            role="tabpanel"
            hidden={value !== index}
            id={`vertical-tabpanel-${index}`}
            aria-labelledby={`vertical-tab-${index}`}
            {...other}
        >
            {value === index && (
                <Box sx={{ p: 3 }}>
                    <Typography>{children}</Typography>
                </Box>
            )}
        </div>
    );
}

TabPanel.propTypes = {
    children: PropTypes.node,
    index: PropTypes.number.isRequired,
    value: PropTypes.number.isRequired,
};

function a11yProps(index) {
    return {
        id: `vertical-tab-${index}`,
        "aria-controls": `vertical-tabpanel-${index}`,
    };
}

const tabsData = [
    { label: "Dashboard", icon: "house-door" },
    { label: "Inventory", icon: "archive" },
    { label: "Reports", icon: "file-text" },
    { label: "Suppliers", icon: "people" },
    { label: "Orders", icon: "cart-plus" },
    { label: "Manage Store", icon: "shop" },
];

export default function VerticalTabs() {
    const [value, setValue] = React.useState(0);

    const handleChange = (event, newValue) => {
        setValue(newValue);
    };

    return (
        <Box
            className="mt-5"
            sx={{
                bgcolor: "background.paper",
                height: "65vh",
                display: "flex",
                marginTop: "20px", // Use Flexbox for horizontal arrangement
            }}
        >
            <Tabs
                orientation="vertical"
                variant="scrollable"
                value={value}
                onChange={handleChange}
                aria-label="Vertical tabs example"
            >
                {tabsData.map((tab, index) => (
                    <Tab
                        key={index}
                        icon={<i className={`bi bi-${tab.icon}`} />}
                        label={tab.label}
                        {...a11yProps(index)}
                        sx={{
                            flexDirection: "row",
                            alignItems: "center",
                            justifyContent: "flex-start", // Align content to the left
                            margin: 0, // Set margin to 0 to reduce the gap
                        }}
                    />
                ))}
            </Tabs>
        </Box>
    );
}
