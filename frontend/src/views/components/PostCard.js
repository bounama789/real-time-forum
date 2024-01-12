import { Div, Image, Text } from "../elements/index.js";

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
            padding: "15px",
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
          children: [new Text({ text: postObject.title })],
        }),
        new Div({
            className: "post-content",
            style:{
              lineBreak:"anywhere"
            },
            children: [new Text({ text: postObject.body })],
          }),
      ],
    });
  }
}
