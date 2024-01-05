import { Div } from "../../elements/index.js";

export class SuggestionPanel{
    constructor(){
        return new Div({
            className:'suggestion-panel',
            style:{
                flex: 2,
                display: 'flex',
                flexDirection: 'column',
                gap:"1.5rem",
                padding:"20px 0",
                backgroundColor: 'var(--bs-white)',
                height:"100%",
            },
            children: [
            ],
        })
    }
}