import { Div, Text } from "../elements/index.js"

export class Snackbar {
    constructor(error){
        this.error = error
    }

    get element(){
        return new Div({
            className:"snackbar",
            children:[
                new Div({
                    className: "snackbar-message",
                    children:[
                        new Div({
                            children:[
                                new Text({text:this.error.code})
                            ]
                        }),
                        new Div({
                            children:[
                                new Text({text:this.error.msg})
                            ]
                        })
                    ]
                })
            ]
        })
    }

    show() {
        // Append the snackbar to the body
        document.body.appendChild(this.element);

        // After 3 seconds, remove the snackbar
        setTimeout(() => {
            this.element.remove();
        }, 3000);
    }
}