import { View } from "../../common/types/index.js"

export class TextField extends View {
    constructor(props) {
        super(props);
        this.element = document.createElement("input");
        this.type = props.type || "text";
        this._setAttributes()
        this._applyStyles()
        this._setConstraints()

    }
}