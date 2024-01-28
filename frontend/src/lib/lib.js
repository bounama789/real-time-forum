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

export function remView(viewId){
    ViewElement.views.delete(viewId)
}

export function setPage(page){
    ViewElement.pages.set(page.id,page)
}

export function getPage(pagesId){
    return ViewElement.pages.get(pagesId)
}

export function getPageByPath(path){
    const entry =Array.from(ViewElement.pages.entries()).find(([,page])=>page.path === path)
    return entry && entry[1]
}

export function setViewPager(pager){
    ViewElement.pagers.set(pager.id,pager)
}

export function getViewPager(pagerId){
    return ViewElement.pagers.get(pagerId)
}

export function throttle(func, limit) {
  let inThrottle;
  return function() {
    const args = arguments;
    if (!inThrottle) {
      func.apply(this, args);
      inThrottle = true;
      setTimeout(() => inThrottle = false, limit);
    }
  };
}

export function debounce(func, delay){
  let debounceTimer;
  return function(){
    const context = this;
    const args = arguments;
    clearTimeout(debounceTimer);
    debounceTimer = setTimeout(() => func.apply(context, args), delay);
  };
}