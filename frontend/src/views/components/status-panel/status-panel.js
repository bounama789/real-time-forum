import { getChats, getUsers } from "../../../api/api.js";
import { Div } from "../../elements/index.js";
import { StatusItemView } from "./status-item-view.js";
import { ListView } from "../ListView.js";

export class StatusPanel {
    constructor() {
        const chatList = new ListView({
            id: "chatList",
            itemView: StatusItemView,
            provider: getChats,
          });
          const otherUser = new ListView({
            id: "userList",
            itemView: StatusItemView,
            provider: getUsers,
          }); 

        return new Div({
            className: 'status-panel',
            style: {
                flex: 1,
                display: 'flex',
                flexDirection: 'column',
                gap: "1.5rem",
                padding: "20px 10px",
                height: "100%",
                backgroundColor: 'var(--bs-white)',
                boxShadow: "20px 0px 15px -23px rgba(0,0,0,0.1)",
                // top:`${getView("header").clientHeight}`
            },
            children: [
                new Div({
                    style: {
                        width: "fit-content",
                        height: "100%",
                        position: "fixed",
                        display: "flex",
                        flexDirection: "column",
                        gap: "1rem",
                        overflowY:"scroll"


                    },
                    children: [
                        new Div({
                            className: "chats",
                            children: [
                                chatList.listContainer,
                                otherUser.listContainer,

                            ]
                        }),
                    ],
                })
            ],
        })
    }
}