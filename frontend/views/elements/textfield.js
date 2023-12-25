import { View } from "../../common/types/index.js"

export class TextField extends View {
    constructor(props) {
        super(props);
        this.element = document.createElement("input");
        this._setAttributes()
        this._applyStyles()
        this._setConstraints()

    }
}