import { Div, Text } from "../elements/index.js";

export class MessageView {
    constructor(message) {
        this.message = message;
    }

    get element() {
        return new Div({
            className: "message-container outgoing",
            style: {
                display: "flex",
                gap:"0.5rem",
                maxWidth: "80%",

            },
            children: [
                new Div({
                    style: {
                        color: "var(--bs-gray)",
                        fontWeight: "bold",
                        alignSelf: "end",
                    },
                    children: [
                        new Text({
                            text: "17:30",

                        })
                    ]
                }),
                new Div({
                    className:"msg-content",
                    style: {
                        fontWeight: "bold",
                        color: "var(--bs-white)",
                        padding: "1rem 1.2rem",
                    },
                    children: [
                        new Text({
                            text: "Hello dear, qdfdf fqsf sdfdfqdf sdsdfdsf sdfsdfqsf fzefaze tetazt rgfe ggre rer er erer r'r ",

                        })
                    ]
                }),
               
            ]
        }).element
    }
}