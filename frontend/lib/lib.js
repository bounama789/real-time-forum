export const ViewElement = new Map()


export function getView(viewId){
    console.log(ViewElement);
    return ViewElement.get(viewId);
}

export function setView(view){
    ViewElement.set(view.id,view)
}