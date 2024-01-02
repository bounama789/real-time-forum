import { Page } from "../../common/types/index.js"
import { getView, setPage } from "../../lib/lib.js";
import { Divider } from "../components/divider.js";
import { LoginPanel, SignupPanel } from "../components/index.js";
import { Div, Text } from "../elements/index.js";
import { AuthLayout } from "../layout/auth-layout.js";


export class AuthPage {
    constructor(options) {
        // super(options)
        this.id = "auth"
        this.pagerId = options.pagerId
        this.title = "Authentication"
        this.loginPanel = new LoginPanel()
        this.signupPanel = new SignupPanel()
        this.currentPanel = this.loginPanel

        setPage(this)


    }
    get element() {
        return new AuthLayout({
            style: {
                width: '100M',
                height: '100vh',
            },
            children: [
                new Div({
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
                                    id: "form-wrapper",
                                    style: {
                                        display: "flex",
                                        flexDirection: "row",
                                        // justifyContent: "space-between",
                                        alignItems: "center",
                                        // gap: "1rem"
                                    },
                                    children: [
                                        new Div({
                                            id: "login-label",
                                            className: "auth-label",
                                            style: {
                                                flex: 2,
                                                textAlign: "center",
                                                padding: "1rem"
                                            }, children: [new Text({ text: "Login" })],
                                            listeners: {
                                                onclick: () => this.switchPanel("loginPanel")
                                            }
                                        }),
                                        new Divider({
                                            width: "1px",
                                            height: "100%",
                                        }),
                                        new Div({
                                            id: "register-label",
                                            className: "auth-label",
                                            style: {
                                                flex: 2,
                                                textAlign: "center",
                                                padding: "1rem"
                                            }, children: [new Text({ text: "Register" })],
                                            listeners: {
                                                onclick: () => this.switchPanel("signupPanel")
                                            }
                                        })
                                    ]
                                }),
                                new Divider({
                                    width: "100%",
                                }),
                                this.loginPanel,
                                this.signupPanel
                            ]
                        })
                    ]
                })
            ]
        }).element


    }

    switchPanel(panelId) {
        if (panelId !== this.currentPanel.id) {
            const view = getView(panelId)
            this.currentPanel.element.style.display = "none";
            view.element.style.display = "flex";

            this.currentPanel = view
        }
    }
}