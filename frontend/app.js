// import { Menu, Navbar } from "./views/components/index.js";
// import { DefaultLayout } from "./views/layout/default_layout.js";
import { Div } from "./views/elements/index.js";
// import { MainPage } from "./views/pages/index.js";
import { ViewPager } from "./lib/pager/view-pager.js";
import { ContentPage } from "./views/pages/ContentPage.js";
import { AuthPage } from "./views/pages/index.js";

export class App {
  constructor({container}) {
    this.container = container
         
  }
  setContentView(){
    return new ViewPager({
      id:"pager1",
      defaultPage: AuthPage,
      pages:[AuthPage,ContentPage],
      container: new Div()
    })
  }

  run(){
    const content = this.setContentView()
    this.container.appendChild( content.container.element)
  }
}
