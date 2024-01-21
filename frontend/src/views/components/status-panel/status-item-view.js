import { setView } from "../../../lib/lib.js";
import { Div, Image, Text } from "../../elements/index.js";

export class StatusItemView {
    constructor(user) {
        this.user = user
        this.id = `user-status${user.username}`

        setView(this)
        return new Div({
            id:`status-item-${this.user.username}`,
            className: "status-item",
            style: {
                display: "flex",
                flexDirection: "row",
                alignItems: "center",
                width: "fit-content",
                height: "fit-content",
                gap: "1rem",
            },
            children: [
                new Div({
                    style: {
                        position: "relative",
                    },
                    children: [
                        new Image({
                            src: `https://api.dicebear.com/7.x/avataaars/svg?seed=${this.user.username}`,
                            alt: "Author avatar",
                            style: {
                                width: "48px",
                                height: "48px",
                                borderRadius: "50%",
                                backgroundColor: "var(--bs-gray)",
                            },
                        }),
                        new Div({
                            className: "dot",
                            id: `${this.user.username}status-dot`,
                            style: {
                                position: "absolute",
                                bottom: "5px",
                                left: "3px",
                                width: "10px",
                                height: "10px",
                                borderRadius: "50%",
                                backgroundColor: user.status === "online" ? "green" : "gray",
                            },
                        }),
                    ]
                }),

                new Div({
                    style: {
                        display: "flex",
                        flexDirection: "column",
                        alignItems: "start",
                        justifyContent: "center"

                    },
                    children: [
                        new Div({
                            style: {
                                color: "var(--bs-blue)"
                            },
                            children: [
                                new Text({
                                    text: this.user.username
                                }),
                            ]
                        }),
                        new Div({
                            id: `${this.user.username}-status-text`,
                            style: {
                                className: "user-status-text",
                                color: "var(--bs-blue)"
                            },
                            children: [
                                new Text({
                                    text: this.user.status
                                }),
                            ]
                        })
                    ]
                })
            ],
            listeners: {
                onclick: () => {
                    const newEvent = new CustomEvent("chatOpened",{detail:this.user})
                    dispatchEvent(newEvent)
                }
            } 
        })
    }
}