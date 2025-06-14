import { getPosts } from "../../../api/api.js";
import { Page } from "../../../common/types/index.js";
import { ListView, PostCard } from "../../components/index.js";
import { Div, Image, Text, TextField } from "../../elements/index.js";

export class MainPage extends Page {
  constructor() {
    super();
    this.title = "Main Page";

    const postList = new ListView({
      id:"postList",
      itemView: PostCard,
      provider: getPosts,
    });

    window.addEventListener('scroll', ()=>{
        if (window.innerHeight + window.scrollY >= document.body.offsetHeight) {
            postList.fetchMoreItems();
        }
    });


    return new Div({
      className: "main-page",
      style: {
        display: "flex",
        flexDirection: "column",
        padding: "20px 15px",
        height: "100%",
        // overflow: "scroll",
      },
      children: [
        new Div({
          className: "new-post-card",
          style: {
            backgroundColor: "var(--bs-white)",
            width: "100%",
            display: "flex",
            flexDirection: "column",
            alignItems: "start",
            padding: "1rem",
            borderRadius: "10px",
          },
          children: [
            new Div({
              children: [
                new Text({
                  text: "Create Post",
                }),
              ],
            }),
            new Div({
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
                  alt: "New Post",
                  style: {
                    width: "48px",
                    height: "48px",
                    borderRadius: "50%",
                    backgroundColor: "var(--bs-gray)",
                  },
                }),
                new TextField({
                  placeholder: "What's on your mind?",
                  style: {
                    width: "100%",
                    height: "3.5rem",
                  },
                }),
              ],
            }),
          ],
        }),
        postList.listContainer
      ],
    });
  }
}
