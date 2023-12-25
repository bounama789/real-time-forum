
export class ViewPager {
    constructor(pages) {
        this.pages = pages
        this.currentPath

    }

    goTo = (path) => {
        this.currentPath = path
        window.history.pushState("object or string", "Title", path);
    }
}