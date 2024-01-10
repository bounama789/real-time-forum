import { View } from "../../../common/types/index.js";

export class Text extends View{
    constructor(props){
        super(props)
        this.element = document.createTextNode(props.text)
        this._applyStyles()
        this._setConstraints()

    }
}