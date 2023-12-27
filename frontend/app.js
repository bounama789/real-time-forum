// import { Menu, Navbar } from "./views/components/index.js";
// import { DefaultLayout } from "./views/layout/default_layout.js";
// import { Div, Header } from "./views/elements/index.js";
// import { MainPage } from "./views/pages/index.js";
import { AuthLayout } from "./views/layout/auth-layout.js";
import { SignInPage } from "./views/pages/signin-page.js";

export class App {
  constructor() {
    // return new DefaultLayout({
    //   style: {
    //     width: '100%',
    //     height: '100%',
    //   },
    //   children: [
    //     new Header({
    //       id:"header",
    //       className: "layout-header",
    //       style: {
    //         position: "sticky",
    //         top: 0,
    //         zIndex: 1000,
    //       },
    //       children: [
    //         new Navbar()
    //       ]
    //     }),

    //     new Div({
    //       className: "layout-main",
    //       style: {
    //         width: '100%',
    //         height: '100%',
    //         display: "flex",
    //         flexDirection: "row",
    //         gap:"1.5rem"
    //       },
    //       children: [
    //         new Div({
    //           className: "leftPanel",
    //           style: {
    //             position:"sticky",
    //             flex: 1.5,
    //             display: "flex",
    //             flexDirection: "column",
    //             alignItems: "center",
    //             backgroundColor: 'var(--bs-white)',
    //             boxShadow: "20px 0px 15px -23px rgba(0,0,0,0.1)",
    //             // top:`${getView("header").clientHeight}`
    //           },
    //           children: [
    //             new Menu()
    //           ],
    //         }),
    //         new Div({
    //           className: "mainPanel",
    //           style: {
    //             flex: 6,
    //             height: "100%",

    //           },
    //           children:[
    //             new MainPage(),

    //           ]
    //         }),
    //         new Div({
    //           className: "rightPanel",
    //           style: {
    //             position:"sticky",
    //             flex: 4,
    //             backgroundColor: "var(--bs-white)",
    //             boxShadow: "-20px 0 15px -23px rgba(0,0,0,0.1)"

    //           }
    //         })
    //       ],
    //     })
    //   ],
    // })
    return new AuthLayout({
      style:{
        width: '100M',
        height: '100vh',
      },
      children: [
        new SignInPage()
      ]
    })
  }
}
