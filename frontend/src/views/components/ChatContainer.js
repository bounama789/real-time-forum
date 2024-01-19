import { getChatByUser } from "../../api/api.js"
import { ChatView } from "./chatView.js"
import { Div } from "../elements/index.js"
import { getView, remView } from "../../lib/lib.js"

export class ChatContainer {
    constructor() {
        this.chatViews = []

        addEventListener("chatOpened", (event) => {
            const user = event.detail

            if (!getView('chat' + user.username)) {
                if (this.chatViews.length == 2) {
                    const view = this.chatViews.shift()
                    view.parentNode.removeChild(view)
                    remView(view.id)
                }
                getChatByUser(user.username).then(chat => {
                    const chatView =new ChatView({ chat, user })
                    const chatViewElement = chatView.element
                    const container = getView("chats-container").element
                    container.appendChild(chatViewElement)
                    this.chatViews.push(chatViewElement)
                })

            }

        })
    }

    get element() {
        return new Div({
            className: "chats-container",
            id: "chats-container",
            style: {
                display: "flex",
                flexDirection: "row-reverse",
                alignItems: "center",
                position: "fixed",
                bottom: 0,
                right: '8%',
                minWidth: 'fit-content',
                maxWidth: '45%',
                maxHeight: '55%',
                gap:"1rem"

            },
            children: [

            ]
        }).element
    }
}