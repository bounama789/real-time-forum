import { getViewPager, getPage } from "../lib.js"

export function goTo (pageId) {
    const page = getPage(pageId)
    const vp = getViewPager(page.pagerId)
    vp.setCurrentPage(page)
}
