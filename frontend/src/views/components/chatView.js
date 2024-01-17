import { getView } from "../../lib/lib.js";
import { Div, Image, MaterialIcon, Text } from "../elements/index.js";
import { MessageView } from "./MessageView.js";

export class chatView {
  constructor(chat) {
    this.chat = chat;
  }

  get element() {
    return new Div({
      id: this.chat.id,
      className: "chat-container",
      style: {
        display: "flex",
        flexDirection: "column",
        width: "360px",
        boxShadow: "20px 0px 15px -23px rgba(0,0,0,0.1)",
      },
      children: [
        new Div({
          id:`chatHeader${this.chat.id}`,
          className: 'chat-header',
          style: {
            backgroundColor: 'rgb(190, 217, 236)',
            display: "flex",
            flexDirection: "row",
            width: "100%",
            padding: "0.2rem .5rem",
            justifyContent: "space-between",
            alignItems: "center",
            borderTopLeftRadius: "10px",
            borderTopRightRadius: "10px",
            transition: "max-height 0.5s ease-out"

          },
          children: [
            new Div({
              className: 'img-dot',
              style: {
                position: "relative",
              },
              children: [
                new Image({
                  src: "https://api.dicebear.com/7.x/avataaars/svg",
                  alt: "Author avatar",
                  style: {
                    width: "32px",
                    height: "32px",
                    borderRadius: "50%",
                    backgroundColor: "var(--bs-gray)",
                  },
                }),
                new Div({
                  className: "dot",
                  style: {
                    position: "absolute",
                    bottom: "5px",
                    left: "3px",
                    width: "10px",
                    height: "10px",
                    borderRadius: "50%",
                    backgroundColor: 'green'
                  },
                }),
              ]
            }),
            new Div({
              className: 'chat-infos',
              children: [
                new Text({ text: "username" })
              ]
            }),
            new MaterialIcon({
              iconName: 'close'
            })
          ],
          listeners:{
            onclick:()=>{
              this.toggleDisplay()
            }
          }
        }),
        new Div({
          id:`chatContainer${this.chat.id}`,
          className:"messages-container",
          style: {
            display: "flex",
            backgroundColor:"aliceblue",
            flexDirection: "column",
            width: "100%",
            padding: "0.2rem 0.5rem",
            height: "100vh",
            transition: "max-height 0.5s ease-out"
          },
          children:[
            new MessageView({})
          ]
        })
      ],
    }).element
  }

  toggleDisplay() {
    let div = getView(`chatContainer${this.chat.id}`).element;
    console.log(getView(`chatContainer${this.chat.id}`));
    div.style.maxHeight = div.style.maxHeight === '0px' ? "100vh" : '0';
}
}
