import { getView } from "../../lib/lib.js";
import { Div, Image, MaterialIcon, Text, TextField } from "../elements/index.js";
import { MessageView } from "./MessageView.js";

export class chatView {
  constructor(chat) {
    this.chat = chat;
  }

  get element() {
    return new Div({
      id: 'chat' + this.chat.id,
      className: "chat-container",
      style: {
        display: "flex",
        flexDirection: "column",
        position:'relative',
        width: "360px",
        boxShadow: "20px 0px 15px -23px rgba(0,0,0,0.1)",
        maxHeight: '55vh',
        backgroundColor: "aliceblue",

      },
      children: [
        new Div({
          id: `chatHeader${this.chat.id}`,
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
              style: {
                display: 'flex',
                flexDirection: 'row',
                width: '100%',
                position: "relative",
                alignItems:"center",
                gap:'.5rem'

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
                        bottom: "5px",
                        left: "3px",
                        width: "10px",
                        height: "10px",
                        borderRadius: "50%",
                        backgroundColor: 'green',
                        position:'absolute'
                      },
                    }),
                  ],
                }),
                new Div({
                  className: 'chat-infos',
                  children: [
                    new Text({ text: "username" })
                  ]
                }),
              ],

              listeners: {
                onclick: () => {
                  this.toggleDisplay()
                }
              }

            }),
            new MaterialIcon({
              iconName: 'close',
              className: "chat-close",
              listeners: {
                onclick: () => {
                  const view = getView('chat' + this.chat.id).element
                  view.parentNode.removeChild(view)
                }
              }
            })
          ],

        }),
        new Div({
          id: `chatContainer${this.chat.id}`,
          className: "messages-container",
          style: {
            display: "flex",
            width: "100%",
            padding: "0.2rem 0.5rem",
            height: "55vh",
            maxHeight: '55vh',
            transition: "max-height 0.5s ease-out",
            flexDirection: "column",
            gap:"2rem"

          },
          children: [
            new Div({
              style: {
                display: "flex",
                flexDirection: "column",
                width: "100%",
                height:"100%",
                justifyContent:'bottom',
                marginBottom:'"3rem',
                justifyContent: 'end'

              },
              children:[
                new MessageView({})
              ]
            }),
            new Div({
              id:'msgtyperWrapper',
              style: {
                width: '100%',
                bottom:"10px",
                padding:"5px 1rem",
                alignSelf:'end',
                display:'flex',
                flexDirection:'row',
                gap:'1rem',
                justifyContent:'center',
                alignItems:'center',
              },
              children: [
                new TextField({
                  placeholder: 'type your message',
                  style: {
                    height:'34px',
                    width: '100%',
                    border:'none',
                    outline:'none',
                    borderRadius:'15px',
                    border: '1px solid var(--bs-blue)',
                    padding:"10px"
                  },  
                }),
                new MaterialIcon({
                  iconName:'send',
                  style:{
                    color:'var(--bs-white)',
                    backgroundColor:'var(--bs-blue)',
                    borderRadius:'10px',
                    padding:'5px'
                  },
                  listeners:{
                    onclick:()=>{
                      //todo handle click
                    }
                  }

                })
              ]
            })
          ]
        }),
      
      ],
    }).element
  }

  toggleDisplay() {
    let div = getView(`chatContainer${this.chat.id}`).element;
    console.log(getView(`chatContainer${this.chat.id}`));
    if (div.style.maxHeight === '0px' ) {
      div.style.maxHeight = "55vh"
      div.style.height = "55vh"
      setTimeout(()=>{
        div.style.visibility = "visible"

      },300)

    } else {
      div.style.maxHeight = "0px"
      div.style.height = "0px"
      // div.style.display = "none"
      div.style.visibility = "hidden"
      

    }
  }
}
