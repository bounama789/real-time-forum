const ViewElement = {
    views: new Map(),
    pages: new Map(),
    pagers: new Map()
}


export function getView(viewId){
    console.log(ViewElement);
    return ViewElement.views.get(viewId);
}

export function setView(view){
    ViewElement.views.set(view.id,view)
}

export function setPage(page){
    ViewElement.pages.set(page.id,page)
}

export function getPage(pagesId){
    return ViewElement.pages.get(pagesId)
}

export function setViewPager(pager){
    ViewElement.pagers.set(pager.id,pager)
}

export function getViewPager(pagerId){
    return ViewElement.pagers.get(pagerId)
}