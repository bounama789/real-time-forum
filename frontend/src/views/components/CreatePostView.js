import { setView } from "../../lib/lib";
import { Button, Div, Form, Input, Text } from "../elements/index.js";

export class CreatePost {
  constructor() {
    this.path = "/create-post";
    setView(this);
    this.render();
  }
  render() {
    return new Div({
      id: "create-post-container",
      className: "create-post-container",
      style: {
        display: "none",
        flexDirection: "column",
        alignItems: "center",
        justifyContent: "center",
        width: "100%",
        height: "100%",
        gap: "1rem",
      },
      children: [
        new Div({
          id: "create-post-header",
          className: "create-post-header",
          style: {
            display: "flex",
            flexDirection: "row",
            alignItems: "center",
            justifyContent: "center",
            width: "100%",
            height: "100%",
            gap: "1rem",
          },
          children: [
            new Text({
              text: "Create a post",
              style: {
                fontSize: "1.5rem",
                fontWeight: "bold",
                color: "var(--bs-blue)",
              },
            }),
          ],
        }),
        new Form({
          id: "create-post-form",
          className: "create-post-form",
          style: {
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
            justifyContent: "center",
            width: "100%",
            height: "100%",
            gap: "1rem",
          },
          children: [
            new Input({
              id: "create-post-title",
              className: "create-post-title",
              placeholder: "Title",
              style: {
                width: "100%",
                height: "100%",
                padding: "0.5rem",
                borderRadius: "10px",
                border: "1px solid var(--bs-gray)",
              },
            }),
            new Input({
              id: "create-post-body",
              className: "create-post-body",
              placeholder: "Body",
              style: {
                width: "100%",
                height: "100%",
                padding: "0.5rem",
                borderRadius: "10px",
                border: "1px solid var(--bs-gray)",
              },
            }),
            new Button({
              id: "create-post-submit",
              className: "create-post-submit",
              text: "Create Post",
              style: {
                width: "100%",
                height: "100%",
                padding: "0.5rem",
                borderRadius: "10px",
                border: "1px solid var(--bs-gray)",
                backgroundColor: "var(--bs-blue)",
                color: "white",
                fontWeight: "bold",
              },
              listeners: {
                click: () => {
                  createPost();
                },
              },
            }),
          ],
        }),
      ],
    }).element;
  }
}

function createPost() {
  console.log("Create post button clicked");
}
