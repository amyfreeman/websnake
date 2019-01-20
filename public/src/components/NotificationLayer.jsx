import React from "react";
import Button from "@material-ui/core/Button"
import Snackbar from "@material-ui/core/Snackbar";
import IconButton from '@material-ui/core/IconButton';
import CloseIcon from '@material-ui/icons/Close';

class NotificationLayer extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        var style = {
            position:"fixed",
            zIndex:2,
            left:"0",
            top:"0",
        };
        return (
            <div id="notification-layer" style={style}>
                <Snackbar
                    anchorOrigin={{
                        vertical: 'bottom',
                        horizontal: 'right',
                    }}
                    open={true}
                    autoHideDuration={6000}
                    ContentProps={{
                        'aria-describedby': 'message-id',
                    }}
                    message={<span id="message-id">Note archived</span>}
                    action={[
                        <Button key="undo" color="secondary" size="small" onClick={this.handleClose}>
                        UNDO
                        </Button>,
                        <IconButton
                        key="close"
                        aria-label="Close"
                        color="inherit"
                        >
                        <CloseIcon />
                        </IconButton>,
                    ]}
                />
            </div>
        );
    }
}

export default NotificationLayer;