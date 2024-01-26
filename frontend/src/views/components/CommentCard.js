import {Div,Image, Text} from "../elements";

export class CommentCard {
constructor(commentObject) {
    return new Div({
        id: commentObject.comment_id,
        classname: "comment-card",
        style: {
            backgroundColor: "var(--bs-white)",
            width: "100%",
            display: "flex",
            flexDirection: "column",
            alignItems: "start",
            padding:"0.5rem",
            borderRadius: "5px"
        },
        children: [
            new Div({
                classname:"comment-header",
                style:{
                    display:"flex",
                    flexDirection:"row",
                    alignItems:"center",
                    padding:"2px",
                    width: "100%",
                    height:"100%",
                    gap:"0.5rem",
                },
                children: [
                    new Image({
                        src: "https://api.dicebear.com/7.x/avataaars/svg",
                        alt: "Author avatar",
                        style: {
                            width: "24px",
                            height: "24px",
                            borderRadius: "50%",
                            backgroundColor: "var(--bs-gray)",
                        },
                    }),
                    new Div({
                        className: "comment-info",
                        style: {
                            display: "flex",
                            flexDirection: "column",
                            alignItems: "start",
                            padding: "5px",
                            width: "100%",
                            height: "100%",
                            gap: "0.5rem",
                        },
                        children: [
                            new Div({
                                children: [new Text({ text: commentObject.username })],
                            }),
                            new Div({
                                children: [new Text({ text: commentObject.Age })],
                            }),
                        ],
                    }),
                ],
            }),

            new Div({
                className: "comment-title",
                style:{
                    fontSize:"large",
                    fontWeight:"bold",
                    padding: "2px 5px",
                    width: "100%",
                    gap: "0.5rem",
                    textAlign:"left",
                    marginBottom: "0.2rem",
                },
                children: [new Text({ text: commentObject.title })],
            }),
            new Div({
                className: "comment-content",
                style:{
                    overflowWrap:"anywhere"
                },
                children: [new Text({ text: commentObject.body })],
            }),
        ],
    })
}
}