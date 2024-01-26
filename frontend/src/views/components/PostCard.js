import {Button, Div, Image, MaterialIcon, Text, TextField} from "../elements/index.js";
import {Form} from "../elements/form";

export class PostCard {
  constructor(postObject) {
    return new Div({
      id: postObject.post_id,
      className: "post-card", 
      style: {
        backgroundColor: "var(--bs-white)",
        width: "100%",
        display: "flex",
        flexDirection: "column",
        alignItems: "start",
        padding: "1rem",
        borderRadius: "10px",
        boxShadow: "0px 10px 15px -3px rgba(0,0,0,0.1)",
      },
      children: [
        new Div({
          className: "post-header",
          style: {
            display: "flex",
            flexDirection: "row",
            alignItems: "center",
            padding: "5px",
            width: "100%",
            height: "100%",
            gap: "1rem",
          },
          children: [
            new Image({
              src: "https://api.dicebear.com/7.x/avataaars/svg",
              alt: "Author avatar",
              style: {
                width: "48px",
                height: "48px",
                borderRadius: "50%",
                backgroundColor: "var(--bs-gray)",
              },
            }),
            new Div({
              className: "post-info",
              style: {
                display: "flex",
                flexDirection: "column",
                alignItems: "start",
                padding: "15px",
                width: "100%",
                height: "100%",
                gap: "1rem",
              },
              children: [
                new Div({
                  children: [new Text({ text: postObject.username })],
                }),
                new Div({
                  children: [new Text({ text: postObject.Age })],
                }),
              ],
            }),
          ],
        }),

        new Div({
          className: "post-title",
          style:{
            fontSize:"large",
            fontWeight:"bold",
            padding: "5px 15px",
            width: "100%",
            gap: "1rem",
            textAlign:"left",
            marginBottom: "0.5rem",
          },
          children: [new Text({ text: postObject.title })],
        }),
        new Div({
            className: "post-content",
            style:{
              overflowWrap:"anywhere"
            },
            children: [new Text({ text: postObject.body })],
          }),
          new Div ({
            className: "post-reactions",
            style:{
              display: "flex",
              flexDirection: "row",
              alignItems: "start",
              justifyContent: "space-between"
            },
            children: [ new Button({name: "like"}), new Button({name:"comments", className:"post-comment"})]
          }),
          new Div({
            className:"CommentPost",
            children:[
                new Form({
                  method: "post",
                  className: "comment-created",
                  id: "comment-created",
                  children:[
                      new Image({
                        src:"https://api.dicebear.com/7.x/avataaars/svg",
                        alt: "New Post",
                        style: {
                          width: "48px",
                          height: "48px",
                          borderRadius: "50%",
                          backgroundColor: "var(--bs-gray)",
                        },}),
                      new TextField({
                        id:"commentInput",
                        className:"comment-input",
                        placeholder: "Enter your comment",
                        maxlenght:"280"
                      }),
                      new Button({
                        type: "submit",
                        className:"confirmCommentInput",
                        children: [new MaterialIcon({iconName:"send"})]
                      })
                  ]
                }),
                new Div({})

            ]
          })
      ],
    });
  }
}
