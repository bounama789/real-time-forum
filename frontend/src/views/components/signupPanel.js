import { post } from "../../api/api.js";
import { getView } from "../../lib/lib.js";
import { goTo } from "../../lib/pager/navigation.js";
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
                                width: "100%",
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
                    id: "submitButton",
                    className: "auth-button",
                    children: [
                        new Text({ text: "Register" })
                    ],
                    listeners: {
                        onclick: () => {
                            post("/auth/signup", this.formData).then((response) => {
                                if (response.msg === "success") {
                                    localStorage.setItem("auth-token", response.authToken)
                                    goTo("contentPage")
                                }
                            }).catch((error) => console.log(error))
                        }
                    }
                })
            ]
        })
    }

    get formData() {
        return {
            firstname: getView('fname').element.value,
            lastname: getView('lname').element.value,
            username: getView('uname').element.value,
            email: getView('email').element.value,
            password: getView('password').element.value,
        };
    }
}

