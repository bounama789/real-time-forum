import { setViewPager } from "../lib.js"

export class ViewPager {
    constructor(options) {
        this.id = options.id
        this.pages = options.pages.map((page) => new page({pagerId:this.id}))
        this.currentPage = options.defaultPage || this.pages[0]
        this.container = options.container
        this.locPathname = options.locPathname

        addEventListener("pageChanged", (event) => {
            const data = event.detail
            if (this.id == data.pagerId) {
                this.container.replaceContent(data.page)
                // window.history.pushState(data.pagerId, data.page.title, data.page.path);
                document.title = data.page.title
            }
        })

        if (this.currentPage) {
            this.setCurrentPage( new this.currentPage({pagerId:this.id,pathname:this.locPathname}))
        }
        setViewPager(this)
    }

    setCurrentPage(page) {
        this.currentPage = page
        const pageChangedEvnt = new CustomEvent("pageChanged", { detail: { pagerId: this.id, page } })
        dispatchEvent(pageChangedEvnt)
    }

}
