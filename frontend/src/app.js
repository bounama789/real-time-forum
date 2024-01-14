// import { Menu, Navbar } from "./views/components/index.js";
// import { DefaultLayout } from "./views/layout/default_layout.js";
import { Div } from "./views/elements/index.js";
// import { MainPage } from "./views/pages/index.js";
import { ViewPager } from "./lib/pager/view-pager.js";
import { ContentPage } from "./views/pages/ContentPage.js";
import { AuthPage } from "./views/pages/index.js";
import { setWSConnection } from "./api/api.js";

export class App {
  constructor(options) {
    this.container = options.container
    this.currentPath = options.currentPath
    this.wsConnection

    history.pushState("app", "", this.currentPath)
  }
  setContentView() {
    return new ViewPager({
      id: "pager1",
      defaultPage: /\/auth\/\w+/.test(this.currentPath) ? { id: "auth", page: AuthPage } : { page: ContentPage, id: "contentPage" },
      pages: [{ id: "auth", page: AuthPage }, { page: ContentPage, id: "contentPage" }],
      container: new Div(),
      locPathname: location.pathname
    })
  }

  run() {

    // addEventListener("load", (event) => {
    const content = this.setContentView()
    this.container.appendChild(content.container.element)
    // })

    addEventListener("logged",()=>{
      setWSConnection()

      this.wsConnection.addEventListener("open", (event) => {
        console.log("ws connection done");
      })
    })

      
    


  }
}
