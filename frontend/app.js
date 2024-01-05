// import { Menu, Navbar } from "./views/components/index.js";
// import { DefaultLayout } from "./views/layout/default_layout.js";
import { Div } from "./src/views/elements/index.js";
// import { MainPage } from "./views/pages/index.js";
import { ViewPager } from "./src/lib/pager/view-pager.js";
import { ContentPage } from "./src/views/pages/ContentPage.js";
import { AuthPage } from "./src/views/pages/index.js";
import { goTo, router } from "./src/lib/pager/navigation.js";
import { getPageByPath } from "./src/lib/lib.js";

export class App {
  constructor({ container }) {
    this.container = container
    this.currentPath = "/"


    if (location.pathname!== "/") {
      this.currentPath = /\/auth\/\w+/.test(location.pathname) ? "/auth":location.pathname
    }
    // window.addEventListener("hashchange",  () => {
    //   const queryString = window.location.search;
    //   const urlParams = new URLSearchParams(queryString);
    //   const path = urlParams.get('path');
    //   this.currentPath = path
    // })
    // addEventListener("popstate",)
  }
  setContentView() {
    return new ViewPager({
      id: "pager1",
      defaultPage: this.currentPath === "/auth" ? AuthPage : ContentPage,
      pages: [AuthPage, ContentPage],
      container: new Div(),
      locPathname: location.pathname
    })
  }

  run() {
    const content = this.setContentView()
    this.container.appendChild(content.container.element)

    // addEventListener("hashchange", (event) => {
      // event.preventDefault()
      
    // })

  }
}
