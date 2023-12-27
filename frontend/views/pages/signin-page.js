import { Page } from "../../common/types/index.js"
import { Divider } from "../components/divider.js";
import { Button, Div, Text, TextField } from "../elements/index.js";

export class SignInPage extends Page {
    constructor() {
        super()
        this.title = "Authentication"

        return new Div({
            className: "auth-page",
            style: {
                display: "flex",
                flexDirection: "column",
                justifyContent: "center",
                alignItems: "center",
                paddingTop: "1rem",
                height: "100%",
                gap: "1rem",
            },
            children: [
                new Div({
                    style: {
                        width: "45%",
                        backgroundColor: "white",
                        borderRadius: "5px",
                        paddingBottom: "2rem",
                        display: "flex",
                        flexDirection: "column",
                        justifyContent: "center",
                        boxShadow: "0px 10px 15px -3px rgba(0,0,0,0.1)",
                        maxWidth: "520px",
                    },
                    children: [
                        new Div({
                            style: {
                                display: "flex",
                                flexDirection: "row",
                                // justifyContent: "space-between",
                                alignItems: "center",
                                // gap: "1rem"
                            },
                            children: [
                                new Div({
                                    className: "auth-label",
                                    style: {
                                        flex: 2,
                                        textAlign: "center",
                                        padding: "1rem"
                                    }, children: [new Text({ text: "Login" })]
                                }),
                                new Divider({
                                    width: "1px",
                                    height: "100%",
                                }),
                                new Div({
                                    className: "auth-label",
                                    style: {
                                        flex: 2,
                                        textAlign: "center",
                                        padding: "1rem"
                                    }, children: [new Text({ text: "Register" })]
                                })
                            ]
                        }),
                        new Divider({
                            width: "100%",
                        }),
                        new Div({
                            style:{
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
                        }),

                    ]
                })
            ]

        });
    }

    toggleAuthForm() {

    }
}