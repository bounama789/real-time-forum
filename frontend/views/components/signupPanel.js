import { Div, TextField, Button, Text } from "../elements/index.js"

export class SignupPanel {
    constructor() {
        return new Div({
            id: "signupPanel",
            className: "panel-transition",
            style: {
                display: "none",
                flexDirection: "column",
                justifyContent: "center",
                // transition: "visibility 0.7s, left 0.7s",
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
                        new Div({
                            style: {
                                display: "flex",
                                flexDirection: "row",
                                justifyContent: "space-between",
                                alignItems: "center",
                                gap: "1rem"
                            },
                            children: [
                                new TextField({
                                    id: "fname",
                                    className: "auth-field",
                                    placeholder: "firstname",
                                    type: "text",
                                }),
                                new TextField({
                                    id: "lname",
                                    className: "auth-field",
                                    placeholder: "lastname",
                                    type: "text",
                                }),

                            ]
                        }),
                        new TextField({
                            id: "uname",
                            className: "auth-field",
                            placeholder: "username",
                            type: "text",
                        }),

                        new TextField({
                            id: "email",
                            className: "auth-field",
                            placeholder: "email",
                            type: "email",
                        }),
                        new TextField({
                            id: "password",
                            className: "auth-field",
                            placeholder: "password",
                            type: "password",
                        }),
                        new TextField({
                            id: "rpass",
                            className: "auth-field",
                            placeholder: "repeat password",
                            type: "password",
                        }),
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