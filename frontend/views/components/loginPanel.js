import { Div,TextField,Button,Text } from "../elements/index.js"

export class LoginPanel {
    constructor() {
        return new Div({
            id: "loginPanel",
            className: "panel-transition",
            style: {
                display: "flex",
                flexDirection: "column",
                justifyContent: "center",
            },
            children: [
                new Div({
                    style: {
                        display: "flex",
                        flexDirection: "column",
                        justifyContent: "center",
                        alignItems: "center",
                        gap: "2rem",
                        padding: "2rem",
                    },
                    children: [
                        new TextField({
                            id: "login",
                            className: "auth-field",
                            placeholder: "username/email",
                            type: "text",
                        }),
                        new TextField({
                            id: "login",
                            className: "auth-field",
                            placeholder: "password",
                            type: "text",
                        })
                    ]
                }),
                new Button({
                    id: "login",
                    className: "auth-button",
                    children: [
                        new Text({ text: "Login" })
                    ]
                })
            ]
        })
    }
}