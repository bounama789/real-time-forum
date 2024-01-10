import { Div, Nav, MaterialIcon, TextField } from "../../static/scripts/elements/index.js";
import { Logo } from "./logo.js";
export class Navbar {
  constructor() {
    return new Nav({
      className: "navbar",
      children: [
        new Logo(),
        new Div({
          className: "logo-hamburger",
          children: [
            new MaterialIcon({
              iconName: "menu",
            })
          ]
        }),

        new Div({
          className: "search-bar",
          style: {
            width: "35%",
            backgroundColor:"#bed9ec",
            borderRadius: "8px",
            padding: "3px",
          },
          children: [
            new TextField({
              placeholder: "Search here...",
              className: "search-input",
              style: {
                width: "100%",
                border:'none',
                backgroundColor:"transparent",
                outline:'none',
                padding:"0 15px",
                color:"var(--bs-gray-dark)"
              }
            }),
            new MaterialIcon({ iconName: "search" }),
          ]
        }),

        new Div({
          className: "icon-bar",
          children: [
            new MaterialIcon({ iconName: "home"}),
            new MaterialIcon({ iconName: "notifications" }),
            new MaterialIcon({ iconName: "message" }),
            new MaterialIcon({ iconName: "person" }),
          ]
        }),

      ]
    })
  }
}
